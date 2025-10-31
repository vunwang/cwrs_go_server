package cwrs_utils

import (
	"fmt"
	"time"
)

var weekDayMap = map[string]string{
	"Sunday":    "星期天",
	"Monday":    "星期一",
	"Tuesday":   "星期二",
	"Wednesday": "星期三",
	"Thursday":  "星期四",
	"Friday":    "星期五",
	"Saturday":  "星期六",
}

var monthMap = map[string]string{
	"January":   "一月",
	"February":  "二月",
	"March":     "三月",
	"April":     "四月",
	"May":       "五月",
	"June":      "六月",
	"July":      "七月",
	"August":    "八月",
	"September": "九月",
	"October":   "十月",
	"November":  "十一月",
	"December":  "十二月",
}

const (
	YYYY_MM_DD_HH_MM_SS = "2006-01-02 15:04:05"
	YYYY_MM_DD          = "2006-01-02"
)

// GetNowDateTime 获取当前: 年-月-日 时:分:秒	返回类型:string
func GetNowDateTime() string {
	var now = time.Now().Format(YYYY_MM_DD_HH_MM_SS)
	return now
}

// GetClock 获取当前天的 时、分、秒。	返回类型: int ,int ,int
func GetClock() (hour, min, sec int) {
	return time.Now().Clock()
}

// GetNowYear 获取当前年	返回类型:int
func GetNowYear() int {
	var now int = time.Now().Year()
	return now
}

// GetNowMonth 获取当前月	返回类型:int
func GetNowMonth() int {
	var now int = int(time.Now().Month())
	return now
}

// GetNowDay 获取当前日	返回类型:int
func GetNowDay() int {
	var now int = time.Now().Day()
	return now
}

// GetNowHour 获取当前时	返回类型:int
func GetNowHour() int {
	var now int = time.Now().Hour()
	return now
}

// GetNowMinute 获取当前分	返回类型:int
func GetNowMinute() int {
	var now int = time.Now().Minute()
	return now
}

// GetNowSecond 获取当前秒	返回类型:int
func GetNowSecond() int {
	var now int = time.Now().Second()
	return now
}

// GetWeekday 获取当前是星期几	返回类型:string , int
func GetWeekday() (strValue string, intValue int) {
	weekday := time.Now().Weekday()
	str := weekday.String()
	i := int(weekday)
	return weekDayMap[str], i
}

// GetMonth 获取当前是几月	返回类型:string , int
func GetMonth() (strValue string, intValue int) {
	month := time.Now().Month()
	str := month.String()
	i := int(month)
	return monthMap[str], i
}

// GetISOWeek 返回时间点t对应的ISO 9601标准下的年份和星期编号。
// 星期编号范围[1,53]，1月1号到1月3号可能属于上一年的最后一周，12月29号到12月31号可能属于下一年的第一周。
func GetISOWeek() (year, week int) {
	return time.Now().ISOWeek()
}

// GetDayOfYear 返回时间点t对应的那一年的第几天，平年的返回值范围[1,365]，闰年[1,366]。
func GetDayOfYear() int {
	return time.Now().YearDay()
}

// GetUnixSecond 获取时间戳:1970-now秒	返回类型:int64
func GetUnixSecond() int64 {
	var unix int64 = time.Now().Unix()
	fmt.Printf("type:%T value=%v\n", unix, unix)
	return unix
}

// GetUnixMilli 获取时间戳:1970-now毫秒	返回类型:int64
func GetUnixMilli() int64 {
	var unix int64 = time.Now().UnixMilli()
	return unix
}

// GetUnixMicro 获取时间戳:1970-now微秒	返回类型:int64
func GetUnixMicro() int64 {
	var unix int64 = time.Now().UnixMicro()
	return unix
}

// GetUnixNano 获取时间戳:1970-now纳秒	返回类型:int64
func GetUnixNano() int64 {
	var unix int64 = time.Now().UnixNano()
	return unix
}

// SetDate_yyyy_MM_dd 设置日期:年月日
func SetDate_yyyy_MM_dd(yyyy, MM, dd int) time.Time {
	local := time.Local
	date := time.Date(yyyy, time.Month(MM), dd, 0, 0, 0, 0, local)
	return date
}

