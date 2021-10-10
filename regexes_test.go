package koffing

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexes(t *testing.T) {
	t.Parallel()
	assert.True(t, Team.MatchString("=== [gen7] Folder 1/Example Team ==="))
	assert.False(t, Team.MatchString("======"))

	assert.True(t, Gender.MatchString("(F)") && Gender.MatchString("(M)"))
	assert.False(t, Gender.MatchString("F") || Gender.MatchString("M"))

	assert.True(t, Item.MatchString("@ Focus Sash"))
	assert.Equal(t, "Focus Sash", Item.FindStringSubmatch("Weezing-Gmax @ Focus Sash")[1])

	assert.True(t, Name.MatchString("Weezing @ Black Sludge"))
	assert.Equal(t, "Weezing-Gmax ", Name.FindString("Weezing-Gmax @ Black Sludge"))

	assert.True(t, NicknameWithName.MatchString("Smogon (Koffing) (F) @ Eviolite"))
	assert.False(t, NicknameWithName.MatchString("Weezing @ Black Sludge"))
	assert.Equal(t, "Tapu Koko", NicknameWithName.FindStringSubmatch("Tapu Koko (Weezing-Gmax) (F) @ Eviolite")[1])
	assert.Equal(t, "Tapu Koko", NicknameWithName.FindStringSubmatch("Smogon (Tapu Koko) (F) @ Eviolite")[2])

	assert.True(t, Ability.MatchString("Ability: Levitate"))
	assert.False(t, Ability.MatchString("Ability Levitate"))
	assert.Equal(t, "Levitate", Ability.FindStringSubmatch("Ability: Levitate")[1])

	assert.True(t, Level.MatchString("Level: 5"))
	assert.False(t, Level.MatchString("Level: 1000"))
	assert.Equal(t, "5", Level.FindStringSubmatch("Level: 5")[1])

	assert.True(t, Shiny.MatchString("Shiny: Yes"))
	assert.False(t, Shiny.MatchString("Shiny: True"))
	assert.Equal(t, "Yes", Shiny.FindStringSubmatch("Shiny: Yes")[1])

	assert.True(t, Happiness.MatchString("Happiness: 255"))
	assert.False(t, Happiness.MatchString("Happiness: -1"))
	assert.Equal(t, "255", Happiness.FindStringSubmatch("Happiness: 255")[1])

	assert.True(t, Nature.MatchString("Bold Nature"))
	assert.False(t, Nature.MatchString("BoldNature"))
	assert.Equal(t, "Bold", Nature.FindStringSubmatch("Bold Nature")[1])

	assert.True(t, EIvs.MatchString("EVs: 36 HP / 236 Def / 236 SpD"))
	assert.True(t, EIvs.MatchString("IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))
	assert.False(t, EIvs.MatchString("IVs31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))
	assert.Equal(t, "EVs", EIvs.FindStringSubmatch("EVs: 36 HP / 236 Def / 236 SpD")[1])

	assert.True(t, Move.MatchString("- Protect"))
	assert.False(t, Move.MatchString("Protect"))
	assert.Equal(t, "Protect", Move.FindStringSubmatch("- Protect")[1])
}

func TestSplitByEmptyNewline(t *testing.T) {
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
	res := SplitByEmptyNewline(s)
	assert.Len(t, res, 7)
}