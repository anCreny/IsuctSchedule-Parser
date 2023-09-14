package service

import (
	"fmt"
	"github.com/anCreny/IsuctSchedule-Packages/logger"
	"github.com/anCreny/IsuctSchedule-Packages/structs"
	"main/config"
	"main/internal/repo"
	"main/internal/store"
	"sync"
	"time"
)

var s *Service

type Service struct {
}

func Init() error {
	s = &Service{}
	return nil
}

func Start() error {
	logger.Log.Info().Msg("Parser successfully started")

	var errorOccurred bool

	for {
		if errorOccurred {
			errorOccurred = false
			time.Sleep(1 * time.Hour)
		}
		logger.Log.Info().Msg("Starting update database...")

		timetableFile, err := store.GetScheduleFromApi()
		if err != nil {
			logger.Log.Error().Err(err).Time("Next try", time.Now().Add(4*time.Hour)).Msg("Aborting database update")
			errorOccurred = true
			continue
		}

		parseWG := &sync.WaitGroup{}
		parseWG.Add(3)

		var groups []structs.Timetable
		go func() {
			defer parseWG.Done()
			groups = parseGroups(timetableFile)
		}()

		var teacherNames structs.TeachersNames
		go func() {
			defer parseWG.Done()
			teacherNames = parseTeachersNames(timetableFile)
		}()

		var teachers []structs.Timetable
		go func() {
			defer parseWG.Done()
			teachers = parseTeachers(timetableFile)
		}()

		parseWG.Wait()

		repoWG := &sync.WaitGroup{}
		repoWG.Add(3)

		var (
			groupWriteError         error
			teachersWriteError      error
			teachersNamesWriteError error
		)
		go func() {
			defer repoWG.Done()
			if err := repo.RewriteTimetables(groups, config.Cfg.RxCfg.Namespaces.Groups); err != nil {
				groupWriteError = fmt.Errorf("faild to rewrite groups in '%s' namespace, error was occurred: %s", config.Cfg.RxCfg.Namespaces.Groups, err.Error())
			}
		}()

		go func() {
			defer repoWG.Done()
			if err := repo.RewriteTeachersNames(teacherNames); err != nil {
				teachersNamesWriteError = fmt.Errorf("faild to rewrite teachers names, error was occurred: %s", err.Error())
			}
		}()

		go func() {
			defer repoWG.Done()
			if err := repo.RewriteTimetables(teachers, config.Cfg.RxCfg.Namespaces.Teachers); err != nil {
				teachersWriteError = fmt.Errorf("faild to rewrite teachers in '%s' namespace, error was occurred: %s", config.Cfg.RxCfg.Namespaces.Teachers, err.Error())
			}
		}()

		repoWG.Wait()

		if groupWriteError != nil {
			return groupWriteError
		}

		if teachersNamesWriteError != nil {
			return teachersNamesWriteError
		}

		if teachersWriteError != nil {
			return teachersWriteError
		}

		logger.Log.Info().Msg("Database successfully updated!")

		//my hands are in hell
		for time.Now().Format(time.TimeOnly) != "00:09:00" {
		}
	}
}
