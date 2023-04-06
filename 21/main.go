package main

import "fmt"

// MetricSystemEnjoyer описывает человека, который
// измеряет расстояние в километрах
type MetricSystemEnjoyer interface {
	TellDistanceInKilometers() float32
}

// MetricSystemGuy представляет человека.
// который вычисляет расстояние в метрической системе
type MetricSystemGuy struct {
	distance float32
}

// TellDistanceInKilometers представляет собой
// метрическую систему, позволяющую определить расстояние в километрах
func (msg MetricSystemGuy) TellDistanceInKilometers() float32 {
	return msg.distance
}

// ImperialSystemGuy представляет человека
// который рассчитывает расстояние в имперской системе
type ImperialSystemGuy struct {
	distance float32
}

// TellDistanceInMiles представляет собой
// систему, позволяющую определить расстояние в милях
func (isf ImperialSystemGuy) TellDistanceInMiles() float32 {
	return isf.distance
}

// Traveler это клиентский тип типа MetricSystemGuy
type Traveler struct {
	// ...
}

// Traveler хочет узнать расстояние в километрах
func (t Traveler) AskDistanceInKilometers(guy MetricSystemEnjoyer) {
	fmt.Printf("Этот парень сказал, что мне осталось пройти %.1f километр\n", guy.TellDistanceInKilometers())
}

// Adapter позволяет использовать ImperialSystemGuy как MetricSystemEnjoyer
type Adapter struct {
	fan *ImperialSystemGuy
}

// TellDistanceInKilometers делает Adapter для реализации
// MetricSystemEnjoyer
func (a Adapter) TellDistanceInKilometers() float32 {
	return a.fan.distance * 1.6
}

func main() {
	daniil := &ImperialSystemGuy{1}   // Даниил говорит расстояние в милях
	anton := &MetricSystemGuy{1}      // Антон говорит расстояние в километрах
	adaptedDaniil := &Adapter{daniil} // адаптированный Даниил сообщает расстояние в километрах

	kirill := new(Traveler) // Кирилл - путешественник, который понимает только метрическую систему

	kirill.AskDistanceInKilometers(anton)
	kirill.AskDistanceInKilometers(adaptedDaniil)
}
