package iter

type Iterator[T any] interface {
	Current() (T, bool)
}

type bls[T any] struct {
	xs Iterator[T]
	f  func(T) bool
}

func BLS[T any](xs Iterator[T], f func(T) bool) Iterator[T] {
	return &bls[T]{xs: xs, f: f}
}

func (b *bls[T]) Current() (m T, ok bool) {
	m, ok = b.xs.Current()
	ok = ok && b.f(m)
	return
}

func Exec(fs Iterator[func()]) {
	m, ok := fs.Current()
	for ok {
		m()
		m, ok = fs.Current()
	}
}

type concat[T any] struct {
	a, b Iterator[T]
	okA  bool
}

func (c *concat[T]) Current() (m T, ok bool) {
	if c.okA {
		m, ok = c.a.Current()
		c.okA = ok
	}
	if !c.okA {
		m, ok = c.b.Current()
	}
	return
}

func Concat[T any](a, b Iterator[T]) (c Iterator[T]) {
	c = &concat[T]{a: a, b: b}
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
