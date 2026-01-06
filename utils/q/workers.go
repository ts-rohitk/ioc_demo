package q

import (
	"fmt"
	"goat/utils/mapping"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
)

func Workers(id int, tasks <-chan Task, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		result := Result{
			Task: task,
			Data: mapToIoc(task.Data),
		}
		results <- result
		fmt.Printf("worker id: %d finished task with id:%d \n", id, task.Id)
	}
}

func mapToIoc(data mapping.RawData) *mapping.IOC {
	normalizedJSON, hashedCode, err := hashNormalize(data)
	if err != nil {
		log.Fatal(err)
	}

	rawJSONString, err := rawJSONString(data)
	if err != nil {
		log.Fatal(err)
	}

	milliSeconds, err := convertToMilliseconds(data.FirstSeen)
	if err != nil {
		log.Fatal(err)
	}

	lastSeen := milliSeconds
	if data.LastSeen != nil {
		ms, err := convertToMilliseconds(*data.LastSeen)
		if err != nil {
			log.Fatal(err)
		}
		lastSeen = ms
	}

	ioc := mapping.IOC{
		UUID:         uuid.New().String(),
		ValueRaw:     rawJSONString,
		ValueNorm:    normalizedJSON,
		HashCode:     hashedCode,
		Type:         mapping.IOCType(data.IocType),
		Title:        buildTitle(data.Ioc, data.ThreatType, data.MalwarePrintable, mapping.IOCType(data.IocType)), // `${ioc} + ${threat_type} + ${type} + ${malware_family_name}`
		Key:          buildKey(mapping.IOCType(data.IocType), normalizedJSON),
		FirstSeen:    milliSeconds, // required f"{Type}|{valueNorm}"
		LastSeen:     lastSeen,
		CreatedAt:    currentDateToMilliseconds(),
		UpdatedAt:    nil,
		ExpiresAt:    nil,
		ThreatType:   &data.ThreatType,
		Tags:         data.Tags,
		Confidence:   calculateConfidenceLevel(data.ConfidenceLevel),
		Malware:      nil, //req
		Network:      nil,
		ThreatActors: nil, //req
		Victims:      nil,
		Sources:      nil, //req
		TTP:          nil,
		Meta:         nil,
	}

	return &ioc
}

func buildTitle(iocName, threatType, familyName string, iocType mapping.IOCType) *string {
	var iocTitle strings.Builder
	const sep = " "
	iocTitle.Grow(len(iocName) + len(threatType) + len(familyName) + len(string(iocType)) + 3*len(sep))

	iocTitle.WriteString(iocName)
	iocTitle.WriteString(sep)

	iocTitle.WriteString(threatType)
	iocTitle.WriteString(sep)

	iocTitle.WriteString(string(iocType))
	iocTitle.WriteString(sep)

	iocTitle.WriteString(familyName)

	result := iocTitle.String()
	return &result
}

func buildKey(iocType mapping.IOCType, valueNorm string) string {
	var KeyBuilder strings.Builder
	const sep = "|"
	KeyBuilder.Grow(len(string(iocType)) + len(valueNorm) + len(sep))

	KeyBuilder.WriteString(string(iocType))
	KeyBuilder.WriteString(sep)
	KeyBuilder.WriteString(valueNorm)

	return KeyBuilder.String()
}

func convertToMilliseconds(date string) (*int64, error) {
	layout := "2006-01-02 15:04:05 MST"
	t, err := time.Parse(layout, date)
	if err != nil {
		return nil, err
	}

	milli := t.UnixMilli()
	return &milli, nil
}

func currentDateToMilliseconds() int64 {
	current := time.Now()
	return current.UnixMilli()
}

func calculateConfidenceLevel(level int) *string {

	switch {
	case level >= 0 && level <= 25:
		l := "limited"
		return &l
	case level > 25 && level <= 50:
		m := "moderate"
		return &m
	case level > 50 && level <= 75:
		e := "elevated"
		return &e
	case level > 75 && level <= 100:
		h := "high"
		return &h
	default:
		l := "low"
		return &l
	}
}
