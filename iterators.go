package iter

type Iterator[T any] interface {
	Current() (T, bool)
}

type filter[T any] struct {
	xs Iterator[T]
	f  func(T) bool
}

func Filter[T any](xs Iterator[T], f func(T) bool) Iterator[T] {
	return &filter[T]{xs: xs, f: f}
}

func (b *filter[T]) Current() (m T, ok bool) {
	m, ok = b.xs.Current()
	found := b.f(m)
	for ok && !found {
		m, ok = b.xs.Current()
		found = ok && b.f(m)
	}
	return
}

type concat[T any] struct {
	xs   Iterator[Iterator[T]]
	curr Iterator[T]
	ok   bool
}

func (c *concat[T]) Current() (m T, ok bool) {
	for c.ok && !ok {
		m, ok = c.curr.Current()
		if !ok {
			c.curr, c.ok = c.xs.Current()
		}
	}
	return
}

func Concat[T any](xs Iterator[Iterator[T]]) (c Iterator[T]) {
	x := &concat[T]{xs: xs}
	x.curr, x.ok = xs.Current()
	c = x
	return
}

func ToSlice[T any](p Iterator[T]) (rs []T) {
	m, ok := p.Current()
	rs = make([]T, 0)
	for ok {
		rs = append(rs, m)
		m, ok = p.Current()
	}

	return
}

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

type dropLast[T any] struct {
	a    Iterator[T]
	prev T
	cn   bool
}

func DropLast[T any](a Iterator[T]) Iterator[T] {
	r := &dropLast[T]{a: a}
	r.prev, r.cn = a.Current()
	return r
}

func (r *dropLast[T]) Current() (m T, ok bool) {
	var curr T
	if r.cn {
		curr, ok = r.a.Current()
		if !ok {
			r.cn = false
		} else {
			m, ok = r.prev, true
			r.prev = curr
		}
	}
	return
}

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

type slice[T any] struct {
	xs []T
	i  int
}

func Slice[T any](xs []T) Iterator[T] {
	return &slice[T]{xs: xs}
}

func Args[T any](xs ...T) Iterator[T] {
	return Slice(xs)
}

func (p *slice[T]) Current() (m T, ok bool) {
	ok = p.i != len(p.xs)
	if ok {
		m, p.i = p.xs[p.i], p.i+1
	}
	return
}

func Intersperse[T any](xs Iterator[T], x T) (rs Iterator[T]) {
	rs = DropLast(Zip(xs, Const(x)))
	return
}

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
