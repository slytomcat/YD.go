package tools

import (
	"log"
	"os"
	"testing"

	"github.com/slytomcat/llog"
)

func init() {
	llog.SetLevel(llog.DEBUG)
	llog.SetFlags(log.Lshortfile | log.Lmicroseconds)
}

func TestShortName(t *testing.T) {
	if len(ShortName("1234567890", 20)) != 10 {
		t.Error("Short string changed")
	}
	if len(ShortName("1234567890", 8)) != 8 {
		t.Error("Long string not changed")
	}
	if len(ShortName("12345678901", 8)) != 8 {
		t.Error("Long string not changed")
	}
	if len([]rune(ShortName("русский текст", 10))) != 10 {
		t.Error("Long string with national symbols not shorten as expected")
	}
	if len([]rune(ShortName("русский текст да", 40))) != 16 {
		t.Error("Short string with national symbols has unexpected length")
	}

}

func TestNotExists(t *testing.T) {
	wd, err := os.Getwd()
	if err != nil {
		t.Error("Test can't get current working directory.")
	}
	if NotExists(wd) {
		t.Error("NotExists reports that current working dir is not exists.")
	}
	if !NotExists("/Unreal path\n") {
		t.Error("NotExists reports that `/Unreal path\n` exists.")
	}

}

// I have no idea how to test XdgOpen...

// Need tests for AppInit
