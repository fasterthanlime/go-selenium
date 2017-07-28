# go-selenium

[![Build Status](https://travis-ci.org/fasterthanlime/go-selenium.svg?branch=master)](https://travis-ci.org/fasterthanlime/go-selenium)
[![GoDoc](https://godoc.org/github.com/fasterthanlime/go-selenium?status.svg)](https://godoc.org/github.com/fasterthanlime/go-selenium)
[![Go Report Card](https://goreportcard.com/badge/github.com/fasterthanlime/go-selenium)](https://goreportcard.com/report/github.com/fasterthanlime/go-selenium)

## Introduction

Yes, yet another Selenium Web Driver library has been brought to the table. This one, however, is slightly different. 

* Easy to understand.
* Full test coverage by unit tests and integration tests.
* Excellent support; we use this in our main project so if you find an issue - it'll likely impact us!
* Idiomatic, structured code with no gimmicks.
* Simple errors that describe what has happened and why.

## Installation

As with all Go libraries, go-selenium is easy to install. Simply run the below command:

`go get github.com/fasterthanlime/go-selenium`

and then import the library in your project:

`import "github.com/fasterthanlime/go-selenium"`

## Getting started

Prior to using this library you need to ensure that you have a Selenium instance running (standalone or grid is fine). If you don't know how to do this, there is a small section called 'Getting Selenium running' below.

Please see the [examples/getting-started/main.go](https://github.com/fasterthanlime/go-selenium/blob/master/examples/getting-started/main.go) file.

## Examples

Further examples, including tests of HackerNews (c), are available within the `examples` directory.

## Documentation

All documentation is available on the godoc.org website: [https://godoc.org/github.com/fasterthanlime/go-selenium](https://godoc.org/github.com/fasterthanlime/go-selenium). 

## Getting Selenium running

### With Docker

1. Choose an image from the following URL: https://github.com/SeleniumHQ/docker-selenium
2. Execute the following Docker command replacing the image with your chosen one: `docker run -d -p 4444:4444 --name selenium selenium/standalone-firefox`.

### Without Docker

1. Download the Selenium standalone server from the following URL: http://www.seleniumhq.org/download/
2. Download the appropriate web driver executable and include it in your path. For Firefox, that will be the Gecko driver. 
3. Run the Selenium server with the following command: `java -jar selenium-server-standalone-3.0.1.jar`.
