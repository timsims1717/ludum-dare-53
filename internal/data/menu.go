package data

import (
	"timsims1717/ludum-dare-53/pkg/typeface"
)

type MenuItem struct {
	Click    func()
	Text     *typeface.Text
	OrigSize float64
}

var MenuItems []*MenuItem
