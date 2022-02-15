# zdpgo_cron
在golang中使用cron表达式并实现定时任务

项目地址：https://github.com/zhangdapeng520/zdpgo_cron

## 功能清单
- GetCronEverySecond 每隔指定秒执行的表达式
- GetCronEveryMinute 每隔指定分钟执行表达式
- GetCronEveryHour 每隔指定小时执行表达式
- GetCronEveryDay 每隔指定天执行表达式
- GetCronEveryMonth 每隔指定月执行表达式
- GetCronEveryDayHour 每月指定日，指定时执行表达式
- GetCronEveryMinuteMany 在每小时的某几个分钟执行任务表达式
- GetCronEveryHourMany 在每天某几个小时执行任务的表达式
- GetCronHourMinute 每天某时某分执行任务的表达式
- GetCronHourEveryMinute 每天某时每隔minute分钟执行任务的表达式
- GetCronHourManyEveryMinute 每天某几时每隔minute分钟执行任务的表达式
- GetCronMinuteManyEverySecond 每小时某几分钟每隔second秒执行任务表达式
- AddFunc 添加定时任务
- Start 开启cron定时任务

## 版本历史
- 版本0.1.0 2022年2月15日 基础常用功能