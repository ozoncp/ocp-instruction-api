package ocp_instruction_api

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFileOpenClose(t *testing.T) {
	failFunc := func(configText string) error {
		return errors.New("config check fail")
	}

	doneFunc := func(configText string) error {
		return nil
	}

	f, err := os.Create("for_testing.cfg")
	if err != nil {
		assert.Fail(t, "error creating test config file")
	}
	f.Close()

	err = FileOpenClose(doneFunc, "not_exist.cfg")
	assert.NotNil(t, err)

	err = FileOpenClose(doneFunc, "for_testing.cfg")
	assert.Nil(t, err)

	err = FileOpenClose(failFunc, "for_testing.cfg")
	assert.NotNil(t, err)

	err = FileOpenClose(doneFunc, "not_exist.cfg", "for_testing.cfg")
	assert.NotNil(t, err)

	err = os.Remove("for_testing.cfg")
	if err != nil {
		assert.Fail(t, "error removing test config file")
	}
}
