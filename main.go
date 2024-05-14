package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)


func askQuestion(questionlabel string, dataCsv [][]string) {
	scanner := bufio.NewScanner(os.Stdin)
	for _, row := range dataCsv {
		questionConcat := fmt.Sprintf("%s %s", questionlabel, row[0])
		fmt.Print(questionConcat, "? ")
		scanner.Scan()
		text := scanner.Text()
		answer, _ := strconv.Atoi(text)
		correctAnswer, _ := strconv.Atoi(row[1])
		checkAnswerIsCorrect(answer, correctAnswer)

	}
}

func checkAnswerIsCorrect(answer int, correctAnswer int) bool {
	if answer == correctAnswer {
		fmt.Println("Correct!")
		return true
	} else {
		fmt.Println("Incorrect!")
		return false
	}
}

func main() {
	file, err := os.Open("problems.csv")

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	askQuestion("What is", csvData)
}
