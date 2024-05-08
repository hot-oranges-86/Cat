package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
)

func readLines(r io.Reader) {
	reader := bufio.NewReader(r)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		print(line)
	}
}

func cat(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	readLines(file)
}

func main() {
	flag.Parse()
	cat(flag.Arg(0))

}
