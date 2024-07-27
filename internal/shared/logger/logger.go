package logger

import (
	"fmt"
	"messenger_service/internal/adapters/ioutil"
	"messenger_service/internal/shared/console"
	"os"
	"os/exec"
	"strconv"
	"time"
)

type Logger struct {
	date_format    string
	path_folder    string
	file_extension string
}

var loggerInfos Logger

type ILogger interface {
	Load()
	Write(message string, details any)
	getPath(now time.Time) string
	resolvePath(path string) bool
}

func (l *Logger) Load() {
	infos := &Logger{
		date_format:    os.Getenv("LOGGER_DATE_FORMAT"),
		path_folder:    os.Getenv("LOGGER_FOLDER"),
		file_extension: os.Getenv("LOGGER_EXTENSION"),
	}

	loggerInfos = *infos
	fmt.Println("LOGGER LOAD SUCESS!")
}

func (l *Logger) Write(message string, details any) {
	ioutil := &ioutil.IoUtil{}
	console := &console.Console{}

	now := time.Now()

	value := fmt.Sprintf("Time: %s\nMessage: %s\nDetails: %s\n", now, message, details)

	pathfile := l.getPath(now)

	ioutil.WriteFile(pathfile, value)

	console.Print(value)
}

func (l *Logger) getPath(now time.Time) string {
	year, month, day := now.Date()
	hour := now.Hour()

	dir, _ := os.Getwd()

	pathFolder := fmt.Sprintf("%s/%s/%s/%s/%s", dir, loggerInfos.path_folder, strconv.Itoa(year), month.String(), strconv.Itoa(day))

	cmd := exec.Command("cd", pathFolder)
	_, err := cmd.CombinedOutput()

	if err != nil {
		l.resolvePath(pathFolder)
	}

	return fmt.Sprintf("%s/%sh00.%s", pathFolder, strconv.Itoa(hour), loggerInfos.file_extension)
}

func (l *Logger) resolvePath(path string) bool {
	err := os.MkdirAll(path, 0777)
	return err == nil
}
