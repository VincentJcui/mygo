package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	r, w := io.Pipe()
	line := "cat file.txt|grep da"
	cmds := strings.Split(line, "|")
	s1 := strings.Fields(cmds[0])
	s2 := strings.Fields(cmds[1])

	cmd1 := exec.Command(s1[0], s1[1:]...)
	cmd2 := exec.Command(s2[0], s2[1:]...)
	cmd1.Stdin = os.Stdin
	cmd1.Stdout = w
	cmd2.Stdin = r
	cmd2.Stdout = os.Stdout
	cmd1.Start()
	cmd2.Start()
	log.Print("start") //调试,打断点输出内容

	cmd1.Wait()
	//cmd2.Wait()

}
