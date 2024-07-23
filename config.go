package main

import "flag"

type Config struct {
	InputPath            string
	CollectionOutputPath string
	FlattenLevels        bool
}

func DefineCmdlineArgs(c *Config) {
	flag.StringVar(&c.InputPath, "InputPath", ".", "Define the path to handle.")
	flag.StringVar(&c.CollectionOutputPath, "CollectionOutputPath", "./out", "")
	flag.BoolVar(&c.FlattenLevels, "FlattenLevels", false, "Defines whether to flatten ./:Collection/:Level to ./:Level")

	flag.Parse()
}
