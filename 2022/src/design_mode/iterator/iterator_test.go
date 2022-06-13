package iterator

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 18:14
 * @Desc:
 */

func ExampleIteratorPrint() {
	var aggeragte Aggregate
	aggeragte = NewNumbers(1, 10)
	IteratorPrint(aggeragte.Iterator())
	// Output:
	// 1
	// 2
	// 3
	// 4
	// 5
	// 6
	// 7
	// 8
	// 9
	// 10

}
