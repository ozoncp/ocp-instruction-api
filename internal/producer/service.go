package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/utils"
	"unsafe"
)

type ProducerService interface {
	CreateMultiV1(context context.Context, instruction []models.Instruction) error
	UpdateV1(context context.Context, instruction models.Instruction) error
	RemoveV1(context context.Context, id uint64) error
}

type Service struct {
	producer  sarama.SyncProducer
	chunkSize int
}

func BuildService(brokers []string, chunkSize int) (*Service, error) {
	config := sarama.NewConfig()
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, err
	}

	return &Service{
		producer:  producer,
		chunkSize: chunkSize,
	}, nil
}

func prepareMessage(ctx context.Context, topic string, message models.KafkaMessage) (*sarama.ProducerMessage, error) {
	span := opentracing.SpanFromContext(ctx)
	m := make(map[string]string)
	carrier := opentracing.TextMapCarrier(m)
	err := opentracing.GlobalTracer().Inject(
		span.Context(),
		opentracing.TextMap,
		carrier,
	)
	if err != nil {
		return nil, err
	}
	message.TraceId = m["uber-trace-id"]

	b, err := json.Marshal(message)
	if err != nil {
		return nil, err
	}

	msg := &sarama.ProducerMessage{
		Topic:     topic,
		Partition: -1,
		Value:     sarama.StringEncoder(string(b)),
	}

	return msg, nil
}

func (s *Service) CreateMultiV1(ctx context.Context, instructions []models.Instruction) error {
	chunks, err := utils.BatchInstructionSlice(instructions, s.chunkSize)
	if err != nil {
		return err
	}

	tracer := opentracing.GlobalTracer()
	span, _ := opentracing.StartSpanFromContext(ctx, "produce CreateMultiV1")
	defer span.Finish()

	for _, chunk := range chunks {
		childSpan := tracer.StartSpan("CreateMultiV1 chunk", opentracing.ChildOf(span.Context()))
		str := fmt.Sprintf("%d", unsafe.Sizeof(chunk))
		childSpan.SetBaggageItem("size", str)
		defer childSpan.Finish()

		kfmsg := models.KafkaMessage{
			MessageType: models.KafkaMessageType_Create,
			Instruction: chunk,
		}

		msg, err := prepareMessage(ctx, "InstructionCUD", kfmsg)
		if err != nil {
			return err
		}

		_, _, err = s.producer.SendMessage(msg)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) UpdateV1(ctx context.Context, instruction models.Instruction) error {
	kfmsg := models.KafkaMessage{
		MessageType: models.KafkaMessageType_Update,
		Instruction: []models.Instruction{instruction},
	}

	msg, err := prepareMessage(ctx, "InstructionCUD", kfmsg)
	if err != nil {
		return err
	}

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}

func (s *Service) RemoveV1(ctx context.Context, id uint64) error {
	kfmsg := models.KafkaMessage{
		MessageType: models.KafkaMessageType_Delete,
		Id:          id,
	}

	msg, err := prepareMessage(ctx, "InstructionCUD", kfmsg)
	if err != nil {
		return err
	}

	_, _, err = s.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	return nil
}
