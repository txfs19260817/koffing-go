package koffing

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPokemon_FromJson(t *testing.T) {
	t.Parallel()
	p := &Pokemon{}
	assert.Error(t, p.FromJson(`"name": "Koffing"`))
	paste := `{
          "name": "Koffing",
          "nickname": "Smogon",
          "gender": "F",
          "item": "Eviolite",
          "ability": "Neutralizing Gas",
          "level": 5,
          "shiny": true,
          "happiness": 255,
          "nature": "Bold",
          "evs": {
            "hp": 36,
            "def": 236,
            "spd": 236
          },
          "ivs": {
            "hp": 31,
            "atk": 30,
            "spa": 31,
            "spd": 30,
            "spe": 31
          },
          "moves": [
            "Will-O-Wisp",
            "Pain Split",
            "Sludge Bomb",
            "Fire Blast"
          ]
	}`
	err := p.FromJson(paste)
	assert.NoError(t, err)
	assert.Equal(t, "Koffing", p.Name)
	assert.Equal(t, "Smogon", p.Nickname)
	assert.Equal(t, "F", p.Gender)
	assert.Equal(t, "Eviolite", p.Item)
	assert.Equal(t, "Neutralizing Gas", p.Ability)
	assert.Equal(t, 5, p.Level)
	assert.Equal(t, true, p.Shiny)
	assert.Equal(t, 255, p.Happiness)
	assert.Equal(t, "Bold", p.Nature)
	assert.Equal(t, []int{36, 0, 236, 0, 236, 0}, []int{p.Evs.Hp, p.Evs.Atk, p.Evs.Def, p.Evs.Spa, p.Evs.Spd, p.Evs.Spe})
	assert.Equal(t, []int{31, 30, 0, 31, 30, 31}, []int{p.Ivs.Hp, p.Ivs.Atk, p.Ivs.Def, p.Ivs.Spa, p.Ivs.Spd, p.Ivs.Spe})
	assert.Equal(t, []string{"Will-O-Wisp", "Pain Split", "Sludge Bomb", "Fire Blast"}, p.Moves)
}

func TestPokemon_ToJson(t *testing.T) {
	t.Parallel()
	p := Pokemon{
		Name:      "Koffing",
		Nickname:  "Smogon",
		Gender:    "F",
		Item:      "Eviolite",
		Ability:   "Neutralizing Gas",
		Level:     5,
		Shiny:     true,
		Happiness: 255,
		Nature:    "Bold",
		Evs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 36, Def: 236, Spd: 236},
		Ivs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 31, Atk: 30, Spa: 31, Spd: 30, Spe: 31},
		Moves: []string{"Will-O-Wisp", "Pain Split", "Sludge Bomb", "Fire Blast"},
	}
	j, err := p.ToJson()
	assert.NoError(t, err)
	assert.Equal(t, `{"name":"Koffing","nickname":"Smogon","gender":"F","item":"Eviolite","ability":"Neutralizing Gas","level":5,"shiny":true,"happiness":255,"nature":"Bold","evs":{"hp":36,"atk":0,"def":236,"spa":0,"spd":236,"spe":0},"ivs":{"hp":31,"atk":30,"def":0,"spa":31,"spd":30,"spe":31},"moves":["Will-O-Wisp","Pain Split","Sludge Bomb","Fire Blast"]}`, j)
}

func TestPokemon_Validate(t *testing.T) {
	t.Parallel()
	// Name
	p := Pokemon{}
	err := p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "name is required")
	// Ability
	p.Name = "Koffing"
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ability is required")
	// Nature
	p.Ability = "Neutralizing Gas"
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nature is required")
	// Happiness
	p.Nature = "Bold"
	p.Happiness = -1
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "happiness")
	// Evs
	p.Happiness = 200
	p.Evs.Def = 255
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ev")
	// Ivs
	p.Evs.Def = 252
	p.Ivs.Spa = 32
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "iv")
	// Moves length
	p.Ivs.Spa = 1
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "moves")
	// Gender
	p.Moves = []string{"Will-O-Wisp", "Pain Split", "Sludge Bomb", "Fire Blast"}
	p.Gender = "m"
	err = p.Validate()
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "gender")

	// OK
	p = Pokemon{
		Name:      "Koffing",
		Nickname:  "Smogon",
		Gender:    "F",
		Item:      "Eviolite",
		Ability:   "Neutralizing Gas",
		Level:     5,
		Shiny:     true,
		Happiness: 255,
		Nature:    "Bold",
		Evs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 36, Def: 236, Spd: 236},
		Ivs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 31, Atk: 30, Spa: 31, Spd: 30, Spe: 31},
		Moves: []string{"Will-O-Wisp", "Pain Split", "Sludge Bomb", "Fire Blast"},
	}
	err = p.Validate()
	assert.NoError(t, err)
}

func TestPokemon_FromShowdown(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name: "nickname",
			s: `Smogon (Koffing) (F) @ Eviolite
				Level: 100
				Ability: Neutralizing Gas
				Shiny: Yes
				Happiness: 255
				EVs: 36 HP / 236 Def / 236 SpD
				Bold Nature
				IVs: 31 HP / 30 SpD / 0 Spe
				- Will-O-Wisp
				- Pain Split
				- Sludge Bomb
				- Fire Blast`,
			wantErr: false,
		},
		{
			name: "name",
			s: `Landorus @ Life Orb  
				Ability: Sheer Force  
				Level: 50  
				EVs: 244 HP / 92 Def / 100 SpA / 4 SpD / 68 Spe  
				Modest Nature  
				- Earth Power  
				- Sludge Bomb  
				- Substitute  
				- Protect
				`,
			wantErr: false,
		},
		{
			name: "no name",
			s: ` @ Sitrus Berry  
				Ability: Misty Surge  
				EVs: 252 HP / 68 Def / 4 SpA / 116 SpD / 68 Spe  
				Calm Nature  
				IVs: 0 Atk  
				- Moonblast  
				- Icy Wind  
				- Haze  
				- Nature's Madness
				`,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Pokemon{}
			err := p.FromShowdown(tt.s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, p.Validate())
			}
		})
	}
}

func TestPokemon_ToShowdown(t *testing.T) {
	t.Parallel()
	p := Pokemon{
		Name:      "Koffing",
		Nickname:  "Smogon",
		Gender:    "F",
		Item:      "Eviolite",
		Ability:   "Neutralizing Gas",
		Level:     100,
		Shiny:     true,
		Happiness: 255,
		Nature:    "Bold",
		Evs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 36, Def: 236, Spd: 236},
		Ivs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 31, Atk: 30, Spa: 31, Spd: 30, Spe: 31},
		Moves: []string{"Will-O-Wisp", "Pain Split", "Sludge Bomb", "Fire Blast"},
	}

	expected := `Smogon (Koffing) (F) @ Eviolite
				Level: 100
				Ability: Neutralizing Gas
				Shiny: Yes
				Happiness: 255
				EVs: 36 HP / 236 Def / 236 SpD
				Bold Nature
				IVs: 30 Atk / 0 Def / 30 SpD
				- Will-O-Wisp
				- Pain Split
				- Sludge Bomb
				- Fire Blast
				`
	showdown, err := p.ToShowdown()
	assert.NoError(t, err)
	expectedSlice, actualSlice := strings.Split(expected, "\n"), strings.Split(showdown, "\n")
	assert.Equal(t, len(expectedSlice), len(actualSlice))
	for i, e := range expectedSlice {
		assert.Equal(t, strings.TrimSpace(e), strings.TrimSpace(actualSlice[i]))
	}
}
