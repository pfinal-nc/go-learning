package builder

import (
	"testing"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 16:17
 * @Desc:
 */

func TestNewDirector(t *testing.T) {
	b := &BuildTest{}
	director := NewDirector(b)
	director.Construct()
	res := b.GetProduct()
	t.Log(res)
	// Output: ABC	// ABC

	c := &BuildProduct{}
	dir := NewDirector(c)
	dir.Construct()
	res = c.GetProduct()
	t.Log(res)
}
