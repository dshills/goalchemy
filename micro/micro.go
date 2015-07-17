// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.
package micro

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

// Micro enpoint constants
const (
	EndpointURL  = "url/URLGetMicroformatData"
	EndpointHTML = "html/HTMLGetMicroformatData"
)

// Micro represents a Micro query result
type Micro struct {
	Status     string   `json:"status"`
	Usage      string   `json:"usage"`
	TT         string   `json:"TotalTransactions"`
	Language   string   `json:"language"`
	StatusInfo string   `json:"statusInfo"`
	Results    []Result `json:"microformats"`

	Transactions int
}

// Result represents a scoring for a category
type Result struct {
	FieldName string `json:"fieldName"`
	FieldData string `json:"fieldData"`
}

// Decode parses json data into Results.
func (t *Micro) Decode(data []byte) error {
	if err := json.Unmarshal(data, t); err != nil {
		return err
	}
	if t.Status != "OK" {
		return errors.New(t.StatusInfo)
	}
	t.Transactions, _ = strconv.Atoi(t.TT)

	return nil
}

// Required checks for required parameters
func (t *Micro) Required(end string, p url.Values) error {
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
