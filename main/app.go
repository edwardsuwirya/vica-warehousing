package main

import (
	"warehousing/config"
	"warehousing/deliveries"
)

type app struct {
	config *config.AppConfig
}

func (a app) cli() {
	a.run(deliveries.NewAppDelivery(a.config))
}
func (a app) run(delivery deliveries.IAppDelivery) {
	delivery.Run()
}

func newApp() app {
	config := config.NewConfig()
	return app{config}
}

func main() {
	newApp().cli()
}
