package watcher

import (
	"context"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"path"
	"testing"
	"time"
	mocks2 "watchAndRun/internal/mocks"
	"watchAndRun/internal/model"
)

func TestApiService_Handler(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	eventRepo := mocks2.NewMockEvent(ctrl)
	launchRepo := mocks2.NewMockLaunch(ctrl)

	wat := Watcher{
		EventRepo:  eventRepo,
		LaunchRepo: launchRepo,
	}

	eventRepo.EXPECT().
		AddEvent(gomock.Any()).
		Return(int64(0), nil).Times(1)

	launchRepo.EXPECT().
		AddLaunch(gomock.Any()).
		Return(nil).Times(1)

	tempDir := t.TempDir()
	watchedDir := path.Join(tempDir, "watched_dir")
	logFilePath := path.Join(tempDir, "logfile")
	_ = os.MkdirAll(watchedDir, 0777)
	go func() {
		wat.Watch(context.Background(), model.Directory{
			Path:     watchedDir,
			Commands: []string{"echo -n Hello"},
			LogFile:  logFilePath,
		}, time.Second/2)
	}()
	time.Sleep(time.Second)
	//trigger event
	err := os.WriteFile(path.Join(watchedDir, "file_1"), []byte(`hello`), 0777)
	require.NoError(t, err)
	//check
	time.Sleep(time.Second)
	str, err := os.ReadFile(logFilePath)
	require.NoError(t, err)
	assert.Equal(t, "Hello", string(str))
}
