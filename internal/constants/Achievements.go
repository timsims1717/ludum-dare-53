package constants

type Achievement struct {
	Name              string
	LabelText         string
	Description       string
	Achieved          bool
	AchievementFamily AchievementFamilies
}

type AchievementFamilies int

const (
	CreateFactronimos = iota //3,2,1,1
	ClearedRows
)

var (
	Achievements = map[string]Achievement{
		"Create5Factronimos":   Achievement{Name: "Create5Factronimos", LabelText: "Congrats on completing 5 Factronimos", Description: "Construct 5 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create8Factronimos":   Achievement{Name: "Create8Factronimos", LabelText: "Congrats on completing 8 Factronimos", Description: "Construct 8 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create13Factronimos":  Achievement{Name: "Create13Factronimos", LabelText: "Congrats on completing 13 Factronimos", Description: "Construct 13 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create21Factronimos":  Achievement{Name: "Create21Factronimos", LabelText: "Congrats on completing 21 Factronimos", Description: "Construct 21 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create55Factronimos":  Achievement{Name: "Create55Factronimos", LabelText: "Congrats on completing 55 Factronimos", Description: "Construct 55 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create89Factronimos":  Achievement{Name: "Create89Factronimos", LabelText: "Congrats on completing 89 Factronimos", Description: "Construct 89 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create144Factronimos": Achievement{Name: "Create144Factronimos", LabelText: "Congrats on completing 144 Factronimos", Description: "Construct 144 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
		"Create233Factronimos": Achievement{Name: "Create233Factronimos", LabelText: "Congrats on completing 233 Factronimos", Description: "Construct 233 Valid Factronimos and deliver them to the Board", AchievementFamily: CreateFactronimos, Achieved: false},
	}
)
