package ocp_instruction_api_test

import (
	"context"
	"database/sql"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	//"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	ocp_instruction_api "github.com/ozoncp/ocp-instruction-api/internal/app/ocp-instruction-api"
	//"github.com/ozoncp/ocp-instruction-api/internal/mocks"
	pkg_db "github.com/ozoncp/ocp-instruction-api/pkg/db"
	//ocp_instruction_api2 "github.com/ozoncp/ocp-instruction-api/internal/app/ocp-instruction-api"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
)

var _ = Describe("Api handlers", func() {
	var (
		db     *sql.DB
		sqlmck sqlmock.Sqlmock
		ctx    context.Context
		api    desc.OcpInstructionServer
	)

	BeforeEach(func() {
		var err error
		db, sqlmck, err = sqlmock.New()
		if err != nil {
			panic(err)
		}

		ctx = pkg_db.NewContext(context.Background(), db)
		api = ocp_instruction_api.BuildOcpInstructionApi()
	})

	AfterEach(func() {
		db.Close()
	})

	Describe("Create", func() {
		BeforeEach(func() {
		})

		It("save to repo", func() {
			data := &desc.CreateV1Request{
				Instruction: &desc.Instruction{
					Id:          1,
					ClassroomId: 1,
					PrevId:      1,
					Text:        "ssss",
				},
			}

			sqlmck.ExpectExec("INSERT INTO instruction").
				WithArgs(data.Instruction.Id, data.Instruction.ClassroomId, data.Instruction.Text, data.Instruction.PrevId).
				WillReturnResult(sqlmock.NewResult(1, 1))

			resp, err := api.CreateV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})

		It("save error", func() {
			data := &desc.CreateV1Request{
				Instruction: &desc.Instruction{
					Id:          2,
					ClassroomId: 1,
					PrevId:      1,
					Text:        "ssss",
				},
			}

			sqlmck.ExpectExec("INSERT INTO instruction").
				WithArgs(data.Instruction.Id, data.Instruction.ClassroomId, data.Instruction.Text, data.Instruction.PrevId).
				WillReturnError(errors.New("some db error"))

			resp, err := api.CreateV1(ctx, data)
			Ω(err).Should(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})

	Describe("List", func() {
		var (
			rows *sqlmock.Rows
		)
		BeforeEach(func() {
			rows = sqlmck.NewRows([]string{"instruction_id", "text", "prev_id", "classroom_id"}).
				AddRow(1, "row 1", 2, 3).
				AddRow(2, "row 2", 6, 7).
				AddRow(3, "row 3", 24, 33).
				AddRow(4, "row 4", 54, 85).
				AddRow(5, "row 5", 13, 52)
		})

		It("select without limit offset", func() {
			data := &desc.ListV1Request{}

			sqlmck.ExpectQuery("SELECT instruction_id, text, prev_id, classroom_id FROM instruction ORDER BY id").
				WillReturnRows(rows)

			resp, err := api.ListV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
			Ω(len(resp.Instruction)).To(Equal(5))
		})

		It("select limit", func() {
			data := &desc.ListV1Request{Limit: 1}

			sqlmck.ExpectQuery("SELECT instruction_id, text, prev_id, classroom_id FROM instruction ORDER BY id LIMIT 1").
				WillReturnRows(rows)

			resp, err := api.ListV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
		It("select offset", func() {
			data := &desc.ListV1Request{Offset: 1}

			sqlmck.ExpectQuery("SELECT instruction_id, text, prev_id, classroom_id FROM instruction ORDER BY id OFFSET 1").
				WillReturnRows(rows)

			resp, err := api.ListV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})
	})

	Describe("Describe", func() {
		BeforeEach(func() {
		})

		It("select 1 from repo", func() {
			data := &desc.DescribeV1Request{Id: 1}

			sqlmck.ExpectQuery("SELECT instruction_id, text, prev_id, classroom_id FROM instruction WHERE instruction_id = \\$1").
				WithArgs(1).
				WillReturnRows(sqlmck.NewRows([]string{"instruction_id", "text", "prev_id", "classroom_id"}).
					AddRow(data.Id, "ddsd", 2, 3))

			resp, err := api.DescribeV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})

	})

	Describe("Rmove", func() {
		BeforeEach(func() {
		})

		It("remove", func() {
			data := &desc.RemoveV1Request{Id: 1}

			sqlmck.ExpectExec("DELETE FROM instruction WHERE instruction_id = \\$1").
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(0, 1))

			resp, err := api.RemoveV1(ctx, data)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())
		})

		It("remove error", func() {
			data := &desc.RemoveV1Request{Id: 1}

			sqlmck.ExpectExec("DELETE FROM instruction WHERE instruction_id = \\$1").
				WithArgs(1).
				WillReturnResult(sqlmock.NewResult(0, 0))

			resp, err := api.RemoveV1(ctx, data)
			Ω(err).Should(HaveOccurred())
			Ω(resp).ShouldNot(BeNil())

			st, ok := status.FromError(err)
			Ω(ok).Should(BeTrue())
			Ω(st.Code()).Should(Equal(codes.NotFound))
		})

	})

})
