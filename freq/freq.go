package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"strings"
)

// Q: What is the most common word in sherlock.txt?
// Word frequency

func main() {
	file, err := os.Open("sherlock.txt")
	if err != nil {
		log.Fatalf("error: %s", err)
	}

	defer file.Close()

	wfreq, err := wordFrequency(file)

	if err != nil {
		log.Fatalf("error %s", err)
	}

	max := 0
	w := ""
	for word, v := range wfreq {
		if v > max {
			w = word
			max = v
		}
	}

	fmt.Printf("most frequent word is `%v` with %v frequency", w, max)

}

func wordFrequency(r io.Reader) (map[string]int, error) {
	s := bufio.NewScanner(r)
	freqs := make(map[string]int)
	for s.Scan() {
		words := wordRe.FindAllString(s.Text(), -1)
		for _, w := range words {
			freqs[strings.ToLower(w)]++
		}
	}

	if err := s.Err(); err != nil {
		return nil, err
	}

	return freqs, nil
}

var wordRe = regexp.MustCompile(`[a-zA-Z]+`)
