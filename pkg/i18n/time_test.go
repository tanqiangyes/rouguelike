package i18n

import (
	"testing"
	"time"
)

const i18nData = `
["time_day"]
	"cn" = "天"
	"tw" = "天"
	"hk" = "天"
	"en" = "days"

["time_hour"]
	"cn" = "小时"
	"tw" = "小時"
	"hk" = "小時"
	"en" = "hrs"

["time_min"]
	"cn" = "分钟"
	"tw" = "分鍾"
	"hk" = "分鍾"
	"en" = "mins"

["time_second"]
	"cn" = "秒"
	"tw" = "秒"
	"hk" = "秒"
	"en" = "secs"
`

func TestGetTimeSubMessage(t *testing.T) {
	// 需要初始化多语言库
	if err := LoadData([]byte(i18nData)); err != nil {
		panic(err)
	}

	type args struct {
		lan Lang
		t1  time.Time
		t2  time.Time
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1d0h", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672588800, 0),
		}, "1天"},
		{"1d0h", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672588800, 0),
		}, "1 day"},
		{"1d1h", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672592400, 0),
		}, "1天1小时"},
		{"60m", args{
			lan: LangCN,
			t2:  time.Unix(1672592400, 0),
			t1:  time.Unix(1672596000, 0),
		}, "1小时"},
		{"61m", args{
			lan: LangCN,
			t2:  time.Unix(1672592400, 0),
			t1:  time.Unix(1672596061, 0),
		}, "1小时"},
		{"119m", args{
			lan: LangCN,
			t2:  time.Unix(1672592400, 0),
			t1:  time.Unix(1672599540, 0),
		}, "1小时"},
		{"59m", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672505940, 0),
		}, "59分钟"},
		{"59m", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672505940, 0),
		}, "59 mins"},
		{"60s", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672502460, 0),
		}, "1分钟"},
		{"60s", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672502460, 0),
		}, "1 min"},
		{"59s", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672502459, 0),
		}, "59秒"},
		{"59s", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672502459, 0),
		}, "59 secs"},
		{"2d23h", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672758000, 0),
		}, "2天23小时"},
		{"2d23h", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672758000, 0),
		}, "2 days 23 hrs"},
		{"2d1h", args{
			lan: LangCN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672678800, 0),
		}, "2天1小时"},
		{"2d1h", args{
			lan: LangEN,
			t2:  time.Unix(1672502400, 0),
			t1:  time.Unix(1672678800, 0),
		}, "2 days 1 hr"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTimeSubMessage(tt.args.lan, tt.args.t1, tt.args.t2); got != tt.want {
				t.Errorf("GetTimeSubMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}

const i18nData1 = `
["time_d"]
"cn" = "天"
"tw" = "天"
"hk" = "天"
"en" = "d"

["time_h"]
"cn" = "小时"
"tw" = "小時"
"hk" = "小時"
"en" = "h"

["time_m"]
"cn" = "分"
"tw" = "分"
"hk" = "分"
"en" = "m"

["time_s"]
"cn" = "秒"
"tw" = "秒"
"hk" = "秒"
"en" = "s"
`

func Test_getDurationText(t *testing.T) {
	// 需要初始化多语言库
	if err := LoadData([]byte(i18nData1)); err != nil {
		panic(err)
	}
	type args struct {
		lan Lang
		d   time.Duration
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1d0h", args{lan: LangCN, d: 24 * time.Hour}, "1天0小时"},
		{"1d0h1m", args{lan: LangCN, d: 24*time.Hour + time.Minute}, "1天0小时"},
		{"1d1h", args{lan: LangCN, d: 24*time.Hour + time.Hour}, "1天1小时"},
		{"60m", args{lan: LangCN, d: 60 * time.Minute}, "1小时0分"},
		{"59m", args{lan: LangCN, d: 59 * time.Minute}, "59分0秒"},
		{"1m1s", args{lan: LangCN, d: time.Minute + time.Second}, "1分1秒"},
		{"1m1s e", args{lan: LangEN, d: time.Minute + time.Second}, "1m1s"},
		{"59秒", args{lan: LangCN, d: 59 * time.Second}, "59秒"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDurationText(tt.args.lan, tt.args.d); got != tt.want {
				t.Errorf("getDurationText() = %v, want %v", got, tt.want)
			}
		})
	}
}
