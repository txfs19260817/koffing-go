package koffing

import (
	"fmt"
	"strconv"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Pokemon struct {
	Name      string `json:"name"`
	Nickname  string `json:"nickname"`
	Gender    string `json:"gender"`
	Item      string `json:"item"`
	Ability   string `json:"ability"`
	Level     int    `json:"level"`
	Shiny     bool   `json:"shiny"`
	Happiness int    `json:"happiness"`
	Nature    string `json:"nature"`
	Evs       struct {
		Hp  int `json:"hp"`
		Atk int `json:"atk"`
		Def int `json:"def"`
		Spa int `json:"spa"`
		Spd int `json:"spd"`
		Spe int `json:"spe"`
	} `json:"evs"`
	Ivs struct {
		Hp  int `json:"hp"`
		Atk int `json:"atk"`
		Def int `json:"def"`
		Spa int `json:"spa"`
		Spd int `json:"spd"`
		Spe int `json:"spe"`
	} `json:"ivs"`
	Moves []string `json:"moves"`
}

func (p *Pokemon) FromJson(j string) error {
	return json.Unmarshal([]byte(j), p)
}

func (p Pokemon) ToJson() (string, error) {
	return json.MarshalToString(p)
}

func (p Pokemon) Validate() error {
	if len(p.Name) == 0 {
		return fmt.Errorf("name is required")
	}
	if len(p.Ability) == 0 {
		return fmt.Errorf("ability is required")
	}
	if len(p.Nature) == 0 {
		return fmt.Errorf("nature is required")
	}
	if p.Happiness < 0 || p.Happiness > 255 {
		return fmt.Errorf("happiness should be in range [0, 255], yours: %d", p.Happiness)
	}
	if p.Evs.Hp < 0 || p.Evs.Hp > 252 {
		return fmt.Errorf("the HP ev should be in range [0, 252], yours: %d", p.Evs.Hp)
	}
	if p.Evs.Atk < 0 || p.Evs.Atk > 252 {
		return fmt.Errorf("the Atk ev should be in range [0, 252], yours: %d", p.Evs.Atk)
	}
	if p.Evs.Def < 0 || p.Evs.Def > 252 {
		return fmt.Errorf("the Def ev should be in range [0, 252], yours: %d", p.Evs.Def)
	}
	if p.Evs.Spa < 0 || p.Evs.Spa > 252 {
		return fmt.Errorf("the Spa ev should be in range [0, 252], yours: %d", p.Evs.Spa)
	}
	if p.Evs.Spd < 0 || p.Evs.Spd > 252 {
		return fmt.Errorf("the Spd ev should be in range [0, 252], yours: %d", p.Evs.Spd)
	}
	if p.Evs.Spe < 0 || p.Evs.Spe > 252 {
		return fmt.Errorf("the Spe ev should be in range [0, 252], yours: %d", p.Evs.Spe)
	}
	if p.Ivs.Hp < 0 || p.Ivs.Hp > 31 {
		return fmt.Errorf("the HP iv should be in range [0, 31], yours: %d", p.Ivs.Hp)
	}
	if p.Ivs.Atk < 0 || p.Ivs.Atk > 31 {
		return fmt.Errorf("the Atk iv should be in range [0, 31], yours: %d", p.Ivs.Atk)
	}
	if p.Ivs.Def < 0 || p.Ivs.Def > 31 {
		return fmt.Errorf("the Def iv should be in range [0, 31], yours: %d", p.Ivs.Def)
	}
	if p.Ivs.Spa < 0 || p.Ivs.Spa > 31 {
		return fmt.Errorf("the Spa iv should be in range [0, 31], yours: %d", p.Ivs.Spa)
	}
	if p.Ivs.Spd < 0 || p.Ivs.Spd > 31 {
		return fmt.Errorf("the Spd iv should be in range [0, 31], yours: %d", p.Ivs.Spd)
	}
	if p.Ivs.Spe < 0 || p.Ivs.Spe > 31 {
		return fmt.Errorf("the Spe iv should be in range [0, 31], yours: %d", p.Ivs.Spe)
	}
	if len(p.Moves) < 1 || len(p.Moves) > 4 {
		return fmt.Errorf("the number of moves should be in range [1, 4], yours: %d", len(p.Moves))
	}
	if len(p.Gender) > 0 && p.Gender != "F" && p.Gender != "M" {
		return fmt.Errorf("invalid gender: [%s]", p.Gender)

	}
	return nil
}

func (p *Pokemon) FromShowdown(s string) error {
	lines := TrimLines(strings.Split(s, "\n"))
	if len(lines) < 3 {
		return fmt.Errorf("invalid pokemon input: %s", s)
	}
	// name line - name/nickname
	if NicknameWithName.MatchString(lines[0]) {
		submatch := NicknameWithName.FindStringSubmatch(lines[0])
		if len(submatch) != 3 {
			return fmt.Errorf("invalid name with nickname: %s", lines[0])
		}
		p.Nickname, p.Name = submatch[1], submatch[2]
	} else if Name.MatchString(lines[0]) {
		p.Name = strings.TrimSpace(Name.FindString(lines[0]))
	} else {
		return fmt.Errorf("invalid name: %s", lines[0])
	}
	// name line - gender
	if Gender.MatchString(lines[0]) {
		p.Gender = string(Gender.FindString(lines[0])[1])
	}
	// name line - item
	if Item.MatchString(lines[0]) {
		p.Item = Item.FindStringSubmatch(lines[0])[1]
	}
	// init with some default values
	p.Happiness = 255
	p.Moves = make([]string, 0, 4)
	p.Ivs.Hp, p.Ivs.Atk, p.Ivs.Def, p.Ivs.Spa, p.Ivs.Spd, p.Ivs.Spe = 31, 31, 31, 31, 31, 31
	// other lines
	for _, line := range lines[1:] {
		switch {
		case Ability.MatchString(line):
			p.Ability = Ability.FindStringSubmatch(line)[1]
		case Level.MatchString(line):
			level, err := strconv.Atoi(Level.FindStringSubmatch(line)[1])
			if err != nil {
				return fmt.Errorf("invalid level: %w", err)
			}
			p.Level = level
		case Shiny.MatchString(line):
			p.Shiny = strings.ToLower(Shiny.FindStringSubmatch(line)[1]) == "yes"
		case Happiness.MatchString(line):
			happiness, err := strconv.Atoi(Happiness.FindStringSubmatch(line)[1])
			if err != nil {
				return fmt.Errorf("invalid happiness: %w", err)
			}
			p.Happiness = happiness
		case Nature.MatchString(line):
			p.Nature = Nature.FindStringSubmatch(line)[1]
		case EIvs.MatchString(line):
			m, prop, err := fromEIvsLineToMap(line)
			if err != nil {
				return fmt.Errorf("error in parsing evs/ivs line: %w", err)
			}
			if strings.Contains(prop, "E") {
				p.Evs.Hp = m["HP"]
				p.Evs.Atk = m["Atk"]
				p.Evs.Def = m["Def"]
				p.Evs.Spa = m["SpA"]
				p.Evs.Spd = m["SpD"]
				p.Evs.Spe = m["Spe"]
			} else if strings.Contains(prop, "I") {
				p.Ivs.Hp = m["HP"]
				p.Ivs.Atk = m["Atk"]
				p.Ivs.Def = m["Def"]
				p.Ivs.Spa = m["SpA"]
				p.Ivs.Spd = m["SpD"]
				p.Ivs.Spe = m["Spe"]
			} else {
				return fmt.Errorf("error in parsing evs/ivs line: invalid prop %s", prop)
			}
		case Move.MatchString(line):
			p.Moves = append(p.Moves, Move.FindStringSubmatch(line)[1])
		}
	}
	return nil
}

func fromEIvsLineToMap(line string) (m map[string]int, prop string, err error) {
	segments := EIvs.FindStringSubmatch(line)
	if len(segments) != 3 {
		return nil, "", fmt.Errorf("invalid evs/ivs line: %s", line)
	}
	prop = segments[1]
	if strings.Contains(prop, "I") {
		m = map[string]int{"HP": 31, "Atk": 31, "Def": 31, "SpA": 31, "SpD": 31, "Spe": 31}
	} else {
		m = map[string]int{"HP": 0, "Atk": 0, "Def": 0, "SpA": 0, "SpD": 0, "Spe": 0}
	}
	parts := strings.Split(segments[2], " / ")
	for _, part := range parts {
		stat := strings.Split(part, " ")
		num, err := strconv.Atoi(stat[0])
		if err != nil {
			return nil, "", err
		}
		m[stat[1]] = num
	}
	return m, segments[1], nil
}

func (p Pokemon) ToShowdown() (string, error) {
	var showdown strings.Builder
	showdown.Grow(300) // estimated string length
	if err := p.Validate(); err != nil {
		return "", err
	}
	// name/nickname
	if len(p.Nickname) > 0 {
		showdown.WriteString(p.Nickname)
		showdown.WriteString(" (")
		showdown.WriteString(p.Name)
		showdown.WriteByte(')')
	} else {
		showdown.WriteString(p.Name)
	}
	// gender
	if len(p.Gender) > 0 {
		showdown.WriteString(" (")
		showdown.WriteString(strings.ToUpper(p.Gender))
		showdown.WriteByte(')')
	}
	// item
	if len(p.Item) > 0 {
		showdown.WriteString(" @ ")
		showdown.WriteString(p.Item)
	}
	showdown.WriteByte('\n')
	// level
	if p.Level > 0 {
		showdown.WriteString("Level: ")
		showdown.WriteString(strconv.Itoa(p.Level))
		showdown.WriteByte('\n')
	}
	// ability
	showdown.WriteString("Ability: ")
	showdown.WriteString(p.Ability)
	showdown.WriteByte('\n')
	// shiny
	if p.Shiny {
		showdown.WriteString("Shiny: Yes\n")
	}
	// happiness
	if p.Level > 0 {
		showdown.WriteString("Happiness: ")
		showdown.WriteString(strconv.Itoa(p.Happiness))
		showdown.WriteByte('\n')
	}
	// evs
	showdown.WriteString("EVs: ")
	evs := []string{
		strconv.Itoa(p.Evs.Hp) + " HP",
		strconv.Itoa(p.Evs.Atk) + " Atk",
		strconv.Itoa(p.Evs.Def) + " Def",
		strconv.Itoa(p.Evs.Spa) + " SpA",
		strconv.Itoa(p.Evs.Spd) + " SpD",
		strconv.Itoa(p.Evs.Spe) + " Spe",
	}
	showdown.WriteString(strings.Join(evs, " / "))
	showdown.WriteByte('\n')
	// nature
	showdown.WriteString(p.Nature)
	showdown.WriteString(" Nature\n")
	// ivs
	showdown.WriteString("IVs: ")
	ivs := []string{
		strconv.Itoa(p.Ivs.Hp) + " HP",
		strconv.Itoa(p.Ivs.Atk) + " Atk",
		strconv.Itoa(p.Ivs.Def) + " Def",
		strconv.Itoa(p.Ivs.Spa) + " SpA",
		strconv.Itoa(p.Ivs.Spd) + " SpD",
		strconv.Itoa(p.Ivs.Spe) + " Spe",
	}
	showdown.WriteString(strings.Join(ivs, " / "))
	showdown.WriteByte('\n')
	// moves
	for _, move := range p.Moves {
		if len(move) > 0 {
			showdown.WriteString("- ")
			showdown.WriteString(move)
			showdown.WriteByte('\n')
		}
	}
	return showdown.String(), nil
}
