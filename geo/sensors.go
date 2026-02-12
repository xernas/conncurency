package geo

import (
	"context"
	"getback/weather"
	"strconv"
	"sync"
)

func DumpSensor(
	wthr weather.Weather,
	ch chan string,
	wg *sync.WaitGroup,
	ctx context.Context) {

	defer wg.Done()
	day := strconv.Itoa(wthr.Day)
	dump := strconv.Itoa(wthr.AirDump)

	result := "Датчик влажности зарегистрировал " + dump + "%, " + day + " дня месяца."

	select {
	case <-ctx.Done():
		return
	case ch <- result:
	}
}

func PressSensor(
	wthr weather.Weather,
	ch chan string,
	wg *sync.WaitGroup,
	ctx context.Context) {

	defer wg.Done()
	pressSlice := []string{}
	day := strconv.Itoa(wthr.Day)
	press := strconv.Itoa(wthr.AirPressure)
	pressSlice = append(pressSlice, day, press)
	result := "Датчик давления зарегистрировал уровень " + press + "%, " + day + " дня месяца."

	select {
	case <-ctx.Done():
		return
	case ch <- result:
	}
}

func SeismSensor(
	wthr weather.Weather,
	ch chan string,
	wg *sync.WaitGroup,
	ctx context.Context) {

	defer wg.Done()
	day := strconv.Itoa(wthr.Day)
	seism := strconv.FormatFloat(wthr.Seism, 'f', 2, 64)
	result := "Сейсмодатчик зарегистрировал уровень " + seism + " " + day + " дня месяца."

	select {
	case <-ctx.Done():
		return
	case ch <- result:
	}
}
