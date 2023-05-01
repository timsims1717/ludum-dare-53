package constants

import (
	"github.com/faiface/pixel"
	"timsims1717/ludum-dare-53/pkg/object"
)

var NoteVec = pixel.V(40., 115.)

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
		"CreateTetrominos": {
			Name:               "CreateTetrominos",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-40, 510),
		},
		"TrashingTheCamp": {
			Name:               "TrashingTheCamp",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-80, 500),
		},
		"AFullBoard": {
			Name:               "AFullBoard",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-120, 520),
		},
		"WhatDoIDoWithThis": {
			Name:               "WhatDoIDoWithThis",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-140, 510),
		},
		"GridFullOBlocks": {
			Name:               "GridFullOBlocks",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-180, 500),
		},
	}

	Achievements = map[string]Achievement{
		"Create5Tetrominos": {
			Name:                   "Create5Tetrominos",
			LabelText:              "You have met the initial quota of 5 Tetrominos, training's over, now don't fall behind!\n-Management",
			Description:            "Construct 5 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "5"},
		},
		"Create8Tetrominos": {
			Name:                   "Create8Tetrominos",
			LabelText:              "Congrats on completing 8 Tetrominos, your new quota is 13.\r\n-Management",
			Description:            "Construct 8 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 1,
			Properties:             map[string]string{"target": "8"},
		},
		"Create13Tetrominos": {
			Name:                   "Create13Tetrominos",
			LabelText:              "Only 13 Tetrominos? Your quota is 21 didn't you get the memo?\r\n-Management",
			Description:            "Construct 13 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "13"},
		},
		"Create21Tetrominos": {
			Name:                   "Create21Tetrominos",
			LabelText:              "Thank you for 21 Tetrominos, can you do 34?\n-Management",
			Description:            "Construct 21 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 3,
			Properties:             map[string]string{"target": "21"},
		},
		"Create34Tetrominos": {
			Name:                   "Create34Tetrominos",
			LabelText:              "Great work! Management has worked very hard to get you to 34 Tetrominos. Now do 55.\r\n-Management",
			Description:            "Construct 34 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 4,
			Properties:             map[string]string{"target": "34"},
		},
		"Create55Tetrominos": {
			Name:                   "Create55Tetrominos",
			LabelText:              "Excellent. You have done 55 Tetrominos, earning a bonus. For Management, that is. We'll see once you get to 89.\n-Management",
			Description:            "Construct 55 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 5,
			Properties:             map[string]string{"target": "55"},
		},
		"Create89Tetrominos": {
			Name:                   "Create89Tetrominos",
			LabelText:              "You have reached 89 Tetrominos. As promised, here is a company pen. Next stop, 144.\n-Management",
			Description:            "Construct 89 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 6,
			Properties:             map[string]string{"target": "89"},
		},
		"Create144Tetrominos": {
			Name:                   "Create144Tetrominos",
			LabelText:              "144 Tetrominos, go for 233.\n-Management",
			Description:            "Construct 144 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 7,
			Properties:             map[string]string{"target": "144"},
		},
		"Create233Tetrominos": {
			Name:                   "Create233Tetrominos",
			LabelText:              "This was a good shift of 233 Tetrominos. I fully expect you to hit 377 before lunch next time.\r\n-Management",
			Description:            "Construct 233 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 8,
			Properties:             map[string]string{"target": "233"},
		},
		"Trash5": {
			Name:                   "Trash5",
			LabelText:              "You have wasted 5 Factromino shapes, please be more careful\r\n-Management",
			Description:            "Put 5 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "5"},
		},
		"Trash10": {
			Name:                   "Trash10",
			LabelText:              "Current waste 10: Factromino shapes, please be better\r\n-Management",
			Description:            "Put 10 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 1,
			Properties:             map[string]string{"target": "10"},
		},
		"Trash20": {
			Name:                   "Trash20",
			LabelText:              "Current waste 20: Factromino shapes, this is coming from your paycheck!\r\n-Management",
			Description:            "Put 20 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "20"},
		},
		"FillingTheBoard": {
			Name:                   "FullBoard",
			LabelText:              "Excellent, you've consumed an entire board's worth of blocks, we appreciate your business.\r\n-Management",
			Description:            "Clear 20 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 10,
			Properties:             map[string]string{"target": "20"},
		},
		"WhatDoIDoWithThis5": {
			Name:                   "WhatDoIDoWithThis",
			LabelText:              "That Tetromino was too big, try not to do that again.\r\n-Management",
			Description:            "Trash a Factromino with >4 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "5"},
		},
		"WhatDoIDoWithThis10": {
			Name:                   "WhatDoIDoWithThis",
			LabelText:              "Look, if you build them all too big we won't ever meet our goals.\r\n-Management",
			Description:            "Trash a Factromino with >9 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 1,
			Properties:             map[string]string{"target": "10"},
		},
		"WhatDoIDoWithThis20": {
			Name:                   "WhatDoIDoWithThis",
			LabelText:              "Wow, I'm actually impressed you fit that thing in the dumpster.\r\n-Management",
			Description:            "Trash a Factromino with >19 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "20"},
		},
		"WhatDoIDoWithThis35": {
			Name:                   "WhatDoIDoWithThis",
			LabelText:              "You know, I'm beginning to think this job may not be the right fit for you.\r\n-Management",
			Description:            "Trash a Factromino with 35 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "35"},
		},
		"GridFullOBlocks": {
			Name:                   "GridFullOBlocks",
			LabelText:              "Please clean up your workspace, this is a health and safety issue.\r\n-Management",
			Description:            "Fill the workspace with blocks",
			MyFamily:               AchievementFamilies["GridFullOBlocks"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "35"},
		},
	}
)
