package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/gousb"
)

const (
	VENDOR_ID  = 0x11ac
	PRODUCT_ID = 0x4269
)

func main() {
	println()
	println("Nike SportBand+ time setting")
	println()
	println("Searching for USB device...")

	ctx := gousb.NewContext()
	defer ctx.Close()
	dev, err := ctx.OpenDeviceWithVIDPID(VENDOR_ID, PRODUCT_ID)

	if err != nil {
		log.Fatalf("Error while opening device: %v", err)
	}

	if dev == nil {
		log.Fatal("Device not found. Are you sure it is plugged in?")
	}

	defer dev.Close()
	dev.SetAutoDetach(true)
	activeCfg, err := dev.ActiveConfigNum()

	dev.Config(activeCfg)
	if err != nil {
		log.Fatalf("Could not read active config: %v", err)
	}

	manufacturer, err := dev.Manufacturer()
	if err != nil {
		log.Fatalf("Could not read manufacturer: %v", err)
	}

	product, err := dev.Product()
	if err != nil {
		log.Fatalf("Could not read product: %v", err)
	}

	println(fmt.Sprintf("Device found: manufacturer: \"%s\", product: \"%s\"", manufacturer, product))

	var customTime string

	if len(os.Args) > 1 {
		customTime = os.Args[1]
	}

	st := time.Now()
	var t time.Time
	if customTime == "" {
		t = time.Now()
	} else {
		t, err = time.Parse("2006-01-02 MST 15:04", fmt.Sprintf("%s %s", st.Format("2006-01-02 MST"), customTime))

		if err != nil {
			log.Fatalf("Could not parse custom time: %v. Please make sure you are using HH:MM format for argument.", err)
		}
	}

	println("Current system date and time: ", st.Format("2006-01-02 15:04:05 MST"))

	if customTime != "" {
		println("Custom date and time: ", t.Format("2006-01-02 15:04:05 MST"))
	}

	println("Setting time on device...")

	ts := t.Unix()

	//converting timestamp into slice of 4 bytes
	tb := [4]byte{
		byte(0xff & ts),
		byte(0xff & (ts >> 8)),
		byte(0xff & (ts >> 16)),
		byte(0xff & (ts >> 24))}

	data := []byte{10, 11, 48, 33, tb[3], tb[2], tb[1], tb[0], 0, 1, 67, 112, 1, 0, 0, 0}

	n, err := dev.Control(33, 9, 522, 0, data)

	if err != nil || n != 16 {
		log.Fatalf("Error while setting time on device: %v", err)
	}

	println("Time on device has been successfully set to ", t.Format("15:04"))
}
