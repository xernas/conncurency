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
	defer cancelMain()
	defer cancelMid()

	mtx := &sync.Mutex{}

	sensorWG := &sync.WaitGroup{}

	readerWG := &sync.WaitGroup{}

	weatherCh := geo.WeatherGenerator(mainContext, 20, sensorWG)

	dumpCh := make(chan string)
	pressCh := make(chan string)
	seismCh := make(chan string)

	readerWG.Add(3)
	go reader.Reader(dumpCh, mtx, readerWG, middleContext)
	go reader.Reader(pressCh, mtx, readerWG, middleContext)
	go reader.Reader(seismCh, mtx, readerWG, middleContext)

	for weather := range weatherCh {
		sensorWG.Add(3)
		go geo.DumpSensor(weather, dumpCh, sensorWG, middleContext)
		go geo.PressSensor(weather, pressCh, sensorWG, middleContext)
		go geo.SeismSensor(weather, seismCh, sensorWG, middleContext)
	}

	sensorWG.Wait()

	close(dumpCh)
	close(pressCh)
	close(seismCh)

	readerWG.Wait()
}
