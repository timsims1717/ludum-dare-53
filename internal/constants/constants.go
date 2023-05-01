package constants

import (
	"image/color"
	"math/rand"
	"time"
	"timsims1717/ludum-dare-53/pkg/world"
)

const (
	Title   = "Well that didn't work"
	Release = 0
	Version = 1
	Build   = 20230428

	TypeFaceSize = 200.

	// Batches
	BlockKey   = "blocks"
	FactoryKey = "factory"

	// Tetris
	TileSize            = 32.
	TetrisWidth         = 10
	TetrisHeight        = 20
	DefaultSpeed        = 0.8
	ScoreCheckPoint     = 10
	SpeedModifier       = 0.05
	SpeedMax            = 2
	SpeedMin            = 0.1
	HighSpeedModifer    = 0.01
	HighSpeedMark       = 0.3
	MinPiecesToFail     = 5
	BalanceStreakTarget = 5

	// Factory
	FactoryTile   = 48.
	FactoryWidth  = 5
	FactoryHeight = 7

	ConvSpdModifier    = 5.
	ConvSpdMax         = 150.
	ConvSpdMin         = 50.
	HighConvSpdModifer = 1.
	HighConvSpdMark    = 125.
)

var (
	TitleText          = "LD53"
	GlobalSeededRandom = rand.New(rand.NewSource(time.Now().UnixNano()))
	BlackColor         = color.RGBA{
		R: 19,
		G: 19,
		B: 19,
		A: 255,
	}
	TVTextColor = color.RGBA{
		R: 15,
		G: 95,
		B: 21,
		A: 255,
	}
	TetrisStart = world.Coords{X: 4, Y: 19}

	IgnoreEmptyConv      = false
	AutoGenTetrominos    = false
	NormalizedTetronimos = map[[4]world.Coords]TetronimoType{
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {3, 0}}: I,
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {0, 3}}: I,
		[4]world.Coords{{0, 0}, {1, 0}, {0, 1}, {1, 1}}: O,
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {1, 1}}: T, //Point up, flat down
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {1, 1}}: T, //Point Right, Flat Left
		[4]world.Coords{{1, 0}, {0, 1}, {1, 1}, {2, 1}}: T, //Point Left, Flat Right
		[4]world.Coords{{1, 0}, {1, 1}, {1, 2}, {0, 1}}: T, //Point Down, Flat Up
		[4]world.Coords{{0, 1}, {1, 0}, {1, 1}, {2, 0}}: Z, //Horizontal
		[4]world.Coords{{0, 0}, {0, 1}, {1, 1}, {1, 2}}: Z, //Vertical
		[4]world.Coords{{0, 0}, {1, 0}, {1, 1}, {2, 1}}: S, //Horizonal
		[4]world.Coords{{0, 1}, {0, 2}, {1, 0}, {1, 1}}: S, //Vertical
		[4]world.Coords{{0, 0}, {1, 0}, {2, 0}, {2, 1}}: L, //flat down point up
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {1, 0}}: L, //flat left point right
		[4]world.Coords{{0, 0}, {0, 1}, {1, 1}, {2, 1}}: L, //flat up point down
		[4]world.Coords{{0, 2}, {1, 0}, {1, 1}, {1, 2}}: L, //flat right point left
		[4]world.Coords{{0, 0}, {0, 1}, {1, 0}, {2, 0}}: J, //flat down point up
		[4]world.Coords{{0, 0}, {0, 1}, {0, 2}, {1, 2}}: J, //flat left point right
		[4]world.Coords{{0, 1}, {1, 1}, {2, 1}, {2, 0}}: J, //flat up point down
		[4]world.Coords{{0, 0}, {1, 0}, {1, 1}, {1, 2}}: J, //flat right point left
	}
	TitleVariants = []string{
		"42 Days Accident Free",
		"Have you ever played a block falling game and thought, 'I like this game but wouldn't it be great if it was more stressful?'",
		"We need more Blocks! Everyone's working weekends!",
		"Your truck just barfed all over the factory floor, clean it up!",
		"We killed a bug, YAY!!!",
		"Stop everything! We need a full factory reset!",
		"Where did you think the blocks you mined in minecraft went?!?",
		"Don't forget to spend your hard earned dollars at the company store!",
		"There are two fail conditions, either you don't produce, or you can't consume",
		"My nose feels bigger",
		"Why do we even have that lever!?!",
		"This is an allegory for Capitalism",
		"Ain’t no fellow who regretted giving it one extra shake, but you can bet every guy has regretted giving one too few.",
	}
)

func RandomTitle() string {
	TitleText = TitleVariants[GlobalSeededRandom.Intn(len(TitleVariants))]
	return TitleText
}

type FailCondition int

const (
	BoardFull = iota
	ConveyorBeltEmpty
)

func (fc FailCondition) String() string {
	switch fc {
	case BoardFull:
		return "Game Over.\nYou cannot consume anymore of our products, the board is full"
	case ConveyorBeltEmpty:
		return "Game Over.\nYou did not meet your quota, the conveyor belt is empty"
	}
	return ""
}

type TetronimoType int

const (
	UndefinedTetronimoType = iota
	I
	O
	T
	S
	Z
	J
	L
)

func (t TetronimoType) String() string {
	switch t {
	case O:
		return "O"
	case I:
		return "I"
	case S:
		return "S"
	case Z:
		return "Z"
	case L:
		return "L"
	case J:
		return "J"
	case T:
		return "T"
	case UndefinedTetronimoType:
		return "Undefined"
	}
	return ""
}

type FactrominoType int

const (
	FacUndefined = 0
	FacOne       = 1
	FacTwo       = 2
	FacThree     = 3
)

func (f FactrominoType) String() string {
	switch f {
	case FacUndefined:
		return "Undefined Factromino"
	case FacOne:
		return "One Block Factromino"
	case FacTwo:
		return "Two Block Factromino"
	case FacThree:
		return "Three Block Factromino"
	}
	return ""
}

type FactrominoVariant int

const (
	FactVariantUndefined = iota
	Vertical
	Horizontal
	BabyR
	BabySeven
	BabyL
	BabyJ
)

func (f FactrominoVariant) String() string {
	switch f {
	case FactVariantUndefined:
		return "Undefined"
	case Vertical:
		return "Vertical"
	case Horizontal:
		return "Horizontal"
	case BabyR:
		return "Baby R"
	case BabySeven:
		return "Baby Seven"
	case BabyL:
		return "Baby L"
	case BabyJ:
		return "Baby J"
	}
	return ""
}
