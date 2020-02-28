package car

import (
    "fmt"
    "testing"
)

func getTestPrefix(testN int) string {
    return fmt.Sprintf("Test [%d]: ", testN)
}

func TestSetParamStr(t *testing.T) {
    tests := []struct {
        name string
        value string
        expectedValue string
    }{
        { name: "test", value: "aaa", expectedValue: "aaa", },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        car := Car{}

        car.SetParamStr(test.name, test.value)

        if test.value != car.Params[test.name] {
            t.Errorf("%sExpected '%s', found '%s'.", testPrefix, test.expectedValue, car.Params[test.name])
        }
    }
}

func TestGetModel(t *testing.T) {
    tests := []struct {
        key string
        model map[string]ModelElement
        expectedOutput map[string]ModelElement
    }{
        {
            key:            "root",
            model:          map[string]ModelElement{
                "root": { Bg: "255", Fg: "red", Fm: "bold", Text: "test", },
            },
            expectedOutput: map[string]ModelElement{
                "root": { Bg: "255", Fg: "red", Fm: "bold", Text: "test", },
            },
        },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        car := Car{
            Model: test.model,
        }

        output := car.GetModel()

        if output[test.key] != test.expectedOutput[test.key] {
            t.Errorf("%sExpected %x, found %x.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestGetDisplay(t *testing.T) {
    tests := []struct {
        display bool
        expectedOutput bool
    }{
        { display: true,  expectedOutput: true,  },
        { display: false, expectedOutput: false, },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        car := Car{
            Display: test.display,
        }

        output := car.GetDisplay()

        if output != test.expectedOutput {
            t.Errorf("%sExpected %t, found %t.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestGetWrap(t *testing.T) {
    tests := []struct {
        wrap bool
        expectedOutput bool
    }{
        { wrap: true,  expectedOutput: true,  },
        { wrap: false, expectedOutput: false, },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        car := Car{
            Wrap: test.wrap,
        }

        output := car.GetWrap()

        if output != test.expectedOutput {
            t.Errorf("%sExpected %t, found %t.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestFormat(t *testing.T) {
    tests := []struct {
        model map[string]ModelElement
        expectedOutput string
        display bool
        shell string
    }{
        {
            model: map[string]ModelElement{
                "root": { Bg: "222", Fg: "red", Fm: "bold", Text: "test", },
            },
            // TODO: This isn't what's actually generated on the command line.
            expectedOutput: "\x1b[48;5;222m\x1b[38;5;1m\x1b[1m\x1b[22mtest",
            display: true,
            shell: "plain",
        },
        {
            model: map[string]ModelElement{
                "root": { Bg: "222", Fg: "red", Fm: "bold", Text: "{{ Unknown }}", },
            },
            // TODO: This isn't what's actually generated on the command line.
            expectedOutput: "\x1b[48;5;222m\x1b[38;5;1m\x1b[1m\x1b[22m{{ Unknown }}",
            display: true,
            shell: "plain",
        },
        {
            model: map[string]ModelElement{
                "root": { Bg: "222", Fg: "red", Fm: "bold",    Text: "{{ User }}", },
                "User": { Bg: "222", Fg: "red", Fm: "default", Text: "text", },
            },
            // TODO: This isn't what's actually generated on the command line.
            expectedOutput: "\x1b[48;5;222m\x1b[38;5;1m\x1b[1m\x1b[22m\x1b[48;5;222m\x1b[38;5;1mtext\x1b[48;5;222m\x1b[38;5;1m\x1b[1m\x1b[22m",
            display: true,
            shell: "plain",
        },
        {
            model: map[string]ModelElement{
                "root": { Bg: "222", Fg: "red", Fm: "bold", Text: "text", },
            },
            expectedOutput: "",
            display: false,
            shell: "plain",
        },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        Shell = test.shell
        car := Car{
            Model: test.model,
            Display: test.display,
        }

        output := car.Format()

        if output != test.expectedOutput {
            t.Errorf("%sExpected %x, found %x.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestDecorateElement(t *testing.T) {
    tests := []struct {
        element string
        model map[string]ModelElement
        expectedOutput string
        display bool
        shell string
    }{
        {
            element: "root",
            model: map[string]ModelElement{
                "root": { Bg: "222", Fg: "red", Fm: "bold", Text: "test", },
            },
            expectedOutput: "\x1b[48;5;222m\x1b[38;5;1m\x1b[1m\x1b[22m",
            display: true,
            shell: "plain",
        },
        {
            element: "User",
            model: map[string]ModelElement{
                "User": { Bg: "222", Fg: "red", Fm: "bold", Text: "test", },
            },
            expectedOutput: "\x1b[48;5;222m\x1b[38;5;1m\x1b[1mtest\x1b[22m",
            display: true,
            shell: "plain",
        },
    }

    for i, test := range tests {
        testPrefix := getTestPrefix(i)
        Shell = test.shell
        car := Car{
            Model: test.model,
            Display: test.display,
        }

        output := car.DecorateElement(
            test.element,
            test.model[test.element].Bg,
            test.model[test.element].Fg,
            test.model[test.element].Fm,
            test.model[test.element].Text,
        )

        if output != test.expectedOutput {
            t.Errorf("%sExpected %x, found %x.", testPrefix, test.expectedOutput, output)
        }
    }
}

func TestGetColor(t *testing.T) {
    tests := []struct {
        name string
        isFg bool
        expectedOutput string
        shell string
    }{
        { name: "red",      isFg: false, expectedOutput: "%{\x1b[48;5;1m%}",           shell: "zsh",   },
        { name: "red",      isFg: false, expectedOutput: "\001\x1b[48;5;1m\002",       shell: "bash",  },
        { name: "red",      isFg: false, expectedOutput: "\x1b[48;5;1m",               shell: "plain", },
        { name: "red",      isFg: false, expectedOutput: "\\[\\e[48;5;1m\\]",          shell: "_bash", },
        { name: "red",      isFg: true,  expectedOutput: "%{\x1b[38;5;1m%}",           shell: "zsh",   },
        { name: "red",      isFg: true,  expectedOutput: "\001\x1b[38;5;1m\002",       shell: "bash",  },
        { name: "red",      isFg: true,  expectedOutput: "\x1b[38;5;1m",               shell: "plain", },
        { name: "red",      isFg: true,  expectedOutput: "\\[\\e[38;5;1m\\]",          shell: "_bash", },
        { name: "222",      isFg: false, expectedOutput: "%{\x1b[48;5;222m%}",         shell: "zsh",   },
        { name: "222",      isFg: false, expectedOutput: "\001\x1b[48;5;222m\002",     shell: "bash",  },
        { name: "222",      isFg: false, expectedOutput: "\x1b[48;5;222m",             shell: "plain", },
        { name: "222",      isFg: false, expectedOutput: "\\[\\e[48;5;222m\\]",        shell: "_bash", },
        { name: "222",      isFg: true,  expectedOutput: "%{\x1b[38;5;222m%}",         shell: "zsh",   },
        { name: "222",      isFg: true,  expectedOutput: "\001\x1b[38;5;222m\002",     shell: "bash",  },
        { name: "222",      isFg: true,  expectedOutput: "\x1b[38;5;222m",             shell: "plain", },
        { name: "222",      isFg: true,  expectedOutput: "\\[\\e[38;5;222m\\]",        shell: "_bash", },
        { name: "0;0;255",  isFg: false, expectedOutput: "%{\x1b[48;2;0;0;255m%}",     shell: "zsh",   },
        { name: "0;0;255",  isFg: false, expectedOutput: "\001\x1b[48;2;0;0;255m\002", shell: "bash",  },
        { name: "0;0;255",  isFg: false, expectedOutput: "\x1b[48;2;0;0;255m",         shell: "plain", },
        { name: "0;0;255",  isFg: false, expectedOutput: "\\[\\e[48;2;0;0;255m\\]",    shell: "_bash", },
        { name: "0;0;255",  isFg: true,  expectedOutput: "%{\x1b[38;2;0;0;255m%}",     shell: "zsh",   },
        { name: "0;0;255",  isFg: true,  expectedOutput: "\001\x1b[38;2;0;0;255m\002", shell: "bash",  },
        { name: "0;0;255",  isFg: true,  expectedOutput: "\x1b[38;2;0;0;255m",         shell: "plain", },
        { name: "0;0;255",  isFg: true,  expectedOutput: "\\[\\e[38;2;0;0;255m\\]",    shell: "_bash", },
        { name: "default",  isFg: false, expectedOutput: "%{\x1b[49m%}",               shell: "zsh",   },
        { name: "default",  isFg: false, expectedOutput: "\001\x1b[49m\002",           shell: "bash",  },
        { name: "default",  isFg: false, expectedOutput: "\x1b[49m",                   shell: "plain", },
        { name: "default",  isFg: false, expectedOutput: "\\[\\e[49m\\]",              shell: "_bash", },
        { name: "default",  isFg: true,  expectedOutput: "%{\x1b[39m%}",               shell: "zsh",   },
        { name: "default",  isFg: true,  expectedOutput: "\001\x1b[39m\002",           shell: "bash",  },
        { name: "default",  isFg: true,  expectedOutput: "\x1b[39m",                   shell: "plain", },
        { name: "default",  isFg: true,  expectedOutput: "\\[\\e[39m\\]",              shell: "_bash", },
        { name: "RESETALL", isFg: false, expectedOutput: "%{\x1b[0m%}",                shell: "zsh",   },
        { name: "RESETALL", isFg: false, expectedOutput: "\001\x1b[0m\002",            shell: "bash",  },
        { name: "RESETALL", isFg: false, expectedOutput: "\x1b[0m",                    shell: "plain", },
        { name: "RESETALL", isFg: false, expectedOutput: "\\[\\e[0m\\]",               shell: "_bash", },
        { name: "RESETALL", isFg: true,  expectedOutput: "%{\x1b[0m%}",                shell: "zsh",   },
        { name: "RESETALL", isFg: true,  expectedOutput: "\001\x1b[0m\002",            shell: "bash",  },
        { name: "RESETALL", isFg: true,  expectedOutput: "\x1b[0m",                    shell: "plain", },
        { name: "RESETALL", isFg: true,  expectedOutput: "\\[\\e[0m\\]",               shell: "_bash", },
        { name: "_unknown", isFg: false, expectedOutput: "%{\x1b[49m%}",               shell: "zsh",   },
        { name: "_unknown", isFg: false, expectedOutput: "\001\x1b[49m\002",           shell: "bash",  },
        { name: "_unknown", isFg: false, expectedOutput: "\x1b[49m",                   shell: "plain", },
        { name: "_unknown", isFg: false, expectedOutput: "\\[\\e[49m\\]",              shell: "_bash", },
        { name: "_unknown", isFg: true,  expectedOutput: "%{\x1b[39m%}",               shell: "zsh",   },
        { name: "_unknown", isFg: true,  expectedOutput: "\001\x1b[39m\002",           shell: "bash",  },
        { name: "_unknown", isFg: true,  expectedOutput: "\x1b[39m",                   shell: "plain", },
        { name: "_unknown", isFg: true,  expectedOutput: "\\[\\e[39m\\]",              shell: "_bash", },
    }

    car := Car{
        Model: make(map[string]ModelElement),
        Display: true,
    }

    for i, test := range tests {
        Shell = test.shell
        testPrefix := getTestPrefix(i)
        output := car.GetColor(test.name, test.isFg)

        if output != test.expectedOutput {
            t.Errorf("%sExpected (%s) %x, found %x.", testPrefix, test.shell, test.expectedOutput, output)
        }
    }
}

func TestGetFormat(t *testing.T) {
    tests := []struct {
        name string
        isEnd bool
        expectedOutput string
        shell string
    }{
        { name: "normal",    isEnd: false, expectedOutput: "%{\x1b[0m%}",      shell: "zsh",   },
        { name: "normal",    isEnd: false, expectedOutput: "\001\x1b[0m\002",  shell: "bash",  },
        { name: "normal",    isEnd: false, expectedOutput: "\x1b[0m",          shell: "plain", },
        { name: "normal",    isEnd: false, expectedOutput: "\\[\\e[0m\\]",     shell: "_bash", },
        { name: "normal",    isEnd: true,  expectedOutput: "%{\x1b[0m%}",      shell: "zsh",   },
        { name: "normal",    isEnd: true,  expectedOutput: "\001\x1b[0m\002",  shell: "bash",  },
        { name: "normal",    isEnd: true,  expectedOutput: "\x1b[0m",          shell: "plain", },
        { name: "normal",    isEnd: true,  expectedOutput: "\\[\\e[0m\\]",     shell: "_bash", },
        { name: "dim",       isEnd: false, expectedOutput: "%{\x1b[2m%}",      shell: "zsh",   },
        { name: "dim",       isEnd: false, expectedOutput: "\001\x1b[2m\002",  shell: "bash",  },
        { name: "dim",       isEnd: false, expectedOutput: "\x1b[2m",          shell: "plain", },
        { name: "dim",       isEnd: false, expectedOutput: "\\[\\e[2m\\]",     shell: "_bash", },
        { name: "dim",       isEnd: true,  expectedOutput: "%{\x1b[22m%}",     shell: "zsh",   },
        { name: "dim",       isEnd: true,  expectedOutput: "\001\x1b[22m\002", shell: "bash",  },
        { name: "dim",       isEnd: true,  expectedOutput: "\x1b[22m",         shell: "plain", },
        { name: "dim",       isEnd: true,  expectedOutput: "\\[\\e[22m\\]",    shell: "_bash", },
        { name: "bold",      isEnd: false, expectedOutput: "%{\x1b[1m%}",      shell: "zsh",   },
        { name: "bold",      isEnd: false, expectedOutput: "\001\x1b[1m\002",  shell: "bash",  },
        { name: "bold",      isEnd: false, expectedOutput: "\x1b[1m",          shell: "plain", },
        { name: "bold",      isEnd: false, expectedOutput: "\\[\\e[1m\\]",     shell: "_bash", },
        { name: "bold",      isEnd: true,  expectedOutput: "%{\x1b[22m%}",     shell: "zsh",   },
        { name: "bold",      isEnd: true,  expectedOutput: "\001\x1b[22m\002", shell: "bash",  },
        { name: "bold",      isEnd: true,  expectedOutput: "\x1b[22m",         shell: "plain", },
        { name: "bold",      isEnd: true,  expectedOutput: "\\[\\e[22m\\]",    shell: "_bash", },
        { name: "underline", isEnd: false, expectedOutput: "%{\x1b[4m%}",      shell: "zsh",   },
        { name: "underline", isEnd: false, expectedOutput: "\001\x1b[4m\002",  shell: "bash",  },
        { name: "underline", isEnd: false, expectedOutput: "\x1b[4m",          shell: "plain", },
        { name: "underline", isEnd: false, expectedOutput: "\\[\\e[4m\\]",     shell: "_bash", },
        { name: "underline", isEnd: true,  expectedOutput: "%{\x1b[24m%}",     shell: "zsh",   },
        { name: "underline", isEnd: true,  expectedOutput: "\001\x1b[24m\002", shell: "bash",  },
        { name: "underline", isEnd: true,  expectedOutput: "\x1b[24m",         shell: "plain", },
        { name: "underline", isEnd: true,  expectedOutput: "\\[\\e[24m\\]",    shell: "_bash", },
        { name: "blink",     isEnd: false, expectedOutput: "%{\x1b[5m%}",      shell: "zsh",   },
        { name: "blink",     isEnd: false, expectedOutput: "\001\x1b[5m\002",  shell: "bash",  },
        { name: "blink",     isEnd: false, expectedOutput: "\x1b[5m",          shell: "plain", },
        { name: "blink",     isEnd: false, expectedOutput: "\\[\\e[5m\\]",     shell: "_bash", },
        { name: "blink",     isEnd: true,  expectedOutput: "%{\x1b[25m%}",     shell: "zsh",   },
        { name: "blink",     isEnd: true,  expectedOutput: "\001\x1b[25m\002", shell: "bash",  },
        { name: "blink",     isEnd: true,  expectedOutput: "\x1b[25m",         shell: "plain", },
        { name: "blink",     isEnd: true,  expectedOutput: "\\[\\e[25m\\]",    shell: "_bash", },
        { name: "invert",    isEnd: false, expectedOutput: "%{\x1b[7m%}",      shell: "zsh",   },
        { name: "invert",    isEnd: false, expectedOutput: "\001\x1b[7m\002",  shell: "bash",  },
        { name: "invert",    isEnd: false, expectedOutput: "\x1b[7m",          shell: "plain", },
        { name: "invert",    isEnd: false, expectedOutput: "\\[\\e[7m\\]",     shell: "_bash", },
        { name: "invert",    isEnd: true,  expectedOutput: "%{\x1b[27m%}",     shell: "zsh",   },
        { name: "invert",    isEnd: true,  expectedOutput: "\001\x1b[27m\002", shell: "bash",  },
        { name: "invert",    isEnd: true,  expectedOutput: "\x1b[27m",         shell: "plain", },
        { name: "invert",    isEnd: true,  expectedOutput: "\\[\\e[27m\\]",    shell: "_bash", },
        { name: "hide",      isEnd: false, expectedOutput: "%{\x1b[8m%}",      shell: "zsh",   },
        { name: "hide",      isEnd: false, expectedOutput: "\001\x1b[8m\002",  shell: "bash",  },
        { name: "hide",      isEnd: false, expectedOutput: "\x1b[8m",          shell: "plain", },
        { name: "hide",      isEnd: false, expectedOutput: "\\[\\e[8m\\]",     shell: "_bash", },
        { name: "hide",      isEnd: true,  expectedOutput: "%{\x1b[28m%}",     shell: "zsh",   },
        { name: "hide",      isEnd: true,  expectedOutput: "\001\x1b[28m\002", shell: "bash",  },
        { name: "hide",      isEnd: true,  expectedOutput: "\x1b[28m",         shell: "plain", },
        { name: "hide",      isEnd: true,  expectedOutput: "\\[\\e[28m\\]",    shell: "_bash", },
        { name: "strikeout", isEnd: false, expectedOutput: "%{\x1b[9m%}",      shell: "zsh",   },
        { name: "strikeout", isEnd: false, expectedOutput: "\001\x1b[9m\002",  shell: "bash",  },
        { name: "strikeout", isEnd: false, expectedOutput: "\x1b[9m",          shell: "plain", },
        { name: "strikeout", isEnd: false, expectedOutput: "\\[\\e[9m\\]",     shell: "_bash", },
        { name: "strikeout", isEnd: true,  expectedOutput: "%{\x1b[29m%}",     shell: "zsh",   },
        { name: "strikeout", isEnd: true,  expectedOutput: "\001\x1b[29m\002", shell: "bash",  },
        { name: "strikeout", isEnd: true,  expectedOutput: "\x1b[29m",         shell: "plain", },
        { name: "strikeout", isEnd: true,  expectedOutput: "\\[\\e[29m\\]",    shell: "_bash", },
        { name: "none",      isEnd: false, expectedOutput: "",                 shell: "zsh",   },
        { name: "none",      isEnd: false, expectedOutput: "",                 shell: "bash",  },
        { name: "none",      isEnd: false, expectedOutput: "",                 shell: "plain", },
        { name: "none",      isEnd: false, expectedOutput: "",                 shell: "_bash", },
        { name: "none",      isEnd: true,  expectedOutput: "",                 shell: "zsh",   },
        { name: "none",      isEnd: true,  expectedOutput: "",                 shell: "bash",  },
        { name: "none",      isEnd: true,  expectedOutput: "",                 shell: "plain", },
        { name: "none",      isEnd: true,  expectedOutput: "",                 shell: "_bash", },
    }

    car := Car{
        Model: make(map[string]ModelElement),
        Display: true,
    }

    for i, test := range tests {
        Shell = test.shell
        testPrefix := getTestPrefix(i)
        output := car.GetFormat(test.name, test.isEnd)

        if output != test.expectedOutput {
            t.Errorf("%sExpected (%s) %x, found %x.", testPrefix, test.shell, test.expectedOutput, output)
        }
    }
}
