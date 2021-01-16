package service

import (
	"context"
	"log"
	"time"
)

type TimeServe struct {
}

func (timeServe *TimeServe) GetCurrtentTime(ctx context.Context) (r int32, err error) {
	currentTime := time.Now()
	log.Println("GetCurrtentTime() ")
	return int32(currentTime.Hour()), nil
}
func NewTimeServerHandle() *TimeServe {
	log.Println("TimeServer() service Start()")
	return &TimeServe{}
}
