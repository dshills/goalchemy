// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

// Package concept supports decoding of Concept Tagging API calls
// The Concepts data struct supports the AlchemyAPIer interface
package concept

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Concept endpoint constants
const (
	EndpointURL  = "url/URLGetRankedConcepts"
	EndpointText = "text/TextGetRankedConcepts"
	EndpointHTML = "html/HTMLGetRankedConcepts"
)

// Concepts represents a Concept query result
type Concepts struct {
	data.QStatus
	Results []data.Concept `json:"concepts"`
}

// Decode parses json data into Results.
func (t *Concepts) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if err := t.Error(); err != nil {
		return err
	}
	t.Clean()
	for i := 0; i < len(t.Results); i++ {
		t.Results[i].Clean()
	}
	return nil
}

// Required checks for required parameters
func (t *Concepts) Required(end string, p url.Values) error {
	var el []string
	switch end {
	case EndpointURL:
		if p.Get("url") == "" {
			el = append(el, "url")
		}
	case EndpointText:
		if p.Get("text") == "" {
			el = append(el, "text")
		}
	case EndpointHTML:
		if p.Get("html") == "" {
			el = append(el, "html")
		}
	}
	if len(el) > 0 {
		es := fmt.Sprintf("Missing required parameters: %v", strings.Join(el, ", "))
		return errors.New(es)
	}
	return nil
}
