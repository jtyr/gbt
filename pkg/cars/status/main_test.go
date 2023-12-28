package status

import (
	"testing"

	ct "github.com/jtyr/gbt/pkg/core/testing"
)

func TestInit(t *testing.T) {
	ct.ResetEnv()

	tests := []struct {
		args            string
		dropArgs        bool
		expectedSignal  string
		expectedDisplay bool
	}{
		{
			args:            "",
			dropArgs:        true,
			expectedSignal:  "?",
			expectedDisplay: false,
		},
		{
			args:            "-1",
			expectedSignal:  "FATAL",
			expectedDisplay: true,
		},
		{
			args:            "0",
			expectedSignal:  "OK",
			expectedDisplay: false,
		},
		{
			args:            "1",
			expectedSignal:  "FAIL",
			expectedDisplay: true,
		},
		{
			args:            "2",
			expectedSignal:  "BLTINMUSE",
			expectedDisplay: true,
		},
		{
			args:            "6",
			expectedSignal:  "UNKADDR",
			expectedDisplay: true,
		},
		{
			args:            "126",
			expectedSignal:  "NOEXEC",
			expectedDisplay: true,
		},
		{
			args:            "127",
			expectedSignal:  "NOTFOUND",
			expectedDisplay: true,
		},
		{
			args:            "129",
			expectedSignal:  "SIGHUP",
			expectedDisplay: true,
		},
		{
			args:            "130",
			expectedSignal:  "SIGINT",
			expectedDisplay: true,
		},
		{
			args:            "131",
			expectedSignal:  "SIGQUIT",
			expectedDisplay: true,
		},
		{
			args:            "132",
			expectedSignal:  "SIGILL",
			expectedDisplay: true,
		},
		{
			args:            "133",
			expectedSignal:  "SIGTRAP",
			expectedDisplay: true,
		},
		{
			args:            "134",
			expectedSignal:  "SIGABRT",
			expectedDisplay: true,
		},
		{
			args:            "135",
			expectedSignal:  "SIGBUS",
			expectedDisplay: true,
		},
		{
			args:            "136",
			expectedSignal:  "SIGFPE",
			expectedDisplay: true,
		},
		{
			args:            "137",
			expectedSignal:  "SIGKILL",
			expectedDisplay: true,
		},
		{
			args:            "138",
			expectedSignal:  "SIGUSR1",
			expectedDisplay: true,
		},
		{
			args:            "139",
			expectedSignal:  "SIGSEGV",
			expectedDisplay: true,
		},
		{
			args:            "140",
			expectedSignal:  "SIGUSR2",
			expectedDisplay: true,
		},
		{
			args:            "141",
			expectedSignal:  "SIGPIPE",
			expectedDisplay: true,
		},
		{
			args:            "142",
			expectedSignal:  "SIGALRM",
			expectedDisplay: true,
		},
		{
			args:            "143",
			expectedSignal:  "SIGTERM",
			expectedDisplay: true,
		},
		{
			args:            "145",
			expectedSignal:  "SIGCHLD",
			expectedDisplay: true,
		},
		{
			args:            "146",
			expectedSignal:  "SIGCONT",
			expectedDisplay: true,
		},
		{
			args:            "147",
			expectedSignal:  "SIGSTOP",
			expectedDisplay: true,
		},
		{
			args:            "148",
			expectedSignal:  "SIGTSTP",
			expectedDisplay: true,
		},
		{
			args:            "149",
			expectedSignal:  "SIGTTIN",
			expectedDisplay: true,
		},
		{
			args:            "150",
			expectedSignal:  "SIGTTOU",
			expectedDisplay: true,
		},
		{
			args:            "256",
			expectedSignal:  "UNK",
			expectedDisplay: true,
		},
	}

	for i, test := range tests {
		car := Car{}

		if test.dropArgs {
			delete(car.Params, "args")
		} else {
			car.SetParamStr("args", test.args)
		}

		car.Init()

		if car.Display != test.expectedDisplay {
			t.Errorf("Test [%d]: Expected %t, found %t.", i, test.expectedDisplay, car.Display)
		}

		if car.Model["Signal"].Text != test.expectedSignal {
			t.Errorf("Test [%d]: Expected signal %s, found %s.", i, test.expectedSignal, car.Model["Signal"].Text)
		}
	}
}
