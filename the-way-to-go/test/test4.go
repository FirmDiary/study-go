package main

import (
	"fmt"
	"time"
)

var week time.Duration

func main() {
	t := time.Time{}
	fmt.Println(t.Format("21 Dec 2011 08:52")) // 21 Dec 2011 08:52
	s := t.Format("20111221")
	fmt.Println(t, "=>", s)

	t = time.Now()
	fmt.Println(t.Format(time.ANSIC))
	fmt.Println(t)
	fmt.Println(time.ANSIC)

	t = time.Now().UTC()
	fmt.Println(t)

	week = 60 * 60 * 24 * 7 * 1e9
	week_form_now := t.Add(week)
	fmt.Println(week_form_now)

	//time.After
	//time.Ticker

}
