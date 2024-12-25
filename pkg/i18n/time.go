package i18n

import (
	"fmt"
	"math"
	"strings"
	"time"
)

// countdown 占位文本
const countdown = "countdown_%s"

// GetAllTimeSubMessage 返回一个多语言的时间相差文本
func GetAllTimeSubMessage(d time.Duration) map[string]string {
	var m = make(map[string]string)
	for _, s := range []Lang{LangCN, LangHK, LangTW, LangEN} {
		m[fmt.Sprintf(countdown, s)] = getDurationText(s, d)
	}
	return m
}

// FormatTime 格式化时间
func FormatTime(l Lang, d time.Duration) string {
	day := int(d.Hours() / 24)
	hour := int(d.Hours()) % 24
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if day > 0 {
		return fmt.Sprintf("%d%s%d%s", day, Translate("time_d", l), hour, Translate("time_h", l))
	}
	if hour > 0 {
		return fmt.Sprintf("%d%s", hour, Translate("time_h", l))
	}

	if m > 0 {
		return fmt.Sprintf("%d%s", m, Translate("time_min", l))
	}

	return fmt.Sprintf("%d%s", s, Translate("time_s", l))
}

// GetDurationText 时间处理
func GetDurationText(lan Lang, d time.Duration) string {
	return getDurationText(lan, d)
}

// getDurationText 时间处理
func getDurationText(lan Lang, d time.Duration) string {
	day := int(d.Hours() / 24)
	hour := int(d.Hours()) % 24
	m := int(d.Minutes()) % 60
	s := int(d.Seconds()) % 60
	if day > 0 {
		return fmt.Sprintf("%d%s%d%s", day, Translate("time_d", lan), hour, Translate("time_h", lan))
	}

	if hour > 0 {
		return fmt.Sprintf("%d%s%d%s", hour, Translate("time_h", lan), m, Translate("time_m", lan))
	}

	if m > 0 {
		return fmt.Sprintf("%d%s%d%s", m, Translate("time_m", lan), s, Translate("time_s", lan))
	}

	return fmt.Sprintf("%d%s", s, Translate("time_s", lan))
}

// GetTimeSubMessage 获取时间相差消息
func GetTimeSubMessage(lan Lang, t1 time.Time, t2 time.Time) string {
	ss := t1.Sub(t2)

	day := int(ss.Hours() / 24)
	hour := int(ss.Hours()) % 24
	m := int(ss.Minutes()) % 60
	s := int(ss.Seconds()) % 60

	mess := ""
	if day > 0 {
		mess = fmt.Sprintf("%d%s", day, parseTimeUnit(day, "time_day", lan))
		if hour > 0 {
			if lan == LangEN {
				mess += " "
			}
			mess += fmt.Sprintf("%d%s", hour, parseTimeUnit(hour, "time_hour", lan))
		}
	} else if hour > 0 {
		mess = fmt.Sprintf("%d%s", hour, parseTimeUnit(hour, "time_hour", lan))
	} else if m > 0 {
		mess = fmt.Sprintf("%d%s", m, parseTimeUnit(m, "time_min", lan))
	} else {
		mess = fmt.Sprintf("%d%s", s, parseTimeUnit(s, "time_second", lan))
	}

	return mess
}

// FormatTimeWithoutDay 沒有天的时间格式化
func FormatTimeWithoutDay(l Lang, t1 time.Time, t2 time.Time) string {
	diff := t1.Sub(t2)

	m := int(math.Abs(diff.Minutes()))
	if m <= 60 {
		if m <= 0 {
			m = 1
		}
		// 2分钟≤[时间]≤1小时：[分]分钟
		// 时间＜2分钟：1分钟
		minuteUnit := handleUnitOfEn(Translate("time_min", l), l, m)
		return fmt.Sprintf("%d%s", m, minuteUnit)
	}
	// 时间＞1小时：[时]小时[分]分
	h := int(math.Abs(diff.Hours()))

	hourUnit := handleUnitOfEn(Translate("time_hour", l), l, h)
	m = m % 60
	minuteUnit := handleUnitOfEn(Translate("time_minute", l), l, m)
	var dot string
	if l == LangEN {
		//xx hrs xx mins
		dot = " "
	}
	return fmt.Sprintf("%d%s%s%d%s", h, hourUnit, dot, m, minuteUnit)

}

func handleUnitOfEn(unit string, l Lang, num int) string {
	if num <= 1 && l == LangEN {
		unit = strings.TrimRight(unit, "s")
	}
	if l == LangEN {
		unit = fmt.Sprintf(" %s", unit)
	}
	return unit
}

func parseTimeUnit(n int, id TextID, lan Lang) string {
	unit := Translate(id, lan)
	if lan == LangEN {
		if n == 1 {
			unit = strings.TrimRight(unit, "s")
		}
		unit = " " + unit
	}
	return unit
}
