// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

// Package goalchemy package supports queries against the IBM Alchemy API
package goalchemy

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

const baseurl = "https://access.alchemyapi.com/calls/"

// TODO add to HTTP header
// Accept-encoding:gzip

// AlchemyAPIer supports calling and decoding an AlchemyAPI query
type AlchemyAPIer interface {
	Decode([]byte) error
	Required(string, url.Values) error
}

// Query represents a call to the Alchemy API
type Query struct {
	endpoint string
	params   url.Values
	lastRead []byte
}

// NewQuery returns a query with endpoint of end using apikey
func NewQuery(end, apikey string) *Query {
	co := Query{params: url.Values{}, endpoint: end}
	co.AddParam("apikey", apikey)
	return &co
}

// Encode returns a url encoded string of params
func (co *Query) Encode() string {
	return co.params.Encode()
}

// AddParam adds a key value pair to the param list
func (co *Query) AddParam(k, v string) {
	co.params.Add(k, v)
}

// SetParam sets a key value pair replacing they key if found.
func (co *Query) SetParam(k, v string) {
	co.params.Set(k, v)
}

// Raw returns the last read byte data
func (co *Query) Raw() []byte {
	return co.lastRead
}

// Payload returns the encoded query string
func (co *Query) Payload() string {
	return co.params.Encode()
}

// Endpoint returns the query endpoint
func (co *Query) Endpoint() string {
	return baseurl + co.endpoint
}

// Run builds a query string, calls the appropriate AlchemyAPI and decodes the results
// into the AlchemyAPIer interface
func (co *Query) Run(api AlchemyAPIer) error {
	co.AddParam("outputMode", "json")
	if err := api.Required(co.endpoint, co.params); err != nil {
		return err
	}
	res, err := http.PostForm(co.Endpoint(), co.params)
	if err != nil {
		return err
	}
	co.lastRead, err = ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return err
	}

	err = api.Decode(co.lastRead)
	return err
}
