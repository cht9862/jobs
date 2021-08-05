package dingshijob

import (
	"fmt"
	"testing"
)

func TestDingShiQi_AddWork(t *testing.T) {
	ds := NewDingShiQi()
	ds.AddWork("2021/08/05-17:13", func() {
		fmt.Println("我是定时器任务，已经开始执行了")
	})
	ds.PollingRun()


}
