package types

import "testing"

func Test(t *testing.T) {
	t.Logf("%T", UndefinedColorInterpolate)
	t.Logf("%T", VoronoiColorInterpolate)
	t.Logf("%T", InverseColorInterpolate)
}
