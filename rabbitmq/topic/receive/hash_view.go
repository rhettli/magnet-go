package main

import (
	"fmt"
	"log"
	"magnet/rabbitmq"
	"github.com/streadway/amqp"
	"magnet/mdb"
	"time"
	"magnet/model"
	"github.com/jinzhu/gorm"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}

func main() {
	conn, err := amqp.Dial(rabbitmq.CRbmqAddr)
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
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
	failOnError(err, "Failed to declare an exchange")

	q, err := ch.QueueDeclare(
		"",    // name
		true, // durable
		false, // delete when usused
		true,  // exclusive
		false, // no-wait
		nil,   // arguments
	)
	failOnError(err, "Failed to declare a queue")


	//for _, s := range os.Args[1:] {
		log.Printf("Binding queue %s to exchange %s with routing key %s",
			q.Name, "logs_topic", rabbitmq.CRabbitMqHashView)
		err = ch.QueueBind(
			q.Name,       // queue name
			rabbitmq.CRabbitMqHashView,            // routing key
			"logs_topic", // exchange
			false,
			nil)
		failOnError(err, "Failed to bind a queue")


	//}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto ack
		false,  // exclusive
		false,  // no local
		false,  // no wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf(" [x] Receive data:%s", d.Body)

			updateView(string(d.Body))

		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}

//更新数据库，requests和last_seen
func updateView(id string){
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	mysqlDB,err:=mdb.GetLocalDB()
	if err!=nil{
		log.Printf(" [x] Receive data:%s", id)
	}
	defer mysqlDB.Close()

	//searHash:=model.SearchHash{}

	//更新request&last_seen
	err=mysqlDB.Model(&model.SearchHash{}).Where(&model.SearchHash{Id:id}).
		UpdateColumn("requests", gorm.Expr("requests + ?", 1)).
	Update(&model.SearchHash{LastSeen:time.Now()}).Error
	if err!=nil {
		fmt.Println("update REQUEST&LAST_SEEN error:", err)
	}

	/*
	k:=map[string]interface{}{"requests":gorm.Expr("requests + ?", 1),"last_seen":time.Now()}
	if err:=mysqlDB.Model(&model.SearchHash{}).Where("id = ?",id).Updates(&k).Error;err!=nil{
		fmt.Println("update requests error:",err)
	}
	*/

	//mysqlDB.Model(&model.SearchHash{}).Where("id = ?",id).
		//Update( map[string]interface{}{"last_seen":	time.Now().Format("2006-01-02 15:04:05")})

	//tm:= time.Now().Format("2006-01-02 15:04:05")

	//if err:=mysqlDB.Model(&model.SearchHash{}).Where(&model.SearchHash{Id:id}).Update(&model.SearchHash{LastSeen:time.Now()}).Error;err!=nil{
	//	fmt.Println("update last_seen error:",err)

	//}

	// UPDATE "products" SET "quantity" = quantity - 1 WHERE "id" = '2' AND quantity > 1;



}