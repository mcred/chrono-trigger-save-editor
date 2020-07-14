package storage

import (
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

type Attribute struct {
	Location int
	ByteLen   uint
}

type Game struct {
	Path string
	Data []byte
}

func (g *Game) GetValue(a Attribute) uint {
	var r uint
	switch a.ByteLen {
	case 3:
		b := []byte{0,0,0,0}
		for i := uint(0); i < a.ByteLen; i++ {
			b[i] = g.Data[uint(a.Location) + i]
		}
		r =  uint(binary.LittleEndian.Uint32(b))
	case 2:
		r = uint(binary.LittleEndian.Uint16(g.Data[a.Location : a.Location+2]))
	default:
		r = uint(g.Data[a.Location])
	}
	return r
}

func (g *Game) SetValue(a Attribute, v uint) {
	switch a.ByteLen {
	case 2:
		b := make([]byte, 2)
		binary.LittleEndian.PutUint16(b, uint16(v))
		g.Data[a.Location] = b[0]
		g.Data[a.Location+1] = b[1]
	default:
		g.Data[a.Location] = byte(v)
	}
}

func Open(path string) Game {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return Game{Path: path, Data: f}
}

func (g *Game) Save() {
	g.generateChecksum()
	err := ioutil.WriteFile(g.Path, g.Data, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println(g.Path + " updated.")
}

/*
Thanks to https://github.com/mikearnos/snessum for help with this
*/
func (g *Game) generateChecksum() {
	for slot := 1; slot <= 3; slot++ {
		var checksum uint = 0
		max := (slot * 0xA00) - 2
		min := (slot - 1) * 0xA00
		for i := max; i >= min; i -= 2 {
			if checksum > 0xFFFF {
				checksum -= 0xFFFF
			}
			checksum += uint(g.GetValue(Attribute{i, 2}))
		}
		checksum &= 0xFFFF
		l := 0x1FF0 + ((slot - 1) * 2)
		g.SetValue(Attribute{l, 2}, checksum)
	}
}
