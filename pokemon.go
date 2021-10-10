package koffing

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
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
