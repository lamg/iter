package iter

func Exec(fs Iterator[func()]) {
	m, ok := fs.Current()
	for ok {
		m()
		m, ok = fs.Current()
	}
}

// I executes all members of fs until the last or until it finds the first false guard
// guard is called before every execution of fs members
func I(guard func() bool, fs ...func()) (reachedLast bool) {
	i := 0
	Exec(BLS(Slice(fs), func(_ func()) bool { i++; return guard() }))
	reachedLast = i == len(fs)
	return
}

// W it executes repeteadly all members of fs until it finds the first false guard
// guard is called before every execution of fs members
func W(guard func() bool, fs ...func()) {
	i, do := 0, len(fs) != 0
	for do {
		if i == len(fs) {
			i = 0
		}
		fs[i]()
		i = i + 1
		do = guard()
	}
}