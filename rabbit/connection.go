package rabbit

import (
 "context"

 amqp "github.com/rabbitmq/amqp091-go"
 "go.uber.org/fx"
)

type RabbitModel struct {
}

var Module = fx.Options(fx.Provide(NewRabbit))

func NewRabbit() *RabbitModel {
 return &RabbitModel{}
}

type RabbitInterface interface {
}

func (rabbitModel *RabbitModel) StartConnect(ctx context.Context) error {
 connection, err := amqp.Dial("amqp://user:123@localhost:7575")
 if err != nil {
  return err
 }
 defer connection.Close()
 ch, err := connection.Channel()
 if err != nil {
  return err
 }
 defer ch.Close()
 mainQueue, err := ch.QueueDeclare(
  "queue",
  false,
  false,
  false,
  false,
  nil,
 )
 if err != nil {
  return err
 }

 err = ch.PublishWithContext(
  ctx,
  "",
  mainQueue.Name,
  false,
  false,
  amqp.Publishing{
   ContentType: "text/plain",
   Body:        []byte("hello world"),
  },
 )

 if err != nil {
  return err
 }
 return nil
}
