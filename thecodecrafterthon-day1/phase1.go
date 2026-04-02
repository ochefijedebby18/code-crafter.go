package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\033[32;1m=====WELCOME TO THE CLI-CALCULATOR=====\033[0m")
	running := true
	for running {
		fmt.Println("\n\033[34;1m===ENTER: NUMBER OPERATOR NUMBER===\033[0m")

		reader := bufio.NewReader(os.Stdin)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println("\033[32;1m===GOOD BYE===\033[0m")
			break
		}
		if input == "help" {
			fmt.Println("\033[31;1mINSTRUCTIONS:\033[0m")
			fmt.Println("multiplying two numbers: use num * num")
			fmt.Println("adding two numbers: use num + num")
			fmt.Println("dividing two numbers: use num / num")
			fmt.Println("subtracting two numbers: use num - num")
			continue
		}

		parts := strings.Fields(input)

		if len(parts) < 3 || len(parts) > 3 {
			fmt.Println("INVALID FORMAT: USE NUMBER OPERATOR NUMBER")
			continue
		}

		num1, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println("INVALID NUMBER")
			fmt.Println("USE NUMBER OPERATOR NUMBER")
			continue
		}
		num2, err := strconv.Atoi(parts[2])
		if err != nil {
			fmt.Println("INVALID NUMBER")
			continue
		}

		operator := parts[1]

		if operator == "+" {
			fmt.Println(num1 + num2)
		} else if operator == "-" {
			fmt.Println(num1 - num2)
		} else if operator == "*" {
			fmt.Println(num1 * num2)
		} else if operator == "/" {
			if num2 == 0 {
				fmt.Println("NOT DIVISIBLE BY ZERO")
				continue
			}
			fmt.Println(num1 / num2)
		} else {
			fmt.Println("INVALID INPUT")
		}
	}
}
