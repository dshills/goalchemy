// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

package micro

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Micro endpoint constants
const (
	EndpointURL  = "url/URLGetMicroformatData"
	EndpointHTML = "html/HTMLGetMicroformatData"
)

// Micros represents a Micro query result
type Micros struct {
	data.QStatus
	Results []data.Micro `json:"microformats"`
}

// Decode parses json data into Results.
func (t *Micros) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if err := t.Error(); err != nil {
		return err
	}
	t.Clean()

	return nil
}

// Required checks for required parameters
func (t *Micros) Required(end string, p url.Values) error {
	var el []string
	switch end {
	case EndpointURL:
		if p.Get("url") == "" {
			el = append(el, "url")
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
