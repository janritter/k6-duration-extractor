# K6 duration extractor

[![CircleCI](https://circleci.com/gh/janritter/k6-duration-extractor/tree/main.svg?style=svg)](https://circleci.com/gh/janritter/k6-duration-extractor/tree/main)
[![Maintainability](https://api.codeclimate.com/v1/badges/a491d029cae01b6b464a/maintainability)](https://codeclimate.com/github/janritter/k6-duration-extractor/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/a491d029cae01b6b464a/test_coverage)](https://codeclimate.com/github/janritter/k6-duration-extractor/test_coverage)

> Tool to extract the duration information per request from the K6 csv output

## Prerequisites
- Golang >= 1.16.6

## Installation

### Clone git repo
```bash
git clone git@github.com:janritter/k6-duration-extractor.git
```

### Open project directory
```bash
cd k6-duration-extractor
```

### Install via go
```bash
go install
```

## Installation via Homebrew (For Mac / Linux)

### Get the formula
```bash
brew tap janritter/k6-duration-extractor https://github.com/janritter/k6-duration-extractor
```

### Install formula
```bash
brew install k6-duration-extractor
```

## Usage

### Execute your K6 test with CSV output 

```bash
k6 run --out csv=result.csv script.js
```

### Run the k6-duration-extractor for the result.csv K6 output file

```bash
k6-duration-extractor result.csv
```

The k6-duration-extractor will save each request duration in a CSV file called `<your-k6-output-filename>-extracted.csv`

In our case this is `result-extracted.csv`

## Result example
```csv
timestamp,duration,check,error_message,error_code,group,method,name,proto,scenario,status,subproto,tls_version,url,extra_tags
1595325560,221.899000,,,,,GET,http://test.k6.io,HTTP/1.1,default,200,,,http://test.k6.io,
1595325560,221.899000,,,,,GET,http://test.k6.io,HTTP/1.1,default,200,,,http://test.k6.io,
1595325560,221.899000,,,,,GET,http://test.k6.io,HTTP/1.1,default,200,,,http://test.k6.io,
1595325560,221.899000,,,,,GET,http://test.k6.io,HTTP/1.1,default,200,,,http://test.k6.io,
1595325560,221.899000,,,,,GET,http://test.k6.io,HTTP/1.1,default,200,,,http://test.k6.io,

```

## License and Author

Author: Jan Ritter

License: MIT
