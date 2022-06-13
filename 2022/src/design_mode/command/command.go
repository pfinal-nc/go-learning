package command

import "fmt"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 17:51
 * @Desc:
 */

type Command interface {
	Execute()
}

type MotherBoard struct{}

type StartCommand struct {
	mb *MotherBoard
}

type RebootCommand struct {
	mb *MotherBoard
}

func NewStartCommand(mb *MotherBoard) *StartCommand {
	return &StartCommand{mb: mb}
}

func NewRebootCommand(mb *MotherBoard) *RebootCommand {
	return &RebootCommand{mb: mb}
}
func (c *RebootCommand) Execute() {
	c.mb.Reboot()
}

func (c *StartCommand) Execute() {
	c.mb.Start()
}

func (c *MotherBoard) Start() {
	fmt.Println("MotherBoard start")
}

func (c *MotherBoard) Reboot() {
	fmt.Println("MotherBoard reboot")
}

type Box struct {
	button1 Command
	button2 Command
}

func NewBox(button1 Command, button2 Command) *Box {
	return &Box{button1: button1, button2: button2}
}

func (b *Box) PressButton1() {
	b.button1.Execute()
}

func (b *Box) PressButton2() {
	b.button2.Execute()
}
