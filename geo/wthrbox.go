package geo

import (
	"context"
	"getback/weather"
	"math/rand"
	"sync"
)

func NewWeather(
	ctx context.Context,
	transferPoint chan weather.Weather,
	wg *sync.WaitGroup, dayNum int) {

	defer wg.Done()
	randPressure := 700 + rand.Intn(100)
	randDump := rand.Intn(100)
	randSeism := rand.Intn(9)

	wthr := weather.Weather{
		AirPressure: randPressure,
		AirDump:     randDump,
		Seism:       float64(randSeism),
		Day:         dayNum,
	}

	select {
	case <-ctx.Done():
		return
	case transferPoint <- wthr:
	}
}

func WeatherGenerator(
	ctx context.Context,
	dayNum int) <-chan weather.Weather {
	trPoint := make(chan weather.Weather)
	wg := &sync.WaitGroup{}

	for i := 1; i <= dayNum; i++ {
		wg.Add(1)
		go NewWeather(ctx, trPoint, wg, i)
	}

	go func() {
		wg.Wait()
		close(trPoint)
	}()

	return trPoint
}
