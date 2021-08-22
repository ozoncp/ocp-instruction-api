package ocp_instruction_api

import (
	"io/ioutil"
	"os"
)

type chkConfFunc func(confText string) error

func FileOpenClose(checkConfig chkConfFunc, fname ...string) (retErr error) {
	for _, confFile := range fname {
		file, err := os.Open(confFile)
		if err != nil {
			return err
		}

		defer func(f *os.File) {
			err := f.Close()
			if err != nil && retErr == nil {
				retErr = err
			}
		}(file)

		b, err := ioutil.ReadAll(file)
		if err != nil {
			return err
		}

		err = checkConfig(string(b))
		if err != nil {
			return err
		}
	}

	return retErr
}
