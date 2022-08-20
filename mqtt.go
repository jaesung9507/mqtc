package main

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const IS_TLS = false
const IP = "test.mosquitto.org"
const PORT = 1883
const TLS_PORT = 1884

var CA = []byte(`-----BEGIN CERTIFICATE-----
-----END CERTIFICATE-----`)

func getServer() string {
	protocol := "tcp"
	port := PORT
	if IS_TLS {
		protocol = "tls"
		port = TLS_PORT
	}
	return fmt.Sprintf("%s://%s:%d", protocol, IP, port)
}

func Connect(h mqtt.ConnectionLostHandler) (*mqtt.Client, error) {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(getServer())
	opts.SetConnectTimeout(3 * time.Second)
	opts.SetAutoReconnect(false)
	opts.SetClientID("test")
	opts.SetConnectionLostHandler(h)
	if IS_TLS {
		certpool := x509.NewCertPool()
		certpool.AppendCertsFromPEM(CA)
		opts.SetTLSConfig(&tls.Config{
			//RootCAs: certpool,
			InsecureSkipVerify: true,
		})
	}
	client := mqtt.NewClient(opts)
	token := client.Connect()
	token.Wait()
	if err := token.Error(); err != nil {
		return nil, err
	}
	return &client, nil
}

func Subscription(client *mqtt.Client, topic string, handler mqtt.MessageHandler) error {
	token := (*client).Subscribe(topic, 0, handler)
	if !token.WaitTimeout(1 * time.Second) {
		return errors.New("Failed to Subscribe, timeout duration: 1sec")
	}
	if err := token.Error(); err != nil {
		return err
	}
	return nil
}

func Publish(client *mqtt.Client, topic, msg string) error {
	token := (*client).Publish(topic, 0, false, msg)
	if !token.WaitTimeout(1 * time.Second) {
		return errors.New("Failed to Publish, timeout duration: 1sec")
	}
	if err := token.Error(); err != nil {
		return err
	}
	return nil
}
