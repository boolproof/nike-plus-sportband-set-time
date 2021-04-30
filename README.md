# Nike+ Sportband time setting
Tiny tool to set time on old Nike Sportband+ (the one that used sensor in the shoe).

After Nike shut down web app few years ago, device became useless. However it can still be used as a simple digital watch. This tiny tool lets you set the time on the device. Once in a while you will need to do it (e.g. discharged battery or switch from/to DST). You can of course use old official desktop app (I guess it should still work, just won't connect to shut down API) for that purpose, but that doesn't sound right (and still, the desktop app was available only for Win and Mac, no Linux etc.).

## Requirements
You will need https://pkg.go.dev/github.com/google/gousb as a dependency, which underneath uses https://github.com/libusb/libusb/wiki

## Running (the tool...)
In order to access device on USB port, I had to run built executable with SUDO.
