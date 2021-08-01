package flusher_test

import (
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-instruction-api/internal/flusher"
	"github.com/ozoncp/ocp-instruction-api/internal/mocks"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/utils"
)

var _ = Describe("Flusher", func() {
	var (
		ctrl     *gomock.Controller
		mockRepo *mocks.MockRepo
		f        flusher.Flusher
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())
		mockRepo = mocks.NewMockRepo(ctrl)
		f = flusher.NewFlusher(3, mockRepo)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("No saves", func() {
		BeforeEach(func() {
		})

		It("Empty slice", func() {
			mockRepo.EXPECT().
				Add(gomock.Any()).
				MaxTimes(0).
				Return(nil)

			data := []models.Instruction{}
			ret, err := f.Flush(data)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ret).Should(Equal(make([]models.Instruction, 0)))
		})
	})

	Context("1 flush", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().
				Add(gomock.Any()).
				MinTimes(1).
				MaxTimes(1).
				Return(nil)
		})

		It("1 item", func() {
			data := utils.GenerateInstructionSlice(1)
			ret, err := f.Flush(data)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ret).Should(Equal(make([]models.Instruction, 0)))
		})

		It("2 items", func() {
			data := utils.GenerateInstructionSlice(2)
			ret, err := f.Flush(data)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ret).Should(Equal(make([]models.Instruction, 0)))
		})

		It("3 items", func() {
			data := utils.GenerateInstructionSlice(3)
			ret, err := f.Flush(data)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ret).Should(Equal(make([]models.Instruction, 0)))
		})
	})

	Context("2 flushes", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().
				Add(gomock.Any()).
				MinTimes(2).
				MaxTimes(2).
				Return(nil)
		})
		It("4 items", func() {
			data := utils.GenerateInstructionSlice(4)
			ret, err := f.Flush(data)

			Ω(err).ShouldNot(HaveOccurred())
			Ω(ret).Should(Equal(make([]models.Instruction, 0)))
		})
	})

	Context("Repo error", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().
				Add(gomock.Any()).
				AnyTimes().
				Return(errors.New("server down"))
		})
		It("4 items", func() {
			data := utils.GenerateInstructionSlice(4)
			ret, err := f.Flush(data)

			Ω(err).Should(HaveOccurred())
			Ω(ret).Should(Equal(data))
		})
	})

	Context("Repo floating error", func() {
		BeforeEach(func() {
			mockRepo.EXPECT().
				Add(gomock.Any()).
				MinTimes(2).
				MaxTimes(2).
				DoAndReturn(func(entities []models.Instruction) error {
					if len(entities) < 3 {
						return errors.New("overflow")
					}
					return nil
				})
		})

		It("4 items", func() {
			dt1 := utils.GenerateInstructionSlice(3)
			dt2 := utils.GenerateInstructionSlice(1)
			data := append(dt1, dt2...)

			ret, err := f.Flush(data)

			Ω(err).Should(HaveOccurred())
			Ω(ret).Should(Equal(dt2))
		})
	})
})
