package koffing

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func ExampleTeam_FromJson() {
	team := &Team{}
	j := `{"name":"Example Team","format":"gen7","folder":"Folder 1","pokemon":[{"name":"Koffing","nickname":"Smogon","gender":"F","item":"Eviolite","ability":"Levitate","level":5,"shiny":true,"happiness":255,"nature":"Bold","evs":{"hp":36,"def":236,"spd":236},"ivs":{"hp":31,"atk":30,"spa":31,"spd":30,"spe":31},"moves":["Will-O-Wisp","Pain Split","Sludge Bomb","Fire Blast"]},{"name":"Weezing","item":"Black Sludge","ability":"Levitate","nature":"Bold","evs":{"hp":252,"def":160,"spe":96},"moves":["Sludge Bomb","Will-O-Wisp","Toxic Spikes","Taunt"]}]}`
	_ = team.FromJson(j)
	fmt.Printf("%+v", team)
	// Output: &{Name:Example Team Format:gen7 Folder:Folder 1 Pokemon:[{Name:Koffing Nickname:Smogon Gender:F Item:Eviolite Ability:Levitate Level:5 Shiny:true Happiness:255 Nature:Bold Evs:{Hp:36 Atk:0 Def:236 Spa:0 Spd:236 Spe:0} Ivs:{Hp:31 Atk:30 Def:0 Spa:31 Spd:30 Spe:31} Moves:[Will-O-Wisp Pain Split Sludge Bomb Fire Blast]} {Name:Weezing Nickname: Gender: Item:Black Sludge Ability:Levitate Level:0 Shiny:false Happiness:0 Nature:Bold Evs:{Hp:252 Atk:0 Def:160 Spa:0 Spd:0 Spe:96} Ivs:{Hp:0 Atk:0 Def:0 Spa:0 Spd:0 Spe:0} Moves:[Sludge Bomb Will-O-Wisp Toxic Spikes Taunt]}]}
}

func TestTeam_FromJson(t *testing.T) {
	t.Parallel()
	team := &Team{}
	assert.Error(t, team.FromJson(`"name": "Koffing"`))
	j := `{"name":"Example Team","format":"gen7","folder":"Folder 1","pokemon":[{"name":"Koffing","nickname":"Smogon","gender":"F","item":"Eviolite","ability":"Levitate","level":5,"shiny":true,"happiness":255,"nature":"Bold","evs":{"hp":36,"def":236,"spd":236},"ivs":{"hp":31,"atk":30,"spa":31,"spd":30,"spe":31},"moves":["Will-O-Wisp","Pain Split","Sludge Bomb","Fire Blast"]},{"name":"Weezing","item":"Black Sludge","ability":"Levitate","nature":"Bold","evs":{"hp":252,"def":160,"spe":96},"moves":["Sludge Bomb","Will-O-Wisp","Toxic Spikes","Taunt"]}]}`
	err := team.FromJson(j)
	assert.NoError(t, err)
	assert.Equal(t, "Example Team", team.Name)
	assert.Equal(t, "gen7", team.Format)
	assert.Equal(t, "Folder 1", team.Folder)
	assert.Len(t, team.Pokemon, 2)
}

func ExampleTeam_FromShowdown() {
	s := `=== [gen7] Folder 1/Example Team ===

		Smogon (Koffing) (F) @ Eviolite
		Level: 5
		Ability: Levitate
		Shiny: Yes
		Happiness: 255
		EVs: 36 HP / 236 Def / 236 SpD
		IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe
		Bold Nature
		- Will-O-Wisp
		- Pain Split
		- Sludge Bomb
		- Fire Blast
		
		Venusaur-Gmax @ Coba Berry  
		Ability: Chlorophyll  
		Level: 50  
		EVs: 156 HP / 4 Def / 252 SpA / 4 SpD / 92 Spe  
		Modest Nature  
		IVs: 0 Atk  
		- Frenzy Plant  
		- Sludge Bomb  
		- Earth Power  
		- Sleep Powder  
		`
	team := new(Team)
	_ = team.FromShowdown(s)
	fmt.Printf("%+v", team)
	// Output: &{Name:Example Team Format:gen7 Folder:Folder 1 Pokemon:[{Name:Koffing Nickname:Smogon Gender:F Item:Eviolite Ability:Levitate Level:5 Shiny:true Happiness:255 Nature:Bold Evs:{Hp:36 Atk:0 Def:236 Spa:0 Spd:236 Spe:0} Ivs:{Hp:31 Atk:30 Def:31 Spa:31 Spd:30 Spe:31} Moves:[Will-O-Wisp Pain Split Sludge Bomb Fire Blast]} {Name:Venusaur-Gmax Nickname: Gender: Item:Coba Berry Ability:Chlorophyll Level:50 Shiny:false Happiness:255 Nature:Modest Evs:{Hp:156 Atk:0 Def:4 Spa:252 Spd:4 Spe:92} Ivs:{Hp:31 Atk:0 Def:31 Spa:31 Spd:31 Spe:31} Moves:[Frenzy Plant Sludge Bomb Earth Power Sleep Powder]}]}
}

