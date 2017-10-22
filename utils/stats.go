package utils

import (
	"bytes"
	"strconv"
	"time"
	"sort"
)

type CorrectWrong struct {
	Correct string
	Wrong   string
}

type Stats struct {
	isRecording  bool
	beginTime    int64
	endTime      int64
	mistakeCount int
	mistakeTrend map[CorrectWrong]int
}

func NewStats() *Stats {
	var stats Stats
	stats.mistakeTrend = make(map[CorrectWrong]int)

	return &stats
}

func (s *Stats) String() string {
	if s.isRecording {
		panic("should not be in recording.")
	}

	var buf bytes.Buffer

	var recordTime = s.endTime - s.beginTime
	buf.WriteString("Time:\n")
	buf.WriteString("\t" + strconv.FormatInt(recordTime, 10) + " seconds.\n")
	buf.WriteString("Mistake:\n")
	buf.WriteString("\t" + strconv.Itoa(s.mistakeCount) + " times.\n")
	buf.WriteString("Mistake Trend:\n")

	if len(s.mistakeTrend) == 0 {
		buf.WriteString("\tHappy NO mistake !!" + "\n")
	} else {
		var list = List{}
		for k, v := range s.mistakeTrend {
			list = append(list, entry{
				key:   "['" + k.Correct + "' -> '" + k.Wrong + "']",
				value: strconv.Itoa(v),
			})
		}
		sort.Sort(sort.Reverse(list))
		for _, v := range list {
			var key = v.key
			var value = v.value
			buf.WriteString("\t" + key + " -- " + value + " times.\n")
		}
	}

	return buf.String()
}

func (s *Stats) Begin() {
	if s.isRecording {
		panic("should not be in recording.")
	}
	s.isRecording = true
	s.beginTime = time.Now().Unix()
}

func (s *Stats) End() {
	if !s.isRecording {
		panic("should be in recording.")
	}
	s.isRecording = false
	s.endTime = time.Now().Unix()
}

func (s *Stats) AddErrorCount(cw CorrectWrong) {
	s.mistakeCount += 1
	s.mistakeTrend[cw] += 1
}
