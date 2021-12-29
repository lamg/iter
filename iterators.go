package iter

type Iterator[T any] interface {
	Current() (T, bool)
	Next()
}

// Slice iterator definition

type slice[T any] struct {
	xs []T
	i  int
	ok bool
}

func Slice[T any](xs []T) Iterator[T] {
	return &slice[T]{xs: xs, ok: len(xs) != 0}
}

func Args[T any](xs ...T) Iterator[T] {
	return Slice(xs)
}

func (p *slice[T]) Current() (m T, ok bool) {
	ok = p.ok
	if ok {
		m = p.xs[p.i]
	}
	return
}

func (p *slice[T]) Next() {
	if p.ok {
		p.i = p.i + 1
	}
	p.ok = p.i != len(p.xs)
}

func ToSlice[T any](p Iterator[T]) (rs []T) {
	m, ok := p.Current()
	rs = make([]T, 0)
	for ok {
		rs = append(rs, m)
		p.Next()
		m, ok = p.Current()
	}

	return
}

// end

// Filter iterator definition

type filter[T any] struct {
	xs Iterator[T]
	f  func(T) bool
}

func Filter[T any](xs Iterator[T], f func(T) bool) Iterator[T] {
	return &filter[T]{xs: xs, f: f}
}

func (b *filter[T]) Current() (m T, ok bool) {
	m, ok = b.xs.Current()
	found := ok && b.f(m)
	for ok && !found {
		b.xs.Next()
		m, ok = b.xs.Current()
		found = ok && b.f(m)
	}
	return
}

func (b *filter[T]) Next() {
	b.xs.Next()
}

// end

// Concat iterator definition

func Concat[T any](xs Iterator[Iterator[T]]) (c Iterator[T]) {
	c = &concat[T]{xs: xs}
	return
}

type concat[T any] struct {
	xs Iterator[Iterator[T]]
}

func (c *concat[T]) Current() (m T, ok bool) {
	curr, okXs := c.xs.Current()
	m, ok = curr.Current()
	for okXs && !ok {
		c.xs.Next()
		curr, okXs = c.xs.Current()
		if okXs {
			m, ok = curr.Current()
		}
	}
	return
}

func (c *concat[T]) Next() {
	curr, ok := c.xs.Current()
	if ok {
		curr.Next()
	}
}

// end

// Map iterator definition

type mapi[T any, U any] struct {
	xs Iterator[T]
	f  func(T) U
}

func Map[T any, U any](xs Iterator[T], f func(T) U) Iterator[U] {
	return &mapi[T, U]{xs: xs, f: f}
}

func (r *mapi[T, U]) Current() (m U, ok bool) {
	n, ok := r.xs.Current()
	if ok {
		m = r.f(n)
	}
	return
}

func (r *mapi[T, U]) Next() {
	r.xs.Next()
}

// end

// DropLast iterator definition

type dropLast[T any] struct {
	xs      Iterator[T]
	m0, m1  T
	hasNext bool // or m0 is last
}

func DropLast[T any](xs Iterator[T]) Iterator[T] {
	r := &dropLast[T]{xs: xs}
	var ok bool
	r.m0, ok = xs.Current()
	if ok {
		xs.Next()
		r.m1, r.hasNext = xs.Current()
	}
	return r
}

func (r *dropLast[T]) Current() (m T, ok bool) {
	if r.hasNext {
		m, ok = r.m0, true
	}
	return
}

func (r *dropLast[T]) Next() {
	if r.hasNext {
		r.m0 = r.m1
		r.xs.Next()
		r.m1, r.hasNext = r.xs.Current()
	}
}

// end

// Zip iterator definition
/*
type zip[T any] struct {
	a, b Iterator[T]
	ca   bool // consume from a
	cn   bool // consume from next
}

func Zip[T any](a, b Iterator[T]) Iterator[T] {
	return &zip[T]{a: a, b: b, ca: true, cn: true}
}

func (r *zip[T]) Current() (curr T, ok bool) {
	if r.cn {
		if r.ca {
			curr, ok = r.a.Current()
			r.ca = false
		} else {
			curr, ok = r.b.Current()
			r.ca = true
		}
		r.cn = ok
	}
	return
}

func (r *zip[T]) Next() {
	//
}
*/

// end

// Const iterator definition

type consti[T any] struct {
	curr T
}

func Const[T any](c T) Iterator[T] {
	return &consti[T]{curr: c}
}

func (r *consti[T]) Current() (x T, ok bool) {
	x, ok = r.curr, true
	return
}

func (r *consti[T]) Next() {

}

// end

// Surround iterator definition
/*
type surround[T any] struct {
	xs          Iterator[T]
	a, b        T
	first, last bool
}

func Surround[T any](xs Iterator[T], a, b T) Iterator[T] {
	return &surround[T]{xs: xs, a: a, b: b}
}

func (p *surround[T]) Current() (m T, ok bool) {
	if !p.first {
		m, ok, p.first = p.a, true, true
	} else if !p.last {
		m, ok = p.xs.Current()
		if !ok {
			m, ok, p.last = p.b, true, true
		}
	}
	return
}

// end

// Intersperse iterator definition
*/

//func Intersperse[T any](xs Iterator[T], x T) (rs Iterator[T]) {
//	rs = DropLast(Zip(xs, Const(x)))
//	return
//}

// end
