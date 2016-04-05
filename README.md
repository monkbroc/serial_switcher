# Serial Switcher

A simple Go program to switch a serial port to DFU or listen mode

## Usage

1. Download the executable for your platform from the dist folder.
2. Run `serial_switcher 14400 <port_name>` for DFU (where port name
    looks like `/dev/ttyACM0` on Linux, `COM1` on Windows, `/dev/tty.serial1` on OSX)
3. Run `serial_switcher 28800 <port_name>` for listen mode

## Build

1. Install the go toolchain for your platform
2. Clone this repository in to `~/go/src/github.com/monkbroc/serial_switcher`
3. Fetch dependencies with `go get`
4. Build an executable with `go build`

For release, build all platform version with `rake build` (you'll need Ruby and rake installed)

## License

(C) 2016 Julien Vanier
MIT license
