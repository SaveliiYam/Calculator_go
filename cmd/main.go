package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/SaveliiYam/Calculator_go/service"
)

func main() {
	query, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	//query := "I + X"
	sliceOperands := strings.Split(query, " ")
	if len(sliceOperands) != 3 {
		log.Fatal("Input error")
		return
	}
	sliceOperands[2] = strings.TrimSpace(sliceOperands[2])

	service := service.NewService(sliceOperands[0], sliceOperands[2], sliceOperands[1])

	checkNumbers, err := service.Convertor.CheckNumbers()
	if err != nil {
		log.Fatal(err)
		return
	}
	if checkNumbers {
		answer, err := service.Calculator.GetAnswer()
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(answer)
	} else {
		first, second, err := service.Convertor.FromArabian()
		if err != nil {
			log.Fatal(err)
			return
		}
		service.Calculator.UpdateValues(first, second)
		resultCalc, err := service.Calculator.GetAnswer()
		if err != nil {
			log.Fatal(err)
			return
		}
		answer, err := service.Convertor.FromRome(resultCalc)
		if err != nil {
			log.Fatal(err)
			return
		}
		fmt.Println(answer)
	}
}
