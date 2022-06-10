package factory_method

import "testing"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 14:27
 * @Desc:
 */

func compute(factory OperatorFactory, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestMinusOperator_Result(t *testing.T) {
	t.Log(compute(&MinusOperatorFactory{}, 1, 2))
}

func TestOperator(t *testing.T) {
	var factory OperatorFactory
	factory = &PlusOperatorFactory{}
	c := compute(factory, 1, 2)
	if c != 3 {
		t.Error("PlusOperatorFactory error")
	}
	t.Log(c)
	factory = &MinusOperatorFactory{}
	e := compute(factory, 1, 2)
	if e != -1 {
		t.Error("MinusOperatorFactory error")
	}
	t.Log(e)
}
