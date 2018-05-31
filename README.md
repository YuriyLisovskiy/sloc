# SLOC - **S**ource **L**ines **O**f **C**ode
[![Build Status](https://travis-ci.org/YuriyLisovskiy/sloc.svg?branch=master)](https://travis-ci.org/YuriyLisovskiy/sloc)
## Download
```
$ git clone https://github.com/YuriyLisovskiy/sloc.git
```
## Build
`make all-os` to build for all supported operating systems. 

`make all` to run tests and build for all supported operating systems.

Available operating systems build:
* **Linux (`make linux` - to build all supported linux architectures):** 
    * `make linux-386`
    * `make linux-amd64`
    * `make linux-arm`
    * `make linux-arm64`
* **Windows (`make windows` - to build all supported windows architectures):**
    * `make windows-386`
    * `make windows-amd64`
* **Darwin (`make darwin` - to build all supported darwin architectures):**
    * `make darwin-386`
    * `make darwin-amd64`
* **FreeBSD (`make freebsd` - to build all supported freebsd architectures):**
* **Solaris (only amd64):** `make solaris` 
## Usage
By default, `sloc` will count lines of code in current directory:
```
$ ./sloc
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Yaml                   1          10           1           0           9
 Xml                    3         781           0           0         781
 Go                    18        1229         117           0        1112
--------------------------------------------------------------------------
 Total                 22        2020         118           0        1902
```
You can also pass a directory to inspect files:
```
$ ./sloc -d src/
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                    11         690          68           0         622
--------------------------------------------------------------------------
 Total                 11         690          68           0         622

```
Or file:
```
$ ./sloc -f src/main.go
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                     1          46           3           0          43
--------------------------------------------------------------------------
 Total                  1          46           3           0          43

```
Or multiple files:
```
$ ./sloc -m "src/main.go src/parser/models.go"
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                     2          82           9           0          73
--------------------------------------------------------------------------
 Total                  2          82           9           0          73

```
Read usage info:
```
$ ./sloc -h
$ ./sloc --h
$ ./sloc -help
$ ./sloc --help
```
Also You can save output to a file:
* **Json**: 
    * `./sloc -j`
    * `./sloc -j -c path/to/result.json`
* **Xml**:
    * `./sloc -x`
    * `./sloc -x -c path/to/result.xml`
* **Yaml**:
    * `./sloc -y`
    * `./sloc -y -c path/to/result.yml`
