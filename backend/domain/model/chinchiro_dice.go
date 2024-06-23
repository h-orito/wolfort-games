package model

import (
	"math/rand"
	"sort"
	"time"
	"wolfort-games/util/array"
)

type ChinchiroDiceRoll struct {
	Dice1       int
	Dice2       int
	Dice3       int
	Combination ChinchiroCombination
}

func RollChinchiroDice() ChinchiroDiceRoll {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dice1 := r.Intn(6) + 1
	dice2 := r.Intn(6) + 1
	dice3 := r.Intn(6) + 1
	return NewChinchiroDiceRoll(dice1, dice2, dice3)
}

func NewChinchiroDiceRoll(dice1, dice2, dice3 int) ChinchiroDiceRoll {
	// 小さい順に並び替える
	dice := []int{dice1, dice2, dice3}
	sort.Slice(dice, func(i, j int) bool {
		return dice[i] < dice[j]
	})

	return ChinchiroDiceRoll{
		Dice1:       dice1,
		Dice2:       dice2,
		Dice3:       dice3,
		Combination: combination(dice1, dice2, dice3),
	}
}

func combination(dice1, dice2, dice3 int) ChinchiroCombination {
	// 全部同じ目
	if dice1 == dice2 && dice2 == dice3 {
		switch dice1 {
		case 1:
			return ChinchiroCombinationPinzoro
		case 2:
			return ChinchiroCombinationNizoro
		case 3:
			return ChinchiroCombinationSanzoro
		case 4:
			return ChinchiroCombinationYonzoro
		case 5:
			return ChinchiroCombinationGozoro
		case 6:
			return ChinchiroCombinationRokuzoro
		}
	} else if dice1 == dice2 || dice2 == dice3 || dice3 == dice1 {
		// 2つが同じ目
		judgeDice := dice1
		if dice1 == dice2 {
			judgeDice = dice3
		} else if dice2 == dice3 {
			judgeDice = dice1
		} else if dice3 == dice1 {
			judgeDice = dice2
		}
		switch judgeDice {
		case 1:
			return ChinchiroCombinationIchinome
		case 2:
			return ChinchiroCombinationNinome
		case 3:
			return ChinchiroCombinationSannome
		case 4:
			return ChinchiroCombinationYonnome
		case 5:
			return ChinchiroCombinationGonome
		case 6:
			return ChinchiroCombinationRokunome
		}
	}

	// ひとつも同じ目がない
	if dice1 == 1 && dice2 == 2 && dice3 == 3 {
		// ヒフミ
		return ChinchiroCombinationHifumi
	} else if dice1 == 4 && dice2 == 5 && dice3 == 6 {
		// シゴロ
		return ChinchiroCombinationShigoro
	}

	return ChinchiroCombinationMenashi
}

func (r ChinchiroDiceRoll) IsArashi() bool {
	switch r.Combination {
	case ChinchiroCombinationPinzoro,
		ChinchiroCombinationNizoro,
		ChinchiroCombinationSanzoro,
		ChinchiroCombinationYonzoro,
		ChinchiroCombinationGozoro,
		ChinchiroCombinationRokuzoro:
		return true
	}
	return false
}

func CanRollDice(rolls []ChinchiroDiceRoll) bool {
	// 振れるのは3回まで
	if len(rolls) >= 3 {
		return false
	}
	// 既にメナシ以外の目が出ている場合は振れない
	if array.Any(rolls, func(r ChinchiroDiceRoll) bool {
		return r.Combination != ChinchiroCombinationMenashi
	}) {
		return false
	}
	return true
}

func DetermineDiceCombination(rolls []ChinchiroDiceRoll) *ChinchiroCombination {
	count := len(rolls)
	if count == 0 {
		return nil
	}
	lastCombination := rolls[count-1].Combination
	if count == 3 {
		return &lastCombination
	}
	if lastCombination == ChinchiroCombinationMenashi {
		return nil
	}
	return &lastCombination
}

// 子の獲得額を返す
func CalculateWinnings(combination ChinchiroCombination, dealerCombination ChinchiroCombination, betAmount int) int {
	ratio := combination.Ratio()
	dealerRatio := dealerCombination.Ratio()

	// 引き分け
	if ratio == dealerRatio {
		return 0
	}

	// 親の勝ち
	if ratio < dealerRatio {
		winnings := betAmount * dealerRatio
		// ヒフミの場合は2倍払い
		if ratio == -2 {
			return winnings * ratio
		} else {
			return -winnings
		}
	}
	// 子の勝ち
	winnings := betAmount * ratio
	// ヒフミの場合は2倍払い
	if dealerRatio == -2 {
		return -winnings * ratio
	} else {
		return winnings
	}
}
