package myecs

import (
	"github.com/bytearena/ecs"
	"timsims1717/ludum-dare-53/pkg/object"
)

var (
	FullCount   = 0
	IDCount     = 0
	LoadedCount = 0
)

type ClearFlag bool

var (
	Manager = ecs.NewManager()

	Temp   = Manager.NewComponent()
	Func   = Manager.NewComponent()
	Update = Manager.NewComponent()

	Drawable  = Manager.NewComponent()
	Animation = Manager.NewComponent()

	Object   = Manager.NewComponent()
	Parent   = Manager.NewComponent()
	ViewPort = Manager.NewComponent()
	Drag     = Manager.NewComponent()
	Click    = Manager.NewComponent()

	Block = Manager.NewComponent()

	IsTemp    = ecs.BuildTag(Temp, Object)
	HasFunc   = ecs.BuildTag(Func)
	HasUpdate = ecs.BuildTag(Update)

	HasAnimation = ecs.BuildTag(Animation, Object)
	IsDrawable   = ecs.BuildTag(Drawable, Object)

	IsObject  = ecs.BuildTag(Object)
	HasParent = ecs.BuildTag(Object, Parent)
	IsDrag    = ecs.BuildTag(Object, Drag, ViewPort)
	CanClick  = ecs.BuildTag(Object, Click, ViewPort)

	IsBlock = ecs.BuildTag(Object, Block)
)

func UpdateManager() {
	LoadedCount = 0
	IDCount = 0
	FullCount = 0
	for _, result := range Manager.Query(IsObject) {
		if t, ok := result.Components[Object].(*object.Object); ok {
			FullCount++
			if t.ID != "" {
				IDCount++
				if t.Load {
					LoadedCount++
				}
			}
		}
	}
}
