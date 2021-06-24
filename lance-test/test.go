package main

func f() error {
	return nil
}

func g(h func() error) error {
	err := h() // SHOULD REPORT
	_ = 1
	return err
}

func h() (int, error) {
	return 0, nil
}

func i(e error) error {
	return e
}

func j() error {
	e := f()
	return e
}

func k() error {
	var err13 error
	err13 = f()
	if err13 == nil {
		err13 = f()
		if err13 == nil {
			err13 = f()
		}
	}
	return i(err13)
}

func l() error {
	return f()
}

func m() (int, error) {
	var err14 error
	a := 1
	switch a {
	case 1:
		err14 = f()
	default:
		err14 = f()
	}
	return a, err14
}

func n() (err error) {
	err = f() // TODO: should ignore
	return
}

func o(error) bool {
	return false
}

func p(int) bool {
	return false
}

func check(e error) {}

func main() {
	// error must be checked before any action!
	err := f() // SHOULD REPORT
	a := 1
	if err != nil {
		panic(err)
	}
	_ = a

	// this is fine with empty line
	err2 := f()

	if err2 != nil {
		panic(err2)
	}

	err3 := g(func() error { // SHOULD REPORT
		// test nested func
		err4 := f() // SHOULD REPORT
		_ = 1
		return err4
	})
	_ = 1
	if err3 != nil {
		panic(err3)
	}

	// could handle error by function
	err5 := f()
	check(err5)

	ch := make(chan error, 10)
	// could handle error by channel
	err6 := f()
	ch <- err6

	// support multiple return value
	iii, err7 := h()
	if err7 != nil {
		panic(err7)
	}
	_ = iii

	// return the second error
	err8 := f()
	err9 := i(err8) // SHOULD REPORT
	_ = 1
	if err9 != nil {
		panic(err9)
	}

	err10 := f()
	ch <- i(err10)

	if err11 := f(); err11 != nil {
		panic(err11)
	}

	var err12 error
	switch iii {
	case 1:
		err12 = f()
	default:
	}
	if err12 != nil {
		panic(err12)
	}

	err15 := f()
	switch err15 {
	case nil:
	}

	err16 := f()
	switch i(err16) {
	case nil:
	}

	if err17 := f(); err17 != nil && o(err17) {
		panic(err17)
	}

	var(
		b bool
		err18 error
	)
	if b {
		err18 = f() // TODO: should ignore
	} else {
		err18 = f()
	}
	check(err18)

	var err19 error
	switch iii {
	case 1:
		err19 = f() // TODO: SHOULD REPORT
	default:
		err19 = f() // SHOULD REPORT
	}
	_ = iii
	if err19 != nil {
		panic(err19)
	}

	var err20 error
	switch iii {
	case 1:
		err20 = f()
	default:
		_ = p(iii)
	}
	check(err20)

	err16 = f() // TODO: SHOULD REPORT
}
