package car

import (
	"fmt"
	"path"
	"regexp"
	"strings"

	"github.com/jtyr/gbt/pkg/core/utils"
)

// ModelElement is an element of which the car model is constructed from.
type ModelElement struct {
	Bg   string
	Fg   string
	Fm   string
	Text string
}

// Car is a type defining the model of the car.
type Car struct {
	Model   map[string]ModelElement
	Display bool
	Sep     string
	Wrap    bool
	Params  map[string]interface{}
}

// Shell type.
var Shell = utils.GetEnv("GBT_SHELL", path.Base(utils.GetEnv("SHELL", "bash")))

// Higher colors convertor flag.
var forceHigherColors = utils.GetEnvBool("GBT_FORCE_HIGHER_COLORS", true)

// List of named Standard and High-intensity colors and their ANSI codes.
var colors = map[string]string{
	"black":         "0",
	"red":           "1",
	"green":         "2",
	"yellow":        "3",
	"blue":          "4",
	"magenta":       "5",
	"cyan":          "6",
	"light_gray":    "7",
	"dark_gray":     "8",
	"light_red":     "9",
	"light_green":   "10",
	"light_yellow":  "11",
	"light_blue":    "12",
	"light_magenta": "13",
	"light_cyan":    "14",
	"white":         "15",
}

// List of named Standard and High-intensity colors represented as 216 and Grayscale colors.
var higherColors = map[string]string{
	"black":         "16",
	"red":           "124",
	"green":         "34",
	"yellow":        "100",
	"blue":          "19",
	"magenta":       "90",
	"cyan":          "30",
	"light_gray":    "248",
	"dark_gray":     "240",
	"light_red":     "196",
	"light_green":   "46",
	"light_yellow":  "226",
	"light_blue":    "63",
	"light_magenta": "201",
	"light_cyan":    "51",
	"white":         "231",
}

// SetParamStr sets string value to a parameter.
func (c *Car) SetParamStr(name, value string) {
	if c.Params == nil {
		c.Params = make(map[string]interface{})
	}

	c.Params[name] = value
}

// GetModel returns the Model value.
func (c *Car) GetModel() map[string]ModelElement {
	return c.Model
}

// GetDisplay returns the Display value.
func (c *Car) GetDisplay() bool {
	return c.Display
}

// GetWrap returns the Wrap value.
func (c *Car) GetWrap() bool {
	return c.Wrap
}

var reTemplating = regexp.MustCompile(`{{\s*(\w+)\s*}}`)

// Format initiates replacement of all templating elements.
func (c *Car) Format() string {
	if !c.Display {
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

	if _, ok := c.Model[match]; !ok {
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

// Patterns to parse the color codes.
var reColorNumber = regexp.MustCompile(`^\d{1,3}$`)
var reRgbTriplet = regexp.MustCompile(`^\d{1,3};\d{1,3};\d{1,3}$`)

// GetColor returns color sequence based on the color name or code.
func (c *Car) GetColor(name string, isFg bool) (ret string) {
	kind := "4"
	seq := ""
	esc := "\x1b"

	if isFg {
		kind = "3"
	}

	if Shell == "_bash" {
		esc = "\\e"
	}

	if name == "RESETALL" {
		seq = fmt.Sprintf("%s[0m", esc)
	} else if name == "default" {
		// Default
		seq = fmt.Sprintf("%s[%s9m", esc, kind)
	} else {
		if val, ok := colors[name]; ok {
			if forceHigherColors {
				val = higherColors[name]
			}

			// Named color
			seq = fmt.Sprintf("%s[%s8;5;%sm", esc, kind, val)
		} else if match := reColorNumber.MatchString(name); match {
			val := name

			if forceHigherColors {
				for k, v := range colors {
					if v == name {
						val = higherColors[k]
					}
				}
			}

			// Color number
			seq = fmt.Sprintf("%s[%s8;5;%sm", esc, kind, val)
		} else if match := reRgbTriplet.MatchString(name); match {
			// RGB color
			seq = fmt.Sprintf("%s[%s8;2;%sm", esc, kind, name)
		} else {
			// If anything else, use default
			seq = fmt.Sprintf("%s[%s9m", esc, kind)
		}
	}

	ret = decorateShell(seq)

	return
}

// GetFormat returns formatting sequence based on the format name.
func (c *Car) GetFormat(name string, end bool) (ret string) {
	seq := ""
	kind := ""
	esc := "\x1b"

	if end {
		kind = "2"
	}

	if Shell == "_bash" {
		esc = "\\e"
	}

	if strings.Contains(name, "normal") {
		seq += fmt.Sprintf("%s[0m", esc)
	}

	if strings.Contains(name, "bold") {
		if end {
			seq += fmt.Sprintf("%s[22m", esc)
		} else {
			seq += fmt.Sprintf("%s[%s1m", esc, kind)
		}
	}

	if strings.Contains(name, "dim") {
		seq += fmt.Sprintf("%s[%s2m", esc, kind)
	}

	if strings.Contains(name, "underline") {
		seq += fmt.Sprintf("%s[%s4m", esc, kind)
	}

	if strings.Contains(name, "blink") {
		seq += fmt.Sprintf("%s[%s5m", esc, kind)
	}

	if strings.Contains(name, "invert") {
		seq += fmt.Sprintf("%s[%s7m", esc, kind)
	}

	if strings.Contains(name, "hide") {
		seq += fmt.Sprintf("%s[%s8m", esc, kind)
	}

	if strings.Contains(name, "strikeout") {
		seq += fmt.Sprintf("%s[%s9m", esc, kind)
	}

	ret = decorateShell(seq)

	return
}

// decorateShell decorates the string with shell-specific closure.
func decorateShell(seq string) (ret string) {
	if len(seq) == 0 {
		ret = ""
	} else if Shell == "zsh" {
		ret = fmt.Sprintf("%%{%s%%}", seq)
	} else if Shell == "_bash" {
		ret = fmt.Sprintf("\\[%s\\]", seq)
	} else if Shell == "plain" {
		ret = fmt.Sprintf("%s", seq)
	} else {
		// bash
		ret = fmt.Sprintf("\001%s\002", seq)
	}

	return
}
