package main

import (
	"regexp"
	"strings"
)

func length(s string) int {
	return len(s)
}

func substr(s string, start, length int) string {
	if start < 0 || start >= len(s) {
		return ""
	}
	end := start + length
	if end > len(s) {
		end = len(s)
	}
	return s[start:end]
}

func index(s, substr string) int {
	return strings.Index(s, substr)
}

func split(s, sep string) []string {
	return strings.Split(s, sep)
}

func sub(regex, replacement, s string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(s, replacement)
}

func gsub(regex, replacement, s string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(s, replacement)
}

func match(regex, s string) int {
	re := regexp.MustCompile(regex)
	loc := re.FindStringIndex(s)
	if loc == nil {
		return -1
	}
	return loc[0]
}

func sprintf(format string, args ...interface{}) string {
	return fmt.Sprintf(format, args...)
}

func tolower(s string) string {
	return strings.ToLower(s)
}

func toupper(s string) string {
	return strings.ToUpper(s)
}