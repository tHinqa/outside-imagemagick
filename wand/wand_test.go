package wand

import (
	I "github.com/tHinqa/outside-imagemagick"
	"testing"
)

func aTestInit(t *testing.T) {
	var w *MagickWand
	t.Log(QueryConfigureOption("NAME"),
		QueryConfigureOption("VERSION"))
	w = NewMagickWand()
	if !IsMagickWand(NewMagickWand()) {
		t.Error("IsMagickWand(NewMagickWand()) failed")
	}
	if w != nil {
		var s uint32
		t.Logf("%x, %s", s, GetVersion(&s))
		w.Destroy()
	}
	t.Log(GetPackageName())
	var d uint32
	t.Log(GetQuantumDepth(&d), d)
	t.Log(GetReleaseDate())
	t.Log(GetCopyright())
	t.Log(GetHomeURL())
}

func TestQuery(t *testing.T) {
	var n uint32
	t.Log(n, QueryFonts("*", &n))
	t.Log(n, QueryFormats("*", &n))
	os := QueryConfigureOptions("*", &n)
	t.Log(n, os)
	for i, o := range os {
		t.Log(i, ":", o, ":", QueryConfigureOption(string(o)))
	}
}

func aTestMagickWand(t *testing.T) {
	w := NewMagickWand()
	defer w.Destroy()
	if w.New(1, 1, w.BackgroundColor()) {
		w.Resize(200, 200, 0, 0)
		w.AddNoise(I.GaussianNoise)
		w.Edge(0)
		w.Enhance()
		w.Despeckle()
		w.Emboss(10, 2)
		w.Equalize()
		for _, f := range []string{"png" /*, "gif", "jpg", "svg", "xbm"*/} {
			t.Log(f, w.Write("_test."+f))
		}
	}
}
