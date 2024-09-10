package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	flag.Parse()

	f, err := os.Open("./log.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	bufReader := bufio.NewReader(f)

	level := flag.String("level", "CRITICAL", "log level to filter for")
	for line, err := bufReader.ReadString('\n'); err == nil; line, err = bufReader.ReadString('\n') {
		if strings.Contains(line, *level) {
			fmt.Println("our critical >>", line)
		} else {
			fmt.Println("standard >>", line)
		}
	}
}
