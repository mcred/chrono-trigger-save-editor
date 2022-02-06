package characters

import (
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
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

// Character : struct with editable game attributes
type Character struct {
	NameID              Attr           `json:"nameID"`
	CharID              Attr           `json:"charID"`
	HP                  Attr           `json:"HP"`
	MaxHP               Attr           `json:"maxHP"`
	MP                  Attr           `json:"MP"`
	MaxMP               Attr           `json:"maxMP"`
	BasePower           Attr           `json:"basePower"`
	BaseStamina         Attr           `json:"baseStamina"`
	BaseSpeed           Attr           `json:"baseSpeed"`
	BaseMagic           Attr           `json:"baseMagic"`
	BaseHit             Attr           `json:"baseHit"`
	BaseEvade           Attr           `json:"baseEvade"`
	BaseMagicDefense    Attr           `json:"baseMagicDefense"`
	Level               Attr           `json:"level"`
	Experience          Attr           `json:"experience"`
	Helmet              Attr           `json:"helmet"`
	Armor               Attr           `json:"armor"`
	Weapon              Attr           `json:"weapon"`
	Relic               Attr           `json:"relic"`
	ExpToLevel          Attr           `json:"expToLevel"`
	CurrentPower        Attr           `json:"currentPower"`
	CurrentStamina      Attr           `json:"currentStamina"`
	CurrentSpeed        Attr           `json:"currentSpeed"`
	CurrentMagic        Attr           `json:"currentMagic"`
	CurrentHit          Attr           `json:"currentHit"`
	CurrentEvade        Attr           `json:"currentEvade"`
	CurrentMagicDefense Attr           `json:"currentMagicDefense"`
	CurrentAttack       Attr           `json:"currentAttack"`
	CurrentDefense      Attr           `json:"currentDefense"`
	CurrentMaxHP        Attr           `json:"currentMaxHP"`
	Name                binding.String `json:"name"`
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

// CreateCharacter : Character based off a root position in the save file
func CreateCharacter(root int) Character {
	return Character{
		NameID:           gen8Bit(root),
		CharID:           gen8Bit(root + 0x01),
		HP:               gen16Bit(root + 0x03),
		MaxHP:            gen16Bit(root + 0x05),
		MP:               gen16Bit(root + 0x07),
		MaxMP:            gen16Bit(root + 0x09),
		BasePower:        gen8Bit(root + 0x0B),
		BaseStamina:      gen8Bit(root + 0x0C),
		BaseSpeed:        gen8Bit(root + 0x0D),
		BaseMagic:        gen8Bit(root + 0x0E),
		BaseHit:          gen8Bit(root + 0x0F),
		BaseEvade:        gen8Bit(root + 0x10),
		BaseMagicDefense: gen8Bit(root + 0x11),
		Level:            gen8Bit(root + 0x12),
		//Experience:          gen8Bit(root + 0x13, 3},
		Helmet:              gen8Bit(root + 0x27),
		Armor:               gen8Bit(root + 0x28),
		Weapon:              gen8Bit(root + 0x29),
		Relic:               gen8Bit(root + 0x2A),
		ExpToLevel:          gen16Bit(root + 0x2B),
		CurrentPower:        gen8Bit(root + 0x36),
		CurrentStamina:      gen8Bit(root + 0x37),
		CurrentSpeed:        gen8Bit(root + 0x38),
		CurrentMagic:        gen8Bit(root + 0x39),
		CurrentHit:          gen8Bit(root + 0x3A),
		CurrentEvade:        gen8Bit(root + 0x3B),
		CurrentMagicDefense: gen8Bit(root + 0x3C),
		CurrentAttack:       gen8Bit(root + 0x3D),
		CurrentDefense:      gen8Bit(root + 0x3E),
		CurrentMaxHP:        gen16Bit(root + 0x3F),
		Name:                binding.NewString(),
	}
}

func (c *Character) Init(card *savetools.Card) {
	c.NameID.init(card)
	c.CharID.init(card)
	c.HP.init(card)
	c.MaxHP.init(card)
	c.MP.init(card)
	c.MaxMP.init(card)
	c.BasePower.init(card)
	c.BaseStamina.init(card)
	c.BaseSpeed.init(card)
	c.BaseMagic.init(card)
	c.BaseHit.init(card)
	c.BaseEvade.init(card)
	c.BaseMagicDefense.init(card)
	c.Level.init(card)
	//c.Experience.init(card)
	c.Helmet.init(card)
	c.Armor.init(card)
	c.Weapon.init(card)
	c.Relic.init(card)
	c.ExpToLevel.init(card)
	c.CurrentPower.init(card)
	c.CurrentStamina.init(card)
	c.CurrentSpeed.init(card)
	c.CurrentMagic.init(card)
	c.CurrentHit.init(card)
	c.CurrentEvade.init(card)
	c.CurrentMagicDefense.init(card)
	c.CurrentAttack.init(card)
	c.CurrentDefense.init(card)
	c.CurrentMaxHP.init(card)
	c.Name.Set(c.DecodeName(card))
}

func (c *Character) GenerateForm(card *savetools.Card) *fyne.Container {
	f := container.NewVBox(
		widget.NewLabelWithData(c.Name),
		widget.NewLabel("Current HP: "),
		widget.NewEntryWithData(binding.IntToString(c.HP.Value)),
		widget.NewLabel("Max Hp: "),
		widget.NewEntryWithData(binding.IntToString(c.MaxHP.Value)),
		widget.NewLabel("Current MP: "),
		widget.NewEntryWithData(binding.IntToString(c.MP.Value)),
		widget.NewLabel("Base Max HP:"),
		widget.NewEntryWithData(binding.IntToString(c.MaxMP.Value)),
		widget.NewLabel("Base Power: "),
		widget.NewEntryWithData(binding.IntToString(c.BasePower.Value)),
		widget.NewLabel("Base Stamina: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseStamina.Value)),
		widget.NewLabel("Base Speed: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseSpeed.Value)),
		widget.NewLabel("Base Magic: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseMagic.Value)),
		widget.NewLabel("Base Hit: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseHit.Value)),
		widget.NewLabel("Base Evade: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseEvade.Value)),
		widget.NewLabel("Base Magic Defence: "),
		widget.NewEntryWithData(binding.IntToString(c.BaseMagicDefense.Value)),
		widget.NewLabel("Level: "),
		widget.NewEntryWithData(binding.IntToString(c.Level.Value)),
		//widget.NewLabel("Experience: "),
		//widget.NewEntryWithData(binding.IntToString(c.Experience.Value)),
		widget.NewLabel("Helmet: "),
		widget.NewEntryWithData(binding.IntToString(c.Helmet.Value)),
		widget.NewLabel("Armors: "),
		widget.NewEntryWithData(binding.IntToString(c.Armor.Value)),
		widget.NewLabel("Weapon: "),
		widget.NewEntryWithData(binding.IntToString(c.Weapon.Value)),
		widget.NewLabel("Relic: "),
		widget.NewEntryWithData(binding.IntToString(c.Relic.Value)),
		widget.NewLabel("Exp To Level: "),
		widget.NewEntryWithData(binding.IntToString(c.ExpToLevel.Value)),
		widget.NewLabel("Current Power: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentPower.Value)),
		widget.NewLabel("Current Stamina: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentStamina.Value)),
		widget.NewLabel("Current Speed: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentSpeed.Value)),
		widget.NewLabel("Current Magic: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentMagic.Value)),
		widget.NewLabel("Current Hit: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentHit.Value)),
		widget.NewLabel("Current Evade: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentEvade.Value)),
		widget.NewLabel("Current Magic Defence: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentMagicDefense.Value)),
		widget.NewLabel("Current Attack: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentAttack.Value)),
		widget.NewLabel("Current Defense: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentDefense.Value)),
		widget.NewLabel("Current Max HP: "),
		widget.NewEntryWithData(binding.IntToString(c.CurrentMaxHP.Value)),
	)
	return f
}

// CharMap : map[uint]string in game int to string mapping for text
var CharMap = map[int]string{
	0xA0: "A",
	0xA1: "B",
	0xA2: "C",
	0xA3: "D",
	0xA4: "E",
	0xA5: "F",
	0xA6: "G",
	0xA7: "H",
	0xA8: "I",
	0xA9: "J",
	0xAA: "K",
	0xAB: "L",
	0xAC: "M",
	0xAD: "N",
	0xAE: "O",
	0xAF: "P",
	0xB0: "Q",
	0xB1: "R",
	0xB2: "S",
	0xB3: "T",
	0xB4: "U",
	0xB5: "V",
	0xB6: "W",
	0xB7: "X",
	0xB8: "Y",
	0xB9: "Z",
	0xBA: "a",
	0xBB: "b",
	0xBC: "c",
	0xBD: "d",
	0xBE: "e",
	0xBF: "f",
	0xC0: "g",
	0xC1: "h",
	0xC2: "i",
	0xC3: "j",
	0xC4: "k",
	0xC5: "l",
	0xC6: "m",
	0xC7: "n",
	0xC8: "o",
	0xC9: "p",
	0xCA: "q",
	0xCB: "r",
	0xCC: "s",
	0xCD: "t",
	0xCE: "u",
	0xCF: "v",
	0xD0: "w",
	0xD1: "x",
	0xD2: "y",
	0xD3: "z",
	0xD4: "0",
	0xD5: "1",
	0xD6: "2",
	0xD7: "3",
	0xD8: "4",
	0xD9: "5",
	0xDA: "6",
	0xDB: "7",
	0xDC: "8",
	0xDD: "9",
	0xDE: "!",
	0xDF: "?",
	0xE0: "/",
	0xE1: "\"",
	0xE2: "\"",
	0xE3: ":",
	0xE4: "%",
	0xE5: "(",
	0xE6: ")",
	0xE7: "'",
	0xE8: ".",
	0xE9: ",",
	0xEA: "=",
	0xEB: "-",
	0xEC: "+",
	0xED: "\\",
	0xFF: " ",
}

// GetCharFromInt : string Get string from CharMap uint
func GetCharFromInt(i int) string {
	return CharMap[i]
}

// GetIntFromChar: uint Get uint value for string character
func GetIntFromChar(c string) int {
	var i int
	for index, element := range CharMap {
		if element == c {
			i = index
		}
	}
	return i
}

// DecodeName : string Get name for character from Save File
func (c *Character) DecodeName(card *savetools.Card) string {
	id, _ := c.NameID.Value.Get()
	ints := []int{
		card.GetValue(savetools.Attribute{0x5B0 + (id * 6), 8, binary.LittleEndian}),
		card.GetValue(savetools.Attribute{0x5B1 + (id * 6), 8, binary.LittleEndian}),
		card.GetValue(savetools.Attribute{0x5B2 + (id * 6), 8, binary.LittleEndian}),
		card.GetValue(savetools.Attribute{0x5B3 + (id * 6), 8, binary.LittleEndian}),
		card.GetValue(savetools.Attribute{0x5B4 + (id * 6), 8, binary.LittleEndian}),
		card.GetValue(savetools.Attribute{0x5B5 + (id * 6), 8, binary.LittleEndian}),
	}
	var s string
	for _, i := range ints {
		s += GetCharFromInt(i)
	}
	return s
}

// EncodeName : []uint Generate uint splice for character string
func EncodeName(n string) []int {
	i := []int{0, 0, 0, 0, 0, 0}
	for p, c := range n {
		i[p] = GetIntFromChar(string(c))
	}
	return i
}
