package model

import "sort"

type ChinchiroDiceRoll struct {
	Dice1       int
	Dice2       int
	Dice3       int
	Combination ChinchiroCombination
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
