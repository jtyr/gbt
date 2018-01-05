package main

import (
    "os"
    "testing"
)

func TestMain(t *testing.T) {
    var ran bool

    myPrint = func(s string) {
        ran = true
    }

    os.Setenv("GBT_CARS", "Status, Os, Time, Custom, Hostname, Dir, PyVirtEnv, Git, Sign")
    os.Setenv("GBT_CAR_SIGN_WRAP", "1")

    main()

    if ! ran {
        t.Error("Expected myPrint() to be called, but it wasn't.")
    }
}
