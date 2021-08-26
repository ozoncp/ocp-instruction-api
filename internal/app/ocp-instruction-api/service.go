package ocp_instruction_api

import (
	"github.com/ozoncp/ocp-instruction-api/internal/producer"
	"github.com/ozoncp/ocp-instruction-api/internal/repoService"
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
)

type OcpInstructionApi struct {
	desc.UnimplementedOcpInstructionServer
	srv      repoService.IRepoService
	producer producer.ProducerService
}

func NewOcpInstructionApi(srv repoService.IRepoService, producer producer.ProducerService) desc.OcpInstructionServer {
	return &OcpInstructionApi{
		srv:      srv,
		producer: producer,
	}
}

func BuildOcpInstructionApi() desc.OcpInstructionServer {
	return &OcpInstructionApi{
		srv: repoService.BuildRequestService(),
		//producer: producer.BuildService(),
	}
}
