package consumer

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
)

type MLService interface {
	StartConsuming() error
}

type Service struct {
	brokers []string //var brokers = []string{"127.0.0.1:9094"}
	groupId string
	topic   string

	srv repoService.IRepoService
}

func BuildService(brokers []string, groupId string, topic string) *Service {
	return &Service{
		brokers: brokers,
		groupId: groupId,
		topic:   topic,

		srv: repoService.BuildRequestService(),
	}
}

func (s *Service) StartConsuming(ctx context.Context) error {
	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		return err
	}

	config := sarama.NewConfig()

	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(s.brokers, s.groupId, config)

	if err != nil {
		return err
	}

	consumer := Consumer{}

	ctxRepo := context.WithValue(ctx, "Repo", s.srv)

	go func() {
		for {
			if err := consumerGroup.Consume(ctxRepo, []string{s.topic}, &consumer); err != nil {
				fmt.Printf("Error from consumer: %v", err)
			}
			if ctxRepo.Err() != nil {
				return
			}
		}
	}()

	return nil
}
