package ocp_instruction_api

import (
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
)

type OcpInstructionApi struct {
	desc.UnimplementedOcpInstructionServer
	srv repoService.IRepoService
}

func NewOcpInstructionApi(srv repoService.IRepoService) desc.OcpInstructionServer {
	return &OcpInstructionApi{
		srv: srv,
	}
}

func BuildOcpInstructionApi() desc.OcpInstructionServer {
	return &OcpInstructionApi{
		srv: repoService.BuildRequestService(),
	}
}
