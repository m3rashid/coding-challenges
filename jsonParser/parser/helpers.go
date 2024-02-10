package parser

import (
	"errors"
	"internal/stack"
	"strings"
)

const (
	SPACE                = ' '
	OPEN_BRACKET         = '{'
	CLOSE_BRACKET        = '}'
	OPEN_SQUARE_BRACKET  = '['
	CLOSE_SQUARE_BRACKET = ']'
	COMMA                = ','
	COLON                = ':'
	DOUBLE_QUOTE         = '"'
	SINGLE_QUOTE         = '\''
	TRUE                 = "true"
	FALSE                = "false"
	NULL                 = "null"
)

var BRACKET_PAIRS = map[rune]rune{
	OPEN_BRACKET:        CLOSE_BRACKET,
	OPEN_SQUARE_BRACKET: CLOSE_SQUARE_BRACKET,
}

func handleTrue(json string, idx int) int {
	if json[idx:idx+4] == TRUE {
		idx += 3
	}

	return idx
}

func hanldeFalse(json string, idx int) int {
	if json[idx:idx+5] == FALSE {
		idx += 4
	}

	return idx
}

func handleNull(json string, idx int) int {
	if json[idx:idx+4] == NULL {
		idx += 3
	}

	return idx
}

func handleNumbers(json string, idx int) int {
	for j := idx + 1; j < len(json); j++ {
		if json[j] == COMMA || json[j] == CLOSE_BRACKET || json[j] == CLOSE_SQUARE_BRACKET {
			idx = j
			break
		}
	}

	return idx
}

func handleDoubleQuotes(json string, idx int) int {
	for j := idx + 1; j < len(json); j++ {
		if json[j] == DOUBLE_QUOTE {
			idx = j
			break
		}
	}

	return idx
}

func handleOpenBracket(json string, idx int) (string, int) {
	nestedContents := ""
	for j := idx; j < len(json); j++ {
		nestedContents += string(json[j])
		if json[j] == CLOSE_BRACKET {
			idx = j
			break
		}
	}

	return nestedContents, idx
}

func handleOpenSquareBracket(json string, idx int, allContents []string) (int, error) {
	currentStr := ""
	bracketStack := stack.New[rune]() // stack of ] & [

	for j := idx; j < len(json); j++ {
		if rune(json[j]) == OPEN_SQUARE_BRACKET {
			bracketStack.Push(rune(json[j]))
			continue
		}

		if rune(json[j]) == CLOSE_SQUARE_BRACKET {
			if bracketStack.IsEmpty() {
				return 0, errors.New("invalid json")
			}
			bracketStack.Pop()

			if bracketStack.IsEmpty() {
				if strings.Trim(currentStr, " ") != "" {
					allContents = append(allContents, currentStr)
					currentStr = ""
				}

				idx = j
				break
			}
			continue
		}

		currentStr += string(json[j])
	}

	if !bracketStack.IsEmpty() {
		return 0, errors.New("invalid json")
	}

	if len(allContents) == 0 {
		return idx + 1, nil
	}

	for _, content := range allContents {
		if content[0] == '"' { // array of strings
			// !TODO
		} else if content[0] == 't' { // array of booleans
			// !TODO
		} else if content[0] == 'f' { // array of booleans
			// !TODO
		} else if content[0] == 'n' { // array of nulls
			// !TODO
		} else if content[0] >= '0' && content[0] <= '9' { // array of numbers
			// !TODO
		} else if content[0] == OPEN_BRACKET { // array of objects
			// !TODO
		} else {
			return 0, errors.New("invalid json")
		}

	}

	return idx, nil
}
