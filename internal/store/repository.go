package store

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	cmap "github.com/fatihsimsek/go-case-study/pkg/concurrentmap"
	"github.com/fatihsimsek/go-case-study/pkg/util"
)

type Repository interface {
	Get(key string) (string, bool)
	Put(key string, value string)
	Remove(key string)
	Flush() error
	Init() error
}

type repository struct {
	store *cmap.ConcurrentMap
}

func NewRepository() Repository {
	return repository{store: cmap.New()}
}

func (r repository) Get(key string) (string, bool) {
	value, found := r.store.Get(key)
	if found {
		return fmt.Sprintf("%v", value), true
	} else {
		return "", false
	}
}

func (r repository) Put(key string, value string) {
	r.store.Put(key, value)
}

func (r repository) Remove(key string) {
	r.store.Remove(key)
}

func (r repository) Flush() error {
	filename := getFilename()

	copyMap := r.store.CopyTo()
	newMap := make(map[string]string)
	for k, v := range copyMap {
		key := fmt.Sprintf("%v", k)
		value := fmt.Sprintf("%v", v)
		newMap[key] = value
	}

	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if file != nil {
		defer file.Close()
	}
	if err != nil {
		log.Printf("Flush open file error: %v", err)
		return err
	}

	jsonData, err := json.Marshal(newMap)
	if err != nil {
		log.Printf("Flush json error: %v", err)
		return err
	}

	if _, err := file.Write(jsonData); err != nil {
		log.Printf("Flush write file error: %v", err)
		return err
	}
	return nil
}

func (r repository) Init() error {
	filename := getFilename()
	if util.FileIsExists(filename) {
		file, err := os.Open(filename)
		if file != nil {
			defer file.Close()
		}
		if err != nil {
			log.Printf("Init open file error: %v", err)
			return err
		}
		byteValue, _ := ioutil.ReadAll(file)

		var dataMap map[string]string
		err = json.Unmarshal([]byte(byteValue), &dataMap)
		if err != nil {
			log.Printf("Init unmarshall error: %v", err)
			return err
		}
		for k, v := range dataMap {
			r.store.Put(k, v)
		}
	}
	return nil
}

func getFilename() string {
	value, exists := os.LookupEnv("FILENAME")
	if !exists {
		value = "temp-store.json"
	}
	return value
}
