package core

import (
	T "github.com/tHinqa/outside-imagemagick/types"
	"testing"
)

func TestInit(t *testing.T) {
	var v T.Size
	t.Logf("GetMagickVersion: %s (%x)\n", GetMagickVersion(&v), v)
}

func TestMime(t *testing.T) {
	t.Log(GetMimeList("*", new(T.Size), new(ExceptionInfo)))
}
