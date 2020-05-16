# pretty
[![Build Status](https://travis-ci.org/ackneal/pretty.svg?branch=master)](https://travis-ci.com/github/ackneal/pretty)

It's very simple and easy way to pretty print JSON on terminal.

### Installation

Install and update this go package with `go get -u github.com/ackneal/pretty`

### Usage
```sh
> curl -s http://date.jsontest.com | pretty
{
    "date": "05-05-2020",
    "milliseconds_since_epoch": 1588686468047,
    "time": "01:47:48 PM"
}%

> pretty '{"pretty": true}' 
{
    "pretty": true
}%        
```
