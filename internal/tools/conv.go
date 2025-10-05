package tools

import (
	"fmt"
	"strconv"
)

func AnyToInt(value any) int {
	return StrToInt(AnyToStr(value))
}

func AnyToStr(value any) string {
	return fmt.Sprint(value)
}

func StrToInt(value string) int {
	result, err := strconv.Atoi(value)
	if err != nil {
		return 0
	}
	return result
}

func IntToStr(value int) string {
	return strconv.Itoa(value)
}

func AnyToMapStr(value any) map[string]any {
	if val, ok := value.(map[string]any); ok {
		return val
	}

	return nil
}

func Conv[T any](value any) (T, bool) {
	if val, ok := value.(T); ok {
		return val, true
	}
	var res T
	return res, false
}

func AnyToArray(value any) []any {
	if val, ok := value.([]any); ok {
		return val
	}

	return nil
}
