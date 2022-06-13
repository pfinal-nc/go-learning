package proxy

import (
	"testing"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 10:57
 * @Desc:
 */

func TestProxy(t *testing.T) {
	var sub Subject
	sub = &Proxy{}

	res := sub.Do()

	if res != "pre:real:after" {
		t.Fail()
	}
}
