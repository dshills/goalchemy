// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

// Package relation supports decoding of Relation Extraction API calls
// The Relations data struct supports the AlchemyAPIer interface
package relation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Relation enpoint constants
const (
	EndpointURL  = "url/URLGetRelations"
	EndpointText = "text/TextGetRelations"
	EndpointHTML = "html/HTMLGetRelations"
)

// Relations represents a Relation query result
type Relations struct {
	data.QStatus
	Results []data.Relation `json:"relations"`
}

// Decode parses json data into Results.
func (t *Relations) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if err := t.Error(); err != nil {
		return err
	}
	for i := 0; i < len(t.Results); i++ {
	}

	return nil
}

// Required checks for required parameters
func (t *Relations) Required(end string, p url.Values) error {
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
