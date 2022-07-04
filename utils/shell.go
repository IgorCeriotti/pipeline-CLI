package utils

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func ExecuteShell(command string) string {
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s;%s", "set -x", command))
	var outb, call bytes.Buffer
	cmd.Stdout = &outb
	cmd.Stderr = &call
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(call.String(), "SAÍDA:", outb.String())

	return outb.String()
}

//TODO: verificar esse metodo
func ExecuteShellHide(command string, descriptor string) string {
	fmt.Println(descriptor)
	cmd := exec.Command("/bin/bash", "-c", fmt.Sprintf("%s; %s", command, "&>/dev/null"))
	var outb bytes.Buffer
	cmd.Stdout = &outb
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	return outb.String()
}
