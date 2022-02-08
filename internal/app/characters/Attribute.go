package characters

import (
	"ChronoTrigger/internal/app/inventory"
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mcred/savetools"
)

// Attr : struct that embeds and extends the savetools attribute
type Attr struct {
	savetools.Attribute
	Value binding.Int
	Entry *widget.Entry
	List  *widget.List
	Items inventory.Inventory
}

type AttrType string

const (
	Entry  AttrType = "Entry"
	Select AttrType = "Select"
)

func (a *Attr) init(c *savetools.Card) {
	a.getValue(c)
	a.Value.AddListener(binding.NewDataListener(func() {
		a.setValue(c)
	}))
	if a.List != nil {
		for c, d := range a.Items.GetVals() {
			item, _ := a.Value.Get()
			s := a.Items.GetValByID(item)
			if s == d {
				a.List.Select(c)
			}
		}
	}
}

func (a *Attr) getValue(c *savetools.Card) {
	aa := savetools.Attribute{a.Location, a.Bits, a.Endianness}
	a.Value.Set(c.GetValue(aa))
}

func (a *Attr) setValue(c *savetools.Card) {
	aa := savetools.Attribute{a.Location, a.Bits, a.Endianness}
	v, _ := a.Value.Get()
	c.SetValue(aa, v)
}

func gen8Bit(l int) Attr {
	a := Attr{}
	a.Location = l
	a.Bits = 8
	a.Endianness = binary.LittleEndian
	a.Value = binding.NewInt()
	return a
}

func gen16Bit(l int) Attr {
	a := Attr{}
	a.Location = l
	a.Bits = 16
	a.Endianness = binary.LittleEndian
	a.Value = binding.NewInt()
	return a
}

func generate(l int, b int, t AttrType, i inventory.Inventory) Attr {
	var a Attr
	switch b {
	case 16:
		a = gen16Bit(l)
	default:
		a = gen8Bit(l)
	}
	switch t {
	case Select:
		items := i.GetVals()
		data := binding.BindStringList(&items)
		list := widget.NewListWithData(data,
			func() fyne.CanvasObject {
				return widget.NewLabel("template")
			},
			func(r binding.DataItem, o fyne.CanvasObject) {
				o.(*widget.Label).Bind(r.(binding.String))
			},
		)
		list.OnSelected = func(id widget.ListItemID) {
			a.Value.Set(i.GetIDByVal(items[id]))
		}
		a.List = list
		a.Items = i
	default:
		a.Entry = widget.NewEntryWithData(binding.IntToString(a.Value))
	}
	return a
}
