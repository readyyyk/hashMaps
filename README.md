# Hash maps

## Usage
**http://`host`:`port`/render?seed=`any`&w=`[1..100]`&h=`[1..100]`**
> width and height are not required, default is 7x7

> returns `svg` image

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
