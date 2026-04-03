// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: [Squad Name]
// ───────────────────────────────────────────
// Input line types:
//   [line 1: ALL CAP]
//   [line 2: lower]
//   [line 3: Trimspace]
//   [line 4: TODO ]
//
// Transformation rules (in order):
//   1. Trim all leading and trailing whitespace
//   2. Replace TODO: with ✦ ACTION
//   3. Convert ALL CAPS lines to Title Case
//   4. Convert all lowercase lines to uppercase
//   5. Remove lines that are only dashes or blanks
//
// Output format:
//   Header: SENTINEL FIELD REPORT — PROCESSED
//   Line numbering format: 1.
//
// Terminal summary fields:
//   ✦ Lines read
//   ✦ Lines written
//   ✦ Lines removed
//   ✦ Rules applied
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TrimSpace(text string) string {
	return strings.TrimSpace(text)
}

func replaceTODO(text string) string {
	if strings.HasPrefix(text, "TODO:") {
		return strings.Replace(text, "TODO:", "ACTION:", 1)
	}
	return text
}

func capsToTitle(text string) string {
	if text == strings.ToUpper(text) && text != "" {
		words := strings.Fields(strings.ToLower(text))
		for i, w := range words {
			if len(w) > 0 {
				words[i] = strings.ToUpper(string(w[0])) + w[1:]
			}
		}
		return strings.Join(words, " ")
	}
	return text
}

func lowerToUpper(text string) string {
	if text == strings.ToLower(text) && text != "" {
		return strings.ToUpper(text)
	}
	return text
}

func DashLine(text string) bool {
	if text == "" {
		return false
	}
	for _, ch := range text {
		if ch != '-' {
			return false
		}
	}
	return true
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if input == output {
		fmt.Println("Input and output cannot be the same file.")
		return
	}

	file, err := os.Open(input)
	if err != nil {
		fmt.Printf("File not found: %s\n", input)
		return
	}
	defer file.Close()

	if info, err := os.Stat(output); err == nil && info.IsDir() {
		fmt.Println("Cannot write to output: path is a directory, not a file.")
		return
	}

	outFile, err := os.Create(output)
	if err != nil {
		fmt.Println("Error writing output file")
		return
	}
	defer outFile.Close()

	scanner := bufio.NewScanner(file)
	writer := bufio.NewWriter(outFile)

	lineNumber := 1
	linesRead := 0
	linesWritten := 0
	linesRemoved := 0

	writer.WriteString("SENTINEL FIELD REPORT — PROCESSED\n\n")

	for scanner.Scan() {
		text := scanner.Text()
		linesRead++

		text = TrimSpace(text)
		text = replaceTODO(text)
		text = capsToTitle(text)
		text = lowerToUpper(text)

		if text == "" || DashLine(text) {
			linesRemoved++
			continue
		}

		output := fmt.Sprintf("%d. %s\n", lineNumber, text)
		writer.WriteString(output)

		lineNumber++
		linesWritten++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file")
		return
	}

	if linesRead == 0 {
		fmt.Println("Input file is empty. Nothing to process.")
	}

	writer.Flush()

	fmt.Printf(" - Lines read    : %d\n", linesRead)
	fmt.Printf(" - Lines written : %d\n", linesWritten)
	fmt.Printf(" - Lines removed : %d\n", linesRemoved)
	fmt.Println(" - Rules applied :")
	fmt.Println("  1. Trim all leading and trailing whitespace")
	fmt.Println("  2. Replace TODO: with ✦ ACTION")
	fmt.Println("  3. Convert ALL CAPS lines to Title Case")
	fmt.Println("  4. Convert all lowercase lines to uppercase")
	fmt.Println("  5. Remove lines that are only dashes or blanks")
}
