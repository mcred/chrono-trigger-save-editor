package characters_test

import (
	"ChronoTrigger/internal/app/characters"
	"reflect"
	"testing"
)

func TestCanGetCrono(t *testing.T) {
	c := characters.Crono()
	if reflect.TypeOf(c) != reflect.TypeOf(characters.Character{}) {
		t.Error("characters.Crono Assertion Failed")
	}
}
