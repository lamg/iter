package iter

import (
	"strings"
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
		xsi := DropLast[string](Slice(j.xs))
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
		xsi := Slice(j.xs)
		ct := Const(",")
		zi := Zip[string](xsi, ct)
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
		xsi := Intersperse[string](Slice(j.xs), ",")
		ms := ToSlice(xsi)
		r.Equal(j.rs, ms)
	}
}

func TestMap(t *testing.T) {
	r := require.New(t)
	const tail = "kkkk"
	xs := Slice([]string{"aeo", "uu"})
	p := Map[string](xs, func(s string) string { return s + tail })
	sl := ToSlice(p)
	r.Equal([]string{"aeo" + tail, "uu" + tail}, sl)
}

func TestSurround(t *testing.T) {
	r := require.New(t)
	xs := Slice([]string{"aeo", "uu"})
	p := Surround[string](xs, "(", ")")
	sl := ToSlice[string](p)
	r.Equal([]string{"(", "aeo", "uu", ")"}, sl)
}

func TestCompose(t *testing.T) {
	r := require.New(t)
	c0 := []string{"aeo", "uu"}
	xs := Slice(c0)
	p0 := Intersperse[string](xs, ",")
	p1 := Surround(p0, "(", ")")
	sl := ToSlice[string](p1)
	r.Equal([]string{"(", "aeo", ",", "uu", ")"}, sl)
	ss := strings.Join(c0, ",")
	t.Log(ss)
}
