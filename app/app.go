package app

import (
	"bufio"
	"compress/gzip"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/olivere/elastic"
	log "github.com/sirupsen/logrus"
)

func createDocument(jsonLine string) map[string]interface{} {
	var anyJSON map[string]interface{}
	json.Unmarshal([]byte(jsonLine), &anyJSON)
	tsEpochMillis, _ := strconv.ParseInt(anyJSON["__REALTIME_TIMESTAMP"].(string), 10, 64)
	ts := timeFromEpochMicroseconds(tsEpochMillis)
	anyJSON["@timestamp"] = ts.Format("2006-01-02T15:04:05.000000Z07:00")
	doc := make(map[string]interface{})
	// strip '_' field prefix, fields starting with underscore doesn't show in Kibana
	for k, v := range anyJSON {
		for strings.HasPrefix(k, "_") {
			k = k[1:]
		}
		doc[k] = v
	}

	return doc
}

func timeFromEpochMicroseconds(msec int64) time.Time {
	sec := msec / 1000000
	nsec := msec - sec*1000000
	return time.Unix(sec, nsec)
}

func saveInBulk(client *elastic.Client, typeName string, indexName string, docs []map[string]interface{}) (*elastic.BulkResponse, error) {
	bulk := client.
		Bulk().
		Index(indexName).
		Type(typeName)

	for _, d := range docs {
		docBytes, _ := json.Marshal(d)
		docHash := sha256.Sum256(docBytes)
		docID := hex.EncodeToString(docHash[:])
		bulk.Add(elastic.NewBulkIndexRequest().Doc(string(docBytes)).Id(docID))
	}
	return bulk.Do(context.Background())
}

const mapping = `
{
	"mappings": {
	 "log": {
	  "properties": {
	   "PRIORITY": {
		"type": "keyword"
	   },
	   "STREAM_ID": {
		"type": "keyword"
	   },
	   "PID": {
		"type": "keyword"
	   },
	   "SYSTEMD_UNIT": {
		"type": "keyword"
	   },
	   "UID": {
		"type": "keyword"
	   },
	   "SYSLOG_FACILITY": {
		"type": "keyword"
	   },
	   "HOSTNAME": {
		"type": "keyword"
	   },
	   "CURSOR": {
		"type": "keyword"
	   },
	   "SYSTEMD_CGROUP": {
		"type": "keyword"
	   },
	   "MONOTONIC_TIMESTAMP": {
		"type": "keyword"
	   },
	   "EXE": {
		"type": "keyword"
	   },
	   "BOOT_ID": {
		"type": "keyword"
	   },
	   "SELINUX_CONTEXT": {
		"type": "keyword"
	   },
	   "SYSLOG_IDENTIFIER": {
		"type": "keyword"
	   },
	   "SYSTEMD_SLICE": {
		"type": "keyword"
	   },
	   "CMDLINE": {
		"type": "keyword"
	   },
	   "TRANSPORT": {
		"type": "keyword"
	   },
	   "REALTIME_TIMESTAMP": {
		"type": "keyword"
	   },
	   "MACHINE_ID": {
		"type": "keyword"
	   },
	   "CAP_EFFECTIVE": {
		"type": "keyword"
	   },
	   "GID": {
		"type": "keyword"
	   },
	   "COMM": {
		"type": "keyword"
	   },
	   "@timestamp": {
		"type": "date"
	   },
	   "file": {
		"type": "keyword"
	   }
	  }
	 }
	}
   }
`

func hasErrors(indexed []*elastic.BulkResponseItem) bool {
	for _, i := range indexed {
		if i.Error != nil {
			return true
		}
	}
	return false
}

func SaveInElasticsearch(ctx context.Context, client *elastic.Client, filename string, typeName string, indexName string, batchSize int) error {
	createIndex, err := client.CreateIndex(indexName).BodyString(mapping).Do(ctx)
	if err != nil {
		err, _ := err.(*elastic.Error)
		if err.Details.Type != "resource_already_exists_exception" {
			log.Fatalf("Failed to create index %v", err)
		}

	} else {
		if !createIndex.Acknowledged {
			log.Fatalf("Failed to create index %v", err)
		}
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var reader io.Reader
	if strings.HasSuffix(filename, ".gz") {
		reader, err = gzip.NewReader(file)
	} else {
		reader = bufio.NewReader(file)
	}

	if err != nil {
		log.Fatal(err)
	}

	batches := NewBatchIterator(batchSize, bufio.NewScanner(reader))

	for batches.Next() {
		jsonLineBatch := batches.Value()
		docs := make([]map[string]interface{}, 0)
		for _, jsonLine := range jsonLineBatch {
			doc := createDocument(jsonLine)
			docs = append(docs, doc)
		}
		for _, d := range docs {
			d["file"] = filename
		}

		bulkResponse, err := saveInBulk(client, typeName, indexName, docs)
		indexed := bulkResponse.Indexed()
		if hasErrors(indexed) || err != nil {
			log.Panic("elasticsearch write failed", err)
		}
	}

	if batches.Err() != nil {
		log.Fatalf("error: %s\n", batches.Err())
	}

	return nil
}
