package termlog

import (
	"testing"
)

func TestBad(t *testing.T) {
	t.Run("TestBad", func(t *testing.T) {
		Bad("TestBad")
		SetTimestampEnabled(true)
		Bad("TestBad with timestamp")
		SetColorsEnabled(false)
		Bad("TestBad without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestEmph(t *testing.T) {
	t.Run("TestEmph", func(t *testing.T) {
		Emph("TestEmph")
		SetTimestampEnabled(true)
		Emph("TestEmph with timestamp")
		SetColorsEnabled(false)
		Emph("TestEmph without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestErr(t *testing.T) {
	t.Run("TestErr", func(t *testing.T) {
		Err("TestErr")
		SetTimestampEnabled(true)
		Err("TestErr with timestamp")
		SetColorsEnabled(false)
		Err("TestErr without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestGood(t *testing.T) {
	t.Run("TestGood", func(t *testing.T) {
		Good("TestGood")
		SetTimestampEnabled(true)
		Good("TestGood with timestamp")
		SetColorsEnabled(false)
		Good("TestGood without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestInfo(t *testing.T) {
	t.Run("TestInfo", func(t *testing.T) {
		Info("TestInfo")
		SetTimestampEnabled(true)
		Info("TestInfo with timestamp")
		SetColorsEnabled(false)
		Info("TestInfo without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}

func TestWarn(t *testing.T) {
	t.Run("TestWarn", func(t *testing.T) {
		Warn("TestWarn")
		SetTimestampEnabled(true)
		Warn("TestWarn with timestamp")
		SetColorsEnabled(false)
		Warn("TestWarn without colors")
		SetTimestampEnabled(false)
		SetColorsEnabled(true)
	})
}
