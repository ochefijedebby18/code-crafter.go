package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("\033[32;1mSENTINENTAL STRING TRANSFORMATION CLI-TOOL\033[0m")
	fmt.Println(" \033[0;31m=====================================\033[0m ")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("INPUT TEXT TO CONVERT:")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "" {
			continue
		}
		if input == "quit" {
			fmt.Println("\033[34mSHUTTING DOWN STRING TRANSFORMATION\033[0m")
			fmt.Println("\033[32;1m===GOOD BYE===\033[0m")
			break
		}
		parts := strings.SplitN(input, " ", 2)
		command := strings.ToLower(parts[0])
		if len(parts) == 1 {
			fmt.Println("✗ No text provided. Usage: <command> <text>")
			continue
		}

		text := parts[1]

		if command == "upper" {
			fmt.Println(toUpper(text))
		} else if command == "lower" {
			fmt.Println(toLower(text))
		} else if command == "cap" {
			fmt.Println(toCap(text))
		} else if command == "title" {
			fmt.Println(toTitle(text))
		} else if command == "snake" {
			fmt.Println(toSnake(text))
		} else if command == "reverse" {
			fmt.Println(toReverse(text))
		} else {
			fmt.Println("\033[31;1mINVALID COMMAND\033[0m")
		}
	}
}

func toUpper(s string) string {
	return strings.ToUpper(s)
}

func toLower(s string) string {
	return strings.ToLower(s)
}

func toCap(s string) string {
	words := strings.Fields(s)
	for i, w := range words {
		if len(w) > 0 {
			words[i] = strings.ToUpper(string(w[0])) + strings.ToLower(w[1:])
		}
	}
	return strings.Join(words, " ")
}

func toTitle(s string) string {
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true, "but": true,
		"or": true, "for": true, "nor": true, "on": true, "at": true,
		"to": true, "by": true, "in": true, "of": true, "up": true,
		"as": true, "is": true, "it": true,
	}

	words := strings.Fields(strings.ToLower(s))

	for i, w := range words {
		if i == 0 || !smallWords[w] {
			if len(w) > 0 {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
		}
	}

	return strings.Join(words, " ")
}

func toSnake(s string) string {
	s = strings.ToLower(s)
	var result []rune
	prev := true

	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= '0' && r <= '9') {
			result = append(result, r)
			prev = false
		} else if r == ' ' {
			if !prev {
				result = append(result, '_')
			}
		}
	}
	return string(result)
}

func toReverse(s string) string {
	words := strings.Split(s, " ")

	for i, w := range words {
		r := []rune(w)
		for a, b := 0, len(r)-1; a < b; a, b = a+1, b-1 {
			r[a], r[b] = r[b], r[a]
		}
		words[i] = string(r)
	}

	return strings.Join(words, " ")
}
