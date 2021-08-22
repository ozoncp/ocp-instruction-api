package ocp_instruction_api

import (
	"context"
	"github.com/ozoncp/ocp-instruction-api/internal/models"
	"github.com/ozoncp/ocp-instruction-api/internal/repo"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (api *OcpInstructionApi) CreateV1(ctx context.Context, req *desc.CreateV1Request) (*desc.CreateV1Response, error) {
	//return nil, status.Errorf (codes.Unimplemented, "method CreateV1 not implemented")
	log.Debug().Msg("CreateV1")

	if err := req.Validate(); err != nil {
		return &desc.CreateV1Response{}, err
	}

	instr := req.GetInstruction()
	entities := []models.Instruction{{
		Id:          instr.GetId(),
		Text:        instr.GetText(),
		ClassroomId: instr.GetClassroomId(),
		PrevId:      instr.GetPrevId(),
	}}

	if err := api.srv.Add(ctx, entities); err != nil {
		return &desc.CreateV1Response{}, err
	}

	return &desc.CreateV1Response{}, nil
}

func (api *OcpInstructionApi) DescribeV1(ctx context.Context, req *desc.DescribeV1Request) (*desc.DescribeV1Response, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method DescribeV1 not implemented")
	log.Debug().Msg("DescribeV1")

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
	//return nil, status.Errorf(codes.Unimplemented, "method ListV1 not implemented")
	log.Debug().Msg("ListV1")

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
	//return nil, status.Errorf(codes.Unimplemented, "method RemoveV1 not implemented")
	log.Debug().Msg("RemoveV1")

	if err := req.Validate(); err != nil {
		return &desc.RemoveV1Response{}, err
	}

	err := api.srv.Remove(ctx, req.GetId())
	if err != nil {
		if err == repo.ErrNotFound {
			log.Info().Err(err).Msg("id not found")
			return &desc.RemoveV1Response{}, status.Error(codes.NotFound, "id not found")
		}

		log.Error().Err(err).Msg("Remove error")
		return &desc.RemoveV1Response{}, status.Error(codes.Internal, "internal error")
	}

	return &desc.RemoveV1Response{}, nil
}