// SetDateYear_Month_Day 设置日期:年月日
func SetDateYear_Month_Day(yyyy int, MM time.Month, dd int) time.Time {
	local := time.Local
	date := time.Date(yyyy, MM, dd, 0, 0, 0, 0, local)
	return date
}

// SetDate_yyyy_MM_dd_hh_mm_ss  设置日期:年月日 时分秒
func SetDate_yyyy_MM_dd_hh_mm_ss(yyyy, MM, dd, hour, min, sec int) time.Time {
	local := time.Local
	date := time.Date(yyyy, time.Month(MM), dd, hour, min, sec, 0, local)
	return date
}

// ParseStrToDate 传参日期和格式必须一致:
// "2006-01-02 15:04:05" 格式:YYYY_MM_DD_HH_MM_SS
// "2006-01-02 15:04" 格式:YYYY_MM_DD_HH_MM
// "2006-01-02":YYYY_MM_DD
// 艹, 这里有坑 先把拿到的string日期打印出来, 看看格式. , 参考:			location, err := time.ParseInLocation("2006-01-02T15:04:05+08:00", *item.StartTime, time.Local)
func ParseStrToDate(dateTime string, format string) (date time.Time, err error) {
	parse, err := time.Parse(format, dateTime)
	return parse, err
}

// ChangeDateTime 增加或减少日期时间:  正数+  , 负数-  , 0不处理
func ChangeDateTime(now time.Time, year, month, day, hour, minute, second int) time.Time {
	return now.AddDate(year, month, day).Add(time.Duration(hour) * time.Hour).Add(time.Duration(minute) * time.Minute).Add(time.Duration(second) * time.Second)
}

// GetBetweenStartAndEndDates 根据开始日期和结束日期计算出时间段内所有日期
// 参数为日期格式，如：2020-01-01
func GetBetweenStartAndEndDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

// GetStartAndEndDaysOfMonth 获得当前月的初始和结束日期
func GetStartAndEndDaysOfMonth() (string, string) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)
	f := firstOfMonth.Unix()
	l := lastOfMonth.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetStartAndEndDaysOfWeek 获得当前周的初始和结束日期
func GetStartAndEndDaysOfWeek() (string, string) {
	now := time.Now()
	offset := int(time.Monday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if offset > 0 {
		offset = -6
	}

	lastoffset := int(time.Saturday - now.Weekday())
	//周日做特殊判断 因为time.Monday = 0
	if lastoffset == 6 {
		lastoffset = -1
	}

	firstOfWeek := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, offset)
	lastOfWeeK := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local).AddDate(0, 0, lastoffset+1)
	f := firstOfWeek.Unix()
	l := lastOfWeeK.Unix()
	return time.Unix(f, 0).Format("2006-01-02") + " 00:00:00", time.Unix(l, 0).Format("2006-01-02") + " 23:59:59"
}

// GetStartAndEndDaysOfSeason 获得当前季度的初始和结束日期
func GetStartAndEndDaysOfSeason() (string, string) {
	year := time.Now().Format("2006")
	month := int(time.Now().Month())
	var firstOfQuarter string
	var lastOfQuarter string
	if month >= 1 && month <= 3 {
		//1月1号
		firstOfQuarter = year + "-01-01 00:00:00"
		lastOfQuarter = year + "-03-31 23:59:59"
	} else if month >= 4 && month <= 6 {
		firstOfQuarter = year + "-04-01 00:00:00"
		lastOfQuarter = year + "-06-30 23:59:59"
	} else if month >= 7 && month <= 9 {
		firstOfQuarter = year + "-07-01 00:00:00"
		lastOfQuarter = year + "-09-30 23:59:59"
	} else {
		firstOfQuarter = year + "-10-01 00:00:00"
		lastOfQuarter = year + "-12-31 23:59:59"
	}
	return firstOfQuarter, lastOfQuarter
}

// Equal 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。
func Equal(t time.Time, u time.Time) bool {
	return t.Equal(u)
}

// Before 如果t代表的时间点在u之前，返回真；否则返回假。
func Before(t time.Time, u time.Time) bool {
	return t.Before(u)
}

// After 如果t代表的时间点在u之后，返回真；否则返回假。
func After(t time.Time, u time.Time) bool {
	return t.After(u)
}
