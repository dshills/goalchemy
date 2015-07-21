# GoAlchemy [![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://godoc.org/github.com/dshills/goalchemy)

## Overview
Simple Go SDK for the IBM Alchemy API

## Features

* Concept Tagging: URL,Text, HTML
* Entity Extraction: URL, Text, HTML
* Keyword Extraction: URL, Text, HTML
* Microformats Parsing: URL, HTML
* Taxonomy: URL, Text, HTML
* Relation Extraction: URL, Text, HTML
* Sentiment Analysis: URL, Text, HTML (and targeted)

## Installation
	go get github.com/dshills/goalchemy
	
## Concepts

### Example Query

```go
const apikey = "<MY API KEY>"
const aurl = "http://somerandom.com/url/about/stuff"

tax := &taxonomy.Taxonomies{}
query := goalchemy.NewQuery(taxonomy.EndpointURL, apikey)
query.AddParam("url", aurl)
if err := query.Run(tax); err != nil {
	panic(err)
}

// Shiny list of taxonomies for aurl
fmt.Println(tax)
```

### Query
Query is used to setup a connection the Alchemy API, pick the end point, and return the query results.

```go
// NewQuery returns a query with endpoint of end using apikey
func NewQuery(end, apikey string) *Query

// AddParam adds a key value pair to the param list
func (co *Query) AddParam(k, v string)

// SetParam sets a key value pair replacing they key if found.
func (co *Query) SetParam(k, v string) {

// Run builds a query string, calls the appropriate AlchemyAPI and returns the decoded results
// into the AlchemyAPIer interface
func (co *Query) Run(api AlchemyAPIer) error
```

### AlchemyAPIer
All of the supported APIs are built as AlchemyAPIer interfaces.

```go
// AlchemyAPIer supports calling and decoding an AlchemyAPI query
type AlchemyAPIer interface {
	Decode([]byte) error
	Required(string, url.Values) error
}
```

### Example AlchemyAPIer for Taxonomy

```go
// Taxonomies is a list of Taxonomy results
type Taxonomies struct {
	Status     string 
	Usage      string 
	Language   string 
	StatusInfo string 
	Transactions int
	Results []Taxonomy
}

// Taxonomy represents a Taxonomy query result
type Taxonomy struct {
	Category string 
	Confident bool
	Score     float32
}
```

## To Do
* Support all of the API

## Alternatives

* [lineback/alchemyapi_go](https://github.com/lineback/alchemyapi_go)
* [elvuel/alchemyapi_go](http://github.com/elvuel/alchemyapi_go)

## License
Copyright 2015 Davin Hills. All rights reserved.
MIT license. License details can be found in the LICENSE file.
