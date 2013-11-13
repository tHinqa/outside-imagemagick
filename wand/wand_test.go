package magickwand

import "testing"

func TestInit(t *testing.T) {
	t.Log(MagickQueryConfigureOption("NAME"),
		MagickQueryConfigureOption("VERSION"))
	if !IsMagickWand(NewMagickWand()) {
		t.Error("IsMagickWand(NewMagickWand()) failed")
	}
}

func aTestQuery(t *testing.T) {
	// MagickWandGenesis()
	// defer MagickWandTerminus()
	var n size_t
	t.Log(n, MagickQueryFonts("*", &n))
	t.Log(n, MagickQueryFormats("*", &n))
	os := MagickQueryConfigureOptions("*", &n)
	t.Log(n, os)
	for _, o := range os {
		t.Log(o, ":", MagickQueryConfigureOption(o))
	}
}

func TestMagickWand(t *testing.T) {
	w := NewMagickWand()
	defer w.Destroy()
	var n size_t
	t.Log(n, MagickQueryFormats("*", &n))
	if w.NewImage(20, 20, w.BackgroundColor()) {
		var x, y float64
		w.ImageResolution(&x, &y)
		t.Log(x, y)
		for _, f := range []string{"gif", "png", "svg", "xbm"} {
			t.Log(f, w.WriteImage("test."+f))
		}
	}
}
