package simple_factory

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 14:15
 * @Desc:
 */

// API is interface
type API interface {
	Say(name string) string
}

// NewAPI is function
func NewAPI(name string) API {
	if name == "A" {
		return &implA{}
	} else if name == "B" {
		return &implB{}
	}
	return nil
}

type implA struct {
}

func (*implA) Say(name string) string {
	return "Hello " + name
}

type implB struct {
}

func (*implB) Say(name string) string {
	return "Hi " + name
}
