package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

type problem struct {
	q string
	a string
}

func main() {
	filename, timeLimit := readArguments()

	file, err := openFile(*filename)

	if err != nil {
		exit(fmt.Sprintf("Failed to open the CSV file: %s\n", *filename))
	}

	problems, err := readCSV(file)

	if err != nil {
		exit(fmt.Sprintf("Failed to parse the provided CSV file. %s\n", *filename))
	}

	score, err := startQuiz(problems, timeLimit)
	if err != nil {
		exit(fmt.Sprintf("Error starting the quiz: %s\n", err))
	}

	fmt.Printf("You scored %d out of %d.\n", score, len(problems))
}

func readArguments() (*string, *int) {
	csvFilename := flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	timeLimit := flag.Int("limit", 30, "the time limit for the quiz in seconds")
	flag.Parse()
	return csvFilename, timeLimit
}

func openFile(filename string) (io.Reader, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func readCSV(file io.Reader) ([]problem, error) {
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	if err != nil {
		return nil, err
	}
	return parseLines(lines), nil
}

func parseLines(lines [][]string) []problem {
	ret := make([]problem, len(lines))
	for i, line := range lines {
		ret[i] = problem{
			q: strings.TrimSpace(line[0]),
			a: strings.TrimSpace(line[1]),
		}
	}
	return ret
}

func exit(msg string) {
	fmt.Println(msg)
	os.Exit(1)
}

func startQuiz(problems []problem, timeLimit *int) (int, error) {
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)
	correct := 0

problemloop:
	for i, p := range problems {
		fmt.Printf("Problem #%d: %s = ", i+1, p.q)

		answerCh := make(chan string)
		go func() {
			var answer string
			fmt.Scanf("%s\n", &answer)
			answerCh <- answer
		}()

		select {
		case <-timer.C:
			fmt.Println()
			break problemloop
		case answer := <-answerCh:
			if answer == p.a {
				correct++
			}
		}
	}
	return correct, nil
}
