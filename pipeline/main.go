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
  //   1. [Trim all leading and trailing whitespace ]
  //   2. [Replace TODO: with ✦ ACTION]
  //   3. [Convert ALL CAPS lines to Title Case ]
  //   4. [Convert all lowercase lines to uppercase ]
  //   5. [Remove lines that are only dashes or blanks 	]
  //
  // Output format:
  //   [Header: yes/no — SENTINEL FIELD REPORT — PROCESSED]
  //   [Line numbering format- 1]
  //
  // Terminal summary fields:
  //    ✦ Lines read    : [number]                  						
   //   ✦ Lines written : [number]                  						
   //  ✦ Lines removed : [number]                  					
  //   ✦ Rules applied :   
  //   1. [Trim all leading and trailing whitespace ]
  //   2. [Replace TODO: with ✦ ACTION]
  //   3. [Convert ALL CAPS lines to Title Case ]
  //   4. [Convert all lowercase lines to uppercase ]
  //   5. [Remove lines that are only dashes or blanks 	]
  // ═══════════════════════════════════════════


package main 

import (
	"fmt" 
	"os"
	"bufio"
)


func TrimSpace(text string) string {
	return string.Trimspace(text)
}

func replaceTODO(text string) string {
	if strings.HasPrefix(text, "TODO:") {
		return strings.Replace(text, "TODO:", "✦ ACTION:", 1)
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
		return strings.Join(text, " ")
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

func shouldRemove(text string) bool {
	if text == "" {
		return true
	}
	if isDashLine(text) {
		return true
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("usage: go run . <input.txt> <output.txt>")
		return
	}

	input := os.Args[1] 
	input := os.Args[2]

	if input == output {
		fmt.Println("command should not be the same")
		return
	}

	fmt.Printf("Lines read    : %d\n", linesRead)
	fmt.Printf("Lines written : %d\n", linesWritten)
	fmt.Printf("Lines removed : %d\n", linesRemoved)
	fmt.Println(" Rules applied :")
	fmt.Println("   - Trim all leading and trailing whitespace")
	fmt.Println("   - Replace TODO: with ✦ ACTION:")
	fmt.Println("   - Convert ALL CAPS lines to Title Case")
	fmt.Println("   - Convert all lowercase lines to uppercase")
	fmt.Println("   - Remove lines that are only dashes or blanks")
}
