package simple_factory

import "testing"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 14:18
 * @Desc:
 */

// API is interface

func TestNewAPI(t *testing.T) {
	api := NewAPI("A")
	t.Log(api.Say("PFinal"))

	api = NewAPI("B")
	t.Log(api.Say("PFinal"))
}