func TestTeam_FromShowdown(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name    string
		s       string
		wantErr bool
	}{
		{
			name: "name with folder",
			s: `=== [gen7] Folder 1/Example Team ===

				Smogon (Koffing) (F) @ Eviolite
				Level: 5
				Ability: Levitate
				Shiny: Yes
				Happiness: 255
				EVs: 36 HP / 236 Def / 236 SpD
				IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe
				Bold Nature
				- Will-O-Wisp
				- Pain Split
				- Sludge Bomb
				- Fire Blast
				
				Venusaur-Gmax @ Coba Berry  
				Ability: Chlorophyll  
				Level: 50  
				EVs: 156 HP / 4 Def / 252 SpA / 4 SpD / 92 Spe  
				Modest Nature  
				IVs: 0 Atk  
				- Frenzy Plant  
				- Sludge Bomb  
				- Earth Power  
				- Sleep Powder  
				`,
			wantErr: false,
		},
		{
			name: "name without folder",
			s: `=== [gen8vgc2021] Untitled 10 ===

				Charizard-Gmax @ Wacan Berry  
				Ability: Solar Power  
				Level: 50  
				EVs: 4 HP / 252 SpA / 252 Spe  
				Timid Nature  
				IVs: 0 Atk  
				- Blast Burn  
				- Hurricane  
				- Ancient Power  
				- Protect  
				`,
			wantErr: false,
		},
		{
			name: "no team tag",
			s: `Charizard-Gmax @ Wacan Berry  
				Ability: Solar Power  
				Level: 50  
				EVs: 4 HP / 252 SpA / 252 Spe  
				Timid Nature  
				IVs: 0 Atk  
				- Blast Burn  
				- Hurricane  
				- Ancient Power  
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
			team := new(Team)
			err := team.FromShowdown(tt.s)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, team.Validate())
			}
		})
	}
}

func ExampleTeam_ToJson() {
	team := Team{
		Name:    "Test",
		Format:  "gen8",
		Folder:  "Folder 0",
		Pokemon: []Pokemon{{Name: "Koffing", Ability: "Neutralizing Gas", Nature: "Bold", Moves: []string{"Haze"}}},
	}
	j, _ := team.ToJson()
	fmt.Print(j)
	// Output: {"name":"Test","format":"gen8","folder":"Folder 0","pokemon":[{"name":"Koffing","nickname":"","gender":"","item":"","ability":"Neutralizing Gas","level":0,"shiny":false,"happiness":0,"nature":"Bold","evs":{"hp":0,"atk":0,"def":0,"spa":0,"spd":0,"spe":0},"ivs":{"hp":0,"atk":0,"def":0,"spa":0,"spd":0,"spe":0},"moves":["Haze"]}]}
}

func TestTeam_ToJson(t *testing.T) {
	t.Parallel()
	team := Team{
		Name:    "Test",
		Format:  "gen8",
		Folder:  "Folder 0",
		Pokemon: []Pokemon{{Name: "Koffing", Ability: "Neutralizing Gas", Nature: "Bold", Moves: []string{"Haze"}}},
	}
	j, err := team.ToJson()
	assert.NoError(t, err)
	assert.Equal(t, j, `{"name":"Test","format":"gen8","folder":"Folder 0","pokemon":[{"name":"Koffing","nickname":"","gender":"","item":"","ability":"Neutralizing Gas","level":0,"shiny":false,"happiness":0,"nature":"Bold","evs":{"hp":0,"atk":0,"def":0,"spa":0,"spd":0,"spe":0},"ivs":{"hp":0,"atk":0,"def":0,"spa":0,"spd":0,"spe":0},"moves":["Haze"]}]}`)
}

func ExampleTeam_ToShowdown() {
	team := Team{
		Name:   "Test",
		Format: "gen8",
		Folder: "Folder 0",
		Pokemon: []Pokemon{{Name: "Koffing", Ability: "Neutralizing Gas", Nature: "Bold", Moves: []string{"Haze"}, Ivs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 31, Def: 31, Spd: 31}}},
	}
	s, _ := team.ToShowdown()
	fmt.Print(s)
	// Output: === [gen8] Folder 0/Test ===
	//
	//Koffing
	//Ability: Neutralizing Gas
	//Bold Nature
	//IVs: 0 Atk / 0 SpA / 0 Spe
	//- Haze
}

func TestTeam_ToShowdown(t *testing.T) {
	t.Parallel()
	team := Team{
		Name:   "Test",
		Format: "gen8",
		Folder: "Folder 0",
		Pokemon: []Pokemon{{Name: "Koffing", Ability: "Neutralizing Gas", Nature: "Bold", Moves: []string{"Haze"}, Ivs: struct {
			Hp  int `json:"hp"`
			Atk int `json:"atk"`
			Def int `json:"def"`
			Spa int `json:"spa"`
			Spd int `json:"spd"`
			Spe int `json:"spe"`
		}{Hp: 31, Def: 31, Spd: 31}}},
	}
	expected := `=== [gen8] Folder 0/Test ===
        
        Koffing
        Ability: Neutralizing Gas
        Bold Nature
        IVs: 0 Atk / 0 SpA / 0 Spe
        - Haze
		`
	s, err := team.ToShowdown()
	assert.NoError(t, err)
	expectedSlice, actualSlice := strings.Split(expected, "\n"), strings.Split(s, "\n")
	assert.Equal(t, len(expectedSlice), len(actualSlice))
	for i, e := range expectedSlice {
		assert.Equal(t, strings.TrimSpace(e), strings.TrimSpace(actualSlice[i]))
	}
}

func ExampleTeam_Validate() {
	team := Team{
		Name:    "Test",
		Format:  "gen8",
		Folder:  "Folder 0",
		Pokemon: []Pokemon{}, // empty Pokemon is not allowed
	}
	err := team.Validate()
	fmt.Print(err.Error())
	// Output: empty team members
}

func TestTeam_Validate(t *testing.T) {
	t.Parallel()
	team := Team{
		Name:    "Test",
		Format:  "gen8",
		Folder:  "Folder 0",
		Pokemon: []Pokemon{{Name: "Koffing", Ability: "Neutralizing Gas", Nature: "Bold", Moves: []string{"Haze"}}},
	}
	assert.NoError(t, team.Validate())
	team.Pokemon = nil
	assert.Error(t, team.Validate())
}
