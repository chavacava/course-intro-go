package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("proverbs.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close() // new

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		fmt.Printf("%s\n", scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
