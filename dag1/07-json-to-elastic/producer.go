package main

import (
	"context"
	"fmt"
	"strings"

	elasticsearch "github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
	"github.com/sirupsen/logrus"
)

type ElasticPublisher struct {
	client *elasticsearch.Client
	url    string
	index  string
}

func (p *ElasticPublisher) Connect(URL string, index string) error {
	logrus.Infof("Connecting to elastic server on %s", URL)

	cfg := elasticsearch.Config{
		Addresses: []string{
			URL,
		},
	}

	es, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	res, err := es.Info()
	if err != nil {
		return err
	}

	logrus.Info("ES Info:%#v", es)

	// Check response status
	if res.IsError() {
		return fmt.Errorf("error: %s", res.String())
	}

	p.client = es
	p.url = URL
	p.index = index

	return nil
}

// Publish to elastic index
func (p *ElasticPublisher) Publish(index, payload string) error {

	req := esapi.IndexRequest{
		Index:   index,
		Body:    strings.NewReader(payload),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), p.client)
	if err != nil {
		return err
	}
	logrus.Debugf("Result : %+v\n", res)

	res.Body.Close()

	return nil

}
