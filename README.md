# Go REST-SOAP Relay from WSDL 

Forked from [hooklift/gowsdl](https://github.com/hooklift/gowsdl) Generates Go code from a WSDL file.
Uses [gin-gonic](https://gin-gonic.com/) as REST server library

### Install 
To be updated
* [Download release](https://github.com/hooklift/gowsdl/releases)
* Download and build locally
    * 1.15: `go get github.com/hooklift/gowsdl/...`
    * 1.20: `go install github.com/hooklift/gowsdl/cmd/gowsdl@latest`
* Install from Homebrew: `brew install gowsdl`

### Goals
* Generate idiomatic Go code as much as possible
* Support gSOAP servers with REST clients
* Resolve external XML Schemas
* Support external and local WSDL

### Caveats
* Please keep in mind that the generated code is just a reflection of what the WSDL is like. If your WSDL has duplicated type definitions, your Go code is going to have the same and may not compile.

### Usage
```
Usage: gowsdl [options] myservice.wsdl
  -o string
        File where the generated code will be saved (default "myservice.go")
  -p string
        Package under which code will be generated (default "myservice")
  -i    Skips TLS Verification
  -v    Shows gowsdl version
  ```
