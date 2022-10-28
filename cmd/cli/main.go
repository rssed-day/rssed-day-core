package main

import "github.com/rssed-day/rssed-day-core/cli"

func main() {
	if err := cli.Execute(); err != nil {
		panic(err)
	}
}
