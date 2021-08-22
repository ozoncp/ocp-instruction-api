package ocp_instruction_api_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestOcpInstructionApi(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OcpInstructionApi Suite")
}
