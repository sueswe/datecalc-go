package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

var Version string = "0.3.2"
var REV string = "DEV"

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func usage() {
	fmt.Println("datecalc, Version " + Version + ", written Go")
	fmt.Println("Commit: " + REV)
	fmt.Println("use -h for help.")
	fmt.Println(`
	
	## Add or subtract days, month or years to a date

	~~~
	datecalc -v 20250529 -d +1
	datecalc -v 20250202 -m 1
	datecalc -v 20241231 -y -1
	~~~

	## Number of week of given date

	~~~
	datecalc -v 20250314 -w
	~~~

	## nth weekday of a Month

	* Example:

	~~~
	datecalc -v 20250601 -nth 2 -dow Monday
	~~~

	Calculated the 2nd Monday of the given month. Given date will always start with 01 while calculating.


	## Return number of days of given date

	~~~
	datecalc -v 20250624 -c
	~~~
	`)
	os.Exit(1)
}

// https://stackoverflow.com/questions/73880828/list-the-number-of-days-in-current-date-month
func daysInMonth(t time.Time) int {
	t = time.Date(t.Year(), t.Month(), 32, 0, 0, 0, 0, time.UTC)
	daysInMonth := 32 - t.Day()
	days := make([]int, daysInMonth)
	for i := range days {
		days[i] = i + 1
	}
	tage := days[len(days)-1]
	return tage
}

func addSubtr(theDate string, day int, month int, year int) string {
	format := "20060102"
	d, err := time.Parse(format, theDate)
	CheckErr(err)

	targetDate := d.AddDate(year, month, day)
	result := targetDate.Format(time.DateOnly)

	dow := targetDate.Weekday()
	//dow_num := int(dow)

	erg := dow.String() + "," + result
	return erg
}

// Calculating nth DOW of a Month
// 1,0 = first,sunday
func nth_dow_of_month(theDate string, nth int, weekday string) string {
	format := "20060102"
	yymm01 := strings.Split(theDate, "")
	f := strings.Join(yymm01[0:6], "")
	first := f + "01"
	// fmt.Println("Startdate: " + first)
	// timeObject:
	given, err := time.Parse(format, first)
	CheckErr(err)

	for i := 0; i < 7; i++ {
		given_dow := given.Weekday().String()
		// fmt.Println("debug: given is: " + given_dow)
		if given_dow != weekday {
			given = given.AddDate(0, 0, 1)
		} else {
			// fmt.Println("Identisch " + given_dow)
			// fmt.Println("Sind nun bei: " + given.Format(time.DateOnly))
			break
		}

	}
	if nth > 5 {
		fmt.Println("Not allowed. Nth has to be between 1 and 5.")
		os.Exit(2)
	}
	result := "none"

	switch {
	case nth == 1:
		//fmt.Println(given.Format(time.DateOnly))
		result = given.Format(time.DateOnly)
	case nth == 2:
		g2 := given.AddDate(0, 0, 7)
		//fmt.Println(g2.Format(time.DateOnly))
		result = g2.Format(time.DateOnly)
	case nth == 3:
		g3 := given.AddDate(0, 0, 14)
		//fmt.Println(g3.Format(time.DateOnly))
		result = g3.Format(time.DateOnly)
	case nth == 4:
		g4 := given.AddDate(0, 0, 21)
		//fmt.Println(g4.Format(time.DateOnly))
		result = g4.Format(time.DateOnly)
	case nth == 5:
		g5 := given.AddDate(0, 0, 28)
		//fmt.Println(g5.Format(time.DateOnly))
		result = g5.Format(time.DateOnly)
	}
	return result

}

func main() {

	if len(os.Args) == 1 {
		usage()
	}

	yyyymmdd := flag.String("v", "n/a", "Parameter. Given date in format yyyymmdd")
	dayFlag := flag.Int("d", 0, "Parameter. Calculates +/-days.")
	monthFlag := flag.Int("m", 0, "Parameter. Calculates +/-months.")
	yearFlag := flag.Int("y", 0, "Parameter. Calculates +/-years.")
	weekFlag := flag.Bool("w", false, "Bool. Returns weeknumber of given date.")
	nthFlag := flag.Int("nth", 1, "Parameter. Nth occurance of weekday.")
	tFlag := flag.String("dow", "none", "Parameter. Day of Week")
	cFlag := flag.Bool("c", false, "Bool. Returns number of days in month.")
	flag.Parse()

	switch {
	case *cFlag:
		format := "20060102"
		c, e := time.Parse(format, *yyyymmdd)
		CheckErr(e)
		mnum := daysInMonth(c)
		fmt.Println(mnum)
	case *tFlag != "none":
		//is :=
		is := nth_dow_of_month(*yyyymmdd, *nthFlag, *tFlag) // 1,0 = first,sunday
		fmt.Println(is)
	case *weekFlag:
		format := "20060102"
		isoweek, err := time.Parse(format, *yyyymmdd)
		CheckErr(err)
		_, w := isoweek.ISOWeek()
		fmt.Println("Week:" + strconv.Itoa(w))
	case *dayFlag != 0:
		fmt.Println(addSubtr(*yyyymmdd, *dayFlag, 0, 0))
	case *monthFlag != 0:
		fmt.Println(addSubtr(*yyyymmdd, 0, *monthFlag, 0))
	case *yearFlag != 0:
		fmt.Println(addSubtr(*yyyymmdd, 0, 0, *yearFlag))
	default:
		fmt.Println(*yyyymmdd)
	}
}
