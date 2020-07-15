package numcover

import (
	"regexp"
	"strconv"
	"strings"
)

/*ConvertNumToCap 财务金额转最大单位 */
func ConvertNumToCap(num float64) string {
	strnum := strconv.FormatFloat(num, 'f', 2, 64)
	capitalSlice := []string{"万", "亿", "兆"}
	index := 0
	result := ""
	sdivision := strings.Split(strnum, ".")
	sl := sdivision[0]
	if len(sdivision) > 1 {
		result = "." + sdivision[1]
	}
	// slength := len(sl)
	for len(sl) > 4 {
		result = capitalSlice[index] + sl[len(sl)-4:] + result
		index = index + 1
		sl = sl[0 : len(sl)-4]
	}
	result = sl + result
	result = strings.Replace(result, "万0000", "万", -1)
	result = strings.Replace(result, "亿0000", "亿", -1)
	result = strings.Replace(result, "兆0000", "兆", -1)
	result = strings.Replace(result, "亿万", "亿", -1)
	result = strings.Replace(result, "兆亿", "兆", -1)
	return result
}

/*ConvertNumToCny 财务金额大写 */
func ConvertNumToCny(num float64) string {
	strnum := strconv.FormatFloat(num*100, 'f', 0, 64)
	sliceUnit := []string{"仟", "佰", "拾", "亿", "仟", "佰", "拾", "万", "仟", "佰", "拾", "元", "角", "分"}
	// log.Println(sliceUnit[:len(sliceUnit)-2])
	s := sliceUnit[len(sliceUnit)-len(strnum) : len(sliceUnit)]
	upperDigitUnit := map[string]string{"0": "零", "1": "壹", "2": "贰", "3": "叁", "4": "肆", "5": "伍", "6": "陆", "7": "柒", "8": "捌", "9": "玖"}
	str := ""
	for k, v := range strnum[:] {
		str = str + upperDigitUnit[string(v)] + s[k]
	}
	reg, err := regexp.Compile(`零角零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零角`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零分$`)
	str = reg.ReplaceAllString(str, "整")

	reg, err = regexp.Compile(`零[仟佰拾]`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零{2, }`)
	str = reg.ReplaceAllString(str, "零")

	reg, err = regexp.Compile(`零亿`)
	str = reg.ReplaceAllString(str, "亿")

	reg, err = regexp.Compile(`零万`)
	str = reg.ReplaceAllString(str, "万")

	reg, err = regexp.Compile(`零*元`)
	str = reg.ReplaceAllString(str, "元")

	reg, err = regexp.Compile(`亿零{0, 3}万`)
	str = reg.ReplaceAllString(str, "^元")

	reg, err = regexp.Compile(`零元`)
	str = reg.ReplaceAllString(str, "零")
	if err != nil {
		return "/"
	}
	return str
}
