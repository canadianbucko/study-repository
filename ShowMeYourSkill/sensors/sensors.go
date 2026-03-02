package sensors

import (
	"context"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"

	"github.com/k0kubun/pp"
)

type sensorsInfo struct {
	mainInfo    string
	coordinates string
}

func NewSensorsInfo(mainInfo string, coordinates string) sensorsInfo {
	return sensorsInfo{
		mainInfo:    mainInfo,
		coordinates: coordinates,
	}
}

func airPressure(ctx context.Context, airPressureChannel chan<- sensorsInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сенсор Air Pressure отключен!")
			close(airPressureChannel)
			return
		default:
			airP := rand.Intn(500)
			convertAirP := strconv.Itoa(airP) + "давление воздуха"
			coords := "100.40*, 50.59*"

			airPressureChannel <- NewSensorsInfo(convertAirP, coords)

			time.Sleep(5 * time.Second)

		}
	}

}

func airHumidity(ctx context.Context,
	airHumidityChannel chan<- sensorsInfo,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сенсор Air Humidity отключен!")
			close(airHumidityChannel)
			return
		default:
			airHum := rand.Intn(500)
			convertAirH := strconv.Itoa(airHum) + "влажность воздуха"
			coords := "145.14*, 14.23*"

			airHumidityChannel <- NewSensorsInfo(convertAirH, coords)

			time.Sleep(5 * time.Second)

		}
	}

}
func airTemperature(ctx context.Context,
	airTemperatureChannel chan<- sensorsInfo,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Сенсор Air Temperature отключен!")
			close(airTemperatureChannel)
			return
		default:
			airTemp := rand.Intn(500)
			convertAirTemp := strconv.Itoa(airTemp) + "температура воздуха"
			coords := "124.44*, 56.69*"

			airTemperatureChannel <- NewSensorsInfo(convertAirTemp, coords)

			time.Sleep(5 * time.Second)

		}
	}

}

var airPressContext, AirPressCancel = context.WithCancel(context.Background())
var airHumidityContext, AirHumCancel = context.WithCancel(context.Background())
var airTempContext, AirTempCancel = context.WithCancel(context.Background())

func AirPresssureKill() {
	AirPressCancel()
}

func AirHumidityKill() {
	AirHumCancel()
}

func AirTemperatureKill() {
	AirTempCancel()
}

func SensorsSimulationPool(WaityGroupy *sync.WaitGroup) { // а что мне в канал передавать? и где?
	defer WaityGroupy.Done()
	wg := &sync.WaitGroup{}

	airPressureChannel := make(chan sensorsInfo)
	airHumidityChannel := make(chan sensorsInfo)
	airTemperatureChannel := make(chan sensorsInfo)

	wg.Add(1)
	go airPressure(airPressContext, airPressureChannel, wg)

	wg.Add(1)
	go airHumidity(airHumidityContext, airHumidityChannel, wg)

	wg.Add(1)
	go airTemperature(airTempContext, airTemperatureChannel, wg)

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range airPressureChannel { // struct string--string
			pp.Println(v)
			time.Sleep(1 * time.Second)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range airHumidityChannel { // struct string--string
			pp.Println(v)
			time.Sleep(2 * time.Second)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for v := range airTemperatureChannel { // struct string--string
			pp.Println(v)
			time.Sleep(1 * time.Second)
		}
	}()
	wg.Wait()

}
