package main

import (
	"context"
	"getback/geo"
	"sync"
)

func main() {
	mainContext, cancelMain := context.WithCancel(context.Background())
	midContext, cancelMid := context.WithCancel(mainContext)

	defer cancelMain()
	defer cancelMid()

	sWG := &sync.WaitGroup{}
	rWG := &sync.WaitGroup{}

	weatherCh := geo.WeatherGenerator(midContext, 20, sWG)

	templateCh := make(chan string, 30)
	dumpCh := make(chan string, 30)
	pressCh := make(chan string, 30)
	seismCh := make(chan string, 30)

	template := geo.Msg{
		Ch:  templateCh,
		Ctx: midContext}

	dump := template
	dump.Ch = dumpCh

	press := template
	press.Ch = pressCh

	seism := template
	seism.Ch = seismCh

	rWG.Add(3)
	go dump.Reader(rWG)
	go press.Reader(rWG)
	go seism.Reader(rWG)

	for weather := range weatherCh {
		sWG.Add(3)
		go dump.Sensor(1, weather, sWG)
		go press.Sensor(2, weather, sWG)
		go seism.Sensor(3, weather, sWG)
	}

	sWG.Wait()

	close(dumpCh)
	close(pressCh)
	close(seismCh)

	rWG.Wait()
}
