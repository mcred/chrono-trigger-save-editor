package characters

import (
	"ChronoTrigger/internal/app/inventory"
	"encoding/binary"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/mcred/savetools"
)

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
	Weapons             inventory.Inventory
}

// CreateCharacter : Character based off a root position in the save file
// Creates form widget and sets up binding
func CreateCharacter(root int, weapons inventory.Inventory) Character {
	return Character{
		NameID:           generate(root, 8, entry, nil),
		CharID:           generate(root+0x01, 8, entry, nil),
		HP:               generate(root+0x03, 16, entry, nil),
		MaxHP:            generate(root+0x05, 16, entry, nil),
		MP:               generate(root+0x07, 16, entry, nil),
		MaxMP:            generate(root+0x09, 16, entry, nil),
		BasePower:        generate(root+0x0B, 8, entry, nil),
		BaseStamina:      generate(root+0x0C, 8, entry, nil),
		BaseSpeed:        generate(root+0x0D, 8, entry, nil),
		BaseMagic:        generate(root+0x0E, 8, entry, nil),
		BaseHit:          generate(root+0x0F, 8, entry, nil),
		BaseEvade:        generate(root+0x10, 8, entry, nil),
		BaseMagicDefense: generate(root+0x11, 8, entry, nil),
		Level:            generate(root+0x12, 8, entry, nil),
		//Experience:          gen8Bit(root + 0x13, 3},
		Helmet:              generate(root+0x27, 8, list, inventory.Helmets()),
		Armor:               generate(root+0x28, 8, list, inventory.Armors()),
		Weapon:              generate(root+0x29, 8, list, weapons),
		Relic:               generate(root+0x2A, 8, list, inventory.Relics()),
		ExpToLevel:          generate(root+0x2B, 16, entry, nil),
		CurrentPower:        generate(root+0x36, 8, entry, nil),
		CurrentStamina:      generate(root+0x37, 8, entry, nil),
		CurrentSpeed:        generate(root+0x38, 8, entry, nil),
		CurrentMagic:        generate(root+0x39, 8, entry, nil),
		CurrentHit:          generate(root+0x3A, 8, entry, nil),
		CurrentEvade:        generate(root+0x3B, 8, entry, nil),
		CurrentMagicDefense: generate(root+0x3C, 8, entry, nil),
		CurrentAttack:       generate(root+0x3D, 8, entry, nil),
		CurrentDefense:      generate(root+0x3E, 8, entry, nil),
		CurrentMaxHP:        generate(root+0x3F, 16, entry, nil),
		Name:                binding.NewString(),
	}
}

// Init loads values from card during load
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

// GenerateForm creates fyne widgets per character
func (c *Character) GenerateForm() *fyne.Container {
	f := container.NewVBox(
		widget.NewLabelWithData(c.Name),
		widget.NewLabel("Current HP: "),
		c.HP.Entry,
		widget.NewLabel("Max Hp: "),
		c.MaxHP.Entry,
		widget.NewLabel("Current MP: "),
		c.MP.Entry,
		widget.NewLabel("Base Max HP:"),
		c.MaxMP.Entry,
		widget.NewLabel("Base Power: "),
		c.BasePower.Entry,
		widget.NewLabel("Base Stamina: "),
		c.BaseStamina.Entry,
		widget.NewLabel("Base Speed: "),
		c.BaseSpeed.Entry,
		widget.NewLabel("Base Magic: "),
		c.BaseMagic.Entry,
		widget.NewLabel("Base Hit: "),
		c.BaseHit.Entry,
		widget.NewLabel("Base Evade: "),
		c.BaseEvade.Entry,
		widget.NewLabel("Base Magic Defence: "),
		c.BaseMagicDefense.Entry,
		widget.NewLabel("Level: "),
		c.Level.Entry,
		//widget.NewLabel("Experience: "),
		//widget.NewEntryWithData(binding.IntToString(c.Experience.Value)),
		widget.NewLabel("Helmet: "),
		c.Helmet.List,
		widget.NewLabel("Armors: "),
		c.Armor.List,
		widget.NewLabel("Weapon: "),
		c.Weapon.List,
		widget.NewLabel("Relic: "),
		c.Relic.List,
		widget.NewLabel("Exp To Level: "),
		c.ExpToLevel.Entry,
		widget.NewLabel("Current Power: "),
		c.CurrentPower.Entry,
		widget.NewLabel("Current Stamina: "),
		c.CurrentStamina.Entry,
		widget.NewLabel("Current Speed: "),
		c.CurrentSpeed.Entry,
		widget.NewLabel("Current Magic: "),
		c.CurrentMagic.Entry,
		widget.NewLabel("Current Hit: "),
		c.CurrentHit.Entry,
		widget.NewLabel("Current Evade: "),
		c.CurrentEvade.Entry,
		widget.NewLabel("Current Magic Defence: "),
		c.CurrentMagicDefense.Entry,
		widget.NewLabel("Current Attack: "),
		c.CurrentAttack.Entry,
		widget.NewLabel("Current Defense: "),
		c.CurrentDefense.Entry,
		widget.NewLabel("Current Max HP: "),
		c.CurrentMaxHP.Entry,
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

// GetIntFromChar : uint Get uint value for string character
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
