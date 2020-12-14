#学习笔记
总结:
本周学习内容二叉树相关问题较为熟悉，比较轻松。重点还在重复刷以前的题。
学会了go语言sort及heap接口
//sort.interface
`type Interface interface {
    // Len方法返回集合中的元素个数
    Len() int
    // Less方法报告索引i的元素是否比索引j的元素小
    Less(i, j int) bool
    // Swap方法交换索引i和j的两个元素
    Swap(i, j int)
}`

heap.interface
`type Interface interface {
    sort.Interface
    Push(x interface{}) // 向末尾添加元素
    Pop() interface{}   // 从末尾删除元素
}`