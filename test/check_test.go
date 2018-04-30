package test

import (
	"testing"
	"github.com/MaJloe3Jlo/mapisacard_test/lib"
)

func TestCheckHolder(t *testing.T)  {
	holder := "DMITRIY KLESTOV"
	val := lib.CheckHolder(holder)
	if val != true {
		t.Error("Cardholder isn't correct")
	}
}
