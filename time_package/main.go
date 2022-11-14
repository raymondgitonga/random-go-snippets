package main

import (
	"fmt"
	"time"
)

func main() {
	//parseTime()

	yom := fmt.Sprintf("2022-04%s", "-01")
	year, err := time.Parse("2006-01-02", yom)

	year.Month()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(year)

	//timeDifference()

}

func timeFormats() {
	currentTime := time.Now()
	fmt.Println("time.now()")
	fmt.Println("The time is ", currentTime)
	fmt.Println("--------------")

	fmt.Println("Time break down")
	fmt.Println("The year is", currentTime.Year())
	fmt.Println("The month is", currentTime.Month())
	fmt.Println("The day is", currentTime.Day())
	fmt.Println("The hour is", currentTime.Hour())
	fmt.Println("The minute is", currentTime.Minute())
	fmt.Println("The second is", currentTime.Second())
	fmt.Println("--------------")

	fmt.Println("Pre defined time")
	theTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.Local)
	fmt.Println("The time is", theTime)
	fmt.Println("--------------")

	fmt.Println("Format")
	fmt.Println(currentTime.Format("2006-1-2 15:4:5 pm"))
	fmt.Println("--------------")

	fmt.Println("Predefined Format")
	fmt.Println(currentTime.Format(time.RFC822))
	fmt.Println("--------------")

	fmt.Println("Format custom")
	fmt.Println(currentTime.Format("2006-Jan-Mon 15:4 pm"))
	fmt.Println("--------------")
}

func parseTime() {
	timeString := "2021-08-15 02:30:45"
	theTime, err := time.Parse("2006-01-02 03:04:05", timeString)
	if err != nil {
		fmt.Println("Could not parse time:", err)
		return
	}
	fmt.Println("The time is", theTime)

	fmt.Println(theTime.Format(time.RFC3339Nano))
	fmt.Println("--------------")

	fmt.Println("Time with Utc time zone")
	utcTime := time.Now().UTC()
	fmt.Println(utcTime.Format(time.RFC3339Nano))
	fmt.Println("--------------")

	fmt.Println("Time with specific time zone")
	timeLoc, err := time.LoadLocation("Africa/Nairobi")
	if err != nil {
		fmt.Println("Could not load timezone:", err)
		return
	}
	fmt.Println(utcTime.In(timeLoc))
	fmt.Println("--------------")
}

func timeDifference() {
	firstTime := time.Date(2021, 8, 15, 14, 30, 45, 100, time.UTC)
	fmt.Println("The first time is", firstTime)

	secondTime := time.Date(2021, 12, 25, 16, 40, 55, 200, time.UTC)
	fmt.Println("The second time is", secondTime)

	fmt.Println("Duration between first and second time is", firstTime.Sub(secondTime).Hours())
	fmt.Println("Duration between second and first time is", secondTime.Sub(firstTime).Hours())

}
