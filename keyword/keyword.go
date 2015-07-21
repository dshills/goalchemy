// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

package keyword

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Keyword endpoint constants
const (
	EndpointURL  = "url/URLGetRankedKeywords"
	EndpointText = "text/TextGetRankedKeywords"
	EndpointHTML = "html/HTMLGetRankedKeywords"
)

// Keywords represents a Keyword query result
type Keywords struct {
	data.QStatus
	Results []data.Keyword `json:"keywords"`
}

// Decode parses json data into Results.
func (t *Keywords) Decode(data []byte) error {
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
func (t *Keywords) Required(end string, p url.Values) error {
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
