# datecalc-go

## Synopsis


datecalc -v YYYYMMDD [-d|-m|-y +/-int] [-w] [-nth int -dow Weekday]

where 'Weekday' hast to be one of the following:

~~~
const (
	Sunday
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)
~~~

## Usage

~~~
-d int
    	Calculate +/-days.
-dow string
    Day of Week (default "none")
-m int
    Calculate +/-months.
-nth int
    Nth hit of weekday (default 1)
-v string
    Given date in format yyyymmdd (default "n/a")
-w	Get weeknumber of given date.
-y int
    Calculate +/-years.
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

 