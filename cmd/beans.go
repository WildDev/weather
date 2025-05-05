package cmd

type Component interface {
	Init()
	Destroy()
}

type Beans []Component

type Context struct {
	beans Beans
}

func (ctx *Context) Add(c Component) {

	c.Init()

	ctx.beans = append(ctx.beans, c)
}

func (ctx *Context) DestroyAll() {

	beans := ctx.beans

	for i := len(beans) - 1; i >= 0; i-- {
		beans[i].Destroy()
	}

	beans = nil
}
