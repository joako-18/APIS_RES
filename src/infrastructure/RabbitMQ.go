package infrastructure

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/streadway/amqp"
)

type RabbitMQService struct {
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func NewRabbitMQService() (*RabbitMQService, error) {
	conn, err := amqp.Dial("amqp://xk27:mando18D@localhost:5672/")
	if err != nil {
		return nil, fmt.Errorf("Error conectando a RabbitMQ: %v", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("Error abriendo canal en RabbitMQ: %v", err)
	}

	return &RabbitMQService{Connection: conn, Channel: ch}, nil
}

func (r *RabbitMQService) PublishMessage(message string) error {
	return r.Channel.Publish(
		"",
		"ordenes",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(message),
		},
	)
}

func (r *RabbitMQService) ConsumeOrders() {
	msgs, err := r.Channel.Consume(
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
