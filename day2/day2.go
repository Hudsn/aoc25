package day2

import (
	_ "embed"
	"fmt"
	"log"
	"strconv"
	"strings"
)

//go:embed day2.txt
var inputBytes []byte

func solve(input string) int {
	count := 0
	rangeStrings := strings.Split(input, ",")
	for _, idRangeStr := range rangeStrings {
		idRangeStr = strings.TrimSpace(idRangeStr)
		r, ok, err := stringToRangeStruct(idRangeStr)
		if err != nil {
			log.Fatalf("issue converting string to range integers: %s", err)
		}
		if !ok { // skip instances where the range is going to be ignored anyway
			continue
		}
		count += r.sumInvalidIds()
	}
	return count
}

func Solve() {
	fmt.Printf("FLAG: %d\n", solve(string(inputBytes)))
}

type idRange struct {
	min         int
	max         int
	lowestLeft  int
	highestLeft int
}

func (r idRange) sumInvalidIds() int {
	ret := 0
	for i := r.lowestLeft; i <= r.highestLeft; i++ {
		asStr := strconv.Itoa(i)
		test, err := strconv.Atoi(asStr + asStr)
		if err != nil {
			log.Fatalf("unable to convert string %q to number: %s", asStr+asStr, err)
			return -1
		}
		if test >= r.min && test <= r.max {
			ret += test
		}
	}
	return ret
}

func stringToRangeStruct(idRangeStr string) (idRange, bool, error) {
	splits := strings.Split(idRangeStr, "-")
	if len(splits) != 2 {
		return idRange{}, false, fmt.Errorf("unable to split range string %q on dash character", idRangeStr)
	}
	lowStr := splits[0]
	lowLen := len(lowStr)
	highStr := splits[1]
	highLen := len(highStr)
	if lowLen == highLen && lowLen%2 != 0 {
		return idRange{}, false, nil
	}
	min, err := strconv.Atoi(lowStr)
	if err != nil {
		return idRange{}, false, fmt.Errorf("unable to convert left side id to a usable integer: %s", err)
	}
	max, err := strconv.Atoi(highStr)
	if err != nil {
		return idRange{}, false, fmt.Errorf("unable to convert right side id to a usable integer: %s", err)
	}
	lowLeft, err := extractLeftSideInt(lowStr, true)
	if err != nil {
		return idRange{}, false, fmt.Errorf("unable to convert left side id to a usable integer: %s", err)
	}
	highLeft, err := extractLeftSideInt(highStr, false)
	if err != nil {
		return idRange{}, false, fmt.Errorf("unable to convert right side id to a usable integer: %s", err)
	}
	return idRange{min: min, max: max, lowestLeft: lowLeft, highestLeft: highLeft}, true, nil
}

func extractLeftSideInt(s string, isLower bool) (int, error) {
	if len(s)%2 == 0 {
		return strconv.Atoi(s[:len(s)/2])
	}
	if isLower { // ex: 1 - 50
		pad := (len(s) + 1) / 2
		return strconv.Atoi("1" + strings.Repeat("0", pad-1))
	} // else is upper. // ex: 10 - 999
	pad := (len(s) - 1) / 2
	return strconv.Atoi(strings.Repeat("9", pad))
}
