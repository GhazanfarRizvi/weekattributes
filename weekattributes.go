package weekattributes

import "time"

// Weekattribs is to provide firstDateOfThisWeek, weekday and wk_num
func Weekattribs() (firstDateOfThisWeek time.Time, weekday time.Weekday, wk_num int ) {
	todayDate := time.Now().Local()
	dateForCal := todayDate.AddDate(0, 0, 0)
	weekday = dateForCal.Weekday()
	_, wk_num = dateForCal.ISOWeek()
	if weekday == time.Sunday {
		wk_num++
	}
	switch wd := dateForCal.Weekday(); wd {
	case time.Sunday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, 0)

	case time.Monday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -1)

	case time.Tuesday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -2)

	case time.Wednesday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -3)

	case time.Thursday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -4)

	case time.Friday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -5)

	case time.Saturday:
		firstDateOfThisWeek = dateForCal.AddDate(0, 0, -6)
	}
	return firstDateOfThisWeek, weekday, wk_num
}
