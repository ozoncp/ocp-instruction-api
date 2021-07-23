package ocp_instruction_api

import (
	"fmt"
	"os"
)

func FileOpenClose() {
	checkConfig := func(fname string) error {
		file, err := os.Open(fname)
		if err != nil {
			return err
		}

		defer file.Close()

		// some code

		return nil
	}

	files := [...]string{"./config.conf", "/config.conf", "/var/lib/configsStore/config.cfg"}

	for _, confFile := range files {
		err := checkConfig(confFile)
		if err != nil {
			fmt.Println("check config error:", err)
		}
	}
}

