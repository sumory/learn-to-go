package main

import "fmt"
import "time"
import "strconv"

//import "errors"

var TimeFormats = []string{"2006-01-02 15:04:05", "2006-01-02 15:04", "2006-01-02", "01-02", "15:04:05", "15:04", "15", "15:04:05 Jan 2, 2006 MST"}

type TimeUtils struct {
	time time.Time
}

func Init(t time.Time) *TimeUtils {
	return &TimeUtils{t}
}

func (this *TimeUtils) Now() string {
	return time.Now().Format(TimeFormats[0])
}

func (this *TimeUtils) Value() time.Time {
	return this.time
}

func (this *TimeUtils) Format() string {
	return this.time.Format(TimeFormats[0])
}

func (this *TimeUtils) Rebase(t time.Time) {
	this.time = t
}

func (this *TimeUtils) Convert(radix int) string {
	return strconv.FormatInt(this.time.Unix(), radix)
}

func main() {
	var timeUtils *TimeUtils = Init(time.Now())
	print := fmt.Println

	print(timeUtils.Now())

	timeUtils.Rebase(time.Now())
	print(timeUtils.Value().Unix())
	print(timeUtils.Convert(10))
	print(timeUtils.Convert(2))
	print(timeUtils.Convert(8))
	print(timeUtils.Convert(16))

}
