package wand

import (
	I "github.com/tHinqa/outside-imagemagick"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	var w *MagickWand
	t.Log(QueryConfigureOption("NAME"),
		QueryConfigureOption("VERSION"))
	w = NewMagickWand()
	if !IsMagickWand(w) {
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
	t.Log(n, "fonts", QueryFonts("*", &n))
	t.Log(n, "formats", QueryFormats("*", &n))
	opts := QueryConfigureOptions("*", &n)
	t.Log(n, opts)
	for i, o := range opts {
		t.Log(i, ":", o, ":", QueryConfigureOption(o))
	}
}

func TestMagickWand(t *testing.T) {
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
			fn := os.TempDir() + "/_test." + f
			t.Log(f, w.Write(fn))
			if e := os.Remove(fn); e != nil {
				t.Error(e)
			}
		}
	}
}
