package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	filename := flag.String("f", "proverbs.txt", "file containing proverbs")
	flag.Parse()

	err := loadProverbs(*filename)
	if err != nil {
		log.Fatal(err)
	}

	printProverbs()
}

var proverbs = map[string]string{}

func loadProverbs(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("error while opening file %q: %v", filename, err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) < 2 {
			return fmt.Errorf("malformed line: %q", line)
		}

		key := parts[0]
		proverb := parts[1]
		proverbs[key] = proverb
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error while reading the file: %v", err)
	}

	return nil
}

func printProverbs() {
	for i := 1; i <= len(proverbs); i++ {
		key := strconv.Itoa(i)
		fmt.Printf("#%d\t%s\n", i, proverbs[key])
	}
}
