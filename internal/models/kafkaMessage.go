package models

type KafkaMessageType uint

const (
	KafkaMessageType_Create = iota
	KafkaMessageType_Update
	KafkaMessageType_Delete
)

type KafkaMessage struct {
	TraceId     string
	MessageType KafkaMessageType
	Id          uint64
	Instruction []Instruction
}
