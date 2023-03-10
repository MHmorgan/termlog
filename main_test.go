package termlog

import (
	"os"
	"testing"
)

func Example() {
	SetLogFile(os.Stdout)
	Bad("Bad")
	Emph("Emph")
	Error("Error")
	Good("Good")
	Info("Info")
	Warn("Warn")
	// Output:
	// [✗] Bad
	// [*] Emph
	// [!!] Error
	// [✓] Good
	// [·] Info
	// [!] Warn
}

func TestBad(t *testing.T) {
	t.Run("TestBad", func(t *testing.T) {
		Bad("TestBad")
		SetTimestampEnabled(true)
		Badln("TestBad with timestamp")
		SetColorsEnabled(false)
		Badf("TestBad without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestEmph(t *testing.T) {
	t.Run("TestEmph", func(t *testing.T) {
		Emph("TestEmph")
		SetTimestampEnabled(true)
		Emphln("TestEmph with timestamp")
		SetColorsEnabled(false)
		Emphf("TestEmph without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestErr(t *testing.T) {
	t.Run("TestErr", func(t *testing.T) {
		Error("TestErr")
		SetTimestampEnabled(true)
		Errorln("TestErr with timestamp")
		SetColorsEnabled(false)
		Errorf("TestErr without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestGood(t *testing.T) {
	t.Run("TestGood", func(t *testing.T) {
		Good("TestGood")
		SetTimestampEnabled(true)
		Goodln("TestGood with timestamp")
		SetColorsEnabled(false)
		Goodf("TestGood without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestInfo(t *testing.T) {
	t.Run("TestInfo", func(t *testing.T) {
		Info("TestInfo")
		SetTimestampEnabled(true)
		Infoln("TestInfo with timestamp")
		SetColorsEnabled(false)
		Infof("TestInfo without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestWarn(t *testing.T) {
	t.Run("TestWarn", func(t *testing.T) {
		Warn("TestWarn")
		SetTimestampEnabled(true)
		Warnln("TestWarn with timestamp")
		SetColorsEnabled(false)
		Warnf("TestWarn without %v", "colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}
