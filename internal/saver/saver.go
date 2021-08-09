package saver

import (
	"context"
	"github.com/ozoncp/ocp-instruction-api/internal/flusher"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"log"
	"sync"
	"time"
)

type Saver interface {
	Save(entity models.Instruction)
	Close()
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity int, flusher flusher.Flusher, duration time.Duration) Saver {
	ctx, cancelFunc := context.WithCancel(context.Background())
	s := saver{
		capacity:    capacity,
		flusher:     flusher,
		storage:     make([]models.Instruction, 0, capacity),
		beforeClose: cancelFunc,
	}

	go s.ticker(ctx, duration)

	return &s
}

type saver struct {
	capacity int
	flusher  flusher.Flusher

	storage     []models.Instruction
	mu          sync.Mutex
	beforeClose context.CancelFunc
}

func (s *saver) dump() {
	s.mu.Lock()
	defer s.mu.Unlock()

	if len(s.storage) > 0 {
		ret, err := s.flusher.Flush(s.storage)
		if err != nil {
			s.storage = ret
			log.Println("saver dump error:", err)
			return
		}

		s.storage = make([]models.Instruction, 0, s.capacity)
	}
}

func (s *saver) ticker(ctx context.Context, duration time.Duration) {
	tm := time.NewTicker(duration)
	//defer tm.Stop()

	for {
		select {
		case <-tm.C:
			s.dump()
		case <-ctx.Done():
			tm.Stop()
			s.dump()
			return
		}
	}
}

func (s *saver) Save(entity models.Instruction) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.storage = append(s.storage, entity)
}

func (s *saver) Close() {
	s.beforeClose()
}
