package commonlibs

import (
	"crypto/md5"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"os"
	"strconv"
	"time"
)

//TypePatternUserMessage тип для функции PatternUserMessage
// Section - секция в которой выполнялась задача
// TaskType - тип выполняемой задачи задачи (например, добавление, поиск, удаление и т.д.)
// FinalResult - финальный результат выполненного, над задачей, действия
// Message - информационное сообщение
type PatternUserMessageType struct {
	Section, TaskType, FinalResult, Message string
}

//PatternUserMessage шаблон для сообщения пользователю
func PatternUserMessage(tpum *PatternUserMessageType) string {
	var (
		section     = "?"
		taskType    = "?"
		finalResult = "останов"
	)

	if tpum.Section != "" {
		section = tpum.Section
	}

	if tpum.TaskType != "" {
		taskType = tpum.TaskType
	}

	if tpum.FinalResult != "" {
		finalResult = tpum.FinalResult
	}

	var message string
	if tpum.Message != "" {
		message = fmt.Sprintf(" Сообщение: '%v'", tpum.Message)
	}

	return fmt.Sprintf("Секция: '%v'. Тип: '%v'. Действие: '%v'.%v", section, taskType, finalResult, message)
}

//GetUniqIDFormatMD5 генерирует уникальный идентификатор в формате md5
func GetUniqIDFormatMD5(str string) string {
	rand.Seed(82)
	salt := fmt.Sprint(rand.Intn(10000))

	currentTime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, str+"_"+strconv.FormatInt(currentTime, 10)+"_"+salt)

	hsum := hex.EncodeToString(h.Sum(nil))

	return hsum
}

//GetFileParameters получает параметры файла, его размер и хеш-сумму
func GetFileParameters(filePath string) (int64, string, error) {
	fd, err := os.Open(filePath)
	if err != nil {
		return 0, "", err
	}
	defer fd.Close()

	fileInfo, err := fd.Stat()
	if err != nil {
		return 0, "", err
	}

	h := md5.New()
	if _, err := io.Copy(h, fd); err != nil {
		return 0, "", err
	}

	return fileInfo.Size(), hex.EncodeToString(h.Sum(nil)), nil
}

//GetCountPartsMessage получить количество частей сообщений
func GetCountPartsMessage(list map[string]int, sizeChunk int) int {
	var maxFiles float64
	for _, v := range list {
		if maxFiles < float64(v) {
			maxFiles = float64(v)
		}
	}

	newCountChunk := float64(sizeChunk)
	x := math.Floor(maxFiles / newCountChunk)
	y := maxFiles / newCountChunk

	if (y - x) != 0 {
		x++
	}

	return int(x)
}

//GetChunkListFiles разделяет список файлов на кусочки
func GetChunkListFiles(numPart, sizeChunk, countParts int, listFilesFilter map[string][]string) map[string][]string {
	lff := map[string][]string{}

	for disk, files := range listFilesFilter {
		if numPart == 1 {
			if len(files) < sizeChunk {
				lff[disk] = files[:]
			} else {
				lff[disk] = files[:sizeChunk]
			}

			continue
		}

		num := sizeChunk * (numPart - 1)
		numEnd := num + sizeChunk

		if numPart == countParts {
			if num < len(files) {
				lff[disk] = files[num:]

				continue
			}

			lff[disk] = []string{}
		}

		if numPart < countParts {
			if num > len(files) {
				lff[disk] = []string{}

				continue
			}

			if numEnd < len(files) {
				lff[disk] = files[num:numEnd]

				continue
			}

			lff[disk] = files[num:]
		}

	}
	return lff
}

//MothPrintIntAsString выводит месяц в виде числа как строку
func MothPrintIntAsString(m time.Month) string {
	var moth string

	switch m {
	case time.January:
		moth = "01"

	case time.February:
		moth = "02"

	case time.March:
		moth = "03"

	case time.April:
		moth = "04"

	case time.May:
		moth = "05"

	case time.June:
		moth = "06"

	case time.July:
		moth = "07"

	case time.August:
		moth = "08"

	case time.September:
		moth = "09"

	case time.October:
		moth = "10"

	case time.November:
		moth = "11"

	case time.December:
		moth = "12"

	}

	return moth
}

//MothNameAsString выводит название месяца по числу
func MothNameAsString(num int) string {
	if num == 0 {
		return ""
	}

	if num > 12 {
		return ""
	}

	mothList := map[int]string{
		1:  "January",
		2:  "February",
		3:  "March",
		4:  "April",
		5:  "May",
		6:  "June",
		7:  "July",
		8:  "August",
		9:  "September",
		10: "October",
		11: "November",
		12: "December",
	}

	return mothList[num]
}

func Ip2long(ipAddr string) (uint32, error) {
	ip := net.ParseIP(ipAddr)
	if ip == nil {
		return 0, errors.New("wrong ipAddr format")
	}
	ip = ip.To4()

	return binary.BigEndian.Uint32(ip), nil
}

func Long2ip(ipLong uint32) string {
	ipByte := make([]byte, 4)
	binary.BigEndian.PutUint32(ipByte, ipLong)
	ip := net.IP(ipByte)

	return ip.String()
}

func IntToIP(ip uint32) string {
	result := make(net.IP, 4)
	result[0] = byte(ip)
	result[1] = byte(ip >> 8)
	result[2] = byte(ip >> 16)
	result[3] = byte(ip >> 24)

	return result.String()
}
