package characters

import (
	"fyne.io/fyne/v2/data/binding"
	"github.com/mcred/savetools"
)

// Attr : struct that embeds and extends the savetools attribute
type Attr struct {
	savetools.Attribute
	Value binding.Int
}

func (a *Attr) init(c *savetools.Card) {
	a.getValue(c)
	a.Value.AddListener(binding.NewDataListener(func() {
		a.setValue(c)
	}))
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
