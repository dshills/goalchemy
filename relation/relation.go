// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package relation

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Relation enpoint constants
const (
	EndpointURL  = "url/URLGetRelations"
	EndpointText = "text/TextGetRelations"
	EndpointHTML = "html/HTMLGetRelations"
)

// Relation represents a Relation query result
type Relation struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"language"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"relations"`

	Transactions int
}

// Result represents a scoring for a category
type Result struct {
	Sentence string `json:"sentence"`
	Subject  struct {
		Text      string `json:"text"`
		Sentiment struct {
			Type  string `json:"type"`
			Score string `json:"score"`
			Mixed string `json:"mixed"`
		}
		Entity struct {
			Type          string `json:"type"`
			Text          string `json:"text"`
			Disambiguated struct {
				Name        string `json:"name"`
				SubType     string `json:"subType"`
				Website     string `json:"website"`
				Geo         string `json:"geo"`
				Dbpedia     string `json:"dbpedia"`
				Yago        string `json:"yago"`
				Opencyc     string `json:"opencyc"`
				Umbel       string `json:"umbel"`
				Freebase    string `json:"freebase"`
				CiaFactbook string `json:"ciaFactbook"`
				Census      string `json:"census"`
				Geonames    string `json:"geonames"`
				MusicBrainz string `json:"musicBrainz"`
				Crunchbase  string `json:"crunchbase"`
			}
		}
	}
	Action struct {
		Text       string `json:"text"`
		Lemmatized string `json:"lemmatized"`
		Verb       struct {
			Text    string `json:"text"`
			Tense   string `json:"tense"`
			Negated string `json:"negated"`
		}
	}
	Object struct {
		Text      string `json:"text"`
		Sentiment struct {
			Type  string `json:"type"`
			Score string `json:"score"`
			Mixed string `json:"mixed"`
		}
		SentimentFromSubject struct {
			Type  string `json:"type"`
			Score string `json:"score"`
			Mixed string `json:"mixed"`
		}
		Entity struct {
			Type          string `json:"type"`
			Text          string `json:"text"`
			Disambiguated struct {
				Name        string `json:"name"`
				SubType     string `json:"subType"`
				Website     string `json:"website"`
				Geo         string `json:"geo"`
				Dbpedia     string `json:"dbpedia"`
				Yago        string `json:"yago"`
				Opencyc     string `json:"opencyc"`
				Umbel       string `json:"umbel"`
				Freebase    string `json:"freebase"`
				CiaFactbook string `json:"ciaFactbook"`
				Census      string `json:"census"`
				Geonames    string `json:"geonames"`
				MusicBrainz string `json:"musicBrainz"`
				Crunchbase  string `json:"crunchbase"`
			}
		}
	}
	Location struct {
		Text string `json:"text"`
	}
}

// Decode parses json data into Results.
func (t *Relation) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Status != "OK" {
		return errors.New(t.StatusInfo)
	}
	t.Transactions, _ = strconv.Atoi(t.TT)
	for i := 0; i < len(t.Results); i++ {
	}

	return nil
}

// Required checks for required parameters
func (t *Relation) Required(end string, p url.Values) error {
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
