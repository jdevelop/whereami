# WhereAmI - IP address finder
This small app helps to figure out the IP address of a RaspberryPI plugged into an unknown network.
 
It tries to find the default route to identify the active network and then displays the source IP address.

## Build
`go get github.com/jdevelop/whereami/cli/whereami`

## Usage
```
Usage of ./whereami:
  -interval int
        Status refresh interval, seconds (default 30)
  -iterations int
        Number of iterations (default 30)
  -lcd-data string
        LCD data pins, comma-separated BCM pin numbers
  -lcd-e string
        LCD E pin, PCM pin number
  -lcd-rs string
        LCD RS pin, BCM pin number
  -out string
        Status output format (console or lcd) (default "console")```
