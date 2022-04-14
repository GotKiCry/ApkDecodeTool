package utils

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"log"
	"os/exec"
)

// RunCmd EOF代表程序正常结束，结束不代表正常执行，具体请看输出内容
func RunCmd(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	log.Println("cmd = ", cmd)
	pipe, stdoutErr := cmd.StdoutPipe()
	if stdoutErr != nil {
		return stdoutErr
	}
	cmd.Stderr = cmd.Stdout
	if stdoutErr = cmd.Start(); stdoutErr != nil {
		return stdoutErr
	}

	for {
		temp := make([]byte, 1024)
		_, readErr := pipe.Read(temp)
		decodeTemp, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(temp)
		fmt.Print(string(decodeTemp))
		if readErr != nil {
			return readErr
		}
	}
	if stdoutErr = cmd.Wait(); stdoutErr != nil {
		log.Printf("run on 32 error = %s", stdoutErr)
		return stdoutErr
	}
	return nil
}
