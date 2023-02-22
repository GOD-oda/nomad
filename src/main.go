package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
	"reflect"
)

type Places struct {
	Places []Place `yaml:"places"`
}

type Place struct {
	Prefecture string `yaml:"prefecture"`
	Station    string `yaml:"station"`
	Name       string `yaml:"name"`
	Link       string `yaml:"link"`
	City       string `yaml:"city"`
	Note       string `yaml:"note"`
}

func main() {
	buf, err := os.ReadFile("./japan.yml")
	if err != nil {
		fmt.Println(err)
		return
	}

	places := Places{}
	err = yaml.Unmarshal(buf, &places)
	if err != nil {
		fmt.Println(err)
		return
	}

	f, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}

	var data []byte
	data = append(data, "# nomad\n\n"...)
	data = append(data, "## Japan\n\n"...)
	data = makeTable(data, places.Places)

	_, err = f.Write(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
}

func makeTable(data []byte, places []Place) []byte {
	//headerの作成
	var t Place
	refT := reflect.TypeOf(t)

	var names []string
	var hyphen []string
	for i := 0; i < refT.NumField(); i++ {
		names = append(names, refT.Field(i).Name)
		hyphen = append(hyphen, "---")
	}
	data = makeLine(data, names)
	data = makeLine(data, hyphen)

	// bodyの作成
	for _, place := range places {
		var fields []string
		for _, name := range names {
			refPlace := reflect.ValueOf(place)
			fields = append(fields, refPlace.FieldByName(name).String())
		}
		data = makeLine(data, fields)
	}

	return data
}

func makeLine(data []byte, fields []string) []byte {
	size := len(fields)

	for i := 0; i < size; i++ {
		data = append(data, "|"...)

		value := fields[i]
		if len(value) > 0 {
			data = append(data, value...)
		} else {
			data = append(data, " "...)
		}

	}
	data = append(data, "|\n"...)

	return data
}
