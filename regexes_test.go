package koffing

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_regexes(t *testing.T) {
	t.Parallel()
	assert.True(t, teamTagRegex.MatchString("=== [gen7] Folder 1/Example Team ==="))
	assert.False(t, teamTagRegex.MatchString("======"))
	assert.Equal(t, teamTagRegex.FindStringSubmatch("=== [gen8vgc2021] Untitled 10 ===")[1], "gen8vgc2021")
	assert.Equal(t, teamTagRegex.FindStringSubmatch("=== [gen7] Folder 1/Example Team ===")[2], "Folder 1/Example Team")

	assert.True(t, genderRegex.MatchString("(F)") && genderRegex.MatchString("(M)"))
	assert.False(t, genderRegex.MatchString("F") || genderRegex.MatchString("M"))

	assert.True(t, itemRegex.MatchString("@ Focus Sash"))
	assert.Equal(t, "Focus Sash", itemRegex.FindStringSubmatch("Weezing-Gmax @ Focus Sash")[1])

	assert.True(t, nameRegex.MatchString("Weezing @ Black Sludge"))
	assert.Equal(t, "Weezing-Gmax ", nameRegex.FindString("Weezing-Gmax @ Black Sludge"))

	assert.True(t, nicknameWithNameRegex.MatchString("Smogon (Koffing) (F) @ Eviolite"))
	assert.False(t, nicknameWithNameRegex.MatchString("Weezing @ Black Sludge"))
	assert.Equal(t, "Tapu Koko", nicknameWithNameRegex.FindStringSubmatch("Tapu Koko (Weezing-Gmax) (F) @ Eviolite")[1])
	assert.Equal(t, "Tapu Koko", nicknameWithNameRegex.FindStringSubmatch("Smogon (Tapu Koko) (F) @ Eviolite")[2])

	assert.True(t, abilityRegex.MatchString("Ability: Levitate"))
	assert.False(t, abilityRegex.MatchString("Ability Levitate"))
	assert.Equal(t, "Levitate", abilityRegex.FindStringSubmatch("Ability: Levitate")[1])

	assert.True(t, levelRegex.MatchString("Level: 5"))
	assert.False(t, levelRegex.MatchString("Level: 1000"))
	assert.Equal(t, "5", levelRegex.FindStringSubmatch("Level: 5")[1])

	assert.True(t, shinyRegex.MatchString("Shiny: Yes"))
	assert.False(t, shinyRegex.MatchString("Shiny: True"))
	assert.Equal(t, "Yes", shinyRegex.FindStringSubmatch("Shiny: Yes")[1])

	assert.True(t, happinessRegex.MatchString("Happiness: 255"))
	assert.False(t, happinessRegex.MatchString("Happiness: -1"))
	assert.Equal(t, "255", happinessRegex.FindStringSubmatch("Happiness: 255")[1])

	assert.True(t, natureRegex.MatchString("Bold Nature"))
	assert.False(t, natureRegex.MatchString("BoldNature"))
	assert.Equal(t, "Bold", natureRegex.FindStringSubmatch("Bold Nature")[1])

	assert.True(t, eivsRegex.MatchString("EVs: 36 HP / 236 Def / 236 SpD"))
	assert.True(t, eivsRegex.MatchString("IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))
	assert.False(t, eivsRegex.MatchString("IVs31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))
	assert.Equal(t, "EVs", eivsRegex.FindStringSubmatch("EVs: 36 HP / 236 Def / 236 SpD")[1])

	assert.True(t, moveRegex.MatchString("- Protect"))
	assert.False(t, moveRegex.MatchString("Protect"))
	assert.Equal(t, "Protect", moveRegex.FindStringSubmatch("- Protect")[1])
}

func Test_splitByEmptyNewline(t *testing.T) {
	s := `=== [gen8vgc2021] Untitled 10 ===

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

Tapu Fini @ Sitrus Berry  
Ability: Misty Surge  
EVs: 252 HP / 68 Def / 4 SpA / 116 SpD / 68 Spe  
Calm Nature  
IVs: 0 Atk  
- Moonblast  
- Icy Wind  
- Haze  
- Nature's Madness

Thundurus @ Life Orb  
Ability: Defiant  
Level: 50  
EVs: 4 HP / 252 Atk / 252 Spe  
Jolly Nature  
- Fly  
- Wild Charge  
- Superpower  
- Protect  

Urshifu @ Focus Sash  
Ability: Unseen Fist  
Level: 50  
EVs: 252 Atk / 4 SpD / 252 Spe  
Jolly Nature  
- Close Combat  
- Detect  
- Wicked Blow  
- Sucker Punch  

Zacian @ Rusted Sword  
Ability: Intrepid Sword  
Level: 50  
EVs: 252 HP / 108 Atk / 4 Def / 68 SpD / 76 Spe  
Adamant Nature  
- Iron Head  
- Substitute  
- Sacred Sword  
- Protect  

`
	res := splitByEmptyNewline(s)
	assert.Len(t, res, 7)
	assert.True(t, teamTagRegex.MatchString(res[0]))
}

func Test_trimLines(t *testing.T) {
	type args struct {
		lines []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Zacian",
			args: args{strings.Split("\nZacian @ Rusted Sword  \nAbility: Intrepid Sword  \nLevel: 50  \n", "\n")},
			want: []string{"Zacian @ Rusted Sword", "Ability: Intrepid Sword", "Level: 50"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := trimLines(tt.args.lines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("trimLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
