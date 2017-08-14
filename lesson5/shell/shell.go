package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
)

//一
//func main()  {
//	cmd := exec.Command("ls", "-l")
//	out, err := cmd.CombinedOutput()
//	if err != nil{
//		log.Fatal(err)
//	}
//	fmt.Println(string(out))
//}

//二
func main() {
	cmd := exec.Command("ls", "-l")
	out, _ := cmd.StdoutPipe()
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	f := bufio.NewReader(out)
	for {
		line, err := f.ReadString('\n')
		if err != nil {
			break
		}
		fmt.Println(line)
	}
	cmd.Wait()

}
