package lcd

import (
	"fmt"

	"periph.io/x/periph/conn/gpio"
	"periph.io/x/periph/conn/gpio/gpioreg"
	"periph.io/x/periph/experimental/devices/hd44780"
)

type LCD struct {
	dev *hd44780.Dev
}

func New(rs, e string, data []string) (*LCD, error) {
	rsPinReg := gpioreg.ByName(rs)
	if rsPinReg == nil {
		return nil, fmt.Errorf("Register select not set %s", rs)
	}
	ePinReg := gpioreg.ByName(e)
	if ePinReg == nil {
		return nil, fmt.Errorf("Strobe select not set %s", e)
	}

	var dataPins [4]gpio.PinOut

	for i, pinName := range data {
		if dataPins[i] = gpioreg.ByName(pinName); dataPins[i] == nil {
			return nil, fmt.Errorf("Data pin %s can not be found", pinName)
		}
	}

	dev, err := hd44780.New(dataPins[:], rsPinReg, ePinReg)

	if err != nil {
		return nil, err
	}

	return &LCD{
		dev: dev,
	}, nil
}

func (lcd *LCD) Cls() error {
	return lcd.dev.Reset()
}

func (lcd *LCD) Println(message string) error {
	if err := lcd.dev.SetCursor(0, 0); err != nil {
		return err
	}
	return lcd.dev.Print(message)
}
