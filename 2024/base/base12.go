package main

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2024/6/28
 * @Desc:
 * @Project: 2024
 */

// 切片操作常用技巧
// 复制 将切片a 中的元素赋值到切片b中

func main() {
	a := []int{1, 2, 3, 4}
	b := make([]int, len(a))
	copy(b, a) // 最常用的方法是使用内置的 copy 函数
	//fmt.Println(a)
	//fmt.Println(b)
	//b = append(a[:2], a[3:]...)
	//fmt.Println(a)
	// 除了使用内置的 copy 函数外，还可以使用内置的 append 函数
	// append 函数的第一个参数是切片的地址，第二个参数是要添加的元素

	// 剪切 将切片a 中索引i~j位置的元素剪切掉 可以按照下面的方式 使用 append 函数完成
	a = append(a[:3], a[4:]...)
	fmt.Println(b)
	fmt.Println(a)

	// 删除 将切片a 中索引位置为 i 的元素删除 同样可以按照剪切的方式使用 append 函数完成删除操作
	a = append(a[:2], a[3:]...) // 2 3
	fmt.Println(a)
	// 或者搭配 copy 函数使用切片表达式完成删除操作
	// a = a[:i+copy(a[i:], a[i+1:])] // 2 3
	// fmt.Println(a)

	// 剪切或者删除操作可能引起的内存泄漏
	// 需要特别注意的是, 如果切片a 中的元素是一个指针类型或包含指针字段的结构体类型 上面剪切和删除的示例代码会存在一个潜在的内存泄漏问题;一些具有值的元素 仍然被切片a 引用. 因此无法被来及回收机制回收掉.

	// 剪切
	//copy(a[i:],a[j:])
	//for k,n :=len(a)-j+i,len(a);k<n;k++ {
	//	a[k] = nil // 或类型T的零值
	//}
	//a = a[:len(a)-j+i]

	// 内部扩张  在切片a 的索引i之后扩张j个元素
	// 使用 两个 append 函数完成, 即先将索引i之后的元素追加到一个长度为j的切片后,再将这个切片中的所有元素追加到切片a 的索引i之后
	//a = append(a[:i],append(make([]int,j),a[i:]...)...)
	//fmt.Println(a)
	//a = append(a[:1], append([]int{0}, a[1:]...)...)
	//fmt.Println(a)

	// 尾部扩张
	//a = append(a, make([]int, 5)...)
	//fmt.Println(a)

	// 过滤
	// 按照一定的规则将切片a 中的元素进行就地过滤 假设过滤的条件已经封装为 keep 函数,使用 for range 遍历切片a 的所有元素逐一调用 keep 函数进行过滤
	//n :=0
	//for _, x:= range a {
	//	if keep(x) {
	//		a[n] = x // 保留该元素
	//		n++
	//	}
	//}
	//a = a[:n] // 截取切片中需保留的元素

	// 插入  将元素 x 插入切片a 的索引i 处使用两个 append 函数完成
	//a = append(a[:i],append([]int{x},a[i:]...)...) // 第二个 append 函数创建了一个具有自己底层数组的新切片, 并将 a[i:] 中的水元素复制到该切片, 然后由第一个append 函数将这些元素赋值回切片 a
	//a := append(a,0) // 这里应使用元素类型的零值
	//copy(a[i+1:],a[i:])
	//a[i] = x
	// 追加
	// 将元素x 追加到切片a的最后. 这里使用 append 函数即可
	//a = append(a, x)
	//fmt.Println(a)

	// 交换
	// 交换切片a 中索引i 和索引j 的元素
	//a[i],a[j] = a[j],a[i]
	//fmt.Println(a)

	// 弹出
	// 将切片 a 的最后一个元素弹出 这里使用切片表达式完成
}
