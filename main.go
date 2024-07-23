package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var config Config

type Level struct {
	Name string
	Path string
}

type Collection struct {
	Name   string
	Levels []Level
}

func main() {
	DefineCmdlineArgs(&config)

	fmt.Println("WorkDir: " + config.InputPath)

	var Collections []Collection

	FS_collections, err := os.ReadDir(config.InputPath)
	if err != nil {
		panic(err.Error())
	}
	for _, FS_collection := range FS_collections {
		if !FS_collection.IsDir() {
			continue
		}
		FS_levels, err := os.ReadDir(config.InputPath + "/" + FS_collection.Name())
		if err != nil {
			panic(err.Error())
		}
		IsMaidataDir := false
		Collections = append(Collections, Collection{Name: FS_collection.Name(), Levels: []Level{}})
		for _, FS_level := range FS_levels {
			if IsPathExist(config.InputPath + "/" + FS_collection.Name() + "/" + FS_level.Name() + "/" + "maidata.txt") {
				IsMaidataDir = true
				Collections[len(Collections)-1].Levels = append(Collections[len(Collections)-1].Levels, Level{Name: FS_level.Name(), Path: config.InputPath + "/" + FS_collection.Name() + "/" + FS_level.Name()})
			}
		}
		if !IsMaidataDir {
			Collections = Collections[:len(Collections)-1]
		}
	}

	// Output manifest.json

	err = os.MkdirAll(config.CollectionOutputPath+"/"+"collections", 0755)
	if err != nil {
		fmt.Println(err.Error())
	}
	if !IsPathExist(config.CollectionOutputPath + "/" + "collections") {
		panic(config.CollectionOutputPath + "/" + "collections" + " does not exist.")
	}
	for _, collection := range Collections {
		err := os.MkdirAll(config.CollectionOutputPath+"/"+"collections"+"/"+collection.Name, 0755)
		if err != nil {
			fmt.Println(err.Error())
		}
		if !IsPathExist(config.CollectionOutputPath + "/" + "collections" + "/" + collection.Name) {
			panic(config.CollectionOutputPath + "/" + "collections" + "/" + collection.Name + " does not exist.")
		}
		var levelIds []string
		for _, Level := range collection.Levels {
			levelIds = append(levelIds, Level.Name)
		}
		var manifest_json struct {
			Name     string   `json:"name"`
			LevelIds []string `json:"levelIds"`
		} = struct {
			Name     string   `json:"name"`
			LevelIds []string `json:"levelIds"`
		}{Name: collection.Name, LevelIds: levelIds}

		manifest_json_byte, err := json.Marshal(manifest_json)
		if err != nil {
			panic(err.Error())
		}
		if err := os.WriteFile(config.CollectionOutputPath+"/"+"collections"+"/"+collection.Name+"/"+"manifest.json", manifest_json_byte, 0666); err != nil {
			panic(err.Error())
		}

	}
	fmt.Println("Done.")
}
