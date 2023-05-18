package main

import (
	"context"
	_ "database/sql"
	gotoenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"watchAndRun/configs"
	"watchAndRun/internal/repository"
	repository2 "watchAndRun/internal/repository"
	watcher "watchAndRun/internal/watcher"
)

func main() {

	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Println("Reading configs")
	config, err := configs.ParseConfig("./configs/config.yaml")
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := gotoenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository2.NewPostgresDB(repository2.Config{
		Host:     os.Getenv("host"),
		Port:     config.DBConfig.Port,
		Username: config.DBConfig.Username,
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   config.DBConfig.DBName,
		SSLMode:  config.DBConfig.SSLMode,
	})
	if err != nil {
		logrus.Fatalf("failed to inititalize db: %s", err.Error())
	}

	eventRepo := repository.NewEventRepository(db)
	launchRepo := repository.NewLaunchRepository(db)
	watcher := watcher.Watcher{EventRepo: *eventRepo, LaunchRepo: *launchRepo}
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	wg := sync.WaitGroup{}
	for _, path := range config.PathAndCommands {
		wg.Add(1)
		go func(path configs.PathAndCommands) {
			defer wg.Done()
			watcher.Watch(ctx, configs.ImplementDirectoryStructure(path), config.ChangeCheckFrequency)
		}(path)
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("App Shutting Down")

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	wg.Wait()
	logrus.Println("Finished")
}
