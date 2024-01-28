package main

import "fmt"

type Monitor struct{}

// Пусть монитор выводит изображение только через кабель с HDMI коннектором
func (monitor Monitor) ConnectCable(cable HDMIConnector) {
	cable.PlugHDMI()
}

// Тип данных удовлетворяет интерфейс HDMIConnector, если у него определен метод PlugHDMI()
type HDMIConnector interface {
	PlugHDMI()
}

// Структура HDMI кабеля
type HDMI struct{}

func (hdmi HDMI) PlugHDMI() {
	fmt.Println("HDMI Plugged")
}

// Структура DVI кабеля
type DVI struct{}

func (dvi DVI) PlugDVI() {
	fmt.Println("DVI Plugged")
}

// Структура адаптера DVI-to-HDMI, которая удовлетворяет интерфейс HDMIConnector
type DVIToHDMIAdapter struct {
	dvi *DVI
}

func (adapter DVIToHDMIAdapter) PlugHDMI() {
	fmt.Print("DVI-To-HDMI: ")
	adapter.dvi.PlugDVI()
}

func main() {
	monitor := &Monitor{}

	// Подключение HDMI кабеля к монитору напрямую
	hdmi := &HDMI{}
	monitor.ConnectCable(hdmi)

	// Подключение DVI кабеля к монитору через адаптер
	dvi := &DVI{}
	adapter := &DVIToHDMIAdapter{dvi: dvi}
	monitor.ConnectCable(adapter)
}
