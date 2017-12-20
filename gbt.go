package main

import (
    "flag"
    "fmt"
    "os"
    "regexp"

    customCar    "github.com/jtyr/gbt/gbt/cars/custom"
    dirCar       "github.com/jtyr/gbt/gbt/cars/dir"
    exectimeCar  "github.com/jtyr/gbt/gbt/cars/exectime"
    gitCar       "github.com/jtyr/gbt/gbt/cars/git"
    hostnameCar  "github.com/jtyr/gbt/gbt/cars/hostname"
    osCar        "github.com/jtyr/gbt/gbt/cars/os"
    pyvirtenvCar "github.com/jtyr/gbt/gbt/cars/pyvirtenv"
    signCar      "github.com/jtyr/gbt/gbt/cars/sign"
    statusCar    "github.com/jtyr/gbt/gbt/cars/status"
    timeCar      "github.com/jtyr/gbt/gbt/cars/time"

    "github.com/jtyr/gbt/gbt/core/car"
    "github.com/jtyr/gbt/gbt/core/utils"
)

// Cars interface for methods from the core.car package.
type Cars interface {
    Init()
    Format() string
    SetParamStr(string, string)
    GetColor(string, bool) string
    DecorateElement(element, bg, fg, fm, text string) string
    GetModel() map[string]car.ModelElement
    GetDisplay() bool
    GetSep() string
    GetWrap() bool
}

const version = "1.1.5"

func printCars(cars []Cars, right bool) {
    prevBg := "\000"
    prevDisplay := true
    fakeCar := car.Car{}
    defaultSeparator := utils.GetEnv("GBT_SEPARATOR", "")

    if right {
        defaultSeparator = utils.GetEnv("GBT_RSEPARATOR", "")
    }

    if ! right && utils.GetEnv("GBT_BEGINNING_TEXT", "") != "" {
        fmt.Print(
            fakeCar.DecorateElement(
                "",
                fakeCar.GetColor(utils.GetEnv("GBT_BEGINNING_BG", "default"), false),
                fakeCar.GetColor(utils.GetEnv("GBT_BEGINNING_FG", "default"), true),
                fakeCar.GetColor(utils.GetEnv("GBT_BEGINNING_FM", "none"), false),
                utils.GetEnv("GBT_BEGINNING_TEXT", "")))
    }

    for _, c := range cars {
        c.Init()

        cModel := c.GetModel()
        cDisplay := c.GetDisplay()
        cSep := c.GetSep()
        cWrap := c.GetWrap()

        separator := defaultSeparator

        if cSep != "\000" {
            separator = cSep
        }

        if cDisplay {
            fmt.Print(fakeCar.GetColor("RESETALL", false))

            if prevBg != "\000" && prevDisplay {
                bg := c.GetColor(cModel["root"].Bg, false)
                fg := c.GetColor(cModel["root"].Bg, true)

                if cWrap {
                    bg = c.GetColor("default", false)
                    fg = c.GetColor("default", true)
                }

                if right {
                    fmt.Print(
                        c.DecorateElement(
                            "",
                            c.GetColor(prevBg, false),
                            fg,
                            "",
                            separator))
                } else {
                    fmt.Print(
                        c.DecorateElement(
                            "",
                            bg,
                            c.GetColor(prevBg, true),
                            "",
                            separator))
                }

                if cWrap {
                    fmt.Print("\n")
                }
            }

            prevBg = cModel["root"].Bg
            prevDisplay = cDisplay

            fmt.Print(c.Format())
        }
    }

    fmt.Print(fakeCar.GetColor("RESETALL", false))
}

// For the test
var run = flag.PrintDefaults

func main() {
    var argsHelp, argsVersion, argsRight bool

    flag.BoolVar(&argsHelp, "help", false, "show this help message and exit")
    flag.BoolVar(&argsVersion, "version", false, "show version and exit")
    flag.BoolVar(&argsRight, "right", false, "compose right hand site prompt")
    flag.Parse()

    if argsHelp {
        fmt.Printf("Usage of %s:\n", os.Args[0])
        run()
        os.Exit(0)
    }
    if argsVersion {
        fmt.Printf("GBT v%s\n", version)
        os.Exit(0)
    }

    carsStr := utils.GetEnv("GBT_CARS", "Status, Os, Hostname, Dir, Git, Sign")

    if argsRight {
        carsStr = utils.GetEnv("GBT_RCARS", "Time")
    }

    reCarSplit := regexp.MustCompile(`\s*,\s*`)
    carsNames := reCarSplit.Split(carsStr, -1)
    carsFactory := map[string]Cars{
        "Custom":    &customCar.Car{},
        "Dir":       &dirCar.Car{},
        "ExecTime":  &exectimeCar.Car{},
        "Git":       &gitCar.Car{},
        "Hostname":  &hostnameCar.Car{},
        "Os":        &osCar.Car{},
        "PyVirtEnv": &pyvirtenvCar.Car{},
        "Sign":      &signCar.Car{},
        "Status":    &statusCar.Car{},
        "Time":      &timeCar.Car{},
    }
    cars := []Cars{}

    for _, cn := range carsNames {
        custom := "\000"

        if len(cn) >= 6 && cn[:6] == "Custom" {
            custom = cn[6:]
            cn = cn[:6]
        }

        if val, ok := carsFactory[cn]; ok {
            if cn == "Status" && len(flag.Args()) > 0 {
                val.SetParamStr("args", flag.Args()[0])
            } else if custom != "\000" {
                val = &customCar.Car{}
                val.SetParamStr("name", custom)
            }

            cars = append(cars, val)
        }
    }

    printCars(cars, argsRight)
}
