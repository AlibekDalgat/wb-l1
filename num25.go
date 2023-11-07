package main

import (
	"fmt"
	"time"
)

// ожидание реализованное каналом time.After, которое закроется с течением времени
func sleep(d time.Duration) {
	<-time.After(d)
}

// ожидание реализованное с помощью time.NewTimer(d) для создания нового таймера с заданным временем задержки d
// используя <-timer.C ожидается пока таймер завершит свою работу
func sleep2(d time.Duration) {
	timer := time.NewTimer(d)
	<-timer.C
}

// time.Tick(d), которая создает канал, отправляющий значения времени с интервалом d
// <-ticker выдаёт первое значение из канала ticker которое представляет текущее время после прохождения первого интервала d
func sleep3(d time.Duration) {
	ticker := time.Tick(d)
	<-ticker
}

func main() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
		sleep(1 * time.Second)
	}

	for i := 4; i <= 6; i++ {
		fmt.Println(i)
		sleep2(2 * time.Second)
	}

	for i := 7; i <= 9; i++ {
		fmt.Println(i)
		sleep3(3 * time.Second)
	}
}
