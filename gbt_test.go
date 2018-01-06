package main

import (
    "os"
    "testing"

    signCar "github.com/jtyr/gbt/gbt/cars/sign"
    timeCar "github.com/jtyr/gbt/gbt/cars/time"
)

func TestMain(t *testing.T) {
    var ran bool

    os.Setenv("GBT_CARS", "Status, Os, Time, Custom, Hostname, Dir, PyVirtEnv, Git, Sign")
    os.Setenv("GBT_CAR_SIGN_WRAP", "1")
    os.Setenv("GBT_BEGINNING_TEXT", "test")
    os.Setenv("GBT_CAR_CUSTOM_SEP", ">")

    os.Args = append(os.Args, "-help")
    os.Args = append(os.Args, "-version")
    os.Args = append(os.Args, "0")

    myPrint("")

    exit = func(i int) {}
    printDefaults = func() {}
    myPrint = func(s string) {
        ran = true
    }

    main()

    cars := []Cars{}
    cars = append(cars, &signCar.Car{})
    cars = append(cars, &timeCar.Car{})
    printCars(cars, true)

    argsHelp = true
    argsVersion = true
    argsRight = true

    main()

    if ! ran {
        t.Error("Expected myPrint() to be called, but it wasn't.")
    }
}
