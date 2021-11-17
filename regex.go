package thixo

import (
	"regexp"
)

func regexMatch(regex, s string) bool {
	match, _ := regexp.MatchString(regex, s)
	return match
}

func mustRegexMatch(regex, s string) (bool, error) {
	return regexp.MatchString(regex, s)
}

func regexFindAll(regex, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.FindAllString(s, n)
}

func mustRegexFindAll(regex, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.FindAllString(s, n), nil
}

func regexFind(regex, s string) string {
	r := regexp.MustCompile(regex)
	return r.FindString(s)
}

func mustRegexFind(regex, s string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.FindString(s), nil
}

func regexReplaceAll(regex, s, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllString(s, repl)
}

func mustRegexReplaceAll(regex, s, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllString(s, repl), nil
}

func regexReplaceAllLiteral(regex, s, repl string) string {
	r := regexp.MustCompile(regex)
	return r.ReplaceAllLiteralString(s, repl)
}

func mustRegexReplaceAllLiteral(regex, s, repl string) (string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return "", err
	}
	return r.ReplaceAllLiteralString(s, repl), nil
}

func regexSplit(regex, s string, n int) []string {
	r := regexp.MustCompile(regex)
	return r.Split(s, n)
}

func mustRegexSplit(regex, s string, n int) ([]string, error) {
	r, err := regexp.Compile(regex)
	if err != nil {
		return []string{}, err
	}
	return r.Split(s, n), nil
}

func regexQuoteMeta(s string) string {
	return regexp.QuoteMeta(s)
}
