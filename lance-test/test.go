package main

func f() error {
    return nil
}

func g(h func() error) error {
    err := h()
    _ = 1
    return err
}

func h() (int, error) {
    return 0, nil
}

func i(e error) error {
    return e
}

func check(e error) {}

func main() {
    // error must be checked before any action!
    err := f()
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

    err3 := g(func() error {
        // test nested func
        err4 := f()
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
    err9 := i(err8)
    _ = 1
    if err9 != nil {
        panic(err9)
    }

    err10 := f()
    ch <- i(err10)
}
