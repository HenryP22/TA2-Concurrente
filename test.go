package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("dataset.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, "|")

		fmt.Printf("Test: %s | %s | %s | %s | %s | %s | %s | %s\n", items[1], items[2], items[3], items[4], items[5], items[6], items[7], items[8])
		fmt.Println("-------------")
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
