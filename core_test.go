package core

import (
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
	defer MagickCoreTerminus()
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
	defer MagickCoreTerminus()
	ii := AcquireImageInfo() // ~ ii := CloneImageInfo(nil)
	defer ii.Destroy()
	e := AcquireExceptionInfo()
	e.Get()
	r := "wizard.jpg"
	StringToArray(&ii.Filename, r)
	i := ii.ReadImage(e)
	if i == nil {
		t.Fatal(*e.Reason)
	}
	defer i.Destroy()
	w1 := "_.png"
	StringToArray(&i.Filename, w1)
	if !ii.WriteImage(i) {
		t.Error(i.Exception_.Reason)
	} else if e := os.Remove(w1); e != nil {
		t.Error(e)
	}
	w2 := "_.bmp"
	if !ii.WriteImages(i, w2, e) {
		t.Error(*e.Reason)
	} else if e := os.Remove(w2); e != nil {
		t.Error(e)
	}
}

func aTestMime(t *testing.T) {
	l := GetMimeList("image/*", new(Size), new(ExceptionInfo))
	for _, m := range l {
		t.Log(m)
	}
}

func aTestOffset(t *testing.T) {
	var i Image
	t.Logf("%X\n", unsafe.Offsetof(i.ProgressMonitor))
}

// Weird order problems results (6.8.0 bin)
func TestCommandOptions(t *testing.T) {
	for i := AlignOptions; i <= IntensityOptions; i++ {
		t.Logf("%s %d: %s", i, int(i), i.Options())
	}
}

func TestListFiles(t *testing.T) {
	var n Size
	t.Log(ListFiles(".", "*", &n))
}
