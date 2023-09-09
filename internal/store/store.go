package store

import (
	"encoding/json"
	"fmt"
	"log"
	"main/config"
	"net/http"
)

var s *Store

type Store struct {
	Url string
}

func Init() error {
	s = &Store{Url: config.Cfg.ParseUrl}
	return nil
}

func GetScheduleFromApi() (ScheduleFile, error) {
	log.Printf("Send GET request to %s\n", s.Url)
	resp, err := http.Get(s.Url)
	if err != nil {
		return ScheduleFile{}, fmt.Errorf("error with response from %s: %s", s.Url, err.Error())
	}

	defer resp.Body.Close()

	var ans ScheduleFile

	err = json.NewDecoder(resp.Body).Decode(&ans)
	if err != nil {
		return ScheduleFile{}, fmt.Errorf("error with unmarshal: %s", err.Error())
	}

	log.Printf("Catch a result from GET request to %s\n", s.Url)
	return ans, nil
}
