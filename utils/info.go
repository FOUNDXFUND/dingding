package utils

import (
	"reflect"
	"unsafe"
	"os"
	"time"
)

func SetVersion(ver string) {
	argv0str := (*reflect.StringHeader)(unsafe.Pointer(&os.Args[0]))
	argv0 := (*[1 << 30]byte)(unsafe.Pointer(argv0str.Data))[:]
	line := os.Args[0]
	for i := 1; i < len(os.Args); i++ {
		line += (" " + os.Args[i])
	}
	line += (" " + ver)
	copy(argv0, line)
	argv0[len(line)] = 0
}

// 获取当天0点和24点时间戳
func GetTodayUnix() (todayBegin,todayEnd int64)   {
	t := time.Now()
	tm1 := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	todayBegin = tm1.Unix()
	todayEnd = todayBegin + (23 * 60 * 60 + 59 * 60 + 59)
	return
}