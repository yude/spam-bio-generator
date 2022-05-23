package main

import (
	"bufio"
	"fmt"
	"os"
	"math/rand"
	"time"
	"strings"
	"strconv"
)

func main() {
	templates := load_template("template.txt")
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(templates)
	
	picked := templates[n]
	picked = strings.Replace(picked, "{height}", strconv.Itoa(rand.Intn(1000)), -1)
	picked = strings.Replace(picked, "{age}", strconv.Itoa(rand.Intn(1000)), -1)
	picked = strings.Replace(picked, "{bust}", string('A' + rune(rand.Intn(26))), -1)
	
	fmt.Fprintln(os.Stdout, picked)
}

func load_template(file_path string) []string {
	f, err := os.Open(file_path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to read file %s: %v\n", file_path, err)
		os.Exit(1)
	}
	
	defer f.Close()
	
	lines := make([]string, 0, 25)
	scanner := bufio.NewScanner(f)
	
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if serr := scanner.Err(); serr != nil {
		fmt.Fprintf(os.Stderr, "Error: Failed to scan file %s: %v\n", file_path, err)
	}
	
	return lines
}