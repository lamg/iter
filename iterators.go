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

type MapIter[T any, U any] struct {
	xs Iterator[T]
	f  func(T) U
}

func (r *MapIter[T, U]) Current() (m U, ok bool) {
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

func newDropLast[T any](a Iterator[T]) *DropLast[T] {
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

type ZipIter[T any] struct {
	a, b Iterator[T]
	ca   bool // consume from a
	cn   bool // consume from next
}

func NewZipIter[T any](a, b Iterator[T]) *ZipIter[T] {
	return &ZipIter[T]{a: a, b: b, ca: true, cn: true}
}

func (r *ZipIter[T]) Current() (curr T, ok bool) {
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

type ConstIter[T any] struct {
	curr T
}

func (r *ConstIter[T]) Current() (x T, ok bool) {
	x, ok = r.curr, true
	return
}

type SliceIter[T any] struct {
	xs []T
	i  int
}

func NewSliceIter[T any](xs []T) *SliceIter[T] {
	return &SliceIter[T]{xs: xs}
}

func (p *SliceIter[T]) Current() (m T, ok bool) {
	ok = p.i != len(p.xs)
	if ok {
		m, p.i = p.xs[p.i], p.i+1
	}
	return
}

func Intersperse[T any](xs Iterator[T], x T) (rs Iterator[T]) {
	rs = newDropLast[T](NewZipIter[T](xs, &ConstIter[T]{curr: x}))
	return
}
