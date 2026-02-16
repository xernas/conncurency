package geo

import (
	"context"
	"fmt"
	"getback/weather"
	"strconv"
	"sync"
)

type Msg struct {
	Ch  chan string
	Ctx context.Context
}

func (m Msg) Reader(wg *sync.WaitGroup) {
	defer wg.Done()
	for v := range m.Ch {
		fmt.Println(string(v))
	}
}

func (m *Msg) Sensor(
	sensNum int,
	wthr weather.Weather,
	wg *sync.WaitGroup) {

	defer wg.Done()
	var result = ""

	day := strconv.Itoa(wthr.Day)
	dump := strconv.Itoa(wthr.AirDump)
	press := strconv.Itoa(wthr.AirPressure)
	seism := strconv.Itoa(int(wthr.Seism))

	switch sensNum {
	case 1:
		result = fmt.Sprintf("Датчик влажности зарегистрировал %s%%, %s дня", dump, day)
	case 2:
		result = fmt.Sprintf("Датчик давления зарегистрировал %s%%, %s дня", press, day)
	case 3:
		result = fmt.Sprintf("Сейсмодатчик зарегистрировал %s баллов, %s дня", seism, day)
	}

	select {
	case <-m.Ctx.Done():
		return
	case m.Ch <- result:
	}
}
