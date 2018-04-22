package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/jdevelop/whereami/display"
	"github.com/jdevelop/whereami/display/console"
	"github.com/jdevelop/whereami/display/lcd"
	"github.com/jdevelop/whereami/resolver"
	"periph.io/x/periph/host"
)

func main() {

	var (
		out        = flag.String("out", "console", "Status output format (console or lcd)")
		interval   = flag.Int("interval", 30, "Status refresh interval, seconds")
		iterations = flag.Int("iterations", 30, "Number of iterations")

		dataPins = flag.String("lcd-data", "", "LCD data pins, comma-separated BCM pin numbers")
		rsPin    = flag.String("lcd-rs", "", "LCD RS pin, BCM pin number")
		ePin     = flag.String("lcd-e", "", "LCD E pin, PCM pin number")
	)

	flag.Parse()

	rs, err := resolver.New()
	if err != nil {
		log.Fatal(err)
	}

	var d display.Display

	switch *out {
	case "console":
		d, err = console.New(os.Stdout)
		if err != nil {
			log.Fatal(err)
		}
	case "lcd":
		_, err = host.Init()
		if err != nil {
			log.Fatal(err)
		}
		if *rsPin == "" || *ePin == "" || *dataPins == "" {
			fmt.Fprintf(os.Stderr, "LCD is not configured properly\n")
			flag.Usage()
			os.Exit(1)
		}
		dPins := strings.Split(*dataPins, ",")
		if len(dPins) != 4 {
			fmt.Fprintf(os.Stderr, "LCD is not configured properly\n")
			flag.Usage()
			os.Exit(1)
		}
		d, err = lcd.New(*rsPin, *ePin, dPins)
		if err != nil {
			log.Fatal(err)
		}
	default:
		flag.Usage()
		os.Exit(1)
	}

	t := time.NewTicker(time.Duration(*interval) * time.Second)

	var found bool

	d.Println("Network init..")

	for i := 0; i < *iterations; i++ {
		_ = <-t.C
		rt, err := rs.GetDefaultRouteSrc()
		d.Cls()
		if err == nil {
			d.Println(rt.String())
			found = true
		} else {
			d.Println(fmt.Sprintf("No net %d/%d", i, *iterations))
		}
	}

	t.Stop()

	if !found {
		d.Cls()
		d.Println("Network failure")
	}

}
