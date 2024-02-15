package parser

import (
	"internal/stack"
)

func ParseJson(json string) bool {
	if len(json) == 0 || len(json) == 1 {
		return false
	}

	res := false
	isCurrentKey := false
	expectNextEntry := false
	parenthesisStack := stack.New[rune]()

	for i := 0; i < len(json); i++ {
		if json[i] == SPACE {
			continue
		}

		if json[i] == 't' || json[i] == 'f' || json[i] == 'n' { // (true, false, null) cases
			expectNextEntry = false
			if isCurrentKey {
				return false
			}

			if json[i] == 't' { // must be true and nothing else
				if json[i:i+4] == TRUE {
					i = i + 3
					continue
				}

				return false
			}

			if json[i] == 'f' {
				if json[i:i+5] == FALSE {
					i = i + 4
					continue
				}

				return false
			}

			if json[i] == 'n' {
				if json[i:i+4] == NULL {
					i = i + 3
					continue
				}

				return false
			}

			continue
		}

		if json[i] >= '0' && json[i] <= '9' {
			expectNextEntry = false
			for j := i + 1; j < len(json); j++ {
				if json[j] == COMMA || json[j] == CLOSE_CURLY_BRACKET || json[j] == CLOSE_SQUARE_BRACKET || json[j] == SPACE {
					i = j - 1
					break
				}
			}
			continue
		}

		if json[i] == DOUBLE_QUOTE { // it is a string, escape everything till string ends
			expectNextEntry = false
			j := i + 1
			for ; j < len(json); j++ {
				if json[j] == DOUBLE_QUOTE {
					i = j
					break
				}
			}

			continue
		}

		if json[i] == COLON {
			expectNextEntry = false
			if parenthesisStack.IsEmpty() || parenthesisStack.Top() != OPEN_CURLY_BRACKET {
				return false
			}

			isCurrentKey = false
			continue
		}

		if json[i] == COMMA {
			if parenthesisStack.IsEmpty() {
				return false
			}

			isCurrentKey = true
			expectNextEntry = true
			continue
		}

		if json[i] == OPEN_CURLY_BRACKET {
			expectNextEntry = false
			parenthesisStack.Push(OPEN_CURLY_BRACKET)
			isCurrentKey = true
			continue
		}

		if json[i] == CLOSE_CURLY_BRACKET {
			if expectNextEntry {
				return false
			}

			if parenthesisStack.IsEmpty() || parenthesisStack.Top() != OPEN_CURLY_BRACKET {
				return false // paranthesis does not match
			}

			parenthesisStack.Pop()
			continue
		}

		if json[i] == OPEN_SQUARE_BRACKET {
			expectNextEntry = false
			if isCurrentKey {
				return false
			}

			parenthesisStack.Push(OPEN_SQUARE_BRACKET)
			// handle nested contents
			continue
		}

		if json[i] == CLOSE_SQUARE_BRACKET {
			if expectNextEntry {
				return false
			}

			if parenthesisStack.IsEmpty() || parenthesisStack.Top() != OPEN_SQUARE_BRACKET {
				return false
			}

			parenthesisStack.Pop()
			continue
		}

		return false
	}

	// fmt.Println("complete", parenthesisStack, expectNextEntry)
	if parenthesisStack.IsEmpty() {
		res = true
	}

	if expectNextEntry {
		res = false
	}

	return res
}
