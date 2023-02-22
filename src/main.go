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
	f, err := os.Create("README.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	var data []byte
	data = append(data, "# nomad\n\n"...)

	japanPlaces, err := readPlaces("./japan.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	data = append(data, "## Japan\n\n"...)
	data = makeTable(data, japanPlaces.Places)

	worldPlaces, err := readPlaces("./world.yml")
	if err != nil {
		fmt.Println(err)
		return
	}
	data = append(data, "\n## World\n\n"...)
	data = makeTable(data, worldPlaces.Places)

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

func readPlaces(name string) (Places, error) {
	places := Places{}

	buf, err := os.ReadFile(name)
	if err != nil {
		return places, err
	}

	if err = yaml.Unmarshal(buf, &places); err != nil {
		fmt.Println(err)
		return places, nil
	}

	return places, nil
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
	data = makeTableLine(data, names)
	data = makeTableLine(data, hyphen)

	// bodyの作成
	for _, place := range places {
		var fields []string
		for _, name := range names {
			refPlace := reflect.ValueOf(place)
			fields = append(fields, refPlace.FieldByName(name).String())
		}
		data = makeTableLine(data, fields)
	}

	return data
}

func makeTableLine(data []byte, fields []string) []byte {
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
