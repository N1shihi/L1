package main

// Адаптер - это паттерн который позволяет использовать несовместимый клиентский интерфейс, без изменения существующего кода,
// путем оборачивания несовместимого интерфейса в совместимый
// плюсы: повторное использование существующего кода без изменений
// минусы: усложнение кода и его читабельности
// пример использования: интеграция с внешними библиотеками

import "fmt"

// клиентский интерфейс
type CelsiusThermometer interface {
	GetTemperatureCelsius() float64
}

type FahrenheitThermometer struct {
	temperature float64
}

func (t *FahrenheitThermometer) GetTemperatureFahrenheit() float64 {
	return t.temperature
}

// адаптер
type ThermometerAdapter struct {
	fahrenheitThermometer *FahrenheitThermometer
}

func (a *ThermometerAdapter) GetTemperatureCelsius() float64 {
	f := a.fahrenheitThermometer.GetTemperatureFahrenheit()
	return (f - 32) * 5 / 9
}

func main() {
	oldThermometer := &FahrenheitThermometer{temperature: 98.6}
	adapter := &ThermometerAdapter{fahrenheitThermometer: oldThermometer}

	fmt.Printf("Температура в фаренгейтах: %.1f\n", oldThermometer.GetTemperatureFahrenheit())
	fmt.Printf("Температура в цельсиях через адаптер: %.1f", adapter.GetTemperatureCelsius())
}
