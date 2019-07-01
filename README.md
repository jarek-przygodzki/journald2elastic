# journald2elastic

Ship logs exported from systemd in json format to Elasticsearch in Kibana-friendly format

## Installing

```
go get -u github.com/jarek-przygodzki/journald2elastic
go install github.com/jarek-przygodzki/journald2elastic
```

## Running
```
journalctl -u docker --since -7d --output=json > logs.log
journald2elastic --url [ELASTICSEARCH_API] --file ./logs.log 

```

##  How this tool came to be
Some time ago, I was investigating sporadic healtcheck failures in Docker Swarm (which turned out to be caused by [this runc issue](https://github.com/opencontainers/runc/issues/1884)). The first problem I ran into was the fact that I needed to collect systemd service logs from many hosts in one place (there was no centralized logging system in place to aggregate such logs for review). This tool was created to import logs exported from systemd/journald to Elasticsearch.