package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"

	awsCar "github.com/jtyr/gbt/pkg/cars/aws"
	azureCar "github.com/jtyr/gbt/pkg/cars/azure"
	customCar "github.com/jtyr/gbt/pkg/cars/custom"
	dirCar "github.com/jtyr/gbt/pkg/cars/dir"
	exectimeCar "github.com/jtyr/gbt/pkg/cars/exectime"
	gcpCar "github.com/jtyr/gbt/pkg/cars/gcp"
	gitCar "github.com/jtyr/gbt/pkg/cars/git"
	hostnameCar "github.com/jtyr/gbt/pkg/cars/hostname"
	kubectlCar "github.com/jtyr/gbt/pkg/cars/kubectl"
	osCar "github.com/jtyr/gbt/pkg/cars/os"
	pyvirtenvCar "github.com/jtyr/gbt/pkg/cars/pyvirtenv"
	signCar "github.com/jtyr/gbt/pkg/cars/sign"
	statusCar "github.com/jtyr/gbt/pkg/cars/status"
	timeCar "github.com/jtyr/gbt/pkg/cars/time"

	"github.com/jtyr/gbt/pkg/core/car"
	"github.com/jtyr/gbt/pkg/core/utils"
)

// Cars interface for methods from the core.car package.
type Cars interface {
	Init()
	Format() string
	SetParamStr(string, string)
	GetColor(string, bool) string
	GetFormat(string, bool) string
	DecorateElement(element, bg, fg, fm, text string) string
	GetModel() map[string]car.ModelElement
	GetDisplay() bool
	GetWrap() bool
}

var (
	build   string
	version string
)

func printCars(cars []Cars, right bool) {
	prevBg := "\000"
	prevDisplay := true
	fakeCar := car.Car{}
	defaultSeparator := utils.GetEnv("GBT_SEPARATOR", "\ue0b0")

	if right {
		defaultSeparator = utils.GetEnv("GBT_RSEPARATOR", "\ue0b2")
	}

	if !right && utils.GetEnv("GBT_BEGINNING_TEXT", "") != "" {
		myPrint(
			fakeCar.DecorateElement(
				"",
				fakeCar.GetColor(utils.GetEnv("GBT_BEGINNING_BG", "default"), false),
				fakeCar.GetColor(utils.GetEnv("GBT_BEGINNING_FG", "default"), true),
				fakeCar.GetFormat(utils.GetEnv("GBT_BEGINNING_FM", "none"), false),
				utils.GetEnv("GBT_BEGINNING_TEXT", "")))
	}

	for _, c := range cars {
		c.Init()

		cModel := c.GetModel()
		cDisplay := c.GetDisplay()

		if cDisplay {
			cWrap := c.GetWrap()
			cSep := cModel["Sep"]

			separator := defaultSeparator

			if cSep.Text != "\000" {
				separator = cSep.Text
			}

			myPrint(fakeCar.GetColor("RESETALL", false))

			if prevBg != "\000" && prevDisplay {
				bg := c.GetColor(cModel["root"].Bg, false)
				fg := c.GetColor(cModel["root"].Bg, true)
				fm := ""

				if cWrap {
					bg = c.GetColor("default", false)
					fg = c.GetColor("default", true)
				}

				if right {
					bg = c.GetColor(prevBg, false)
				} else {
					fg = c.GetColor(prevBg, true)
				}

				if cSep.Bg != "\000" {
					bg = c.GetColor(cSep.Bg, false)
				}

				if cSep.Fg != "\000" {
					fg = c.GetColor(cSep.Fg, true)
				}

				if cSep.Fm != "\000" {
					fm = c.GetFormat(cSep.Fm, false)
				}

				myPrint(
					c.DecorateElement(
						"",
						bg,
						fg,
						fm,
						separator))

				if cWrap {
					myPrint("\n")
				}
			}

			prevBg = cModel["root"].Bg
			prevDisplay = cDisplay

			myPrint(c.Format())
		}
	}

	myPrint(fakeCar.GetColor("RESETALL", false))
}

// For the test
var myPrint = func(s string) {
	fmt.Print(s)
}
var printDefaults = flag.PrintDefaults
var exit = os.Exit
var argsHelp, argsVersion, argsRight bool

func main() {
	if len(flag.Args()) == 0 {
		flag.BoolVar(&argsHelp, "help", false, "show this help message and exit")
		flag.BoolVar(&argsVersion, "version", false, "show version and exit")
		flag.BoolVar(&argsRight, "right", false, "compose right hand site prompt")
		flag.Parse()
	}

	if argsHelp {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		printDefaults()
		exit(0)
	}

	if argsVersion {
		if version == "" || build == "" {
			myPrint("GBT version wasn't provided at the build time.\n")
		} else {
			myPrint(fmt.Sprintf("GBT version %s, build %s\n", version, build))
		}
		exit(0)
	}

	carsStr := strings.ToLower(utils.GetEnv("GBT_CARS", "status, os, hostname, dir, git, sign"))

	if argsRight {
		carsStr = strings.ToLower(utils.GetEnv("GBT_RCARS", "time"))
	}

	reCarSplit := regexp.MustCompile(`\s*,\s*`)
	carsNames := reCarSplit.Split(carsStr, -1)
	carsFactory := map[string]Cars{
		"aws":       &awsCar.Car{},
		"azure":     &azureCar.Car{},
		"custom":    &customCar.Car{},
		"dir":       &dirCar.Car{},
		"exectime":  &exectimeCar.Car{},
		"gcp":       &gcpCar.Car{},
		"git":       &gitCar.Car{},
		"hostname":  &hostnameCar.Car{},
		"kubectl":   &kubectlCar.Car{},
		"os":        &osCar.Car{},
		"pyvirtenv": &pyvirtenvCar.Car{},
		"sign":      &signCar.Car{},
		"status":    &statusCar.Car{},
		"time":      &timeCar.Car{},
	}
	cars := []Cars{}

	for _, cn := range carsNames {
		cn = strings.TrimSpace(cn)
		custom := "\000"

		if len(cn) >= 6 && cn[:6] == "custom" {
			custom = cn[6:]
			cn = cn[:6]
		}

		if val, ok := carsFactory[cn]; ok {
			if cn == "status" && len(flag.Args()) > 0 {
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
