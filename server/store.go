package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

func Load() (Data, error) {
	b, err := ioutil.ReadFile("data.json")
	if err != nil {
		return createCleanDataSet()
	}
	var data Data
	err = json.Unmarshal(b, &data)
	if err != nil {
		log.Printf("could not load store: %v\n", err)
		return Data{}, err
	}
	return data, nil
}

func Save(data Data) error {
	b, err := json.Marshal(data)
	if err != nil {
		log.Printf("could not save data: %v\n", err)
		return err
	}
	err = ioutil.WriteFile("data.json", b, os.ModePerm)
	if err != nil {
		log.Printf("could not write data.json: %v\n", err)
		return err
	}
	return nil
}

func createCleanDataSet() (Data, error) {
	var data Data
	cleanData, err := json.Marshal(data)
	if err != nil {
		log.Printf("could not create clean data.json file: %v", err)
		return Data{}, err
	}
	err = ioutil.WriteFile("data.json", cleanData, os.ModePerm)
	if err != nil {
		log.Printf("could not write data.json: %v\n", err)
		return Data{}, err
	}
	return data, nil
}
