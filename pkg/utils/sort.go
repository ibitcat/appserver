// 排序

package utils

// 通用排序
type SortSlice struct {
	Slice [][]interface{} // 需要排序的数据
	Index int             // 按哪个字段排序
}

func (a SortSlice) Len() int {
	return len(a.Slice)
}

func (a SortSlice) Swap(i, j int) {
	a.Slice[i], a.Slice[j] = a.Slice[j], a.Slice[i]
}

func (a SortSlice) Less(i, j int) bool {
	switch a.Slice[i][a.Index].(type) {
	case int:
		return a.Slice[i][a.Index].(int) < a.Slice[j][a.Index].(int)
	case string:
		return a.Slice[i][a.Index].(string) < a.Slice[j][a.Index].(string)
	}
	return false
}
