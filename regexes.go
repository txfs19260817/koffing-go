package koffing

import (
	"regexp"
	"strings"
)

// Thanks to https://regexr.com/ to convert these regexes in Go manner.
var (
	teamTagRegex          = regexp.MustCompile(`^===\s+\[(.*)\]\s+(.*)\s+===$`)
	genderRegex           = regexp.MustCompile(`\([FM]\)`)
	itemRegex             = regexp.MustCompile(`@\s?(.*)$`)
	nameRegex             = regexp.MustCompile(`(?i)^([^()=@]{2,})`)
	nicknameWithNameRegex = regexp.MustCompile(`(?i)^([^()=@]*)\s+\(([^()=@]{2,})\)`)
	abilityRegex          = regexp.MustCompile(`^Ability:\s?(.*)$`)
	levelRegex            = regexp.MustCompile(`^Level:\s?([0-9]{1,3})$`)
	shinyRegex            = regexp.MustCompile(`^(?i)Shiny:\s?(Yes|No)$`)
	happinessRegex        = regexp.MustCompile(`^Happiness:\s?([0-9]{1,3})$`)
	eivsRegex             = regexp.MustCompile(`(?i)^([EI]Vs):\s?(.*)$`)
	natureRegex           = regexp.MustCompile(`^(.*)\s+Nature$`)
	moveRegex             = regexp.MustCompile(`^[-~]\s?(.*)$`)
)

// SplitByEmptyNewline splits a multi-line string into parts.
// Each part does not contain any empty line (\n, \r\n or \r).
func SplitByEmptyNewline(s string) []string {
	strNormalized := regexp.MustCompile("\r\n").ReplaceAllString(strings.TrimSpace(s), "\n")
	parts := regexp.MustCompile(`\n\s*\n`).Split(strNormalized, -1)
	return TrimLines(parts)
}

// TrimLines trims space for each element in input string slice,
// and only keep non-empty strings.
func TrimLines(lines []string) []string {
	res := make([]string, 0, len(lines))
	for _, line := range lines {
		if p := strings.TrimSpace(line); len(p) > 0 {
			res = append(res, p)
		}
	}
	return res
}
