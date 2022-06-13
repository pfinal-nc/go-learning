package command

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/13 18:02
 * @Desc:
 */

func ExampleCommand() {
	mb := &MotherBoard{}
	StartCommand := NewStartCommand(mb)
	RebootCommand := NewRebootCommand(mb)
	box1 := NewBox(StartCommand, RebootCommand)
	box1.PressButton1()
	box1.PressButton2()

	box2 := NewBox(RebootCommand, StartCommand)
	box2.PressButton1()
	box2.PressButton2()

	// Output:
	// MotherBoard start
	// MotherBoard reboot
	// MotherBoard reboot
	// MotherBoard start

}
