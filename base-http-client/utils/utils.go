package utils

import (
	"net/url"
	"strconv"
	"time"
)

func AddOptionalStringParam(params url.Values, key string, value string) {
	if value != "" {
		params.Add(key, value)
	}
}

func AddOptionalIntParam(params url.Values, key string, value int) {
	if value != 0 {
		params.Add(key, strconv.Itoa(value))
	}
}

func AddOptionalInt64Param(params url.Values, key string, value int64) {
	if value != 0 {
		params.Add(key, strconv.FormatInt(value, 10))
	}
}

func AddOptionalFloat64Param(params url.Values, key string, value float64) {
	if value != 0 {
		params.Add(key, strconv.FormatFloat(value, 'f', -1, 64))
	}
}

func AddOptionalTimeParam(params url.Values, key string, value *time.Time, layout string) {
	if value != nil {
		params.Add(key, value.Format(layout))
	}
}

func AddOptionalDurationParam(params url.Values, key string, value *time.Duration) {
	if value != nil {
		params.Add(key, value.String())
	}
}
