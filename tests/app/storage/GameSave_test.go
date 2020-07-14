package storage_test

import (
	"ChronoTrigger/internal/app/storage"
	"reflect"
	"testing"
)

var path1 = "../../files/save1.sav"
var path2 = "../../files/save2.sav"
var game1 = storage.Open(path1)
var game2 = storage.Open(path2)

func TestCanLoad(t *testing.T) {
	if reflect.TypeOf(game1) != reflect.TypeOf(storage.Game{}) {
		t.Error("storage.Game Assertion Failed")
	}
}

func TestCanGet8BitValue(t *testing.T) {
	v := game1.GetValue(storage.Attribute{0x203, 1})
	if v != 70 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 70, v)
	}
}

func TestCanGet16BitValue(t *testing.T) {
	v := game2.GetValue(storage.Attribute{0x1FF0, 2})
	if v != 25819 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 25819, v)
	}
}

func TestCanSet8BitValue(t *testing.T) {
	game1.SetValue(storage.Attribute{0x203, 1}, 255)
	v := game1.GetValue(storage.Attribute{0x203, 1})
	if v != 255 {
		t.Errorf("GetValue(0x203, false): expected %d, actual %d", 255, v)
	}
}

func TestCanSet16BitValue(t *testing.T) {
	game2.SetValue(storage.Attribute{0x203, 2}, 2003)
	v := game2.GetValue(storage.Attribute{0x203, 2})
	if v != 2003 {
		t.Errorf("GetValue(0x203, true): expected %d, actual %d", 2003, v)
	}
}

func TestCanGenerateCheckSum(t *testing.T) {
	g1 := storage.Open(path1)
	g1.Save()
	c := g1.GetValue(storage.Attribute{0x1FF0, 2})
	if c != 46426 {
		t.Errorf("Checksum invalid: expected %d, actual %d", 46426, c)
	}

	g2 := storage.Open(path2)
	g2.Save()
	c = g2.GetValue(storage.Attribute{0x1FF0, 2})
	if c != 25819 {
		t.Errorf("Checksum invalid: expected %d, actual %d", 25819, c)
	}
	c = g2.GetValue(storage.Attribute{0x1FF2, 2})
	if c != 25060 {
		t.Errorf("Checksum invalid: expected %d, actual %d", 25060, c)
	}
	c = g2.GetValue(storage.Attribute{0x1FF4, 2})
	if c != 16873 {
		t.Errorf("Checksum invalid: expected %d, actual %d", 16873, c)
	}
}
