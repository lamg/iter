package iter

type Iterator[T any] interface {
	Current() (T, bool)
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

type Map[T any, U any] struct {
	xs Iterator[T]
	f  func(T) U
}

func NewMap[T any, U any](xs Iterator[T], f func(T) U) *Map[T, U] {
	return &Map[T, U]{xs: xs, f: f}
}

func (r *Map[T, U]) Current() (m U, ok bool) {
	n, ok := r.xs.Current()
	if ok {
		m = r.f(n)
	}
	return
}

type DropLast[T any] struct {
	a    Iterator[T]
	prev T
	cn   bool
}

func NewDropLast[T any](a Iterator[T]) *DropLast[T] {
	r := &DropLast[T]{a: a}
	r.prev, r.cn = a.Current()
	return r
}

func (r *DropLast[T]) Current() (m T, ok bool) {
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

type Zip[T any] struct {
	a, b Iterator[T]
	ca   bool // consume from a
	cn   bool // consume from next
}

func NewZip[T any](a, b Iterator[T]) *Zip[T] {
	return &Zip[T]{a: a, b: b, ca: true, cn: true}
}

func (r *Zip[T]) Current() (curr T, ok bool) {
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

type Const[T any] struct {
	curr T
}

func NewConst[T any](c T) *Const[T] {
	return &Const[T]{curr: c}
}

func (r *Const[T]) Current() (x T, ok bool) {
	x, ok = r.curr, true
	return
}

type Slice[T any] struct {
	xs []T
	i  int
}

func NewSlice[T any](xs []T) *Slice[T] {
	return &Slice[T]{xs: xs}
}

func (p *Slice[T]) Current() (m T, ok bool) {
	ok = p.i != len(p.xs)
	if ok {
		m, p.i = p.xs[p.i], p.i+1
	}
	return
}

func Intersperse[T any](xs Iterator[T], x T) (rs Iterator[T]) {
	rs = NewDropLast[T](NewZip[T](xs, NewConst(x)))
	return
}
