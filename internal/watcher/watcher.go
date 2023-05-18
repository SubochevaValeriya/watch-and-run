package watcher

import (
	"context"
	"fmt"
	"github.com/radovskyb/watcher"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"strings"
	"time"
	"watchAndRun/internal/model"
	"watchAndRun/internal/repository"
)

type Watcher struct {
	EventRepo  repository.EventRepository
	LaunchRepo repository.LaunchRepository
}

func (wat *Watcher) Watch(ctx context.Context, dir model.Directory, changeCheckFrequency time.Duration) {
	w := watcher.New()

	// Only notify create, write and remove events.
	w.FilterOps(watcher.Create, watcher.Write, watcher.Remove)

	// Watch only differences in files (not hidden):
	w.AddFilterHook(filterOnlyFiles)

	// Watch folder recursively for changes.
	if err := w.AddRecursive(dir.Path); err != nil {
		log.Fatalln(err)
	}

	go func() { //у тебя Watch и так в горутине запускается
		for {
			select {
			case <-ctx.Done():
				log.Println("Canceling")
				w.Close()
				return
			case event := <-w.Event:
				if isNeedHandle(dir, event) {
					err := wat.Handle(dir, event)
					if err != nil {
						logrus.Errorf("Can't handle event: %s", err)
					} else {
						logrus.Printf("Event was handled: %s", dir)
					}
				}
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	// Wait after watcher started.
	go func() {
		w.Wait()
	}()

	if err := w.Start(changeCheckFrequency); err != nil {
		log.Fatalln(err)
	}

}

func filterOnlyFiles(info os.FileInfo, fullPath string) error {
	if info.IsDir() {
		return watcher.ErrSkip
	}

	if strings.Index(fullPath, "\\.") != -1 {
		return watcher.ErrSkip
	}

	return nil
}

func isNeedHandle(dir model.Directory, event watcher.Event) bool {
	if len(dir.IncludeRegexp) == 0 {
		for _, regexpExcl := range dir.ExcludeRegexp {
			if regexpExcl.MatchString(event.Name()) {
				return false
			}
		}
		return true
	}

	for _, regexpIncl := range dir.IncludeRegexp {
		if regexpIncl.MatchString(event.Name()) {
			for _, regexpExcl := range dir.ExcludeRegexp {
				if regexpExcl.MatchString(event.Name()) {
					return false
				}
			}
			return true
		}
	}
	return false
}

func (wat *Watcher) Handle(dir model.Directory, event watcher.Event) error {
	currentEvent := model.Event{
		Path:      dir.Path,
		FileName:  event.Name(),
		EventType: event.Op.String(),
		Time:      time.Now(),
	}

	eventId, err := wat.EventRepo.AddEvent(currentEvent)
	if err != nil {
		return fmt.Errorf("can't save event data: %w", err)
	}
	for _, command := range dir.Commands {
		launch := model.Launch{
			Command:   command,
			StartTime: time.Now(),
			EndTime:   time.Time{},
			Result:    "",
			EventId:   int(eventId),
		}
		err := executeCommand(command, dir.LogFile)
		launch.EndTime = time.Now()
		if err == nil {
			launch.Result = "success"
			err = wat.LaunchRepo.AddLaunch(launch)
			if err != nil {
				return fmt.Errorf("can't save launch data: %w", err)
			}
		} else {
			launch.Result = "failure"
			errDB := wat.LaunchRepo.AddLaunch(launch)
			if errDB != nil {
				return fmt.Errorf("can't save launch data: %w", errDB)
			}
			return fmt.Errorf("can't execute command: %w", err)
		}
	}
	return nil
}
