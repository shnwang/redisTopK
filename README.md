# redisTopK
A tool to find redis top n keys by key serialized length

## brief

Keys from all db are considered.
This tool based on scan command, so when run against a redis which key change from time to time ,the result may be inaccurate.
It is recommand to run against a static redis server.

## todo
* enable db specified from cmdline
* enable key type specified from cmdline

## build
setup correct GOPATH, run `go build` at root path of the project

## run
Usage of ./redisTopK: -h
