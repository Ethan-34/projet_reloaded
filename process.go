package main

import (
	"strconv"
	"strings"
)

func Process(text string) string {
	words := strings.Fields(text) 
	words = joinCommands(words)   
	words = applyCommands(words) 
	words = fixArticles(words)  
	result := strings.Join(words, " ")
	result = fixPunctuation(result)
	result = fixQuotes(result)
	return result
}


func joinCommands(words []string) []string {
	var result []string
	for i := 0; i < len(words); i++ {
		word := words[i]
		if strings.HasPrefix(word, "(") && !strings.HasSuffix(word, ")") && i+1 < len(words) {
			word = word + " " + words[i+1]
			i++
		}
		result = append(result, word)
	}
	return result
}

func readCommand(word string) (string, int, bool) {
	if !strings.HasPrefix(word, "(") || !strings.HasSuffix(word, ")") {
		return "", 0, false
	}
	inside := word[1 : len(word)-1]
	parts := strings.Split(inside, ",")
	name := strings.TrimSpace(parts[0])
	count := 1
	if len(parts) == 2 {
		n, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			return "", 0, false
		}
		count = n
	}
	if name != "hex" && name != "bin" && name != "up" && name != "low" && name != "cap" {
		return "", 0, false
	}
	return name, count, true
}
func capitalize(word string) string {
	letters := []rune(word)
	if len(letters) == 0 {
		return word
	}
	return strings.ToUpper(string(letters[0])) + strings.ToLower(string(letters[1:]))
}

func applyCommands(words []string) []string {
	var result []string
	for _, word := range words {
		name, count, ok := readCommand(word)
		if !ok {
			result = append(result, word)
			continue
		}
		if len(result) == 0 {

		}
		last := len(result) - 1

		if name == "hex" || name == "bin" {
			base := 16
			if name == "bin" {
				base = 2
			}
			number, err := strconv.ParseInt(result[last], base, 64)
			if err == nil {
				result[last] = strconv.FormatInt(number, 10)
			}
			continue
		}

		start := len(result) - count
		if start < 0 {
			start = 0
		}
		for i := start; i < len(result); i++ {
			if name == "up" {
				result[i] = strings.ToUpper(result[i])
			} else if name == "low" {
				result[i] = strings.ToLower(result[i])
			} else {
				result[i] = capitalize(result[i])
			}
		}
	}
	return result
}

func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ {
		if words[i] != "a" && words[i] != "A" {
			continue
		}
		first := strings.ToLower(string([]rune(words[i+1])[0]))
		if strings.Contains("aeiouh", first) {
			words[i] = words[i] + "n"
		}
	}
	return words
}

func isPunctuation(r rune) bool {
	return strings.ContainsRune(".,!?:;", r)
}

func fixPunctuation(text string) string {
	var result []rune
	for _, r := range text {
		if isPunctuation(r) {

			for len(result) > 0 && result[len(result)-1] == ' ' {
				result = result[:len(result)-1]
			}
			result = append(result, r)
			continue
		}
		if len(result) > 0 && isPunctuation(result[len(result)-1]) && r != ' ' {
			result = append(result, ' ')
		}
		result = append(result, r)
	}
	return string(result)
}

func isLetter(r rune) bool {
	return strings.ContainsRune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", r)
}

func fixQuotes(text string) string {
	letters := []rune(text)
	var result []rune
	open := false

	for i, r := range letters {
		if r == '\'' {
			before := i > 0 && isLetter(letters[i-1])
			after := i+1 < len(letters) && isLetter(letters[i+1])
			if !(before && after) { // ce n'est pas un apostrophe de mot
				if open {

					for len(result) > 0 && result[len(result)-1] == ' ' {
						result = result[:len(result)-1]
					}
				}
				result = append(result, r)
				open = !open
				continue
			}
		}
		if r == ' ' && open && len(result) > 0 && result[len(result)-1] == '\'' {
			continue
		}
		result = append(result, r)
	}
	return string(result)
}