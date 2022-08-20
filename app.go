package main

import (
	"context"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx    context.Context
	client *mqtt.Client
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) MqttConnect(name string) bool {
	client, err := Connect(func(c mqtt.Client, err error) {
		runtime.EventsEmit(a.ctx, "OnMqttDisconnect")
	})
	if err == nil {
		a.client = client
	}
	return err == nil
}

func (a *App) MqttSubscription(topic string) bool {
	return Subscription(a.client, topic, func(c mqtt.Client, m mqtt.Message) {
		t := m.Topic()
		p := m.Payload()
		q := m.Qos()
		r := m.Retained()
		runtime.EventsEmit(a.ctx, "OnMqttMessage", t, string(p), q, r)
	}) == nil
}

func (a *App) MqttPublish(topic, msg string) bool {
	return Publish(a.client, topic, msg) == nil
}
