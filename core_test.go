package core

import (
	"testing"
)

func TestInit(t *testing.T) {
	var v Size
	t.Logf("GetMagickVersion: %s (%x)\n", GetMagickVersion(&v), v)
}

func TestMime(t *testing.T) {
	t.Log(GetMimeList("*", new(Size), new(ExceptionInfo)))
}

func TestEnum(t *testing.T) {
	t.Logf("%T", VoronoiColorInterpolate)
}
