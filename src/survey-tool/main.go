package main

import (
	"bufio"
	"encoding/csv"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {
	helpPtr := flag.Bool("help", false, "Print out help")
	surveyPtr := flag.String("survey", "", "Survey data input file")
	responsesPtr := flag.String("responses", "", "Responses data input file")
	flag.Parse()

	if *helpPtr || *surveyPtr == "" || *responsesPtr == "" {
		print_help()
		return
	}

	surveyFile, err := os.Open(*surveyPtr)
	if err != nil {
		log.Fatal("Can't read survey file")
	}

	responsesFile, err := os.Open(*responsesPtr)
	if err != nil {
		log.Fatal("Can't read responses")
	}

	var questions []question

	surveyReader := csv.NewReader(bufio.NewReader(surveyFile))
	for {
		line, err := surveyReader.Read()
		if err == io.EOF {
			break
		}

		if line[0] == "theme" {
			continue
		}

		q := question{
			theme:        line[0],
			questionType: line[1],
			text:         line[2],
		}

		questions = append(questions, q)
	}

	var responses []response

	responsesReader := csv.NewReader(bufio.NewReader(responsesFile))
	for {
		line, err := responsesReader.Read()
		if err == io.EOF {
			break
		}

		employeeId, _ := strconv.ParseInt(line[1], 10, 0)
		var answers []int64

		// Get the answers
		for i := 3; i < len(line); i++ {
			a, _ := strconv.ParseInt(line[i], 10, 0)
			answers = append(answers, a)
		}

		r := response{
			email:       line[0],
			employeeId:  employeeId,
			submittedAt: line[2],
			answers:     answers,
		}

		responses = append(responses, r)
	}

	// Loop through the questions
	for qIndex, q := range questions {
		fmt.Println("Theme: ", q.theme)
		fmt.Println("Type: ", q.questionType)
		fmt.Println("Question: ", q.text)
		fmt.Println("-----------------------")

		var totalResponses int64 = 0
		var totalRating int64 = 0
		for _, r := range responses {
			if r.submittedAt == "" {
				continue
			}

			totalResponses += 1

			curResponse := r.answers[qIndex]
			totalRating += curResponse
		}

		fmt.Println("Total Responses: ", totalResponses)
		fmt.Println("Total Rating: ", totalRating)
		fmt.Println("Average: ", totalRating/totalResponses)
		fmt.Println("")
	}
}

func print_help() {
	fmt.Println("Example usage:")
	fmt.Println("\t-s, --survey FILE \t Survey data input file")
	fmt.Println("\t-r, --responses FILE \t Responses data input file")
}
