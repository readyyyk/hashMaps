# Random Image API

## Usage
host: https://readyyyk-randimgapi.onrender.com/
**http://`host`:`port`/hashmaps?`...[url params]`**
**http://`host`:`port`/picsum?`...[url params]`**
> width and height defaults:
>  - hashmaps - 7x7
>  - picsum - 64x64

> returns:
> - hashmaps - `svg` image
> - picsum - `jpeg` image

## Installation

### With prebuild binaries
 - Download needed binary from releases
 - Create `.env` with needed `PORT=<int>`
 - Run it

### With custom build
Clone repo with
```bash
$ git clone https://github.com/readyyyk/hashMaps
$ cd hashMaps
```

Create `cmd/.env` file that contains `PORT=<ur port(int)>`
```bash
touch cmd/.env && echo "PORT=8080" >> cmd/.env
```
Create ur build with
```bash
go build -o cmd/hashMaps
```
Run with
```bash
cmd/hashMaps
```
Or run without building
```bash
go run .
```
