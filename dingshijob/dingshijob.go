package dingshijob

import (
	"fmt"
	"time"
)

// WorkData 任务内容
type WorkData struct {
	time string
	work func()
}

// DingShiQi 定时器
type DingShiQi struct {
	workings []*WorkData
}

// PollingRun 轮询执行预定时任务
func (d *DingShiQi) PollingRun() {
	count := 0
	for {
		count ++
		fmt.Printf("第 %d 次轮询 ====》》\n" ,count)
		for i,v := range d.workings{
			if s := time.Now().Format("2006/01/02-15:04");v.time == s {
				v.work()
				res :=  d.workings[i+1:]
				d.workings = d.workings[:i]
				for _,m := range res {
					d.workings = append(d.workings,m)
				}
			}
		}
		fmt.Printf("切片当前长度：%d\n",len(d.workings))
		fmt.Println("轮询结束 ====》》")
		time.Sleep(time.Second * 20)
	}

}

// AddWork 增加任务,datetime 时间格式"2006/01/02-15:04" 最小单位为分
func (d *DingShiQi) AddWork(datetime string,f func()) {
	res := &WorkData{
		time: datetime,
		work: f,
	}
	d.workings = append(d.workings,res)
}

func NewDingShiQi() *DingShiQi{
	return &DingShiQi{
		make([]*WorkData,0),
	}
}