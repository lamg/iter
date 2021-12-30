package iter

import (
	//	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSlice(t *testing.T) {
	r := require.New(t)
	ns := []int{0, 1, 2, 3, 4, 5}
	rs := ToSlice(Slice(ns))
	r.Equal(ns, rs)
}

func TestFilter(t *testing.T) {
	r := require.New(t)
	rs := PipeS(
		[]int{0, 1, 2, 3, 4, 5},
		Filter(func(n int) bool { return n > 2 }),
	)
	r.Equal([]int{3, 4, 5}, rs)
}

func TestConcat(t *testing.T) {
	r := require.New(t)
	its := Args(Slice([]int{0, 1, 2}), Slice([]int{3, 4}), Slice([]int{5}))
	xs := ToSlice(Concat(its))
	r.Equal([]int{0, 1, 2, 3, 4, 5}, xs)
}

func TestMap(t *testing.T) {
	r := require.New(t)
	const tail = "kkkk"
	xs := Slice([]string{"aeo", "uu"})
	p := Map(xs, func(s string) string { return s + tail })
	sl := ToSlice(p)
	r.Equal([]string{"aeo" + tail, "uu" + tail}, sl)
}

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
		{
			xs: []string{},
			rs: []string{},
		},
	}
	for _, j := range ts {
		ms := PipeS(
			j.xs,
			DropLast[string],
		)
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
		ms := PipeI(
			Const(","),
			Zip(Slice(j.xs)),
		)
		r.Equal(j.rs, ToSlice(ms))
	}
}

func TestSurround(t *testing.T) {
	r := require.New(t)
	sl := PipeS(
		[]string{"aeo", "uu"},
		Surround("(", ")"),
	)
	r.Equal([]string{"(", "aeo", "uu", ")"}, sl)
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
		ms := PipeS(j.xs, Intersperse(","))
		r.Equal(j.rs, ms)
	}
}

func TestPipe(t *testing.T) {
	r := require.New(t)
	sl := PipeS(
		[]string{"aeo", "uu"},
		Intersperse(","),
		Surround("(", ")"),
	)
	r.Equal([]string{"(", "aeo", ",", "uu", ")"}, sl)

	rs := PipeS(
		[]int{1, 2, 3},
	)
	r.Equal([]int{1, 2, 3}, rs)
}
