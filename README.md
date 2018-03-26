# redisTopK
A tool to find redis top n keys by key serialized length

## brief

This tool based on scan command, so when run against a redis which key change from time to time ,the result may be inaccurate.
It is recommand to run against a static redis server.

## todo
Currently, only db0 is considered.So I will add other db support later

## build
setup correct GOPATH, run `go build` at root path of the prokect

## run
Usage of ./redisTopK: -h
