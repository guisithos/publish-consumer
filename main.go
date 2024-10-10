package main

import (
 "context"
 "testd/rabbit"

 "go.uber.org/fx"
)

func Register(lifeCycle fx.Lifecycle, rabbit *rabbit.RabbitModel) {
 lifeCycle.Append(fx.Hook{
  OnStart: func(ctx context.Context) error {
   return rabbit.StartConnect(ctx)
  },
  OnStop: nil,
 })
}

func main() {
 app := fx.New(
  rabbit.Module,
  fx.Invoke(Register),
 )
 app.Start(context.Background())
}
