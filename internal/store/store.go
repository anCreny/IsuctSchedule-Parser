package store

import (
	"encoding/json"
	"fmt"
	"main/config"
	"main/logger"
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
	logger.Log.Info().Any("URL", s.Url).Msg("Sending GET request...")
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

	logger.Log.Info().Any("URL", s.Url).Msg("Request was successfully responded")
	return ans, nil
}
