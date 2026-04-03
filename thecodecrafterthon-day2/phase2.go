package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("\033[32;1m=====BASE CONVERTER CLI TOOL=====\033[0m")
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n\033[34;1m===ENTER REQUIREMENT TO CONVERT (or 'quit' to exit)===\033[0m")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("ERROR READING INPUT")
			continue
		}
		input = strings.TrimSpace(input)

		if input == "" {
			continue
		}

		if strings.ToLower(input) == "quit" {
			fmt.Println("\033[32;1m===GOOD BYE===\033[0m")
			break
		}

		parts := strings.Fields(input)
		if len(parts) < 2 {
			fmt.Println("INVALID INPUT: use 'convert <number> <base>'")
			continue
		}

		var conStr, baseStr string

		if parts[0] == "convert" {
			if len(parts) < 3 {
				fmt.Println("INVALID INPUT: use 'convert <number> <base>'")
				continue
			}
			conStr = parts[1]
			baseStr = strings.ToLower(parts[2])
		} else {
			conStr = parts[0]
			baseStr = strings.ToLower(parts[1])
		}

		var baseNum int
		if baseStr == "dec" {
			baseNum = 10
		} else if baseStr == "hex" {
			baseNum = 16
		} else if baseStr == "bin" {
			baseNum = 2
		} else {
			fmt.Println("UNKNOWN BASE: USE DEC, HEX OR BIN")
			continue
		}

		if baseNum == 2 {
			valid := true
			for i, C := range conStr {
				if i == 0 && C == '-' {
					continue
				}
				if C != '0' && C != '1' {
					valid = false
					break
				}
			}
			if !valid {
				fmt.Println("INVALID NUMBER")
				continue
			}
		}

		num, err := strconv.ParseInt(conStr, baseNum, 64)
		if err != nil {
			if baseNum == 10 {
				fmt.Println("INVALID DECIMAL")
			} else if baseNum == 16 {
				fmt.Println("INVALID HEXADECIMAL")
			} else if baseNum == 2 {
				fmt.Println("INVALID BINARY")
			}
			continue
		}

		if baseNum == 10 {
			fmt.Printf("Binary: %b\n", num)
			fmt.Printf("Hexadecimal: %X\n", num)
		} else {
			fmt.Printf("Decimal: %d\n", num)
		}
	}
}
