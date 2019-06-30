package goutil

/**
	map 扩展
**/

//判断值是否在map中
func InMap(value interface{}, m map[interface{}]interface{}) bool {
	for _, v := range m {
		if value == v {
			return true
		}
	}
	return false
}

//仿php array_column 然后 array_unique
//取二维数组中 2维指定字段的值的集合 并且去重
func MapColumn(m []map[string]string, key string) (res map[string]int) {
	for _, v := range m {
		res[v[key]] = 1
	}
	return
}
