// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

// Package sentiment supports decoding of Sentiment Analysis API calls
// The Sentiments data struct supports the AlchemyAPIer interface
package sentiment

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"

	"github.com/dshills/goalchemy/data"
)

// Sentiment enpoint constants
const (
	EndpointURL        = "url/URLGetTextSentiment"
	EndpointText       = "text/TextGetTextSentiment"
	EndpointHTML       = "html/HTMLGetTextSentiment"
	EndpointTargetURL  = "url/URLGetTargetedSentiment"
	EndpointTargetText = "text/TextGetTargetedSentiment"
	EndpointTargetHTML = "html/HTMLGetTargetedSentiment"
)

// Sentiments represents a sentiment query result
type Sentiments struct {
	data.QStatus
	Result data.Sentiment `json:"docSentiment"`
}

// Decode parses json data into Results.
func (t *Sentiments) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if err := t.Error(); err != nil {
		return err
	}
	t.Clean()
	t.Result.Clean()
	return nil
}

// Required checks for required parameters
func (t *Sentiments) Required(end string, p url.Values) error {
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
	case EndpointTargetURL:
		if p.Get("url") == "" {
			el = append(el, "url")
		}
		if p.Get("target") == "" {
			el = append(el, "target")
		}
	case EndpointTargetText:
		if p.Get("text") == "" {
			el = append(el, "text")
		}
		if p.Get("target") == "" {
			el = append(el, "target")
		}
	case EndpointTargetHTML:
		if p.Get("html") == "" {
			el = append(el, "html")
		}
		if p.Get("target") == "" {
			el = append(el, "target")
		}
	}
	if len(el) > 0 {
		es := fmt.Sprintf("Missing required parameters: %v", strings.Join(el, ", "))
		return errors.New(es)
	}
	return nil
}
