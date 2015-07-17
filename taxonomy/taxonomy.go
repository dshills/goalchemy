package taxonomy

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Taxonomy enpoint constants
const (
	EndpointURL  = "url/URLGetRankedTaxonomy"
	EndpointText = "text/TextGetRankedTaxonomy"
	EndpointHTML = "html/HTMLGetRankedTaxonomy"
)

// Taxonomy represents a Taxonomy query result
type Taxonomy struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"language"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"taxonomy"`

	Transactions int
}

// Result represents a scoring for a category
type Result struct {
	Category string `json:"label"`
	C        string `json:"confident"`
	S        string `json:"score"`

	Confident bool
	Score     float32
}

// Decode parses json data into Results.
func (t *Taxonomy) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Status != "OK" {
		return errors.New(t.StatusInfo)
	}
	t.Transactions, _ = strconv.Atoi(t.TT)
	for i := 0; i < len(t.Results); i++ {
		r := &t.Results[i]
		ss, _ := strconv.ParseFloat(r.S, 32)
		r.Score = float32(ss)
		if r.C == "" {
			r.Confident = true
		}
	}

	return nil
}

// Required checks for required parameters
func (t *Taxonomy) Required(end string, p url.Values) error {
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
