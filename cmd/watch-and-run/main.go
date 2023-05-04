package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"sync"
	"syscall"
	worker "watchAndRun/internal/app/watch-and-run"
	"watchAndRun/internal/app/watch-and-run/repository"
)

func main() {

	path, err := os.Getwd()
	fmt.Println(path)
	logrus.SetFormatter(new(logrus.JSONFormatter))
	logrus.Println("Reading configs")
	config, err := parseConfig("./configs/config.yaml")
	if err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	wg := sync.WaitGroup{}

	db, err := repository.NewPostgresDB(repository.Config{
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

	dbTables := repository.DbTables{EventTable: config.DBTables.Event,
		LaunchTable: config.DBTables.Launch}

	repos := repository.NewRepository(db, dbTables)
	service := worker.NewService(repos)
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)

	for i, path := range config.PathAndCommands {
		wg.Add(1)
		fmt.Println(i)
		go func(i PathAndCommands) {
			defer wg.Done()
			service.Watch(ctx, implementDirectoryStructure(path), config.ChangeCheckFrequency)
		}(path)
	}

	// graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Print("App Shutting Down")

	if err := db.Close(); err != nil {
		logrus.Errorf("error occured on db connection close: %s", err.Error())
	}

	wg.Wait()
}

//db, err := repository.NewPostgresDB(repository.Config{
//	Host:     os.Getenv("host"),
//	Port:     config.DBConfig.Port,
//	Username: config.DBConfig.Username,
//	Password: os.Getenv("DB_PASSWORD"),
//	DBName:   config.DBConfig.DBName,
//	SSLMode:  config.DBConfig.SSLMode,
//})
//if err != nil {
//	logrus.Fatalf("failed to inititalize db: %s", err.Error())
//}
//
//dbTables := repository.DbTables{EventTable: config.DBTables.Event,
//	LaunchTable: config.DBTables.Launch}

//repos := repository.NewRepository(db, dbTables)

////app := "echo"
////
////arg0 := "-e"
////arg1 := "Hello world"
////arg2 := "\n\tfrom"
////arg3 := "golang"
//
//cmd := exec.Command("cmd", "/c", "echo %PROCESSOR_ARCHITECTURE%", "hehe")
////cmd.Run()
//stdout, err := cmd.Output()
//if err != nil {
//	fmt.Println(err.Error())
//	return
//}
//// Print the output
//fmt.Println(string(stdout))

//
//func main() {
//	// Create new watcher.
//	watcher, err := fsnotify.NewWatcher()
//	if err != nil {
//		log.Fatal(err)
//	}
//	defer watcher.Close()
//
//	// Start listening for events.
//	go func() {
//		for {
//			select {
//			case event, ok := <-watcher.Events:
//				if !ok {
//					return
//				}
//				log.Println("event:", event)
//				if event.Has(fsnotify.Write) {
//					log.Println("modified file:", event.Name)
//				}
//			case err, ok := <-watcher.Errors:
//				if !ok {
//					return
//				}
//				log.Println("error:", err)
//			}
//		}
//	}()
//
//	// Add a path.
//	err = watcher.Add("\\Users\\MSI-PC\\GolandProjects\\watch-and-run\\a")
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println("added")
//
//	// Block main goroutine forever.
//	<-make(chan struct{})
//
//	//psth cmds
//}
