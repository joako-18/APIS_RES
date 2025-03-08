package infrastructure

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

func ConsumeOrders(rabbit *NewRabbitMQService) {
	msgs, err := rabbit.Channel.Consume(
		"ordenes",
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Error consumiendo la cola:", err)
	}

	go func() {
		for msg := range msgs {
			fmt.Println("Recibido mensaje:", string(msg.Body))

			sendToNotificationAPI(msg.Body)
		}
	}()

	fmt.Println("Consumidor de RabbitMQ iniciado y esperando mensajes...")
}

func sendToNotificationAPI(data []byte) {
	url := "http://segunda-api.com/notificacion"

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Println("Error enviando notificación:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Notificación enviada, respuesta:", resp.Status)
}
