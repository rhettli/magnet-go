package send

import (
	"log"
	"os"
	"strings"
	"magnet/rabbitmq"

	"github.com/streadway/amqp"

	"errors"
)


//发送rbmq消息
func SendRbmqMsg(body string)error {
	/*
	defer func() {
		if err := recover(); err != nil {
			//做异常处理
			if err != nil {
				log.Println(err)
			}
		}
	}()
	*/

	conn, err := amqp.Dial(rabbitmq.CRbmqAddr)

	if err!=nil{
		log.Println("Failed to connect to RabbitMQ",err)
		return errors.New("interal error.1")
	}
	defer conn.Close()

	ch, err := conn.Channel()

	if err!=nil{
		log.Println("Failed to open a channel",err)
		return errors.New("interal error.2")
	}

	defer ch.Close()

	err = ch.ExchangeDeclare(
		"logs_topic", // name
		"topic",      // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)

	if err!=nil{
		log.Println("Failed to declare an exchange",err)
		return errors.New("interal error.3")
	}

	err = ch.Publish(
		"logs_topic",  	// exchange
		rabbitmq.CRabbitMqHashView, 			// routing key
		false, 	// mandatory
		false, 	// immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	if err!=nil{
		log.Println("Failed to publish a message",err)
		return errors.New("interal error.3")
	}

	log.Printf(" [x] Sent %s", body)

	return nil
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "anonymous.info"
	} else {
		s = os.Args[1]
	}
	return s
}