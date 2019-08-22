# License Server 🔐 &middot; ![Status](https://img.shields.io/badge/Status-Pre--Alpha-9cf)

Go based license server.
This README will be populated as development continues to progress. 

## Development
### Requirments
* Docker
* Go
### Setup
* Run make 
  ```sh
  $ src/license-server: make
  ```
* Spin up license server
  ```sh
  $ src/license-server: go run main.go
  ```

If changes are made to the db please make sure to generate xo models. To do so please run `xo_generate.go`
```sh
$ src/license-server/cmd/xo: go run xo_generate.go
```
