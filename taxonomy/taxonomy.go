// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

package taxonomy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Taxonomy enpoint constants
const (
	EndpointURL  = "url/URLGetRankedTaxonomy"
	EndpointText = "text/TextGetRankedTaxonomy"
	EndpointHTML = "html/HTMLGetRankedTaxonomy"
)

// Taxonomies represents a Taxonomy query result
type Taxonomies struct {
	data.QStatus
	Results []data.Taxonomy `json:"taxonomy"`
}

// Decode parses json data into Results.
func (t *Taxonomies) Decode(data []byte) error {
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
func (t *Taxonomies) Required(end string, p url.Values) error {
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
		if p.Get("url") == "" {
			el = append(el, "url")
		}
	}
	if len(el) > 0 {
		es := fmt.Sprintf("Missing required parameters: %v", strings.Join(el, ", "))
		return errors.New(es)
	}
	return nil
}
