package main

import (
	"encoding/json"
	// "fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

var C Config

func init() {
	name := strings.TrimSuffix(os.Args[0], path.Ext(os.Args[0])) + ".json"
	f, err := os.Open(name)
	if err != nil {
		log.Fatalf("Ack!! looks like I couldn't open the config file called %v.\nThe system says %v\n", name, err)
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatalf("Ack!! looks like I couldn't read the config file called %v.\nThe system says %v\n", name, err)
	}

	err = json.Unmarshal(data, &C)
	if err != nil {
		var ec Config
		ex, _ := json.MarshalIndent(ec, "", "  ")
		log.Fatalf("Ack!! %v can't be unmarshalled.\nThe system says %v\nConsider making it look more like\n\n%s\n", name, err, ex)
	}

}
