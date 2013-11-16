package core

import (
	// . "fmt"
	// . "github.com/tHinqa/outside"
	"os"
	"testing"
	"unsafe"
)

func aTestInit(t *testing.T) {
	var v Size
	t.Logf("GetMagickVersion: %s (%x)\n", GetMagickVersion(&v), v)
	if !IsMagickTrue("1") || !IsMagickTrue("true") ||
		!IsMagickTrue("on") || !IsMagickTrue("yes") ||
		IsMagickTrue("0") || IsMagickTrue("false") ||
		IsMagickTrue("off") || IsMagickTrue("no") ||
		IsMagickTrue("") {
		t.Error("IsMagickTrue failed")
	}
	// InitializeMagick("") // Deprecated
	MagickCoreGenesis("", true)
	if !IsMagickInstantiated() {
		t.Log("IsMagickInstantiated() returned false; expected true")
	}
}

func StringToArray(o *[MaxTextExtent]Char, i string) {
	var n int
	for ; n < len(i); n++ {
		o[n] = Char(i[n])
	}
	o[n] = 0
}

func TestCopy(t *testing.T) {
	// InitializeMagick("")     // Deprecated or
	MagickCoreGenesis("", true)
	ii := AcquireImageInfo() // ~ ii := CloneImageInfo(nil)
	var e = new(ExceptionInfo)
	GetExceptionInfo(e)
	StringToArray(&ii.Filename, "wizard.jpg")
	i := ii.ReadImage(e)
	if i == nil {
		t.Error("read failed\n")
	}
	w1 := "_.png"
	StringToArray(&i.Filename, w1)
	if !ii.WriteImage(i) {
		t.Error("write failed")
	} else if e := os.Remove(w1); e != nil {
		t.Error(e)
	}
	w2 := "_.bmp"
	if !ii.WriteImages(i, w2, e) {
		t.Error("write failed")
	} else if e := os.Remove(w2); e != nil {
		t.Error(e)
	}
	var (
		tv = new(Timeval)
		tz = new(Timezone)
	)
	Gettimeofday(tv, tz)
	t.Logf("%+v %+v\n", tv, tz)
	Gettimeofday(tv, tz)
	t.Logf("%+v %+v\n", tv, tz)
}

func aTestMime(t *testing.T) {
	l := GetMimeList("image/*", new(Size), new(ExceptionInfo))
	for _, m := range l {
		t.Log(m)
	}
}

func aTestOffset(t *testing.T) {
	var i Image
	t.Logf("%X\n", unsafe.Offsetof(i.Extent_))
}
