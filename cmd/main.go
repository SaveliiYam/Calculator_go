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
	//query := "1 + 111"
	sliceOperands := strings.Split(query, " ")
	if len(sliceOperands) != 3 {
		log.Panic("Input error")
		return
	}
	sliceOperands[2] = strings.TrimSpace(sliceOperands[2])

	service := service.NewService(sliceOperands[0], sliceOperands[2], sliceOperands[1])

	checkNumbers, err := service.Convertor.CheckNumbers()
	if err != nil {
		log.Panic(err)
		return
	}
	if checkNumbers {
		answer, err := service.Calculator.GetAnswer()
		if err != nil {
			log.Panic(err)
			return
		}
		fmt.Println(answer)
	} else {
		fmt.Println("Nothing")
	}
}
