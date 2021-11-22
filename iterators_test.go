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
		xsi := newDropLast[string](NewSliceIter(j.xs))
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
		xsi := NewSliceIter(j.xs)
		ct := &ConstIter[string]{curr: ","}
		zi := NewZipIter[string](xsi, ct)
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
		xsi := Intersperse[string](NewSliceIter(j.xs), ",")
		ms := ToSlice(xsi)
		r.Equal(j.rs, ms)
	}
}

func TestMap(t *testing.T) {
	xs := NewSliceIter([]string{"aeo", "uu"})
	mi := &MapIter[string, string]{xs: xs, f: func(s string) string { return s + "kkkk" }}
	var xi Iterator[string]
	xi = mi
	c, ok := xi.Current()
	for ok {
		t.Log(c)
		c, ok = xi.Current()
	}
}
