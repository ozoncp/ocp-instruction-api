package ocp_instruction_api

import (
	"context"
	"errors"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
	"github.com/rs/zerolog/log"
)

func (OcpInstructionApi) CreateV1(ctx context.Context, req *desc.CreateV1Request) (*desc.CreateV1Response, error) {
	//return nil, status.Errorf (codes.Unimplemented, "method CreateV1 not implemented")
	log.Debug().Msg("CreateV1")

	if err := req.Validate(); err != nil {
		return &desc.CreateV1Response{}, err
	}

	return &desc.CreateV1Response{}, nil
}

func (OcpInstructionApi) DescribeV1(ctx context.Context, req *desc.DescribeV1Request) (*desc.DescribeV1Response, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method DescribeV1 not implemented")
	log.Debug().Msg("DescribeV1")

	if err := req.Validate(); err != nil {
		return &desc.DescribeV1Response{}, err
	}

	return &desc.DescribeV1Response{}, nil
}

func (OcpInstructionApi) ListV1(ctx context.Context, req *desc.ListV1Request) (*desc.ListV1Response, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method ListV1 not implemented")
	log.Debug().Msg("ListV1")

	if err := req.Validate(); err != nil {
		return &desc.ListV1Response{}, err
	}

	//json request: http://localhost:8081/v1/list?limit=10&offset=10
	log.Debug().Msgf("limit=%v offset=%v", req.Limit, req.Offset)

	return &desc.ListV1Response{Instruction: []*desc.Instruction{}}, nil
}

func (OcpInstructionApi) RemoveV1(ctx context.Context, req *desc.RemoveV1Request) (*desc.RemoveV1Response, error) {
	//return nil, status.Errorf(codes.Unimplemented, "method RemoveV1 not implemented")
	log.Debug().Msg("RemoveV1")

	if err := req.Validate(); err != nil {
		return &desc.RemoveV1Response{}, err
	}

	return &desc.RemoveV1Response{}, errors.New("nothing found")

	//return &desc.RemoveV1Response{}, nil
}
