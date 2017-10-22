package car

import (
    "fmt"
    "testing"
)

func getTestPrefix(testN int) string {
    return fmt.Sprintf("Test [%d]: ", testN)
}

func TestGetColor(t *testing.T) {
    tests := []struct {
        name string
        isFg bool
        expectedOutput string
        shell string
    }{
        { name: "red", isFg: false, expectedOutput: "%{\x1b[48;5;1m%}", shell: "zsh", },
        { name: "red", isFg: false, expectedOutput: "\001\x1b[48;5;1m\002", shell: "bash", },
        { name: "red", isFg: true, expectedOutput: "%{\x1b[38;5;1m%}", shell: "zsh", },
        { name: "red", isFg: true, expectedOutput: "\001\x1b[38;5;1m\002", shell: "bash", },
        { name: "222", isFg: false, expectedOutput: "%{\x1b[48;5;222m%}", shell: "zsh", },
        { name: "222", isFg: false, expectedOutput: "\001\x1b[48;5;222m\002", shell: "bash", },
        { name: "222", isFg: true, expectedOutput: "%{\x1b[38;5;222m%}", shell: "zsh", },
        { name: "222", isFg: true, expectedOutput: "\001\x1b[38;5;222m\002", shell: "bash", },
        { name: "0;0;255", isFg: false, expectedOutput: "%{\x1b[48;2;0;0;255m%}", shell: "zsh", },
        { name: "0;0;255", isFg: false, expectedOutput: "\001\x1b[48;2;0;0;255m\002", shell: "bash", },
        { name: "0;0;255", isFg: true, expectedOutput: "%{\x1b[38;2;0;0;255m%}", shell: "zsh", },
        { name: "0;0;255", isFg: true, expectedOutput: "\001\x1b[38;2;0;0;255m\002", shell: "bash", },
        { name: "default", isFg: false, expectedOutput: "%{\x1b[49m%}", shell: "zsh", },
        { name: "default", isFg: false, expectedOutput: "\001\x1b[49m\002", shell: "bash", },
        { name: "default", isFg: true, expectedOutput: "%{\x1b[39m%}", shell: "zsh", },
        { name: "default", isFg: true, expectedOutput: "\001\x1b[39m\002", shell: "bash", },
        { name: "_unknown", isFg: false, expectedOutput: "%{\x1b[49m%}", shell: "zsh", },
        { name: "_unknown", isFg: false, expectedOutput: "\001\x1b[49m\002", shell: "bash", },
        { name: "_unknown", isFg: true, expectedOutput: "%{\x1b[39m%}", shell: "zsh", },
        { name: "_unknown", isFg: true, expectedOutput: "\001\x1b[39m\002", shell: "bash", },
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
        { name: "bold", isEnd: false, expectedOutput: "%{\x1b[01m%}", shell: "zsh", },
        { name: "bold", isEnd: false, expectedOutput: "\001\x1b[01m\002", shell: "bash", },
        { name: "bold", isEnd: true, expectedOutput: "%{\x1b[21m%}", shell: "zsh", },
        { name: "bold", isEnd: true, expectedOutput: "\001\x1b[21m\002", shell: "bash", },
        { name: "underline", isEnd: false, expectedOutput: "%{\x1b[04m%}", shell: "zsh", },
        { name: "underline", isEnd: false, expectedOutput: "\001\x1b[04m\002", shell: "bash", },
        { name: "underline", isEnd: true, expectedOutput: "%{\x1b[24m%}", shell: "zsh", },
        { name: "underline", isEnd: true, expectedOutput: "\001\x1b[24m\002", shell: "bash", },
        { name: "blink", isEnd: false, expectedOutput: "%{\x1b[05m%}", shell: "zsh", },
        { name: "blink", isEnd: false, expectedOutput: "\001\x1b[05m\002", shell: "bash", },
        { name: "blink", isEnd: true, expectedOutput: "%{\x1b[25m%}", shell: "zsh", },
        { name: "blink", isEnd: true, expectedOutput: "\001\x1b[25m\002", shell: "bash", },
        { name: "none", isEnd: false, expectedOutput: "%{%}", shell: "zsh", },
        { name: "none", isEnd: false, expectedOutput: "\001\002", shell: "bash", },
        { name: "none", isEnd: true, expectedOutput: "%{%}", shell: "zsh", },
        { name: "none", isEnd: true, expectedOutput: "\001\002", shell: "bash", },
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
