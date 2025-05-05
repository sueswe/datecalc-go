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

/*
Usage: datecalc.pl DATE [ -d nn ] | [ -m nn] | [ -y nn] [ --cDIM] [ --week]

where:
 DATE 	 Date to calc from (format YYYYMMDD)
 -d 	 Day to calc (+/-)
 -m 	 Month to calc (+/-)
 -y 	 Year to calc (+/-)
 -c 	 Calculates days in given month
 -w 	 Returns Week of year
 -e 	 Returns days since 1900/1/1
 -nth 3 -tdw 5 	 Return  3th friday (5) from given date (Year Month)

 Examples:
 ./datecalc.pl 20210131 -nth 3 -tdw 5
 ./datecalc.pl 20090504 -d -40 	 subtract 40 days
 ./datecalc.pl 20090506 -m 2 	 summate 2 month
 ./datecalc.pl 20090506 -y 1 	 summate 1 year
 ./datecalc.pl 20120205 --cDIM 	 Calculates days for february 2012
 ./datecalc.pl 20111217 -w 	 Returns number of week for given date
 ./datecalc.pl -e 20130419 	 Returns days since 1900/1/1

*/

var Version string = "0.1"

func CheckErr(e error) {
	if e != nil {
		log.Fatalln(e)
	}
}

func usage() {
	fmt.Println("datecalc, Version " + Version + ", written with .ʕ◔ϖ◔ʔ")
	fmt.Println("use -h for help.")
	os.Exit(1)
}

func addSubtr(theDate string, day int, month int, year int) string {
	format := "20060102"
	d, err := time.Parse(format, theDate)
	CheckErr(err)

	targetDate := d.AddDate(year, month, day)
	result := targetDate.Format(time.DateOnly)

	dow := d.Weekday()
	//dow_num := int(dow)

	erg := dow.String() + "," + result
	return erg
}

// Calculating nth DOW of a Month
// 1,0 = first,sunday
func nth_dow_of_month(theDate string, nth int, weekday string) {
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
	switch {
	case nth == 1:
		fmt.Println(given.Format(time.DateOnly))
	case nth == 2:
		g2 := given.AddDate(0, 0, 7)
		fmt.Println(g2.Format(time.DateOnly))
	case nth == 3:
		g3 := given.AddDate(0, 0, 14)
		fmt.Println(g3.Format(time.DateOnly))
	case nth == 4:
		g4 := given.AddDate(0, 0, 21)
		fmt.Println(g4.Format(time.DateOnly))
	case nth == 5:
		g5 := given.AddDate(0, 0, 28)
		fmt.Println(g5.Format(time.DateOnly))
	}

}

func main() {

	if len(os.Args) == 1 {
		usage()
	}

	yyyymmdd := flag.String("v", "n/a", "Given date in format yyyymmdd")
	dayFlag := flag.Int("d", 0, "Calculate +/-days.")
	monthFlag := flag.Int("m", 0, "Calculate +/-months.")
	yearFlag := flag.Int("y", 0, "Calculate +/-years.")
	weekFlag := flag.Bool("w", false, "Get weeknumber of given date.")
	nthFlag := flag.Int("nth", 1, "Nth hit of weekday")
	tFlag := flag.String("dow", "none", "Day of Week")
	flag.Parse()

	switch {
	case *tFlag != "none":
		//is :=
		nth_dow_of_month(*yyyymmdd, *nthFlag, *tFlag) // 1,0 = first,sunday
		//fmt.Println(is)
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
