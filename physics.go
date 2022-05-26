package main

import (
	"log"

	"github.com/jakecoffman/cp"
)

const (
	COLLISION_BOAT   = 1
	COLLISION_SENSOR = 2
)

func InitSpace(s *cp.Space) {
	sensorHandler := s.NewCollisionHandler(
		COLLISION_BOAT, COLLISION_SENSOR)
	sensorHandler.BeginFunc = sensorHandlerBeginFunc
	sensorHandler.SeparateFunc = sensorHandlerSeparateFunc

	log.Println("init space")
}

func sensorHandlerBeginFunc(arb *cp.Arbiter, space *cp.Space, data interface{}) bool {
	boat, sensor := arb.Bodies()

	b, ok_b := boat.UserData.(*Boat)
	i, ok_i := sensor.UserData.(*Island)
	if ok_b && ok_i {
		b.EnterStore()
		game.boatOnIsland[b] = i
	}
	return true
}

func sensorHandlerSeparateFunc(arb *cp.Arbiter, space *cp.Space, data interface{}) {
	boat, sensor := arb.Bodies()

	b, ok_b := boat.UserData.(*Boat)
	_, ok_i := sensor.UserData.(*Island)
	if ok_b && ok_i {
		b.ExitStore()
		delete(game.boatOnIsland, b)
	}
}
