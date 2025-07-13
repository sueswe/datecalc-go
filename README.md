# datecalc, written in go

![build workflow](https://github.com/sueswe/datecalc-go/actions/workflows/go.yml/badge.svg?event=push)

## Synopsis

`datecalc -v YYYYMMDD [-d|-m|-y +/-int] [-w] [-nth int -dow Weekday]`

where 'Weekday' hast to be one of the following:

~~~
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
~~~

## Usage

~~~
  -c	Bool. Returns number of days in month.
  -d int
    	Parameter. Calculates +/-days.
  -dow string
    	Parameter. Day of Week (default "none")
  -m int
    	Parameter. Calculates +/-months.
  -nth int
    	Parameter. Nth occurance of weekday. (default 1)
  -v string
    	Parameter. Given date in format yyyymmdd (default "n/a")
  -w	Bool. Returns weeknumber of given date.
  -y int
    	Parameter. Calculates +/-years.
~~~

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
