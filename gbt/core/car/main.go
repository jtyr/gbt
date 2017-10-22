package car

import (
    "fmt"
    "path"
    "regexp"
    "strings"

    "github.com/jtyr/gbt/gbt/core/utils"
)

// ModelElement is an element of which the car model is constructed from.
type ModelElement struct {
    Bg string
    Fg string
    Fm string
    Text string
}

// Car is a type defining the model of the car.
type Car struct {
    Model map[string]ModelElement
    Display bool
    Sep string
    Wrap bool
}

// Shell type.
var Shell string = utils.GetEnv("GBT_SHELL", path.Base(utils.GetEnv("SHELL", "zsh")))

// List of named colors and their codes.
var colors = map[string]string {
    "black":          "0",
    "red":            "1",
    "green":          "2",
    "yellow":         "3",
    "blue":           "4",
    "magenta":        "5",
    "cyan":           "6",
    "light_gray":     "7",
    "dark_gray":      "8",
    "light_red":      "9",
    "light_green":   "10",
    "light_yellow":  "11",
    "light_blue":    "12",
    "light_magenta": "13",
    "light_cyan":    "14",
    "white":         "15",
}

// GetModel returns the Model value.
func (c *Car) GetModel() map[string]ModelElement {
    return c.Model
}

// GetDisplay returns the Display value.
func (c *Car) GetDisplay() bool {
    return c.Display
}

// GetSep returns the Sep value.
func (c *Car) GetSep() string {
    return c.Sep
}

// GetWrap returns the Wrap value.
func (c *Car) GetWrap() bool {
    return c.Wrap
}

var reTemplating = regexp.MustCompile(`{{\s*(\w+)\s*}}`)

// Format initiates replacement of all templating elements.
func (c *Car) Format() string {
    if ! c.Display {
        return ""
    }

    text := fmt.Sprintf("%s%s", c.DecorateElement("root", "", "", "", ""), c.Model["root"].Text)

    for range make([]int, 10) {
        match := reTemplating.MatchString(text)

        if match {
            text = reTemplating.ReplaceAllStringFunc(text, c.replaceElement)
        } else {
            break
        }
    }

    return text
}

// Replaces the specific templating element.
func (c *Car) replaceElement(format string) string {
    match := reTemplating.FindStringSubmatch(format)[1]

    if _, ok := c.Model[match]; ! ok {
        return format
    }

    return fmt.Sprintf(
        "%s%s",
        c.DecorateElement(match, "", "", "", ""),
        c.DecorateElement("root", "", "", "", ""))
}

// DecorateElement decorates the element text with its colors and formatting.
func (c *Car) DecorateElement(element, bg, fg, fm, text string) string {
    fmEnd := ""

    if element != "" {
        e := c.Model[element]

        if element != "root" {
            text = e.Text
        } else {
            text = ""
        }

        bg = c.GetColor(e.Bg, false)
        fg = c.GetColor(e.Fg, true)
        fm = c.GetFormat(e.Fm, false)

        if fm != c.GetFormat("empty", false) {
            fmEnd = c.GetFormat(e.Fm, true)
        } else {
            fm = ""
        }
    }

    return fmt.Sprintf("%s%s%s%s%s", bg, fg, fm, text, fmEnd)
}

// Patterns to parse the color codes
var reColorNumber = regexp.MustCompile(`^\d{1,3}$`)
var reRgbTriplet = regexp.MustCompile(`^\d{1,3};\d{1,3};\d{1,3}$`)

// GetColor returns color sequence based on the color name or code.
func (c *Car) GetColor(name string, isFg bool) (ret string) {
    kind := 4
    seq := ""

    if isFg {
        kind = 3
    }

    if name == "default" {
        // Default
        seq = fmt.Sprintf("\x1b[%d9m", kind)
    } else {
        if val, ok := colors[name]; ok {
            // Named color
            seq = fmt.Sprintf("\x1b[%d8;5;%sm", kind, val)
        } else if match := reColorNumber.MatchString(name); match {
            // Color number
            seq = fmt.Sprintf("\x1b[%d8;5;%sm", kind, name)
        } else if match := reRgbTriplet.MatchString(name); match {
            // RGB color
            seq = fmt.Sprintf("\x1b[%d8;2;%sm", kind, name)
        } else {
            // If anything else, use default
            seq = fmt.Sprintf("\x1b[%d9m", kind)
        }
    }

    ret = DecorateShell(seq)

    return
}

// GetFormat returns formatting sequence based on the format name.
func (c *Car) GetFormat(name string, end bool) (ret string) {
    seq := ""
    kind := 0

    if end {
        kind = 2
    }

    if strings.Contains(name, "bold") {
        seq += fmt.Sprintf("\x1b[%d1m", kind)
    }

    if strings.Contains(name, "underline") {
        seq += fmt.Sprintf("\x1b[%d4m", kind)
    }

    if strings.Contains(name, "blink") {
        seq += fmt.Sprintf("\x1b[%d5m", kind)
    }

    ret = DecorateShell(seq)

    return
}

// DecorateShell decorates the string with shell-specific closure.
func DecorateShell(seq string) (ret string) {
    if Shell == "zsh" {
        ret = fmt.Sprintf("%%{%s%%}", seq)
    } else {
        ret = fmt.Sprintf("\001%s\002", seq)
    }

    return
}
