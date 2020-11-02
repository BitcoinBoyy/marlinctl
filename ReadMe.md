# Marlinctl

Marlinctl provides a command line interface for setting up beacon and relay node of marlin network

# Cloning
 ```sh
$ git clone https://github.com/marlinprotocol/marlinctl.git
```


# Building
Prerequisites: go >= 1.15.1 and supervisorctl
```
$ go build -o build/marlinctl
```

# Usage
After building,
```
$ cd build
```
To get list of available commands: 
```
$ sudo ./marlinctl help
```
To create, start or stop **Beacon**:
```
$ sudo ./marlinctl beacon command [command options] [arguments...]
```
Check help to get information on each command:
```
$ sudo ./marlinctl beacon help
```
To create, start or stop **Relay**:
```
$ sudo ./marlinctl relay command [command options] [arguments...]
```
Check help to get information on each command:
```
$ sudo ./marlinctl relay help
```
