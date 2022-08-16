// 小知识：
// golang中空结构体不占内存空间，被广泛作为各种场景下的占位符。既能节省资源又有较强的语义，表明这里不需要任何值，仅作为占位符
// golang本身没有Set类型，此处仅是简易实现，旨在说明问题
// 更全实现参考：https://github.com/deckarep/golang-set

package empty

type Set map[string]struct{}

// Has 是否包含元素
func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}

// Add 新增元素
func (s Set) Add(key string) {
	s[key] = struct{}{}
}

// Delete 删除元素
func (s Set) Delete(key string) {
	delete(s, key)
}
