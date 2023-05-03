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
			StickyNotePosition: pixel.V(-75, 500),
		},
		"AFullBoard": {
			Name:               "AFullBoard",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-125, 520),
		},
		"WhatDoIDoWithThis": {
			Name:               "WhatDoIDoWithThis",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-140, 560),
		},
		"GridFullOBlocks": {
			Name:               "GridFullOBlocks",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-170, 500),
		},
		"AchievementProgress": {
			Name:               "AchievementProgress",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-225, 490),
		},
		"CompletedAchievements": {
			Name:               "CompletedAchievements",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-90, 550),
		},
		"ThrowAwayATetromino": {
			Name:               "ThrowAwayATetromino",
			StickyNote:         nil,
			StickyNotePosition: pixel.V(-260, 505),
		},
	}

	Achievements = map[string]Achievement{
		"Create5Tetrominos": {
			Name:                   "Create5Tetrominos",
			LabelText:              "You have met the initial quota of 5 Tetrominos, if you fall behind, it's game over!\n-Management",
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
			Name:      "Create144Tetrominos",
			LabelText: "This was a good shift of 144 Tetrominos. I fully expect you to hit 233 before lunch next time.\r\n-Management",
			//LabelText:              "144 Tetrominos, go for 233.\n-Management",
			Description:            "Construct 144 Valid Tetrominos and deliver them to the Board",
			MyFamily:               AchievementFamilies["CreateTetrominos"],
			AchievementFamilyOrder: 7,
			Properties:             map[string]string{"target": "144"},
		},
		//"Create233Tetrominos": {
		//	Name:                   "Create233Tetrominos",
		//	LabelText:              "This was a good shift of 233 Tetrominos. I fully expect you to hit 377 before lunch next time.\r\n-Management",
		//	Description:            "Construct 233 Valid Tetrominos and deliver them to the Board",
		//	MyFamily:               AchievementFamilies["CreateTetrominos"],
		//	AchievementFamilyOrder: 8,
		//	Properties:             map[string]string{"target": "233"},
		//},
		"Trash5": {
			Name:                   "Trash5",
			LabelText:              "You have wasted 5 shapes, please be more careful\r\n-Management",
			Description:            "Put 5 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "5"},
		},
		"Trash10": {
			Name:                   "Trash10",
			LabelText:              "You've now wasted 10 shapes, please be better\r\n-Management",
			Description:            "Put 10 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 1,
			Properties:             map[string]string{"target": "10"},
		},
		"Trash20": {
			Name:                   "Trash20",
			LabelText:              "20 shapes thrown away! This is coming from your paycheck!\r\n-Management",
			Description:            "Put 20 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "20"},
		},
		"Trash100": {
			Name:                   "Trash100",
			LabelText:              "DO YOU DO ANY WORK!?!?! 100 shapes thrown away! They don't grow on trees, you know!\r\n-Management",
			Description:            "Put 100 Factromino shapes in the recycle",
			MyFamily:               AchievementFamilies["TrashingTheCamp"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "100"},
		},
		"FillingTheBoard1": {
			Name:                   "FillingTheBoard1",
			LabelText:              "Hey, you completed a line! Great job and keep it up.\r\n-Management",
			Description:            "Clear 1 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "1"},
		},
		"FillingTheBoard10": {
			Name:                   "FillingTheBoard10",
			LabelText:              "Ten lines cleared, what a great start!\r\n-Management",
			Description:            "Clear 10 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 5,
			Properties:             map[string]string{"target": "10"},
		},
		"FillingTheBoard20": {
			Name:                   "FillingTheBoard20",
			LabelText:              "Excellent, 20 lines cleared. We appreciate your business.\r\n-Management",
			Description:            "Clear 20 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 10,
			Properties:             map[string]string{"target": "20"},
		},
		"FillingTheBoard50": {
			Name:                   "FillingTheBoard50",
			LabelText:              "50 lines cleared? Nice. I'm clocking out, make sure to stay for your full shift.\r\n-Management",
			Description:            "Clear 50 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 20,
			Properties:             map[string]string{"target": "50"},
		},
		"FillingTheBoard100": {
			Name:                   "FillingTheBoard100",
			LabelText:              "Good Morning, you've cleared 100 lines? Were you here all night? No sleeping on the job.\r\n-Management",
			Description:            "Clear 100 Rows",
			MyFamily:               AchievementFamilies["AFullBoard"],
			AchievementFamilyOrder: 30,
			Properties:             map[string]string{"target": "100"},
		},
		"WhatDoIDoWithThis5": {
			Name:                   "WhatDoIDoWithThis5",
			LabelText:              "That Tetromino was too big, try not to do that again.\r\n-Management",
			Description:            "Trash a Factromino with >4 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "5"},
		},
		"WhatDoIDoWithThis10": {
			Name:                   "WhatDoIDoWithThis10",
			LabelText:              "Look, if you build them all too big we won't ever meet our goals.\r\n-Management",
			Description:            "Trash a Factromino with >9 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 1,
			Properties:             map[string]string{"target": "10"},
		},
		"WhatDoIDoWithThis20": {
			Name:                   "WhatDoIDoWithThis20",
			LabelText:              "Wow, I'm actually impressed you fit that thing in the dumpster.\r\n-Management",
			Description:            "Trash a Factromino with >19 Blocks",
			MyFamily:               AchievementFamilies["WhatDoIDoWithThis"],
			AchievementFamilyOrder: 2,
			Properties:             map[string]string{"target": "20"},
		},
		"WhatDoIDoWithThis35": {
			Name:                   "WhatDoIDoWithThis35",
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
		//"Achievement5Percent": {
		//	Name:                   "Achievement5Percent",
		//	LabelText:              "You found 5% of the Achievements, you are on your way to meeting the goals we have for you.\r\n-Management",
		//	Description:            "Get 5% of the Achievements",
		//	MyFamily:               AchievementFamilies["AchievementProgress"],
		//	AchievementFamilyOrder: 0,
		//	Properties:             map[string]string{"target": "0.05"},
		//},
		//"Achievement10Percent": {
		//	Name:                   "Achievement10Percent",
		//	LabelText:              "10% of the Achievements! Your dedication is astounding, it is definitely reflecting in my paycheck\r\n-Management",
		//	Description:            "Get 10% of the Achievements",
		//	MyFamily:               AchievementFamilies["AchievementProgress"],
		//	AchievementFamilyOrder: 0,
		//	Properties:             map[string]string{"target": "0.10"},
		//},
		"Achievement20Percent": {
			Name:                   "Achievement20Percent",
			LabelText:              "With 20% of the Achievements, I'll make it to VP in no time\r\n-Management",
			Description:            "Get 20% of the Achievements",
			MyFamily:               AchievementFamilies["AchievementProgress"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "0.20"},
		},
		"Achievement40Percent": {
			Name:                   "Achievement40Percent",
			LabelText:              "If you were 50% of the way there, we could talk about a bonus, but you are only at 40% Achievements.\r\n-Management",
			Description:            "Get 40% of the Achievements",
			MyFamily:               AchievementFamilies["AchievementProgress"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "0.40"},
		},
		"Achievement60Percent": {
			Name:                   "Achievement60Percent",
			LabelText:              "You mentioned something about a bonus... I don't recall that conversation, good job on 60% Achievements though.\r\n-Management",
			Description:            "Get 60% of the Achievements",
			MyFamily:               AchievementFamilies["AchievementProgress"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "0.60"},
		},
		"Achievement80Percent": {
			Name:                   "Achievement80Percent",
			LabelText:              "Remember the 80-20 rule, you do 80% of the Achievements and get 20% of the pay\r\n-Management",
			Description:            "Get 80% of the Achievements",
			MyFamily:               AchievementFamilies["AchievementProgress"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "0.80"},
		},
		"Achievement95Percent": {
			Name:                   "Achievement95Percent",
			LabelText:              "Wow, I mean you could have gone home days ago. I'll bring donuts to mark 95% of Achievements found.\r\n-Management",
			Description:            "Get 95% of the Achievements",
			MyFamily:               AchievementFamilies["AchievementProgress"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "0.95"},
		},
		//"Achievement95Percent": {
		//	Name:                   "Achievement95Percent",
		//	LabelText:              "You are so close, with over 95% of Achievements found. I am such a good manager.\r\n-Management",
		//	Description:            "Get 95% of the Achievements",
		//	MyFamily:               AchievementFamilies["AchievementProgress"],
		//	AchievementFamilyOrder: 0,
		//	Properties:             map[string]string{"target": "0.95"},
		//},
		"CompletedAchievements": {
			Name:                   "CompletedAchievements",
			LabelText:              "100% wow we are flattered, thank you for playing!\r\n-Tim & Ben",
			Description:            "Get 100% of the Achievements",
			MyFamily:               AchievementFamilies["CompletedAchievements"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "1"},
		},
		"ThrowAwayATetromino": {
			Name:                   "ThrowAwayATetromino",
			LabelText:              "That was a perfectly good piece, are you trying to get fired?\r\n-Management",
			Description:            "Throwaway a Tetromino",
			MyFamily:               AchievementFamilies["ThrowAwayATetromino"],
			AchievementFamilyOrder: 0,
			Properties:             map[string]string{"target": "4"},
		},
	}
)
