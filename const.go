package zdpgo_cron

const (
	c1 = "*/5 * * * * ? " // 每隔5秒执行一次
	c2 = "0 */1 * * * ? " // 每隔1分钟执行一次
	c3 = "0 0 12 * * ?" //每天中午十二点触发
	c4 = "0 0 23 * * ?" // 每天23点执行一次
	c5 = "0 0 1 * * ?" // 每天凌晨1点执行一次
	c6 = "0 0 1 1 * ?" // 每月1号凌晨1点执行一次
	c7 = "0 0 23 L * ?" // 每月最后一天23点执行一次
	c8 = "0 0 1 ? * L" // 每周星期天凌晨1点实行一次
	c9 = "0 26,29,33 * * * ?" // 在26分、29分、33分执行一次
	c10 = "0 0 0,13,18,21 * * ?" // 每天的0点、13点、18点、21点都执行一次
	c11 = "0 15 10 ? * *" // 每天早上10：15触发
	c12 = "0 15 10 * * ?" // 每天早上10：15触发
	c13 = "0 15 10 * * ? *" // 每天早上10：15触发
	c14 = "0 15 10 * * ? 2005" // 2005年的每天早上10：15触发
	c15 = "0 * 14 * * ?" // 每天从下午2点开始到2点59分每分钟一次触发
	c16 = "0 0/5 14 * * ?" // 每天从下午2点开始到2：55分结束每5分钟一次触发
	c17 = "0 0/5 14,18 * * ?" // 每天的下午2点至2：55和6点至6点55分两个时间段内每5分钟一次触发
	c18 = "0 0-5 14 * * ?" // 每天14:00至14:05每分钟一次触发
	c19 = "0 10,44 14 ? 3 WED" // 三月的每周三的14：10和14：44触发
)
