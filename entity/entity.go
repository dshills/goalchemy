// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package entity

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Entity endpoint constants
const (
	EndpointURL  = "url/URLGetRankedNamedEntities"
	EndpointText = "text/TextGetRankedNamedEntities"
	EndpointHTML = "html/HTMLGetRankedNamedEntities"
)

// Entity represents a Entity query result
type Entity struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"language"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"entities"`

	Transactions int
}

// Result represents an entity query result
type Result struct {
	Type           string `json:"type"`
	R              string `json:"relevance"`
	KnowledgeGraph struct {
		TH string `json:"TypeHierarchy"`
	}
	Count     string `json:"count"`
	Text      string `json:"text"`
	Sentiment struct {
		Type  string `json:"type"`
		S     string `json:"score"`
		Mixed string `json:"mixed"`

		Score float32
	}
	Disambiguated struct {
		Name        string   `json:"name"`
		SubType     []string `json:"subType"`
		Website     string   `json:"website"`
		Geo         string   `json:"geo"`
		Dbpedia     string   `json:"dbpedia"`
		Yago        string   `json:"yago"`
		Opencyc     string   `json:"opencyc"`
		Umbel       string   `json:"umbel"`
		Freebase    string   `json:"freebase"`
		CiaFactbook string   `json:"ciaFactbook"`
		Census      string   `json:"census"`
		Geonames    string   `json:"geonames"`
		MusicBrainz string   `json:"musicBrainz"`
		Crunchbase  string   `json:"crunchbase"`
	}

	Quotations []struct {
		Quotation string `json:"quotation"`
		Sentiment struct {
			Type  string `json:"type"`
			S     string `json:"score"`
			Mixed string `json:"mixed"`

			Score float32
		}
	}

	Relevance     float32
	TypeHierarchy string
}

// Decode parses json data into Results.
func (t *Entity) Decode(data []byte) error {
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
		for ii := 0; ii < len(r.Quotations); ii++ {
			q := &r.Quotations[ii]
			s, _ := strconv.ParseFloat(q.Sentiment.S, 32)
			q.Sentiment.Score = float32(s)
		}
	}

	return nil
}

// Required checks for required parameters
func (t *Entity) Required(end string, p url.Values) error {
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
