package output

import (
"github.com/tidwall/gjson"
"regexp"
"strconv"
"strings"
)

var startWithDot = regexp.MustCompile(`\.[^\s]+`)

type (
	Formatter func(input interface{}) string
)

const devFormat = `.web_url .author.username .commit.username .commit.created_at`
const homFormat = `.web_url .author.username .commit.created_at`

func NewFormatter(fields string) Formatter {

	f := func(input interface{}) string {
		return ""
	}

	return f
}

func FormatDev(input interface{}) string {

	return format(ToJsonStr(input), parseOutputArgs(devFormat, "#.merge")...)
}

func FormatHom(input interface{}) string {

	return format(ToJsonStr(input), parseOutputArgs(homFormat, "#.merge")...)
}

func FormatJson(input interface{}) string {
	return ToJsonStr(input)
}

func FormatAuto(input interface{}) string {
	r := input.([]MRResult)

	sb := strings.Builder{}
	for _, result := range r {
		var fmt string
		if result.Merge.TargetBranch == "desenvolvimento" {
			fmt = devFormat
		} else {
			fmt = homFormat
		}

		str := format(ToJsonStr(result), parseOutputArgs(fmt, "merge")...)

		sb.WriteString(str)
	}
	return sb.String()

}

func format(input string, output ...string) string {
	get := gjson.GetMany(input, output...)

	sb := new(strings.Builder)

	colLen := len(get)

	lines := get[0].Array()

	for i := 0; i < len(lines); i++ {
		for j := 0; j < colLen; j++ {
			lineArray := get[j].Array()
			if cap(lineArray) == 0 {
				continue
			}
			tab, _ := strconv.Unquote(`"` + "\t" + `"`)
			sb.WriteString(lineArray[i].String() + tab)
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func FormatString(input string, formatter Formatter) string {
	if input == "" {
		return ""
	}
	return formatter(input)
}
func FormatOutput(input interface{}, formatter Formatter) string {
	return formatter(input)
}

func parseOutputArgs(arg, startPath string) []string {
	if arg == "" {
		return []string{}
	}

	allStr := startWithDot.FindAllString(arg, -1)

	for i, s := range allStr {
		allStr[i] = strings.Join([]string{startPath, s}, "")

	}

	return allStr
}

