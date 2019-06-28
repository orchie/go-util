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
