package repo

import (
	"fmt"
	"github.com/restream/reindexer/v3"
	_ "github.com/restream/reindexer/v3/bindings/cproto"
	"main/config"
	"main/internal/repo/structs"
	"main/logger"
)

var r *Repo

type Repo struct {
	Rx *reindexer.Reindexer
}

func Init() error {
	var userPath string
	var cfg = config.Cfg.RxCfg
	r = &Repo{}

	if len(cfg.Username) > 0 {
		userPath = fmt.Sprintf("%v:%v@", cfg.Username, cfg.Password)
	}

	connectionPath := fmt.Sprintf("cproto://%v%v:%v/%v", userPath, cfg.Host, cfg.Port, cfg.Database)
	r.Rx = reindexer.NewReindex(connectionPath, reindexer.WithCreateDBIfMissing())

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Teachers, reindexer.DefaultNamespaceOptions(), structs.Timetable{}); err != nil {
		return fmt.Errorf("error occurred while openning teachers namespace: %s", err.Error())
	}

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Groups, reindexer.DefaultNamespaceOptions(), structs.Timetable{}); err != nil {
		return fmt.Errorf("error occured while openning groups namespace:, %s", err.Error())
	}

	if err := r.Rx.OpenNamespace(cfg.Namespaces.Names, reindexer.DefaultNamespaceOptions(), structs.TeachersNames{}); err != nil {
		return fmt.Errorf("error occurred while openning teachers_names namespace:, %s", err.Error())
	}

	return nil
}

func RewriteTimetables(timetables []structs.Timetable, namespace string) error {
	logger.Log.Info().Any("Namespace", namespace).Msg("Starting transaction...")
	tx, err := r.Rx.BeginTx(namespace)
	if err != nil {
		return fmt.Errorf("an error occurred while begin transaction on '%s' namespace: %s", namespace, err.Error())
	}

	for _, timetable := range timetables {
		if err := tx.Upsert(timetable); err != nil {
			tx.Rollback()
			return fmt.Errorf("an error occurred while upserting %v to '%s' namespace: %s", timetable, namespace, err.Error())
		}
	}

	if err = tx.Commit(); err != nil {
		return fmt.Errorf("an error occurred while commiting the transaction: %s", err.Error())
	}
	logger.Log.Info().Any("Namespace", namespace).Msg("Successfully write all data")
	return nil
}

func RewriteTeachersNames(names structs.TeachersNames) error {
	namespace := config.Cfg.RxCfg.Namespaces.Names
	logger.Log.Info().Any("Namespace", namespace).Msg("Starting upsert...")

	err := r.Rx.Upsert(namespace, names)
	if err != nil {
		return fmt.Errorf("an error was occurred while upsert %v to '%s' namespace: %s", names, namespace, err.Error())
	}

	logger.Log.Info().Any("Namespace", namespace).Msg("Successfully write all data")
	return nil
}
