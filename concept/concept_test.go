package concept

import "testing"

const tjson = `
{
                    "status": "OK",
                    "usage": "By accessing AlchemyAPI or using information generated by AlchemyAPI, you are agreeing to be bound by the AlchemyAPI Terms of Use: http://www.alchemyapi.com/company/terms.html",
                    "url": "http://www.desmoinesregister.com/story/news/crime-and-courts/2015/07/14/van-meter-taser-fundraiser-modified/30129893/?hootPostID=%5B%22%5B%27e34bfcfff7cf7090f7c60c8ecd6167a8%27%5D%22%5D",
                    "totalTransactions": "2",
                    "language": "english",
                    "concepts": [
                        {
                            "text": "Taser",
                            "relevance": "0.966182",
                            "knowledgeGraph": {
                                "typeHierarchy": "/products/stun guns/taser"
                            },
                            "dbpedia": "http://dbpedia.org/resource/Taser",
                            "freebase": "http://rdf.freebase.com/ns/m.01jhdv",
                            "yago": "http://yago-knowledge.org/resource/Taser"
                        },
                        {
                            "text": "Police",
                            "relevance": "0.682894",
                            "knowledgeGraph": {
                                "typeHierarchy": "/organizations/agencies/police"
                            },
                            "dbpedia": "http://dbpedia.org/resource/Police",
                            "freebase": "http://rdf.freebase.com/ns/m.05ws7",
                            "opencyc": "http://sw.opencyc.org/concept/Mx4rP7EFZvz0QdiYzKtIVkuYpA"
                        },
                        {
                            "text": "Electroshock weapon",
                            "relevance": "0.565837",
                            "knowledgeGraph": {
                                "typeHierarchy": "/treatments/electroshock/electroshock weapon"
                            },
                            "dbpedia": "http://dbpedia.org/resource/Electroshock_weapon",
                            "freebase": "http://rdf.freebase.com/ns/m.02b92_",
                            "yago": "http://yago-knowledge.org/resource/Electroshock_weapon"
                        },
                        {
                            "text": "Less-lethal weapon",
                            "relevance": "0.49084",
                            "dbpedia": "http://dbpedia.org/resource/Less-lethal_weapon",
                            "opencyc": "http://sw.opencyc.org/concept/Mx4rPAw22r1MQdeYT6OTREPckw"
                        },
                        {
                            "text": "Police car",
                            "relevance": "0.455486",
                            "knowledgeGraph": {
                                "typeHierarchy": "/vehicles/emergency vehicles/police car"
                            },
                            "dbpedia": "http://dbpedia.org/resource/Police_car",
                            "freebase": "http://rdf.freebase.com/ns/m.04qvtq",
                            "opencyc": "http://sw.opencyc.org/concept/Mx4rR1bJJruqQdiN7_Ppq_BL9Q"
                        },
                        {
                            "text": "Dallas County, Iowa",
                            "relevance": "0.444278",
                            "knowledgeGraph": {
                                "typeHierarchy": "/places/cities/dallas/dallas county, iowa"
                            },
                            "geo": "41.68277777777778 -94.035",
                            "website": "http://www.co.dallas.ia.us",
                            "dbpedia": "http://dbpedia.org/resource/Dallas_County,_Iowa",
                            "freebase": "http://rdf.freebase.com/ns/m.0nsbv",
                            "census": "http://www.rdfabout.com/rdf/usgov/geo/us/ia/counties/dallas_county",
                            "yago": "http://yago-knowledge.org/resource/Dallas_County,_Iowa",
                            "geonames": "http://sws.geonames.org/4853335/"
                        },
                        {
                            "text": "According to Jim",
                            "relevance": "0.42305",
                            "knowledgeGraph": {
                                "typeHierarchy": "/family members/friends/jim/according to jim"
                            },
                            "dbpedia": "http://dbpedia.org/resource/According_to_Jim",
                            "freebase": "http://rdf.freebase.com/ns/m.017v5x",
                            "yago": "http://yago-knowledge.org/resource/According_to_Jim"
                        },
                        {
                            "text": "Official",
                            "relevance": "0.418829",
                            "knowledgeGraph": {
                                "typeHierarchy": "/people/official"
                            },
                            "dbpedia": "http://dbpedia.org/resource/Official",
                            "freebase": "http://rdf.freebase.com/ns/m.035y33",
                            "opencyc": "http://sw.opencyc.org/concept/Mx4rwQryspwpEbGdrcN5Y29ycA"
                        }
                    ]
                }
`

func TestConcept(t *testing.T) {
	c := Concept{}
	if err := c.Decode([]byte(tjson)); err != nil {
		t.Error(err)
	}
	if len(c.Results) != 8 {
		t.Errorf("Expected 8 results got %v\n", len(c.Results))
	}
}
