package main

import "flag"

func main() {

	var cacheTTLFlag = flag.Int("CacheTTL", 60, "Cache Time To Live in seconds")

	flag.Parse()

	startRepl(*cacheTTLFlag)
}