package zdpgo_cron

import (
	"strconv"
	"strings"
)

func uintArr2String(arr []uint) string {
	if arr == nil || len(arr) == 0 {
		return "*"
	}
	var strArr []string
	for _, v := range arr {
		t := strconv.Itoa(int(v))
		strArr = append(strArr, t)
	}
	s := strings.Join(strArr, ",")
	return s
}
