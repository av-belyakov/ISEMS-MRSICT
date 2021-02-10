package modulelogginginformationerrors

import (
	"bufio"
	"compress/gzip"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

//commonInformationLoggingFiles содержит общую информацию о логирующих файлах
// locationLogDirectory - путь по которому находится основная директория для хранения лог-файлов приложения
// nameLogDirectory - название директории в которой хранятся лог-файлы приложения
// maxSizeLogFile - максимальный размер лог-файла (в БАЙТАХ), при превышении которого выполняется архивация текущего файла и создание нового
// fileNameType - список типов файлов
// fileDescriptor - список дескрипторов файлов
// chanWriteMessage - канад для записи информационных сообщений и ошибок
type commonInformationLoggingFiles struct {
	locationLogDirectory string
	nameLogDirectory     string
	maxSizeLogFile       int64
	fileNameType         map[string]string
	fileDescriptor       map[string]*os.File
}

//MainHandlerLoggingParameters основные параметры для конструктора mainHandlerLogging
// LocationLogDirectory - путь по которому находится основная директория для хранения лог-файлов приложения
// NameLogDirectory - название директории в которой хранятся лог-файлы приложения
// MaxSizeLogFile - максимальный размер лог-файла (в Мб), при превышении которого выполняется архивация текущего файла и создание нового
type MainHandlerLoggingParameters struct {
	LocationLogDirectory string
	NameLogDirectory     string
	MaxSizeLogFile       int
}

//LogMessageType описание типа для записи логов
type LogMessageType struct {
	TypeMessage, Description, FuncName string
}

//New конструктор для огранизации записи лог-файлов
func New(mhltp *MainHandlerLoggingParameters) (chan LogMessageType, error) {
	chanLogMessage := make(chan LogMessageType)
	cilf := commonInformationLoggingFiles{
		locationLogDirectory: mhltp.LocationLogDirectory,
		nameLogDirectory:     mhltp.NameLogDirectory,
		maxSizeLogFile:       int64(mhltp.MaxSizeLogFile * 1000000),
		fileNameType: map[string]string{
			"error":    "error_message.log",
			"info":     "info_message.log",
			"requests": "api_client_requests.log",
		},
		fileDescriptor: make(map[string]*os.File),
	}

	if err := cilf.createLogsDirectory(); err != nil {
		return chanLogMessage, err
	}

	for n := range cilf.fileNameType {
		fd, err := os.OpenFile(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory, cilf.fileNameType[n]), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			return chanLogMessage, err
		}
		cilf.fileDescriptor[n] = fd
	}

	go func() {
		defer func() {
			for n := range cilf.fileDescriptor {
				cilf.fileDescriptor[n].Close()
			}
		}()

		for msg := range chanLogMessage {
			typeMessage := msg.TypeMessage
			if typeMessage == "" && msg.Description == "" {
				continue
			}

			if typeMessage == "" {
				typeMessage = "error"
			}

			cilf.writeMessage(&msg)

			fi, _ := cilf.fileDescriptor[msg.TypeMessage].Stat()
			if fi.Size() > cilf.maxSizeLogFile {
				cilf.fileDescriptor[msg.TypeMessage].Close()

				cilf.compressFile(msg.TypeMessage)

				delete(cilf.fileDescriptor, msg.TypeMessage)
				_ = os.Remove(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory, cilf.fileNameType[msg.TypeMessage]))

				fd, _ := os.OpenFile(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory, cilf.fileNameType[msg.TypeMessage]), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
				cilf.fileDescriptor[msg.TypeMessage] = fd
			}
		}
	}()

	return chanLogMessage, nil
}

func (cilf *commonInformationLoggingFiles) createLogsDirectory() error {
	files, err := ioutil.ReadDir(cilf.locationLogDirectory)
	if err != nil {
		return err
	}

	for _, fl := range files {
		if fl.Name() == cilf.nameLogDirectory {
			return nil
		}
	}

	err = os.Mkdir(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory), 0777)
	if err != nil {
		return err
	}

	return nil
}

func (cilf *commonInformationLoggingFiles) writeMessage(lmt *LogMessageType) {
	var err error

	timeNowString := time.Now().String()
	tns := strings.Split(timeNowString, " ")
	strMsg := fmt.Sprintf("%s %s [%s %s] - %s ('%s')\n", tns[0], tns[1], tns[2], tns[3], lmt.Description, lmt.FuncName)

	fd := cilf.fileDescriptor[lmt.TypeMessage]
	writer := bufio.NewWriter(fd)
	defer func() {
		if err == nil {
			err = writer.Flush()
		}
	}()

	if _, err = writer.WriteString(strMsg); err != nil {
		log.Printf("func 'writeMessage' ERROR: '%v'\n", err)
	}
}

func (cilf *commonInformationLoggingFiles) compressFile(tm string) {
	timeNowUnix := time.Now().Unix()
	fn := strconv.FormatInt(timeNowUnix, 10) + "_" + strings.Replace(cilf.fileNameType[tm], ".log", ".gz", -1)

	fileIn, err := os.Create(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory, fn))
	if err != nil {
		return
	}
	defer fileIn.Close()

	zw := gzip.NewWriter(fileIn)
	zw.Name = fn

	fileOut, err := ioutil.ReadFile(path.Join(cilf.locationLogDirectory, cilf.nameLogDirectory, cilf.fileNameType[tm]))
	if err != nil {
		return
	}

	if _, err := zw.Write(fileOut); err != nil {
		return
	}

	_ = zw.Close()
}
