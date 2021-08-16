package ocp_instruction_api

import (
	desc "github.com/ozoncp/ocp-instruction-api/pkg/ocp-instruction-api"
)

type OcpInstructionApi struct {
	desc.UnimplementedOcpInstructionServer
}

func NewOcpInstructionApi() desc.OcpInstructionServer {
	return &OcpInstructionApi{}
}
