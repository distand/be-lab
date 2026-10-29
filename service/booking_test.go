package service

import (
	"fmt"
	"testing"
	"time"
)

func TestCheckCrossNight(t *testing.T) {
	tests := []struct {
		name       string
		rule       string
		start      string
		end        string
		crossNight bool
		wantErr    bool
	}{
		{"未配置", "", "2026-10-29 20:00:00", "2026-10-30 10:00:00", false, false},
		{"完整过夜", "21-9", "2026-10-29 21:00:00", "2026-10-30 09:00:00", true, false},
		{"整段白天", "21-9", "2026-10-29 09:00:00", "2026-10-29 21:00:00", false, false},
		{"白天预约", "21-9", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, false},
		{"早晨交叉", "21-9", "2026-10-29 08:00:00", "2026-10-29 10:00:00", false, true},
		{"晚间交叉", "21-9", "2026-10-29 20:00:00", "2026-10-29 22:00:00", false, true},
		{"部分过夜", "21-9", "2026-10-29 22:00:00", "2026-10-30 08:00:00", false, true},
		{"过夜提前开始", "21-9", "2026-10-29 20:00:00", "2026-10-30 09:00:00", false, true},
		{"过夜延后结束", "21-9", "2026-10-29 21:00:00", "2026-10-30 10:00:00", false, true},
		{"秒级交叉", "21-9", "2026-10-29 09:00:00", "2026-10-29 21:00:01", false, true},
		{"多日预约", "21-9", "2026-10-29 09:00:00", "2026-10-31 09:00:00", false, true},
		{"跨月过夜", "21-9", "2026-10-31 21:00:00", "2026-11-01 09:00:00", true, false},
		{"跨年过夜", "21-9", "2026-12-31 21:00:00", "2027-01-01 09:00:00", true, false},
		{"午夜结束", "21-0", "2026-10-29 21:00:00", "2026-10-30 00:00:00", true, false},
		{"格式错误", "21:00-9:00", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, true},
		{"缺少结束时间", "21", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, true},
		{"小时越界", "24-9", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, true},
		{"非跨夜范围", "9-21", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, true},
		{"相同小时", "9-9", "2026-10-29 10:00:00", "2026-10-29 11:00:00", false, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parse := func(value string) int64 {
				v, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
				if err != nil {
					t.Fatal(err)
				}
				return v.Unix()
			}
			got, err := checkCrossNight(tt.rule, parse(tt.start), parse(tt.end))
			if err != nil {
				fmt.Println(err.Error())
			}
			if got != tt.crossNight || (err != nil) != tt.wantErr {
				t.Fatalf("checkCrossNight() = (%v, %v), want (%v, error=%v)", got, err, tt.crossNight, tt.wantErr)
			}
		})
	}
}
