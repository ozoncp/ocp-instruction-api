package consumer

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	"github.com/rs/zerolog/log"
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

var consumerGroup sarama.ConsumerGroup

func BuildService(brokers []string, groupId string, topic string) *Service {
	return &Service{
		brokers: brokers,
		groupId: groupId,
		topic:   topic,

		srv: repoService.BuildRequestService(),
	}
}

func (s *Service) Consuming(ctx context.Context) error {
	version, err := sarama.ParseKafkaVersion("2.1.1")
	if err != nil {
		return err
	}

	config := sarama.NewConfig()

	config.Version = version
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err = sarama.NewConsumerGroup(s.brokers, s.groupId, config)
	if err != nil {
		return err
	}

	consumer := Consumer{}

	ctx = context.WithValue(ctx, "Repo", s.srv)

	for {
		if err := consumerGroup.Consume(ctx, []string{s.topic}, &consumer); err != nil {
			log.Error().Msgf("Error from consumer: %v", err)
		}
		err := ctx.Err()
		if err == context.Canceled {
			return nil
		} else if err != nil {
			log.Error().Msgf("Error from consumer: %v", err)
			return nil
		}
	}

	return nil
}

func (s *Service) Close() error {
	if err := consumerGroup.Close(); err != nil && err != context.Canceled {
		return err
	}

	return nil
}
