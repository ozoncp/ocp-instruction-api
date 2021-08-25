package consumer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	"github.com/rs/zerolog/log"
)

type Consumer struct{}

func (consumer *Consumer) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (consumer *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	ctx := session.Context()
	for message := range claim.Messages() {
		messageReceived(ctx, message)
		session.MarkMessage(message, "")
	}

	return nil
}

func messageReceived(ctx context.Context, message *sarama.ConsumerMessage) {
	fmt.Printf("Analyzing message: %s\n", string(message.Value))
	var msg models.KafkaMessage
	err := json.Unmarshal(message.Value, &msg)
	if err != nil {
		fmt.Printf("Error unmarshalling message: %s\n", err)
	}

	r := ctx.Value("Repo").(*repoService.RepoService)

	switch msg.MessageType {
	case models.KafkaMessageType_Create:
		err := r.Add(ctx, msg.Instruction)
		if err != nil {
			log.Error().Err(err)
		}
	case models.KafkaMessageType_Update:
		err := r.Update(ctx, msg.Instruction[0])
		if err != nil {
			log.Error().Err(err)
		}
	case models.KafkaMessageType_Delete:
		err := r.Remove(ctx, msg.Id)
		if err != nil {
			log.Error().Err(err)
		}
	}
}
