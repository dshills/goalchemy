// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package keyword

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Keyword enpoint constants
const (
	EndpointURL  = "url/URLGetRankedKeywords"
	EndpointText = "text/TextGetRankedKeywords"
	EndpointHTML = "html/HTMLGetRankedKeywords"
)

// Keyword represents a Keyword query result
type Keyword struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"langauge"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"keywords"`

	Transactions int
}

// Result represents the individual keywords
type Result struct {
	Sentiment struct {
		Type  string `json:"type"`
		S     string `json:"score"`
		Mixed string `json:"mixed"`

		Score float32
	}
	KnowledgeGraph struct {
		TH string `json:"typeHierarchy"`
	}
	Text string `json:"text"`
	R    string `json:"relevance"`

	Relevance     float32
	TypeHierarchy string
}

// Decode parses json data into Results.
func (t *Keyword) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Status != "OK" {
		return errors.New(t.StatusInfo)
	}
	t.Transactions, _ = strconv.Atoi(t.TT)
	for i := 0; i < len(t.Results); i++ {
		r := &t.Results[i]
		s, _ := strconv.ParseFloat(r.Sentiment.S, 32)
		r.Sentiment.Score = float32(s)
		rr, _ := strconv.ParseFloat(r.R, 32)
		r.Relevance = float32(rr)
		r.TypeHierarchy = r.KnowledgeGraph.TH
	}

	return nil
}

// Required checks for required parameters
func (t *Keyword) Required(end string, p url.Values) error {
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
