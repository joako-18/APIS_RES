package services

import (
	"api/src/domain/entities"
	"encoding/json"
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

type RabbitNotificador struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Queue   amqp.Queue
}

func NewRabbitNotificador() *RabbitNotificador {
	log.Println("[DEBUG] Intentando conectar con RabbitMQ...")
	conn, err := amqp.Dial("amqp://xk27:mando18D@44.215.213.150:5672/")
	if err != nil {
		log.Fatalf("[ERROR] No se pudo conectar a RabbitMQ: %v", err)
	}
	log.Println("[DEBUG] Conexión establecida con RabbitMQ.")

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("[ERROR] No se pudo crear el canal: %v", err)
	}
	log.Println("[DEBUG] Canal creado correctamente.")

	q, err := ch.QueueDeclare(
		"notificacion_pedidos", // nombre
		true,                   // durable
		false,                  // autoDelete
		false,                  // exclusive
		false,                  // noWait
		nil,                    // args
	)
	if err != nil {
		log.Fatalf("[ERROR] No se pudo declarar la cola: %v", err)
	}
	log.Printf("[DEBUG] Cola declarada correctamente: %s\n", q.Name)

	return &RabbitNotificador{
		Conn:    conn,
		Channel: ch,
		Queue:   q,
	}
}

func (r *RabbitNotificador) EnviarNotificacion(pedido entities.PedidoCompletado) error {
	log.Printf("[DEBUG] Serializando pedido: %+v\n", pedido)
	body, err := json.Marshal(pedido)
	if err != nil {
		log.Printf("[ERROR] Fallo al serializar el pedido: %v\n", err)
		return err
	}

	log.Printf("[DEBUG] Enviando notificación JSON a la cola '%s'...\n", r.Queue.Name)
	err = r.Channel.Publish(
		"",           // exchange
		r.Queue.Name, // routing key
		false,        // mandatory
		false,        // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Printf("[ERROR] Error al publicar mensaje: %v\n", err)
		return err
	}
	log.Println("[DEBUG] Notificación enviada exitosamente.")
	return nil
}

func (r *RabbitNotificador) NotificarPedido(pedidoID int, estado string) error {
	msg := fmt.Sprintf("Pedido %d ha sido %s", pedidoID, estado)
	log.Printf("[DEBUG] Enviando mensaje de texto a la cola '%s': %s\n", "notificacion_pedidos", msg)
	err := r.Channel.Publish(
		"",                     // exchange
		"notificacion_pedidos", // routing key
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		},
	)
	if err != nil {
		log.Printf("[ERROR] Error al publicar mensaje de texto: %v\n", err)
		return err
	}
	log.Println("[DEBUG] Mensaje de texto publicado correctamente.")
	return nil
}
