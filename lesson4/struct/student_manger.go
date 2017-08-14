package main

import "fmt"

func main() {

	var cmd string
	var name string
	var id int
	var line string
	for {
		fmt.Print("> ")
		fmt.Scan(line, &cmd)
		switch cmd {
		case "list":
			fmt.Println("list")
		case "add":
			fmt.Sscan(line, &cmd, &name, &id)
			fmt.Println(name, id, "add is ok")
		}
	}

}
