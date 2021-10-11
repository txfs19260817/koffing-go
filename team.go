package koffing

import (
	"fmt"
	"strings"
)

type Team struct {
	Name    string    `json:"name,omitempty"`
	Format  string    `json:"format,omitempty"`
	Folder  string    `json:"folder,omitempty"`
	Pokemon []Pokemon `json:"pokemon,omitempty"`
}

func (t *Team) FromJson(j string) error {
	return json.Unmarshal([]byte(j), t)
}

func (t Team) ToJson() (string, error) {
	return json.MarshalToString(t)
}

func (t *Team) FromShowdown(s string) error {
	parts := SplitByEmptyNewline(s)
	if TeamTag.MatchString(parts[0]) {
		teamTags := TeamTag.FindStringSubmatch(parts[0])
		t.Format = teamTags[1]
		if name := strings.Split(teamTags[2], "/"); len(name) == 2 {
			t.Folder, t.Name = name[0], name[1]
		} else {
			t.Name = name[0]
		}
		parts = parts[1:]
	}
	t.Pokemon = make([]Pokemon, 0, 6)
	for i, part := range parts {
		var p Pokemon
		if err := p.FromShowdown(part); err != nil {
			return fmt.Errorf("failed to import the Showdown text to a Pokemon: index: %d, error: %w", i, err)
		}
		t.Pokemon = append(t.Pokemon, p)
	}
	return nil
}

func (t Team) ToShowdown() (string, error) {
	var showdown strings.Builder
	if len(t.Name) > 0 {
		if len(t.Folder) > 0 {
			showdown.WriteString(fmt.Sprintf("=== [%s] %s/%s ===\n\n", t.Format, t.Folder, t.Name))
		} else {
			showdown.WriteString(fmt.Sprintf("=== [%s] %s ===\n\n", t.Format, t.Name))
		}
	}

	for i, pokemon := range t.Pokemon {
		p, err := pokemon.ToShowdown()
		if err != nil {
			return "", fmt.Errorf("failed to export a Pokemon to Showdown: index: %d, error: %w", i, err)
		}
		showdown.Grow(len(p))
		showdown.WriteString(p)
	}
	return showdown.String(), nil
}

func (t Team) Validate() error {
	if len(t.Pokemon) == 0 {
		return fmt.Errorf("empty team members")
	}
	for i, pokemon := range t.Pokemon {
		if err := pokemon.Validate(); err != nil {
			return fmt.Errorf("found an invalid Pokemon: index: %d, error: %w", i, err)
		}
	}
	return nil
}
