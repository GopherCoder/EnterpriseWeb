package assistance

import "strconv"

func ToInt64(i interface{}) int64 {
	ID, _ := strconv.Atoi(i.(string))
	return int64(ID)
}
func ToInt(i interface{}) int {
	return int(ToInt64(i))
}

func ToString(i interface{}) string {
	return i.(string)
}
