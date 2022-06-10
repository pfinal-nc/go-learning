package builder

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 16:11
 * @Desc:
 */

// Builder 是构建者接口
type Builder interface {
	BuildPartA()
	BuildPartB()
	BuilderPartC()
}

type Director struct {
	builder Builder
}

// NewDirector 新建一个构建者
func NewDirector(builder Builder) *Director {
	return &Director{
		builder: builder,
	}
}

// Construct 构建
func (d *Director) Construct() {
	d.builder.BuildPartA()
	d.builder.BuildPartB()
	d.builder.BuilderPartC()
}

type BuildTest struct {
	product string
}

func (b *BuildTest) BuildPartA() {
	b.product += "A"
}

func (b *BuildTest) BuildPartB() {
	b.product += "B"
}

func (b *BuildTest) BuilderPartC() {
	b.product += "C"
}

func (b *BuildTest) GetProduct() string {
	return b.product
}

type BuildProduct struct {
	product string
}

func (b *BuildProduct) BuildPartA() {
	b.product += "PA"
}

func (b *BuildProduct) BuildPartB() {
	b.product += "PB"
}

func (b *BuildProduct) BuilderPartC() {
	b.product += "PC"
}

func (b *BuildProduct) GetProduct() string {
	return b.product
}
