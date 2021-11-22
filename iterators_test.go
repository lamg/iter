package iter

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDropLast(t *testing.T) {
	r := require.New(t)
	ts := []struct {
		xs []string
		rs []string
	}{
		{
			xs: []string{"bla"},
			rs: []string{},
		},
		{
			xs: []string{"bla", "Bli"},
			rs: []string{"bla"},
		},
	}
	for _, j := range ts {
		xsi := NewDropLast[string](NewSlice(j.xs))
		ms := ToSlice[string](xsi)
		r.Equal(j.rs, ms)
	}
}

func TestZip(t *testing.T) {
	r := require.New(t)
	ts := []struct {
		xs []string
		rs []string
	}{
		{xs: []string{"a"}, rs: []string{"a", ","}},
		{xs: []string{"a", "b"}, rs: []string{"a", ",", "b", ","}},
	}
	for _, j := range ts {
		xsi := NewSlice(j.xs)
		ct := NewConst(",")
		zi := NewZip[string](xsi, ct)
		ms := ToSlice[string](zi)
		r.Equal(j.rs, ms)
	}
}

func TestIntersperse(t *testing.T) {
	r := require.New(t)
	ts := []struct {
		xs []string
		rs []string
	}{
		{xs: []string{"a"}, rs: []string{"a"}},
		{xs: []string{"a", "b"}, rs: []string{"a", ",", "b"}},
	}
	for _, j := range ts {
		xsi := Intersperse[string](NewSlice(j.xs), ",")
		ms := ToSlice(xsi)
		r.Equal(j.rs, ms)
	}
}

func TestMap(t *testing.T) {
	r := require.New(t)
	const tail = "kkkk"
	xs := NewSlice([]string{"aeo", "uu"})
	p := NewMap[string](xs, func(s string) string { return s + tail })
	sl := ToSlice[string](p)
	r.Equal([]string{"aeo" + tail, "uu" + tail}, sl)
}

func TestSurround(t *testing.T) {
	r := require.New(t)
	xs := NewSlice([]string{"aeo", "uu"})
	p := NewSurround[string](xs, "(", ")")
	sl := ToSlice[string](p)
	r.Equal([]string{"(", "aeo", "uu", ")"}, sl)
}
