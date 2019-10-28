package main

import (
	"fmt"
	"math"
	"sort"
	"time"
)

func main() {

	fmt.Println(time.Date(2019,4,-9,0,0,0, 0,time.Local))

	var att []int

	att = append(att, 1)
	att = append(att, 3)
	att = append(att, 9)
	att = append(att, 5)

	sort.Ints(att)
	fmt.Println(att)

	fmt.Println("--------")
	dtDurationiToNow, _ := time.ParseDuration(fmt.Sprintf("%-vh", 24*1+24))

	fmt.Println(time.Now().Format("2006-01-02"))

	todayStart, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02") + " 00:00:00")
	fmt.Println(todayStart)
	todayStart = todayStart.Add(dtDurationiToNow)
	fmt.Println(todayStart)

	fmt.Println(time.Now().Local().Format("2006-01-02 15:04:05"))

	todayEnd, _ := time.Parse("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 23:59:59")
	fmt.Println(todayEnd)
	todayEnd = todayEnd.Add(dtDurationiToNow)
	fmt.Println(todayEnd)


	fmt.Println(math.Ceil(3.4))


	dayStart, err := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Local().Format("2006-01-02")+" 00:00:00", time.Local)
	if err != nil {
		fmt.Println(err.Error())

	}
	fmt.Println("aa,", dayStart)
	//fmt.Println(time.Now().Format("2006-01-02")+" 00:00:00")

	todayStart, _ = time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02")+" 00:00:00", time.Local)
	todayStart = todayStart.Add(dtDurationiToNow)
	fmt.Println(todayStart)

	fmt.Println(time.Now().Local().Format("2006-01-02 15"))
}
