package state

import (
	"timsims1717/ludum-dare-53/internal/data"
	"timsims1717/ludum-dare-53/internal/myecs"
	"timsims1717/ludum-dare-53/internal/systems"
	"timsims1717/ludum-dare-53/pkg/timing"
)

func CreateConveyor() {
	data.NewConveyor()
	e := myecs.Manager.NewEntity()
	e.AddComponent(myecs.Update, data.NewFn(UpdateConveyor))
	data.Conveyor.Entity = e
	FactoryBGEntities = append(FactoryBGEntities, e)
}

func UpdateConveyor() {
	for i, t := range data.Conveyor.Tets {
		if t != nil {
			//debug.AddText(fmt.Sprintf("Slot %d, (%d,%d)", i, int(t.Object.Pos.X), int(t.Object.Pos.Y)))
			if i == 0 {
				t.Moving = false
				if data.TetrisBoard.NextShape == nil {
					systems.FactoTet(t)
				}
			} else {
				next := data.Conveyor.Tets[i-1]
				if next == nil || next.Moving {
					t.Moving = true
					t.Object.Pos.X -= timing.DT * data.TetrisBoard.ConvSpd
					if t.Object.Pos.X < data.Conveyor.Slots[i-1].X {
						if next == nil {
							data.Conveyor.Tets[i-1] = t
							if i == data.ConveyorLength-1 {
								data.QueuePad.Tet = nil
							}
							data.Conveyor.Tets[i] = nil
						} else {
							t.Object.Pos.X = data.Conveyor.Slots[i-1].X
						}
					}
				} else {
					t.Moving = false
				}
			}
		}
	}
}
