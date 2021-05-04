// Program to find leap year
// on every year that is evenly divisible by 4
// except every year that is evenly divisible by 100
// unless the year is also evenly divisible by 400

package leap

// LeapYear returns bool to check for leap year
func IsLeapYear(a int) bool {
	return a % 4 == 0 && a % 100 != 0 || a % 400 == 0
}
