package koffing

import "regexp"

// Thanks to https://regexr.com/ to convert these regexes in Go manner.
var (
	Team             = regexp.MustCompile(`^===\s+\[(.*)\]\s+(.*)\s+===$`)
	Gender           = regexp.MustCompile(`\([FM]\)`)
	Item             = regexp.MustCompile(`@\s?(.*)$`)
	Name             = regexp.MustCompile(`(?i)^([^()=@]{2,})`)
	NicknameWithName = regexp.MustCompile(`(?i)^([^()=@]*)\s+\(([^()=@]{2,})\)`)
	Ability          = regexp.MustCompile(`^Ability:\s?(.*)$`)
	Level            = regexp.MustCompile(`^Level:\s?([0-9]{1,3})$`)
	Shiny            = regexp.MustCompile(`^Shiny:\s?(Yes|No)$`)
	Happiness        = regexp.MustCompile(`^Happiness:\s?([0-9]{1,3})$`)
	EIvs             = regexp.MustCompile(`(?i)^([EI]Vs):\s?(.*)$`)
	EIvsValue        = regexp.MustCompile(`(?i)^([0-9]+)\s+(hp|atk|def|spa|spd|spe)$`)
	Nature           = regexp.MustCompile(`^(.*)\s+Nature$`)
	Move             = regexp.MustCompile(`^[-~]\s?(.*)$`)
)
