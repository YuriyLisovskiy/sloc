# SLOC - Source Lines Of Code
[![License](https://img.shields.io/badge/BSD-3--Clause-orange.svg)](LICENSE)
[![Language](https://img.shields.io/badge/Go-1.10-blue.svg)](https://golang.org/)
[![Build Status](https://travis-ci.org/YuriyLisovskiy/sloc.svg?branch=master)](https://travis-ci.org/YuriyLisovskiy/sloc)
## Download
```
$ git clone https://github.com/YuriyLisovskiy/sloc.git
```
## Build
> Minimum version of Go language required: `go1.10`, see [golang installation](https://golang.org/doc/install)

Build binaries for all supported operating systems:
```
$ make build
``` 
Run tests and build binaries for all supported operating systems:
```
$ make
```
Available operating systems build:
* **Linux**: `$ make linux` - to build all supported linux architectures, other:
	* `$ make linux-386`
	* `$ make linux-amd64`
	* `$ make linux-arm`
	* `$ make linux-arm64`
* **Windows**: `$ make windows` - to build all supported windows architectures, other:
	* `$ make windows-386`
	* `$ make windows-amd64`
* **Darwin**: `$ make darwin` - to build all supported darwin architectures, other:
	* `$ make darwin-386`
	* `$ make darwin-amd64`
* **FreeBSD**: `$ make freebsd` - to build all supported freebsd architectures, other:
	* `$ make freebsd-386`
	* `$ make freebsd-amd64`
	* `$ make freebsd-arm`
* **Solaris**(only amd64):
	* `$ make solaris`
## Testing
Run tests:
```
$ make test
```
## Usage
By default, `sloc` will count lines of code in current directory:
```
$ sloc count
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Yaml                   1          10           1           0           9
 Xml                    3         781           0           0         781
 Go                    18        1229         117           0        1112
--------------------------------------------------------------------------
 Total                 22        2020         118           0        1902
```
Pass a directory to inspect files:
```
$ sloc count -d src/
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                    11         690          68           0         622
--------------------------------------------------------------------------
 Total                 11         690          68           0         622
```
Or file:
```
$ sloc count -f src/main.go
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                     1          46           3           0          43
--------------------------------------------------------------------------
 Total                  1          46           3           0          43
```
Or multiple file(s) and(or) folder(s):
```
$ sloc count -m ".travis.yml Makefile src/cli/"
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Yaml                   1          10           1           0           9
 Makefile               1          68          28           0          40
 Go                     6         254          29           0         225
--------------------------------------------------------------------------
 Total                  8         332          58           0         274
```
Exclude some file(s) and(or) folder(s):
```
$ sloc count -d src/ -e "src/main.go src/cli/"
--------------------------------------------------------------------------
 Language           Files       Lines       Blank    Comments        Code
--------------------------------------------------------------------------
 Go                    10        1110         100           0        1010
--------------------------------------------------------------------------
 Total                 10        1110         100           0        1010
```
Read usage info:
```
$ sloc help
Usage:
  count
    -f
	set file to count lines
    -d
	set folder to count lines
    -m
	set set files and(or) folders to count lines using ""
    -e
	set file(s) and(or) folder(s) to exclude from counting lines using ""
    -json
	write result to json file
    -xml
	write result to xml file
    -yaml
	write result to yaml file

  help
	read usage
```
Save output to one of an available file format:
* **Json**:
	* `$ ./sloc count -json`
* **Xml**:
	* `$ ./sloc count -xml`
* **Yaml**:
	* `$ ./sloc count -yaml`
### Supported languages
* ActionScript
* Ada
* Agda
* ASP
* ASP.NET
* Assembly
* Autoconf
* Awk
* Batch
* Bourne Shell
* C
* C++
* C#
* C/C++ Header
* C Shell
* CoffeeScript
* ColdFusionScript
* Coq
* CSS
* CUDA
* CUDA Header
* D
* Dart
* DeviceTree
* Erlang
* Forth
* FORTRAN Legacy
* FORTRAN Modern
* F#
* Go
* Groovy
* Haskell
* Html
* Java
* JavaScript
* Json
* Kotlin
* Makefile
* Markdown
* Objective-C/C++
* Pascal
* Perl
* PHP
* Plain text
* Python
* R
* Ruby
* Sass
* SQL
* Swift
* Xml
* Yaml
## Author
* [Yuriy Lisovskiy](https://github.com/YuriyLisovskiy)
## License
This project is licensed under BSD 3-Clause License - see the [LICENSE](LICENSE) file for details.
