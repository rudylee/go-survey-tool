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

	fmt.Println(questions)
	fmt.Println(responses)
}

func print_help() {
	fmt.Println("Example usage:")
	fmt.Println("\t-s, --survey FILE \t Survey data input file")
	fmt.Println("\t-r, --responses FILE \t Responses data input file")
}
