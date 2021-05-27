package service

import (
	"Website_1/util"
	"fmt"
	"net/http"

	"github.com/streadway/amqp"
)

func MessageBrokerService(w http.ResponseWriter, r *http.Request) {
	//publishing message to the rabbit mq
	if r.Method == "POST" {
		if r.FormValue("feedback") != "" {
			ch, err := util.MessageBroker.Channel()
			if err != nil {
				fmt.Println(err)
			}
			defer ch.Close()
			_, err = ch.QueueDeclare("Webapp", false, false, false, false, nil)
			//durable,autodelete,exclusive,nowait..
			if err != nil {
				fmt.Println(err)
			}

			err = ch.Publish("", "Webapp", false, false, amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(r.FormValue("feedback")),
			})

			if err != nil {
				fmt.Println(err)
			}
			fmt.Println("successfully publish message : ", r.FormValue("feedback"))
		}
		http.Redirect(w, r, "/", 301)
	}
}

func ConsumerService() {
	ch, err := util.MessageBroker.Channel()
	if err != nil {
		fmt.Println(err)
	}
	defer ch.Close()

	message, err := ch.Consume("Webapp", "", true, false, false, false, nil)
	//consumerstring,autoack, exclusive,noLocal,nowait
	for d := range message {
		fmt.Println("received message from rabbitmq : ", string(d.Body))
	}
	defer util.MessageBroker.Close()
}
