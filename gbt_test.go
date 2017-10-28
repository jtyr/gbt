package main

import (
    "testing"
    "os"
)

func TestMain(t *testing.T) {
    var ran bool

    run = func() {
        ran = true
    }

    os.Args[1] = "-help"

    main()

    if ! ran {
        t.Error("Expected Run() to be called, but it wasn't.")
    }
}
