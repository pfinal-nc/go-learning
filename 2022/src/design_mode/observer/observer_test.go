package observer

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 17:39
 * @Desc:
 */

func ExampleObserver() {
	subject := NewSubject()
	reader1 := NewReader("reader1")
	reader2 := NewReader("reader2")
	reader3 := NewReader("reader3")
	subject.Attach(reader1)
	subject.Attach(reader2)
	subject.Attach(reader3)

	subject.SetContext("hello")
	// Output:
	// reader1 receive hello
	// reader2 receive hello
	// reader3 receive hello
}
