package koffing

import (
	"regexp"
	"strings"
)

// Thanks to https://regexr.com/ to convert these regexes in Go manner.
var (
	TeamTag          = regexp.MustCompile(`^===\s+\[(.*)\]\s+(.*)\s+===$`)
	Gender           = regexp.MustCompile(`\([FM]\)`)
	Item             = regexp.MustCompile(`@\s?(.*)$`)
	Name             = regexp.MustCompile(`(?i)^([^()=@]{2,})`)
	NicknameWithName = regexp.MustCompile(`(?i)^([^()=@]*)\s+\(([^()=@]{2,})\)`)
	Ability          = regexp.MustCompile(`^Ability:\s?(.*)$`)
	Level            = regexp.MustCompile(`^Level:\s?([0-9]{1,3})$`)
	Shiny            = regexp.MustCompile(`^(?i)Shiny:\s?(Yes|No)$`)
	Happiness        = regexp.MustCompile(`^Happiness:\s?([0-9]{1,3})$`)
	EIvs             = regexp.MustCompile(`(?i)^([EI]Vs):\s?(.*)$`)
	Nature           = regexp.MustCompile(`^(.*)\s+Nature$`)
	Move             = regexp.MustCompile(`^[-~]\s?(.*)$`)
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
