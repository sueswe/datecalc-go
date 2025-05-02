package main

import (
	"log"
	"os"
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

func main() {
	log.Println("datecalc, Version " + Version)
	YYYYMMDD := os.Args[1:]
	log.Println("calculating with " + YYYYMMDD[0])

	log.Println(addSubtr(YYYYMMDD[0], 1, 0, 1))

}
