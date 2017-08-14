package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	f, err := os.Create("fmt.txt")
	if err != nil {
		log.Fatal(err)
	}
	for x := 1; x <= 9; x++ {
		for y := 1; y <= x; y++ {
			fmt.Fprintf(f, "%d * %d = %2d", y, x, (x * y))
			fmt.Fprint(f, "  ")
		}
		fmt.Fprintln(f, "")
	}

	f.Close()

	w, err := os.OpenFile("fmt.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(w)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(line)
	}

	w.Close()
}
