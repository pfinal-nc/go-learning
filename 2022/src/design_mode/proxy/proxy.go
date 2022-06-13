package proxy

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 10:47
 * @Desc:
 */

type Subject interface {
	Do() string
}

type RealSubject struct {
}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// 在调用真实的对象之前，可以做一些预处理
	res += "pre:"
	// 调用真实的香
	res += p.real.Do()

	// 调用之后的操作, 如缓存结果 对结果进行处理
	res += ":after"
	return res
}
