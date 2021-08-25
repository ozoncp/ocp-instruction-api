package ocp_instruction_api

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/ozoncp/ocp-instruction-api/internal/metrics"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
	"github.com/rs/zerolog/log"
)

func (api *OcpInstructionApi) CreateV1(ctx context.Context, req *desc.CreateV1Request) (*desc.CreateV1Response, error) {
	log.Debug().Msg("CreateV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "CreateV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("create")

	if err := req.Validate(); err != nil {
		return &desc.CreateV1Response{}, err
	}

	log.Debug().Msg("CreateV1 validated")

	instr := req.GetInstruction()
	entities := []models.Instruction{{
		Id:          instr.GetId(),
		Text:        instr.GetText(),
		ClassroomId: instr.GetClassroomId(),
		PrevId:      instr.GetPrevId(),
	}}

	log.Debug().Msgf("ents: %v", entities)

	err := api.producer.CreateMultiV1(ctx, entities)
	if err != nil {
		return &desc.CreateV1Response{}, err
	}

	return &desc.CreateV1Response{}, nil
}

func (api *OcpInstructionApi) CreateMultiV1(ctx context.Context, req *desc.CreateMultiV1Request) (*desc.CreateMultiV1Response, error) {
	log.Debug().Msg("CreateMultiV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "CreateMultiV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("create_multi")

	if err := req.Validate(); err != nil {
		return &desc.CreateMultiV1Response{}, err
	}

	var entities []models.Instruction
	for _, instr := range req.GetInstruction() {
		entities = append(entities, models.Instruction{
			Id:          instr.GetId(),
			Text:        instr.GetText(),
			ClassroomId: instr.GetClassroomId(),
			PrevId:      instr.GetPrevId(),
		})
	}

	err := api.producer.CreateMultiV1(ctx, entities)
	if err != nil {
		return &desc.CreateMultiV1Response{}, err
	}

	return &desc.CreateMultiV1Response{}, nil
}

func (api *OcpInstructionApi) DescribeV1(ctx context.Context, req *desc.DescribeV1Request) (*desc.DescribeV1Response, error) {
	log.Debug().Msg("DescribeV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "DescribeV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("describe")

	if err := req.Validate(); err != nil {
		return &desc.DescribeV1Response{}, err
	}

	id := req.GetId()
	instr, err := api.srv.Describe(ctx, id)
	if err != nil {
		return &desc.DescribeV1Response{}, err
	}

	ret := &desc.DescribeV1Response{Instruction: &desc.Instruction{
		Id:          instr.Id,
		Text:        instr.Text,
		ClassroomId: instr.ClassroomId,
		PrevId:      instr.PrevId,
	}}
	return ret, nil
}

func (api *OcpInstructionApi) ListV1(ctx context.Context, req *desc.ListV1Request) (*desc.ListV1Response, error) {
	log.Debug().Msg("ListV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "ListV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("list")

	if err := req.Validate(); err != nil {
		return &desc.ListV1Response{}, err
	}

	//json request: http://localhost:8081/v1/list?limit=10&offset=10
	log.Debug().Msgf("limit=%v offset=%v", req.Limit, req.Offset)

	entities, err := api.srv.List(ctx, req.GetLimit(), req.GetOffset())
	if err != nil {
		return &desc.ListV1Response{}, err
	}

	ret := &desc.ListV1Response{
		Instruction: make([]*desc.Instruction, len(entities)),
	}

	for i, ent := range entities {
		ret.Instruction[i] = &desc.Instruction{
			Id:          ent.Id,
			Text:        ent.Text,
			ClassroomId: ent.ClassroomId,
			PrevId:      ent.PrevId,
		}
	}

	return ret, nil
}

func (api *OcpInstructionApi) RemoveV1(ctx context.Context, req *desc.RemoveV1Request) (*desc.RemoveV1Response, error) {
	log.Debug().Msg("RemoveV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "RemoveV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("remove")

	if err := req.Validate(); err != nil {
		return &desc.RemoveV1Response{}, err
	}

	err := api.producer.RemoveV1(ctx, req.GetId())
	if err != nil {
		return &desc.RemoveV1Response{}, err
	}

	return &desc.RemoveV1Response{}, nil
}

func (api *OcpInstructionApi) UpdateV1(ctx context.Context, req *desc.UpdateV1Request) (*desc.UpdateV1Response, error) {
	log.Debug().Msg("UpdateV1")

	span, _ := opentracing.StartSpanFromContext(ctx, "UpdateV1")
	defer span.Finish()

	metrics.OpsCounter_Inc("update")

	if err := req.Validate(); err != nil {
		return &desc.UpdateV1Response{}, err
	}

	entity := req.GetInstruction()

	instr := models.Instruction{
		Id:          entity.GetId(),
		Text:        entity.GetText(),
		ClassroomId: entity.GetClassroomId(),
		PrevId:      entity.GetPrevId(),
	}

	err := api.producer.UpdateV1(ctx, instr)
	if err != nil {
		return &desc.UpdateV1Response{}, err
	}

	return &desc.UpdateV1Response{}, nil
}
