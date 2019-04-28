package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	var answers []string
	var correctAnswers []string
	correct := 0
	csvFile, err := os.Open("problems.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}
	reader := csv.NewReader(csvFile)

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		fmt.Printf("%s = ", row[0])
		correctAnswers = append(correctAnswers, row[1])

		reader := bufio.NewReader(os.Stdin)
		ans, err := reader.ReadString('\n')
		ans = strings.TrimSuffix(ans, "\n")
		if err != nil {
			log.Println(err)
		}
		answers = append(answers, ans)
	}

	for i := 0; i < len(correctAnswers); i++ {
		if correctAnswers[i] == answers[i] {
			correct++
		}
	}
	fmt.Printf("You Got %d of %d answers correct", correct, len(correctAnswers))
}
