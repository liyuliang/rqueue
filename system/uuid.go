package system

import (
	"github.com/zheng-ji/goSnowFlake"
	"time"
	"errors"
	"log"
)

var idChans chan int64
var idWorker *goSnowFlake.IdWorker

func init() {
	idChans = make(chan int64)

	if idWorker == nil {
		initIdWorker()
	}

	idChans = idChannels()
}

func initIdWorker() {
	wid := 1

	iw, err := goSnowFlake.NewIdWorker(int64(wid))
	if err != nil {
		panic(err)
	}
	idWorker = iw
}

func getNew() (id int64, err error) {

	id, err = idWorker.NextId()
	return id, err
}

func idChannels() chan int64 {

	var out = make(chan int64)

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	go func() {
		for {
			out <- <-generator()
		}
	}()

	return out
}

func generator() chan int64 {
	// 创建通道
	out := make(chan int64)
	// 创建协程
	go func() {
		for {
			//向通道内写入数据，如果无人读取会等待

			id, err := getNew()
			if err == nil && id != 0 {
				out <- id
			}
		}
	}()
	return out
}

func Uid() (id int64, err error) {

	return uuidFromChannel()
}

func uuidFromChannel() (id int64, err error) {

wait:
	for {

		select {

		case id = <-idChans:
			log.Println("uuid from channel ")
			err = nil
			break wait

		case <-time.After(time.Duration(3) * time.Second):
			err = errors.New("uuid generate failed, timeout")
			break wait
		}
	}

	return id, err
}

func genUUid() (id int64, err error) {
	iw, _ := goSnowFlake.NewIdWorker(1)
	id, err = iw.NextId()
	return id, err
}
