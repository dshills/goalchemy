// Copyright 2015 Davin Hills. All rights reserved.
// MIT license. License details can be found in the LICENSE file.

package goalchemy_test

import (
	"os"
	"testing"

	"github.com/dshills/goalchemy"
	"github.com/dshills/goalchemy/sentiment"
)

var apikey string

func init() {
	apikey = os.Getenv("ALCHEMY_APIKEY")
	if len(apikey) == 0 {
		panic("Alchemy API Key is required. Set $ALCHEMY_APIKEY")
	}
}

const testurl = "http://www.desmoinesregister.com/story/news/crime-and-courts/2015/07/14/van-meter-taser-fundraiser-modified/30129893/?hootPostID=%5B%27e34bfcfff7cf7090f7c60c8ecd6167a8%27%5D"
const testurl2 = `http://www.dailymail.co.uk/sciencetech/article-2355833/Apples-iPhone-5-hated-handset--majority-people-love-Samsung-Galaxy-S4-study-finds.html`
const testtext = `One year ago, several hours before cities across the United States started their annual fireworks displays, a different type of fireworks were set off at the European Center for Nuclear Research (CERN) in Switzerland. At 9:00 a.m., physicists announced to the world that they had found something they had been searching for for nearly 50 years: the elusive Higgs boson. Today, on the anniversary of its discovery, are we any closer to figuring out what that particle's true identity is? The Higgs boson is popularly referred to as "the God particle," perhaps because of its role in giving other particles their mass. However, it's not the boson itself that gives mass. Back in 1964, Peter Higgs proposed a theory that described a universal field (similar to an electric or a magnetic field) that particles interacted with.`
const testtext2 = `
In 2009, Elliot Turner launched AlchemyAPI to process the written word, with all of its quirks and nuances, and got immediate traction. That first month, the company's eponymous language-analysis API processed 500,000 transactions. Today it's processing three billion transactions a month, or about 1,200 a second. “That's a growth rate of 6,000 times over three years,” touts Turner. “Context is super-important,” he adds. “'I'm dying' is a lot different than 'I'm dying to buy the new iPhone.'” “As we move into new markets, we're going to be making some new hires," Turner says. "We knocked down some walls and added 2,000 square feet to our office.” “We're providing the ability to translate human language in the form of web pages and documents into actionable data,” Turner says. Clients include Walmart, PR Newswire and numerous publishers and advertising networks. “This allows a news organization to detect what a person likes to read about,” says Turner of publishers and advertisers.`
const tphrase = "samsung"

/*
func TestQueryTaxonomy(t *testing.T) {
	taxurl := &taxonomy.Taxonomies{}
	q := goalchemy.NewQuery(taxonomy.EndpointURL, apikey)
	q.AddParam("url", testurl)
	if err := q.Run(taxurl); err != nil {
		t.Error(err)
	}
	if len(taxurl.Results) < 1 {
		t.Error("Expected taxonomy len > 1 got 0")
	}

	taxtext := &taxonomy.Taxonomies{}
	q = goalchemy.NewQuery(taxonomy.EndpointText, apikey)
	q.AddParam("text", testtext)
	if err := q.Run(taxtext); err != nil {
		t.Error(err)
	}
	if len(taxtext.Results) < 1 {
		t.Error("Expected len > 1 got 0")
	}
}

func TestQueryConcept(t *testing.T) {
	c := &concept.Concepts{}
	q := goalchemy.NewQuery(concept.EndpointURL, apikey)
	q.AddParam("url", testurl)
	q.AddParam("knowledgeGraph", "1")
	if err := q.Run(c); err != nil {
		t.Error(err)
	}
	if len(c.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}

	c = &concept.Concepts{}
	q = goalchemy.NewQuery(concept.EndpointText, apikey)
	q.AddParam("text", testtext)
	q.AddParam("knowledgeGraph", "1")
	if err := q.Run(c); err != nil {
		t.Error(err)
	}
	if len(c.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}
}

func TestQueryKeyword(t *testing.T) {
	key := &keyword.Keywords{}
	q := goalchemy.NewQuery(keyword.EndpointURL, apikey)
	q.AddParam("url", testurl)
	q.AddParam("knowledgeGraph", "1")
	q.AddParam("sentiment", "1")
	if err := q.Run(key); err != nil {
		t.Error(err)
	}
	if len(key.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}

	key = &keyword.Keywords{}
	q = goalchemy.NewQuery(keyword.EndpointText, apikey)
	q.AddParam("text", testtext)
	q.AddParam("knowledgeGraph", "1")
	q.AddParam("sentiment", "1")
	if err := q.Run(key); err != nil {
		t.Error(err)
	}
	if len(key.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}
}

func TestQueryEntity(t *testing.T) {
	e := &entity.Entities{}
	q := goalchemy.NewQuery(entity.EndpointURL, apikey)
	q.AddParam("url", testurl)
	q.AddParam("knowledgeGraph", "1")
	q.AddParam("structuredEntities", "1")
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
	if len(e.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}

	e = &entity.Entities{}
	q = goalchemy.NewQuery(entity.EndpointText, apikey)
	q.AddParam("text", testtext2)
	q.AddParam("knowledgeGraph", "1")
	q.AddParam("quotations", "1")
	q.AddParam("sentiment", "1")
	q.AddParam("structuredEntities", "1")
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
	if len(e.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}
}

func TestQueryMicro(t *testing.T) {
	e := &micro.Micros{}
	q := goalchemy.NewQuery(micro.EndpointURL, apikey)
	q.AddParam("url", testurl2)
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
}

func TestRelation(t *testing.T) {
	e := &relation.Relations{}
	q := goalchemy.NewQuery(relation.EndpointURL, apikey)
	q.AddParam("url", testurl)
	q.AddParam("sentiment", "1")
	q.AddParam("keywords", "1")
	q.AddParam("entities", "1")
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
	if len(e.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}

	e = &relation.Relations{}
	q = goalchemy.NewQuery(relation.EndpointText, apikey)
	q.AddParam("text", testtext)
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
	if len(e.Results) == 0 {
		t.Error("Expected len > 1 got 0")
	}
}
*/

func TestSentiment(t *testing.T) {
	e := &sentiment.Sentiments{}
	q := goalchemy.NewQuery(sentiment.EndpointURL, apikey)
	q.AddParam("url", testurl2)
	if err := q.Run(e); err != nil {
		t.Error(err)
	}

	e = &sentiment.Sentiments{}
	q = goalchemy.NewQuery(sentiment.EndpointTargetURL, apikey)
	q.AddParam("url", testurl2)
	q.AddParam("target", tphrase)
	if err := q.Run(e); err != nil {
		t.Error(err)
	}
}
