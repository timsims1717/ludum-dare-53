package constants

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/object"
)

type Achievement struct {
	Name                   string
	LabelText              string
	Description            string
	Achieved               bool
	MyFamily               AchievementFamily
	AchievementFamilyOrder int
}

type AchievementFamily struct {
	Name               string
	StickyNote         *object.Object
	StickyNotePosition pixel.Vec
}

func (af *AchievementFamily) String() string {
	maxIter := -1
	message := ""
	for _, value := range Achievements {
		if value.MyFamily.Name == af.Name && value.Achieved && value.AchievementFamilyOrder > maxIter {
			maxIter = value.AchievementFamilyOrder
			message = value.LabelText
		}
	}
	return message
}
func (af *AchievementFamily) Achieved() bool {
	for _, value := range Achievements {
		if value.MyFamily.Name == af.Name && value.Achieved {
			return true
		}
	}
	return false
}

var (
	AchievementFamilies = map[string]AchievementFamily{
		"CreateTetronimos": AchievementFamily{Name: "CreateTetronimos", StickyNote: nil, StickyNotePosition: pixel.V(-40, 510)},
	}

	Achievements = map[string]Achievement{
		"Create5Tetronimos":   Achievement{Name: "Create5Tetronimos", LabelText: "You have met the initial quota of 5 Tetronimos, now don't fall behind", Description: "Construct 5 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 0},
		"Create8Tetronimos":   Achievement{Name: "Create8Tetronimos", LabelText: "Congrats on completing 8 Tetronimos", Description: "Construct 8 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 1},
		"Create13Tetronimos":  Achievement{Name: "Create13Tetronimos", LabelText: "Congrats on completing 13 Tetronimos", Description: "Construct 13 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 2},
		"Create21Tetronimos":  Achievement{Name: "Create21Tetronimos", LabelText: "Congrats on completing 21 Tetronimos", Description: "Construct 21 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 3},
		"Create55Tetronimos":  Achievement{Name: "Create55Tetronimos", LabelText: "Congrats on completing 55 Tetronimos", Description: "Construct 55 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 4},
		"Create89Tetronimos":  Achievement{Name: "Create89Tetronimos", LabelText: "Congrats on completing 89 Tetronimos", Description: "Construct 89 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 5},
		"Create144Tetronimos": Achievement{Name: "Create144Tetronimos", LabelText: "Congrats on completing 144 Tetronimos", Description: "Construct 144 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 6},
		"Create233Tetronimos": Achievement{Name: "Create233Tetronimos", LabelText: "Congrats on completing 233 Tetronimos", Description: "Construct 233 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 7},
	}
)
