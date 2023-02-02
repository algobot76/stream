package start

import (
	"fmt"
	"github.com/devstream-io/devstream/internal/pkg/start/tool"
)

func Start() error {
	fmt.Println("Let's get started.")

	err := installToolsIfNotExist()
	if err != nil {
		return err
	}

	fmt.Println("Enjoy it!☺️")
	return nil
}

func installToolsIfNotExist() error {
	for _, t := range tool.Tools {
		if !t.Exists() {
			if err := t.Install(); err != nil {
				return err
			}
		}
		if t.Stopped != nil && !t.Stopped() {
			if err := t.Start(); err != nil {
				return err
			}
		}
	}
	return nil
}
