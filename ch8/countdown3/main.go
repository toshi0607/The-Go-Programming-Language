package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Please return to abort.")
	// カウントダウンの関数が戻ったときtickからのイベント受信をやめるが、ティッカーのゴルーチンは残っていて
	// 受信するゴルーチンがないチャネルに送信を試み続けるのでゴルーチンのリーク状態
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("boom!")
}
