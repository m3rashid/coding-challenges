package parser

func ParseJson(json string) bool {
	res := false

	if len(json) == 0 ||
		len(json) == 1 ||
		json[0] != OPEN_BRACKET ||
		json[len(json)-1] != CLOSE_BRACKET {
		return res
	}

	if json == string(OPEN_BRACKET)+string(CLOSE_BRACKET) {
		return true
	}

	isKey := true

	for i := 1; i < len(json)-1; i++ {
		if json[i] == SPACE { // skip spaces
			continue
		}

		if json[i] == COMMA { // handle comma
			if i+1 <= len(json) && json[i+1] == CLOSE_BRACKET {
				return false
			}

			isKey = true
			continue
		}

		if json[i] == COLON { // handle colon
			isKey = false
			continue
		}

		if json[i] == OPEN_BRACKET { // handle opening brackets
			if isKey { // open bracket may occur inside a key
				continue
			}

			// collect the values inside the brackets and recursively call to the function
			var nestedContents string
			nestedContents, i = handleOpenBracket(json, i)
			if !ParseJson(nestedContents) {
				return false
			}

			continue
		}

		if json[i] == OPEN_SQUARE_BRACKET { // handle opening square brackets
			if isKey { // open square bracket may occur inside a key
				continue
			}

			idx, err := handleOpenSquareBracket(json, i, []string{})
			if err != nil {
				return false
			}
			i = idx
			continue
		}

		if isKey {
			// handle double quotes
			if json[i] != DOUBLE_QUOTE {
				return false
			}

			i = handleDoubleQuotes(json, i)
		}

		if !isKey {
			// handle double quotes
			if json[i] == DOUBLE_QUOTE {
				i = handleDoubleQuotes(json, i)
				continue
			}

			// handle true
			if json[i] == 't' {
				i = handleTrue(json, i)
				continue
			}

			// handle false
			if json[i] == 'f' {
				i = hanldeFalse(json, i)
				continue
			}

			// handle null
			if json[i] == 'n' {
				i = handleNull(json, i)
				continue
			}

			// handle numbers
			if json[i] >= '0' && json[i] <= '9' {
				i = handleNumbers(json, i)
				continue
			}

			return false
		}

		// handle closing brackets (this case never comes)
	}

	return true
}
