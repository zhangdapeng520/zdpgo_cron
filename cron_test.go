package zdpgo_cron

import (
	"fmt"
	"testing"
	"time"
)

func prepareCron() *Cron {
	c := New(CronConfig{})
	return c
}

func TestCron_New(t *testing.T) {
	c := prepareCron()
	if c == nil {
		t.Error("初始化失败")
	}
}

func TestCron_GetCronEveryMonthsDayHourMinuteSecond(t *testing.T) {
	c := prepareCron()

	// 每隔5秒
	second := c.GetCronEveryMonthsDayHourMinuteSecond(0, 0, 0, 0, 5)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("每隔5秒执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 每隔1秒
	second1 := c.GetCronEveryMonthsDayHourMinuteSecond(0, 0, 0, 0, 1)
	fmt.Println(second)

	// 添加方法
	addFunc1, err := c.AddFunc(second1, func() {
		t.Log("每隔1秒执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法1 id", addFunc1)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 30)
}

func TestCron_GetCronEverySecond(t *testing.T) {
	c := prepareCron()

	// 每隔5秒
	second := c.GetCronEverySecond(5)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("每隔5秒执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 每隔1秒
	second1 := c.GetCronEverySecond(1)
	fmt.Println(second1)

	// 添加方法
	addFunc1, err := c.AddFunc(second1, func() {
		t.Log("每隔1秒执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法1 id", addFunc1)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 30)
}

func TestCron_GetCronEveryMinute(t *testing.T) {
	c := prepareCron()

	// 每隔1分钟
	second := c.GetCronEveryMinute(1)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("每隔1分钟执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}

func TestCron_GetCronEveryMinuteMany(t *testing.T) {
	c := prepareCron()

	// 每隔1分钟
	second := c.GetCronEveryMinuteMany([]uint{1, 24, 25})
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("每小时xx分执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}

func TestCron_GetCronEveryHourMinute(t *testing.T) {
	c := prepareCron()

	// 每隔1分钟
	second := c.GetCronHourMinute(23, 30)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("每天xx时xx分执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}

func TestCron_GetCronHourEveryMinute(t *testing.T) {
	c := prepareCron()

	// 23时每隔1分钟
	second := c.GetCronHourEveryMinute(23, 1)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("23时每隔1分钟执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}

func TestCron_GetCronHourManyEveryMinute(t *testing.T) {
	c := prepareCron()

	// 23,1,2时每隔1分钟
	second := c.GetCronHourManyEveryMinute([]uint{23, 1, 2}, 1)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("23,1,2时每隔1分钟执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}

func TestCron_GetCronMinuteManyEverySecond(t *testing.T) {
	c := prepareCron()

	// 38, 39, 40分钟每隔1秒
	second := c.GetCronMinuteManyEverySecond([]uint{38, 39, 40}, 1)
	fmt.Println(second)

	// 添加方法
	addFunc, err := c.AddFunc(second, func() {
		t.Log("38, 39, 40分钟每隔1秒执行一次任务")
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("方法id", addFunc)

	// 启动任务
	c.Start()

	// 停30秒
	time.Sleep(time.Second * 180)
}
