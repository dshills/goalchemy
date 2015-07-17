// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package concept

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Concept enpoint constants
const (
	EndpointURL  = "url/URLGetRankedConcepts"
	EndpointText = "text/TextGetRankedConcepts"
	EndpointHTML = "html/HTMLGetRankedConcepts"
)

// Concept represents a Concept query result
type Concept struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"langauge"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"concepts"`

	Transactions int
}

// Result represents the individual concepts
type Result struct {
	KnowledgeGraph struct {
		TH string `json:"typeHierarchy"`
	}
	Concept     string `json:"concept"`
	R           string `json:"relevance"`
	Website     string `json:"website"`
	Geo         string `json:"geo"`
	Dbpedia     string `json:"dbpedia"`
	Yago        string `json:"dbpedia"`
	Opencyc     string `json:"opencyc"`
	Freebase    string `json:"freebase"`
	CiaFactbook string `json:"ciaFacebook"`
	Census      string `json:"census"`
	GeoNames    string `json:"geonames"`
	MusicBrainz string `json:"musicBrainz"`
	Crunchbase  string `json:"crunchbase"`

	Relevance     float32
	TypeHierarchy string
}

// Decode parses json data into Results.
func (t *Concept) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Status != "OK" {
		return errors.New(t.StatusInfo)
	}
	t.Transactions, _ = strconv.Atoi(t.TT)
	for i := 0; i < len(t.Results); i++ {
		r := &t.Results[i]
		rr, _ := strconv.ParseFloat(r.R, 32)
		r.Relevance = float32(rr)
		r.TypeHierarchy = r.KnowledgeGraph.TH
	}
	return nil
}

// Required checks for required parameters
func (t *Concept) Required(end string, p url.Values) error {
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
