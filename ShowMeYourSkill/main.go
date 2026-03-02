package main

import (
	"study/sensors"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go sensors.SensorsSimulationPool(wg)

	time.Sleep(3 * time.Second)
	sensors.AirHumidityKill()

	time.Sleep(1 * time.Second)
	sensors.AirPresssureKill()

	time.Sleep(1 * time.Second)
	sensors.AirTemperatureKill()

	wg.Wait()

}
