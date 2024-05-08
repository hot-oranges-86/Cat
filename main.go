package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var numberLines = flag.Bool("n", false, "number all output lines")
var squeezeBlanks = flag.Bool("s", false, "suppress repeated empty output lines")

func readLines(r io.Reader) {
	reader := bufio.NewReader(r)
	lineNum := 1
	prevBlank := false

	for {
		line, err := reader.ReadString('\n')

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		currBlank := line == "\n" || line == "\r\n"

		// Handle squeezeBlanks flag
		if *squeezeBlanks && prevBlank && currBlank {
			continue
		}

		// Handle numberLines flag
		if *numberLines {
			fmt.Printf("%d.\t%s", lineNum, line)
		} else {
			fmt.Print(line)
		}

		lineNum++
		prevBlank = currBlank
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
	for i := 0; i < flag.NArg(); i++ {
		cat(flag.Arg(i))
	}
}
