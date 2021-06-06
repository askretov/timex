# Mutable
[![Go Report Card](https://goreportcard.com/badge/github.com/askretov/timex)](https://goreportcard.com/report/github.com/askretov/timex)
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/8f1471f2c592442cab747c39766ed58c)](https://www.codacy.com/gh/askretov/timex/dashboard?utm_source=github.com&amp;utm_medium=referral&amp;utm_content=askretov/timex&amp;utm_campaign=Badge_Grade)
[![codecov](https://codecov.io/gh/askretov/mutable/branch/master/graph/badge.svg)](https://codecov.io/gh/askretov/timex)
[![Build Status](https://travis-ci.org/askretov/timex.svg?branch=master)](https://travis-ci.org/askretov/timex)
[![GoDoc](https://godoc.org/github.com/askretov/timex?status.svg)](https://godoc.org/github.com/askretov/timex)
[![Licenses](https://img.shields.io/badge/license-mit-brightgreen.svg)](https://opensource.org/licenses/BSD-3-Clause)

## Introduction
Time eXtension package providing extra features to work with time for golang

## Usage
### Installation
```go
go get github.com/askretov/timex
```
### Time Interval examples
```go
// Create an interval
start := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
i := NewInterval(start, start.AddDate(0, 0, 9))
fmt.Println(i)
// > 2021-01-01T00:00:00Z - 2021-01-10T00:00:00Z

// Print the interval duration in days
fmt.Println(i.Days())
// > 9

// Check if t inside of i
t := start.AddDate(0, 0, 5)
fmt.Println(i.Contains(t))
// > true

// Get a half-open interval for i
ho := i.HalfOpenEnd()
fmt.Println(ho)
// > 2021-01-01T00:00:00Z - 2021-01-09T23:59:59Z
```
### A lot more features coming soon, stay tuned!