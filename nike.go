package main

import (
	"fmt"
	"log"
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

	t := time.Now()
	println("Current system date and time: ", t.Format("2006-01-02 15:04:05"))
	println("Setting time on device...")

	ts := t.Unix()

	//converting timestamp into slice of 4 bytes
	b := [4]byte{
		byte(0xff & ts),
		byte(0xff & (ts >> 8)),
		byte(0xff & (ts >> 16)),
		byte(0xff & (ts >> 24))}

	st := []byte{10, 11, 48, 33, b[3], b[2], b[1], b[0], 0, 1, 67, 112, 1, 0, 0, 0}

	n, err := dev.Control(33, 9, 522, 0, st)

	if err != nil || n != 16 {
		log.Fatalf("Error while setting time on device: %v", err)
	}

	println("Time on device has been successfully set to ", t.Format("15:04"))
}
