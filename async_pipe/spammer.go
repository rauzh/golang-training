package main

import (
	"fmt"
	"sort"
	"sync"
)

const goroutinesNum = 5
const batchNum = 2

func RunPipeline(cmds ...cmd) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()

	prevOut := make(chan interface{})

	for _, command := range cmds {
		wg.Add(1)
		curOut := make(chan interface{})

		go func(prevOut, curOut chan interface{}, command cmd, wg *sync.WaitGroup) {
			defer wg.Done()
			defer close(curOut) // закрываем, передаем в преваут, больше не пишем туда, теперь оттуда читаем
			command(prevOut, curOut)
		}(prevOut, curOut, command, wg)

		prevOut = curOut
	}
}

func CombineResults(in, out chan interface{}) {
	resStruct := []MsgData{}
	for i := range in {
		resStruct = append(resStruct, i.(MsgData))
	}
	sort.Slice(resStruct, func(i, j int) bool {
		if resStruct[i].HasSpam == resStruct[j].HasSpam {
			return resStruct[i].ID < resStruct[j].ID
		} else {
			return resStruct[i].HasSpam
		}
	})
	for _, dataStruct := range resStruct {
		out <- fmt.Sprintf("%t %d", dataStruct.HasSpam, dataStruct.ID)
	}
}

func SelectUsers(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	mu := &sync.Mutex{}
	res := map[uint64]User{}

	for dataIn := range in {
		wg.Add(1)
		go func(email string, out chan interface{}, wg *sync.WaitGroup) {
			defer wg.Done()
			tmpUser := GetUser(email)

			mu.Lock()                                // внутри есть работа с мапой
			if _, exist := res[tmpUser.ID]; !exist { // "выдаем только уникальных юзеров"
				res[tmpUser.ID] = tmpUser
				out <- tmpUser
			}
			mu.Unlock()
		}(fmt.Sprintf("%v", dataIn), out, wg)
	}
}

func SelectMessages(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	defer wg.Wait()
	var usrs []User

	for {
		tmpUser, ok1 := (<-in).(User)
		if ok1 {
			usrs = []User{tmpUser}
		} else {
			break
		}
		for i, ok2 := 1, true; i < batchNum && ok2; i++ {
			tmpUser, ok2 := (<-in).(User)
			if ok2 {
				usrs = append(usrs, tmpUser)
			}
		}

		wg.Add(1)
		go func(usrs []User, out chan interface{}, wg *sync.WaitGroup) {
			defer wg.Done()

			tmpMsgIds := getMessagesWrapper(usrs)

			for _, msg := range tmpMsgIds {
				out <- msg
			}
		}(usrs, out, wg)
	}
}

func getMessagesWrapper(usrs []User) []MsgID {
	res, err := GetMessages(usrs...)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return res
}

func CheckSpam(in, out chan interface{}) {
	wg := &sync.WaitGroup{}

	workerIn := make(chan MsgID) // ограничение кол-ва вызовов паттерном workerpool
	for i := 0; i < goroutinesNum; i++ {
		go Worker(out, wg, workerIn)
	}

	for input := range in { // можно сразу в воркер, но так лучше воспринимается код и паттерн
		workerIn <- input.(MsgID)
	}
	close(workerIn)

	wg.Wait()
}

func Worker(out chan interface{}, wg *sync.WaitGroup, in <-chan MsgID) {
	for msg := range in {
		wg.Add(1)
		defer wg.Done()
		ch := make(chan interface{})

		go func() {
			hasSpam := hasSpamWrapper(msg)
			ch <- MsgData{HasSpam: hasSpam, ID: msg}
		}()
		out <- <-ch
	}
}

func hasSpamWrapper(msg MsgID) bool {
	res, err := HasSpam(msg)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return res
}
