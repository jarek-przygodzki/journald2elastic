package main

import (
	"context"
	"os"

	"github.com/jarek-przygodzki/journald2elastic/app"

	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	log.SetLevel(log.InfoLevel)

	cliApp := cli.NewApp()
	cliApp.Name = "journald2elastic"
	cliApp.Version = app.Version
	cliApp.Usage = "Ship logs exported from systemd in json format to Elasticsearch in Kibana-friendly format"

	cliApp.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "url",
			Value: "",
			Usage: "Elasticsearch URL",
		},
		cli.StringFlag{
			Name:  "file",
			Value: "",
			Usage: "Data file",
		},
		cli.StringFlag{
			Name:  "index",
			Value: "logs",
			Usage: "Elasticserach index",
		},
		cli.StringFlag{
			Name:  "type",
			Value: "log",
			Usage: "Elasticserach index",
		},
		cli.IntFlag{
			Name:  "batch-size",
			Value: 5000,
			Usage: "Elasticsearch batch size",
		},
	}

	cliApp.Action = func(c *cli.Context) error {
		url := c.String("url")
		filename := c.String("file")
		typeName := c.String("type")
		indexName := c.String("index")
		batchSize := c.Int("batch-size")

		if url == "" || filename == "" {
			cli.ShowAppHelp(c)
			return cli.NewExitError("", 1)
		}
		client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetSniff(false))
		if err != nil {
			return err
		}
		ctx := context.Background()
		log.Infof("Shipping logs from file %s to Elasticsearch at %s (batch size: %d, index: %s, type: %s)",
			filename, url, batchSize, indexName, typeName)
		return app.SaveInElasticsearch(ctx, client, filename, typeName, indexName, batchSize)
	}

	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
