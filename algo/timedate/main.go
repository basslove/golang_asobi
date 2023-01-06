package main

import (
	"fmt"
	"time"
)

// 現在時刻取得
func no1() {
	day := time.Now()
	const l = "2006-01-02"
	fmt.Println(day.Format(l))
}

// 秒の後にナノセカンドまで指定
func no2() {
	t := time.Date(2018, 5, 1, 23, 59, 59, 0, time.UTC)
	fmt.Println(t)
	t = time.Date(20018, 5, 11, 23, 59, 59, 0, time.Local)
	fmt.Println(t)
}

// 時刻を任意のフォーマット
func no3() {
	t := time.Now()
	const l1 = "Now, Monday Jan 02 15:04:05 JST 2006"
	fmt.Println(t.Format(l1))

	const l2 = "2006-01-02 15:04:05"
	fmt.Println(t.Format(l2))
}

// 時刻オブジェクトを文字列に変換
func no4() {
	t := time.Now()
	s := t.String()
	fmt.Println(s)
}

// 時刻に任意の時間を加減
// 時間差を表現する型はtime.Durationで、単位はナノ秒です。 time.Second,time.Hour等を掛け算することで秒単位、時間単位を表現します。
func no5() {
	t := time.Date(2018, 5, 20, 23, 59, 59, 0, time.Local)
	fmt.Println(t.Add(time.Duration(1) * time.Second))

	t1 := time.Date(2018, 12, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(t1.Add(time.Duration(24) * time.Hour))
}

// 2つの時刻の差
func no6() {
	day1 := time.Date(2000, 12, 31, 0, 0, 0, 0, time.Local)
	day2 := time.Date(2001, 1, 2, 12, 30, 0, 0, time.Local)
	duration := day2.Sub(day1)
	fmt.Println(duration)

	hours0 := int(duration.Hours())
	days := hours0 / 24
	hours := hours0 % 24
	mins := int(duration.Minutes()) % 60
	secs := int(duration.Seconds()) % 60
	fmt.Printf("%d days + %d hours + %d minutes + %d seconds\n", days, hours, mins, secs)
}

// 時刻中の曜日を日本語に変換する
func no7() {
	wdays := [...]string{"日", "月", "火", "水", "木", "金", "土"}

	t := time.Now()
	fmt.Println(t.Weekday())
	fmt.Println(wdays[t.Weekday()])
}

// 日付オブジェクトを作成
func no8() {
	day := time.Date(2018, 5, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(day)
}

// 指定の日付が存在?
func no9() {
	isExist := func(year, month, day int) (float64, error) {
		julian := func(t time.Time) float64 {
			const julian = 2453738.4195
			unix := time.Unix(1136239445, 0)
			const oneDay = float64(86400. * time.Second)
			return julian + float64(t.Sub(unix))/oneDay
		}

		date := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.Local)
		if date.Year() == year && date.Month() == time.Month(month) && date.Day() == day {
			return julian(date), nil
		} else {
			return 0, fmt.Errorf("%d-%d-%d is not exist", year, month, day)
		}
	}

	jd, err := isExist(2001, 1, 31)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int(jd))
	}
	// => "2451940"
	jd, err = isExist(2001, 1, 32)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(int(jd))
	}
}

// 何日後、何日前の日付
func no10() {
	t := time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, 1)
	fmt.Println(t)

	t = time.Date(2001, 1, 1, 0, 0, 0, 0, time.Local)
	t = t.AddDate(0, 0, -1)
	fmt.Println(t)
}

// 何ヶ月後、何ヶ月前の日付
func no11() {
	AddMonth := func(t time.Time, d_month int) time.Time {
		getLastDay := func(year, month int) int {
			t := time.Date(year, time.Month(month+1), 1, 0, 0, 0, 0, time.Local)
			t = t.AddDate(0, 0, -1)
			return t.Day()
		}

		year := t.Year()
		month := t.Month()
		day := t.Day()
		newMonth := int(month) + d_month
		newLastDay := getLastDay(year, newMonth)
		var newDay int
		if day > newLastDay {
			newDay = newLastDay
		} else {
			newDay = day
		}
		return time.Date(year, time.Month(newMonth), newDay, t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), t.Location())
	}

	t0 := time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local)
	t1 := t0.AddDate(0, 1, 0)
	t2 := AddMonth(t0, 1)
	fmt.Println(t1)
	fmt.Println(t2)
	t0 = time.Date(2001, 5, 31, 0, 0, 0, 0, time.Local)
	t1 = t0.AddDate(0, -1, 0)
	t2 = AddMonth(t0, -1)
	fmt.Println(t1)
	fmt.Println(t2)
}

// うるう年かどうか
func no12() {
	isLeapYear := func(year int) bool {
		if year%400 == 0 { // 400で割り切れたらうるう年
			return true
		} else if year%100 == 0 { // 100で割り切れたらうるう年じゃない
			return false
		} else if year%4 == 0 { // 4で割り切れたらうるう年
			return true
		} else {
			return false
		}
	}

	fmt.Println(isLeapYear(2000))
	fmt.Println(isLeapYear(2001))
}

// 日付オブジェクトの年月日・曜日を個別
func no13() {
	t := time.Date(2001, 1, 31, 0, 0, 0, 0, time.Local)
	fmt.Println(t.Year())
	fmt.Println(t.Month())
	fmt.Println(t.Day())
	fmt.Println(t.Weekday())
}

// 文字列の日付を日付オブジェクトに変換
func no14() {
	str := "Thu May 24 22:56:30 JST 2001"
	layout := "Mon Jan 2 15:04:05 MST 2006"
	t, _ := time.Parse(layout, str)
	fmt.Println(t)

	str = "2003/04/18"
	layout = "2006/01/02"
	t, _ = time.Parse(layout, str)
	fmt.Println(t)
}

func main() {
	fmt.Println("time date")

	// no1()
	// no2()
	// no3()
	// no4()
	// no5()
	// no6()
	// no7()
	// no8()
	// no9()
	// no10()
	// no11()
	// no12()
	// no13()
	no14()
}
