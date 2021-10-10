package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegexes(t *testing.T) {
	assert.True(t, Team.MatchString("=== [gen7] Folder 1/Example Team ==="))
	assert.False(t, Team.MatchString("======"))

	assert.True(t, Gender.MatchString("(F)") && Gender.MatchString("(M)"))
	assert.False(t, Gender.MatchString("F") || Gender.MatchString("M"))

	assert.True(t, Item.MatchString("@ Focus Sash"))

	assert.True(t, Name.MatchString("Weezing @ Black Sludge"))

	assert.True(t, NicknameWithName.MatchString("Smogon (Koffing) (F) @ Eviolite"))
	assert.False(t, NicknameWithName.MatchString("Weezing @ Black Sludge"))

	assert.True(t, Ability.MatchString("Ability: Levitate"))
	assert.False(t, Ability.MatchString("Ability Levitate"))

	assert.True(t, Level.MatchString("Level: 5"))
	assert.False(t, Level.MatchString("Level: 1000"))

	assert.True(t, Shiny.MatchString("Shiny: Yes"))
	assert.False(t, Shiny.MatchString("Shiny: True"))

	assert.True(t, Happiness.MatchString("Happiness: 255"))
	assert.False(t, Happiness.MatchString("Happiness: -1"))

	assert.True(t, Nature.MatchString("Bold Nature"))
	assert.False(t, Nature.MatchString("BoldNature"))

	assert.True(t, EIvs.MatchString("EVs: 36 HP / 236 Def / 236 SpD"))
	assert.True(t, EIvs.MatchString("IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))
	assert.False(t, EIvs.MatchString("IVs31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))

	assert.True(t, EIvsValue.MatchString("31 HP"))
	assert.False(t, EIvsValue.MatchString("IVs: 31 HP / 30 Atk / 31 SpA / 30 SpD / 31 Spe"))

	assert.True(t, Move.MatchString("- Protect"))
	assert.False(t, Move.MatchString("Protect"))
}
