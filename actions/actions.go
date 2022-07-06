package actions

import (
	"math/rand"
	"time"
)

//For creating "real" random numbers
var randSource = rand.NewSource(time.Now().UnixNano())
var randGenerator = rand.New(randSource)

//Health managment
var curMonsterHealth = MONSTER_HEALTH
var curPlayerHealth = PLAYER_HEALTH

func AttackMonster(isSpecialAttack bool) int {
	minVal := PL_ATK_MIN_DMG
	maxVal := PL_ATK_MAX_DMG

	if isSpecialAttack {
		minVal = PL_SPACIEL_ATK_MIN_DMG
		maxVal = PL_SPACIEL_ATK_MAX_DMG
	}

	dmgVal := randBetween(minVal, maxVal)
	curMonsterHealth -= dmgVal

	return dmgVal
}

func Heal() int {
	healVal := randBetween(PL_HEAL_MIN_VAL, PL_HEAL_MAX_VAL)

	healDiff := PLAYER_HEALTH - curPlayerHealth
	if healDiff >= healVal {
		curPlayerHealth += healVal
		return healVal
	} else {
		curPlayerHealth = PLAYER_HEALTH
		return healDiff
	}

}

func AttackPlayer() int {

	minVal := MON_ATK_MIN_DMG
	maxVal := MON_ATK_MAX_DMG

	dmgVal := randBetween(minVal, maxVal)
	curPlayerHealth -= dmgVal

	return dmgVal
}

func GetHealthAmounts() (int, int) {
	return curPlayerHealth, curMonsterHealth
}

func randBetween(min int, max int) int {
	return randGenerator.Intn(max-min) + min
}
