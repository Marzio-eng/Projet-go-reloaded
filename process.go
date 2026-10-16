package main

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

func Process(line string) string {
	line = applyMarkers(line)
	line = applyPunctuation(line)
	line = applyQuotes(line)
	line = applyArticles(line)
	return line
}

var markerFinder = regexp.MustCompile(`\(\s*(hex|bin|up|low|cap)\s*(?:,\s*(\d+)\s*)?\)`)
var markerToken = regexp.MustCompile(`^\((hex|bin|up|low|cap)(?:,(\d+))?\)$`)

func applyMarkers(line string) string {
	normalized := markerFinder.ReplaceAllStringFunc(line, func(m string) string {
		sub := markerFinder.FindStringSubmatch(m)
		cmd, n := sub[1], sub[2]
		token := "(" + cmd
		if n != "" {
			token += "," + n
		}
		token += ")"
		return " " + token + " "
	})

	tokens := strings.Fields(normalized)
	out := make([]string, 0, len(tokens))

	for _, tok := range tokens {
		sub := markerToken.FindStringSubmatch(tok)
		if sub == nil {
			out = append(out, tok)
			continue
		}

		cmd := sub[1]
		n := 1
		if sub[2] != "" {
			if v, err := strconv.Atoi(sub[2]); err == nil {
				n = v
			}
		}

		for i := len(out) - 1; i >= 0 && n > 0; i-- {
			if !hasLetter(out[i]) {
				continue
			}
			out[i] = transform(cmd, out[i])
			n--
		}
	}

	return strings.Join(out, " ")
}

func hasLetter(s string) bool {
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func transform(cmd, w string) string {
	switch cmd {
	case "up":
		return strings.ToUpper(w)
	case "low":
		return strings.ToLower(w)
	case "cap":
		return capitalize(w)
	case "hex":
		return convertBase(w, 16)
	case "bin":
		return convertBase(w, 2)
	}
	return w
}

func capitalize(w string) string {
	r := []rune(strings.ToLower(w))
	if len(r) == 0 {
		return w
	}
	r[0] = unicode.ToUpper(r[0])
	return string(r)
}

func convertBase(w string, base int) string {
	n, err := strconv.ParseInt(w, base, 64)
	if err != nil {
		return w
	}
	return strconv.FormatInt(n, 10)
}

var punctBefore = regexp.MustCompile(`\s+([.,!?:;]+)`)
var punctAfter = regexp.MustCompile(`([.,!?:;]+)([^\s.,!?:;])`)

func applyPunctuation(line string) string {
	line = punctBefore.ReplaceAllString(line, "$1")
	line = punctAfter.ReplaceAllString(line, "$1 $2")
	return line
}

func applyQuotes(line string) string {
	tokens := strings.Fields(line)

	const openMark = "\x00OPEN\x00"
	const closeMark = "\x00CLOSE\x00"

	marked := make([]string, 0, len(tokens))
	open := false
	for _, tok := range tokens {
		if tok == "'" {
			if !open {
				marked = append(marked, openMark)
			} else {
				marked = append(marked, closeMark)
			}
			open = !open
			continue
		}
		marked = append(marked, tok)
	}

	result := make([]string, 0, len(marked))
	for i := 0; i < len(marked); i++ {
		switch marked[i] {
		case openMark:
			if i+1 < len(marked) {
				result = append(result, "'"+marked[i+1])
				i++
			}
		case closeMark:
			if len(result) > 0 {
				result[len(result)-1] += "'"
			}
		default:
			result = append(result, marked[i])
		}
	}

	return strings.Join(result, " ")
}

func applyArticles(line string) string {
	tokens := strings.Fields(line)

	for i := 0; i < len(tokens)-1; i++ {
		if tokens[i] != "a" && tokens[i] != "A" {
			continue
		}

		next := strings.TrimPrefix(tokens[i+1], "'")
		r := []rune(next)
		if len(r) == 0 {
			continue
		}

		first := unicode.ToLower(r[0])
		if strings.ContainsRune("aeiouh", first) {
			if tokens[i] == "a" {
				tokens[i] = "an"
			} else {
				tokens[i] = "An"
			}
		}
	}

	return strings.Join(tokens, " ")
}
