package zdpgo_cron

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"github.com/zhangdapeng520/zdpgo_zap"
)

type Cron struct {
	log    *zdpgo_zap.Zap // 日志
	Config *CronConfig    // 配置
	c      *cron.Cron     // cron核心对象
}

func New(config CronConfig) *Cron {
	c := Cron{}

	// 日志
	if config.LogFilePath == "" {
		config.LogFilePath = "logs/zdpgo/zdpgo_cron.log"
	}
	c.log = zdpgo_zap.New(zdpgo_zap.ZapConfig{
		Debug:       config.Debug,
		LogFilePath: config.LogFilePath,
	})

	// 配置
	c.Config = &config

	// cron
	c.c = cron.New(cron.WithSeconds())

	return &c
}

// GetCronEverySecond 每隔指定秒执行的表达式
func (c *Cron) GetCronEverySecond(second uint) string {
	s := fmt.Sprintf("*/%d * * * * ?", second)
	return s
}

// GetCronEveryMinute 每隔指定分钟执行表达式
func (c *Cron) GetCronEveryMinute(minute uint) string {
	s := fmt.Sprintf("0 */%d * * * ?", minute)
	return s
}

// GetCronEveryHour 每隔指定小时执行表达式
func (c *Cron) GetCronEveryHour(hour uint) string {
	s := fmt.Sprintf("0 0 */%d * * ?", hour)
	return s
}

// GetCronEveryDay 每隔指定天执行表达式
func (c *Cron) GetCronEveryDay(day uint) string {
	s := fmt.Sprintf("0 0 0 */%d * ?", day)
	return s
}

// GetCronEveryMonth 每隔指定月执行表达式
func (c *Cron) GetCronEveryMonth(month uint) string {
	s := fmt.Sprintf("0 0 0 1 */%d ?", month)
	return s
}

// GetCronEveryDayHour 每月指定日，指定时执行表达式
func (c *Cron) GetCronEveryDayHour(day, hour uint) string {
	// "0 0 1 1 * ?" // 每月1号凌晨1点执行一次
	s := fmt.Sprintf("0 0 %d %d * ?", day, hour)
	return s
}

// GetCronEveryMinuteMany 在每小时的某几个分钟执行任务表达式
func (c *Cron) GetCronEveryMinuteMany(minutes []uint) string {
	// c9 = "0 26,29,33 * * * ?" // 在26分、29分、33分执行一次
	m := uintArr2String(minutes)
	s := fmt.Sprintf("0 %s * * * ?", m)
	return s
}

// GetCronEveryHourMany 在每天某几个小时执行任务的表达式
func (c *Cron) GetCronEveryHourMany(hours []uint) string {
	// c10 = "0 0 0,13,18,21 * * ?" // 每天的0点、13点、18点、21点都执行一次
	m := uintArr2String(hours)
	s := fmt.Sprintf("0 0 %s * * ?", m)
	return s
}

// GetCronHourMinute 每天某时某分执行任务的表达式
func (c *Cron) GetCronHourMinute(hour, minute uint) string {
	// c12 = "0 15 10 * * ?" // 每天早上10：15触发
	s := fmt.Sprintf("0 %d %d * * ?", minute, hour)
	return s
}

// GetCronHourEveryMinute 每天某时每隔minute分钟执行任务的表达式
func (c *Cron) GetCronHourEveryMinute(hour, minute uint) string {
	// c16 = "0 0/5 14 * * ?" // 每天从下午2点开始到2：55分结束每5分钟一次触发
	s := fmt.Sprintf("0 0/%d %d * * ?", minute, hour)
	return s
}

// GetCronHourManyEveryMinute 每天某几时每隔minute分钟执行任务的表达式
func (c *Cron) GetCronHourManyEveryMinute(hours []uint, minute uint) string {
	// c17 = "0 0/5 14,18 * * ?" // 每天的下午2点至2：55和6点至6点55分两个时间段内每5分钟一次触发
	m := uintArr2String(hours)
	s := fmt.Sprintf("0 0/%d %s * * ?", minute, m)
	return s
}

// GetCronMinuteManyEverySecond 每小时某几分钟每隔second秒执行任务表达式
func (c *Cron) GetCronMinuteManyEverySecond(minutes []uint, second uint) string {
	// c17 = "0 0/5 14,18 * * ?" // 每天的下午2点至2：55和6点至6点55分两个时间段内每5分钟一次触发
	m := uintArr2String(minutes)
	s := fmt.Sprintf("0/%d %s * * * ?", second, m)
	return s
}

// AddFunc 添加定时任务
func (c *Cron) AddFunc(expression string, task func()) (int, error) {
	addFunc, err := c.c.AddFunc(expression, task)
	if err != nil {
		c.log.Error("AddFunc 添加方法失败", "error", err.Error())
	}

	return int(addFunc), err
}

// Start 开启cron定时任务
func (c *Cron) Start() {
	c.c.Start()
}
