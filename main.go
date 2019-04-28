package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	correct := 0
	timeLimit := flag.Int("limit", 30, "the time limit for this quiz in seconds")

	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(csvFile)
	lines, err := reader.ReadAll()
	problems := parseLines(lines)
	timer := time.NewTimer(time.Duration(*timeLimit) * time.Second)

	//a label for a loop to easily break out of it
problemLoop:
	for i, p := range problems {
		fmt.Printf("Question %d: %s = ", i+1, p.q)
		//creates a channel to listen for an answer
		answerCh := make(chan string)
		go func() {
			reader := bufio.NewReader(os.Stdin)
			ans, err := reader.ReadString('\n')
			ans = strings.TrimSuffix(ans, "\n")
			if err != nil {
				log.Println(err)
			}
			//send the answer to the listening channel
			answerCh <- ans
		}()
		select {
		case <-timer.C:
			fmt.Printf("\nYou Got %d of %d answers correct", correct, len(problems))
			break problemLoop
		case ans := <-answerCh:
			if ans == p.a {
				correct++
			}
		}

	}
}

func parseLines(lines [][]string) []problems {
	ret := make([]problems, len(lines))
	for i := 0; i < len(lines); i++ {
		ret[i] = problems{
			q: lines[i][0],
			a: strings.TrimSpace(lines[i][1]),
		}
	}
	return ret
}

type problems struct {
	q string
	a string
}
