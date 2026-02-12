package main

import (
	"context"
	"getback/geo"
	"getback/reader"
	"sync"
)

func main() {
	mainContext, cancelMain := context.WithCancel(context.Background())
	middleContext, cancelMid := context.WithCancel(mainContext)

	defer func() {
		cancelMain()
		cancelMid()
	}()

	wg := &sync.WaitGroup{}
	mtx := &sync.Mutex{}

	weatherCh := geo.WeatherGenerator(mainContext, 20)

	dumpCh := make(chan string)
	pressCh := make(chan string)
	seismCh := make(chan string)

	wg.Add(1)
	go reader.Reader(dumpCh, mtx, wg, middleContext)
	wg.Add(1)
	go reader.Reader(pressCh, mtx, wg, middleContext)
	wg.Add(1)
	go reader.Reader(seismCh, mtx, wg, middleContext)

	for weather := range weatherCh {
		wg.Add(1)
		go geo.DumpSensor(weather, dumpCh, wg, middleContext)
		wg.Add(1)
		go geo.PressSensor(weather, pressCh, wg, middleContext)
		wg.Add(1)
		go geo.SeismSensor(weather, seismCh, wg, middleContext)
	}

	go func() {
		wg.Wait()
	}()
}
