// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

package data

import (
	"errors"
	"strconv"
)

// QResult is a query results
type QResult interface {
	Clean()
}

// QStatus represents a query status
type QStatus struct {
	Status     string `json:"status"`
	Usage      string `json:"usage"`
	TT         string `json:"TotalTransactions"`
	Language   string `json:"language"`
	StatusInfo string `json:"statusInfo"`

	Transactions int
}

// Error returns the error from the query or nil
func (q *QStatus) Error() error {
	if q.Status != "OK" {
		return errors.New(q.StatusInfo)
	}
	return nil
}

// Clean casts query string results to more convenient data types
func (q *QStatus) Clean() {
	q.Transactions, _ = strconv.Atoi(q.TT)
}

// Sentiment represents a Sentiment result
type Sentiment struct {
	Type  string `json:"type"`
	S     string `json:"score"`
	Mixed string `json:"mixed"`

	Score float32
}

// Clean casts query string results to more convenient data types
func (s *Sentiment) Clean() {
	s64, _ := strconv.ParseFloat(s.S, 32)
	s.Score = float32(s64)
}

// Entity represents a entity result
type Entity struct {
	Type           string         `json:"type"`
	R              string         `json:"relevance"`
	KnowledgeGraph KnowledgeGraph `json:"knowledgeGraph"`
	Count          string         `json:"count"`
	Text           string         `json:"text"`
	Sentiment      Sentiment      `json:"sentiment"`
	Disambiguated  Concept        `json:"disambiguated"`
	Quotations     []Quotation    `json:"quotations"`

	Relevance     float32
	TypeHierarchy string
}

// Clean casts query string results to more convenient data types
func (e *Entity) Clean() {
	s64, _ := strconv.ParseFloat(e.R, 32)
	e.Relevance = float32(s64)
	e.TypeHierarchy = e.KnowledgeGraph.TypeHierarchy
	e.Sentiment.Clean()
	for i := 0; i < len(e.Quotations); i++ {
		e.Quotations[i].Clean()
	}
	e.Disambiguated.Clean()
	e.KnowledgeGraph.Clean()
}

// Concept represents a concept result
type Concept struct {
	Text           string         `json:"text"`
	R              string         `json:"relevance"`
	Name           string         `json:"name"`
	SubType        []string       `json:"subType"`
	Website        string         `json:"website"`
	Geo            string         `json:"geo"`
	Dbpedia        string         `json:"dbpedia"`
	Yago           string         `json:"yago"`
	Opencyc        string         `json:"opencyc"`
	Umbel          string         `json:"umbel"`
	Freebase       string         `json:"freebase"`
	CiaFactbook    string         `json:"ciaFactbook"`
	Census         string         `json:"census"`
	Geonames       string         `json:"geonames"`
	MusicBrainz    string         `json:"musicBrainz"`
	Crunchbase     string         `json:"crunchbase"`
	KnowledgeGraph KnowledgeGraph `json:"knowledgeGraph"`

	Relevance     float32
	TypeHierarchy string
}

// Clean casts query string results to more convenient data types
func (c *Concept) Clean() {
	s64, _ := strconv.ParseFloat(c.R, 32)
	c.Relevance = float32(s64)
	if c.Name == "" {
		c.Name = c.Text
	}
	if c.Text == "" {
		c.Text = c.Name
	}
	c.TypeHierarchy = c.KnowledgeGraph.TypeHierarchy
	c.KnowledgeGraph.Clean()
}

// Quotation represents a quotation result
type Quotation struct {
	Quotation string    `json:"quotation"`
	Sentiment Sentiment `json:"sentiment"`
}

// Clean casts query string results to more convenient data types
func (q *Quotation) Clean() {
	q.Sentiment.Clean()
}

// KnowledgeGraph represents a knowledge graph result
type KnowledgeGraph struct {
	TypeHierarchy string `json:"TypeHierarchy"`
}

// Clean casts query string results to more convenient data types
func (k *KnowledgeGraph) Clean() {
}

// Keyword represents a keyword result
type Keyword struct {
	Sentiment      Sentiment      `json:"sentiment"`
	KnowledgeGraph KnowledgeGraph `json:"knowledgeGraph"`
	Text           string         `json:"text"`
	R              string         `json:"relevance"`

	Relevance     float32
	TypeHierarchy string
}

// Clean casts query string results to more convenient data types
func (k *Keyword) Clean() {
	s64, _ := strconv.ParseFloat(k.R, 32)
	k.Relevance = float32(s64)
	k.TypeHierarchy = k.KnowledgeGraph.TypeHierarchy
	k.Sentiment.Clean()
}

// Micro represents a microformats query result
type Micro struct {
	FieldName string `json:"fieldName"`
	FieldData string `json:"fieldData"`
}

// Clean casts query string results to more convenient data types
func (m *Micro) Clean() {
}

// Taxonomy represents a scoring for a category
type Taxonomy struct {
	Category string `json:"label"`
	C        string `json:"confident"`
	S        string `json:"score"`

	Confident bool
	Score     float32
}

// Clean casts query string results to more convenient data types
func (t *Taxonomy) Clean() {
	s64, _ := strconv.ParseFloat(t.S, 32)
	t.Score = float32(s64)
	if t.C == "" {
		t.Confident = true
	}
}

// Relation represents relation query
type Relation struct {
	Sentence string   `json:"sentence"`
	Subject  Subject  `json:"subject"`
	Action   Action   `json:"action"`
	Object   Object   `json:"object"`
	Location Location `json:"location"`
}

// Clean casts query string results to more convenient data types
func (r *Relation) Clean() {
	r.Subject.Clean()
	r.Action.Clean()
	r.Object.Clean()
	r.Location.Clean()
}

// Subject represents the subject of a relation
type Subject struct {
	Text      string    `json:"text"`
	Keywords  []Keyword `json:"keywords"`
	Sentiment Sentiment `json:"sentiment"`
}

// Clean casts query string results to more convenient data types
func (s *Subject) Clean() {
	s.Sentiment.Clean()
	for i := 0; i < len(s.Keywords); i++ {
		s.Keywords[i].Clean()
	}
}

// Action is the action of a relation
type Action struct {
	Text       string `json:"text"`
	Lemmatized string `json:"lemmatized"`
	Verb       Verb   `json:"verb"`
}

// Clean casts query string results to more convenient data types
func (a *Action) Clean() {
	a.Verb.Clean()
}

// Object is the object of a relation
type Object struct {
	Text                 string    `json:"text"`
	Sentiment            Sentiment `json:"sentiment"`
	SentimentFromSubject Sentiment `json:"sentimentFromSubject"`
	Entity               Entity    `json:"entity"`
}

// Clean casts query string results to more convenient data types
func (o *Object) Clean() {
	o.Sentiment.Clean()
	o.SentimentFromSubject.Clean()
	o.Entity.Clean()
}

// Location is the location of a relation
type Location struct {
	Text      string    `json:"text"`
	Sentiment Sentiment `json:"sentiment"`
	Entities  []Entity  `json:"entities"`
}

// Clean casts query string results to more convenient data types
func (l *Location) Clean() {
	l.Sentiment.Clean()
	for i := 0; i < len(l.Entities); i++ {
		l.Entities[i].Clean()
	}
}

// Verb is the verb of an action
type Verb struct {
	Text    string `json:"text"`
	Tense   string `json:"tense"`
	Negated string `json:"negated"`
}

// Clean casts query string results to more convenient data types
func (v *Verb) Clean() {
}
