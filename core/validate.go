package core

import (
	"fmt"
	"log"
	"os/exec"
)

func ValidatePath(val interface{}) error {
	str, ok := val.(string)
	if str == "" || !ok {
		return fmt.Errorf("invalid input")
	}
	cmd := exec.Command("ls", str)

	output, err := cmd.Output()

	if err != nil {
		log.Fatal(fmt.Errorf("failed to execute command: '%w'", err))
	}

	fmt.Println(string(output))

	return nil
}
