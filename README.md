# Nike+ Sportband time setting
Tiny tool to set time on old Nike Sportband+ (the one that used sensor in the shoe).

After Nike shut down web app few years ago, device became useless. However it can still be used as a simple digital watch. This tiny tool lets you set the time on the device. Once in a while you will need to do it (e.g. discharged battery or switch from/to DST). You can of course use old official desktop app (I guess it should still work, just won't connect to shut down API) for that purpose, but that doesn't sound right (and still, the desktop app was available only for Win and Mac, no Linux etc.).

## Requirements
You will need https://pkg.go.dev/github.com/google/gousb as a dependency, which underneath uses https://github.com/libusb/libusb/wiki

## Running (the tool...)
In order to access device on USB port, I had to run built executable with SUDO.

### Build
`go build nike.go`

### Run
`sudo ./nike`

This will try to set time to your current system time

### Run with custom time
`sudo ./nike 12:45`

This will try to set time to 12:45 (in your system's timezone)

## Credits
I couldn't find any useful information about Nike+ Sportband on the web.
I was only able to figure out how to set time on the device thanks to information found in this Python code https://github.com/luigigubello/nike-sportband by [@luigigubello ](https://github.com/luigigubello). All the values for the control command come from his code, so thanks a lot Luigi.


