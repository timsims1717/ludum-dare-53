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
	Properties             map[string]string
	Presented              bool
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
		"TrashingTheCamp":  AchievementFamily{Name: "TrashingTheCamp", StickyNote: nil, StickyNotePosition: pixel.V(-80, 500)},
	}

	Achievements = map[string]Achievement{
		"Create5Tetronimos":   Achievement{Name: "Create5Tetronimos", LabelText: "You have met the initial quota of 5 Tetronimos, training's over, now don't fall behind\n-Thanks Management", Description: "Construct 5 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 0, Properties: map[string]string{"target": "5"}},
		"Create8Tetronimos":   Achievement{Name: "Create8Tetronimos", LabelText: "Congrats on completing 8 Tetronimos, your new quota is 13\r\n-Thanks Management", Description: "Construct 8 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 1, Properties: map[string]string{"target": "8"}},
		"Create13Tetronimos":  Achievement{Name: "Create13Tetronimos", LabelText: "Only 13 Tetronimos? Your quota is 21 didn't you get the memo\r\n-Thanks Management", Description: "Construct 13 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 2, Properties: map[string]string{"target": "13"}},
		"Create21Tetronimos":  Achievement{Name: "Create21Tetronimos", LabelText: "Thanks for 21 Tetronimos, try for 34\n-Thanks Management", Description: "Construct 21 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 3, Properties: map[string]string{"target": "21"}},
		"Create34Tetronimos":  Achievement{Name: "Create34Tetronimos", LabelText: "Yay, managemement has worked very hard to get you to 34 Tetronimos, now do 55\r\n-Thanks Management", Description: "Construct 34 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 4, Properties: map[string]string{"target": "34"}},
		"Create55Tetronimos":  Achievement{Name: "Create55Tetronimos", LabelText: "Excellent you have done 55 Tetronimos, this earns us a bonus. Management that is. We can see when you get to 89.\n-Thanks Management", Description: "Construct 55 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 5, Properties: map[string]string{"target": "55"}},
		"Create89Tetronimos":  Achievement{Name: "Create89Tetronimos", LabelText: "As promised you have reached 89 Tetronimos, here is a company pen. We'll do more at 144...\n-Thanks Management", Description: "Construct 89 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 6, Properties: map[string]string{"target": "89"}},
		"Create144Tetronimos": Achievement{Name: "Create144Tetronimos", LabelText: "144 Tetronimos, go for 233\n-Thanks Management", Description: "Construct 144 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 7, Properties: map[string]string{"target": "144"}},
		"Create233Tetronimos": Achievement{Name: "Create233Tetronimos", LabelText: "This was a good shift of 233 Tetronimos, I fully expect you to hit 377 before lunch next time\r\n-Thanks Management", Description: "Construct 233 Valid Tetronimos and deliver them to the Board", MyFamily: AchievementFamilies["CreateTetronimos"], Achieved: false, AchievementFamilyOrder: 8, Properties: map[string]string{"target": "233"}},
		"Trash5":              Achievement{Name: "Trash5", LabelText: "You have wasted 5 Factromino shapes, please be more careful\r\n-Thanks Management", Description: "Put 5 Factromino shapes in the recycle", MyFamily: AchievementFamilies["TrashingTheCamp"], Achieved: false, AchievementFamilyOrder: 0, Properties: map[string]string{"target": "5"}},
		"Trash10":             Achievement{Name: "Trash10", LabelText: "Current waste 10: Factromino shapes, please be better\r\n-Thanks Management", Description: "Put 10 Factromino shapes in the recycle", MyFamily: AchievementFamilies["TrashingTheCamp"], Achieved: false, AchievementFamilyOrder: 1, Properties: map[string]string{"target": "10"}},
	}
)
