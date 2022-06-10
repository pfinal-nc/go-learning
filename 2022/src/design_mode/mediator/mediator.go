package mediator

import (
	"fmt"
	"strings"
)

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 16:50
 * @Desc:
 */

type CDDriver struct {
	Data string
}
type CPU struct {
	VideoCard string
	SoundCard string
}

func (c *CPU) Process(data string) {
	sp := strings.Split(data, ",")
	c.VideoCard = sp[1]
	c.SoundCard = sp[0]
	fmt.Printf("CPU: split data with Sound %s, Video %s\n", c.SoundCard, c.VideoCard)
	GetMediator().changed(c)
}

type VideoCard struct {
	Data string
}

func (v *VideoCard) Display(data string) {
	v.Data = data
	fmt.Printf("VideoCard display data: %s\n", v.Data)
	GetMediator().changed(v)
}

type SoundCard struct {
	Data string
}

func (s *SoundCard) Play(data string) {
	s.Data = data
	fmt.Printf("SoundCard play data: %s\n", s.Data)
	GetMediator().changed(s)
}

type Mediator struct {
	CDDriver  *CDDriver
	CPU       *CPU
	VideoCard *VideoCard
	SoundCard *SoundCard
}

var mediator *Mediator

func GetMediator() *Mediator {
	if mediator == nil {
		mediator = &Mediator{}
	}
	return mediator
}

func (c *CDDriver) ReadData() {
	c.Data = "music,image,video"
	fmt.Printf("CDDriver read data: %s\n", c.Data)
	GetMediator().changed(c)
}

func (m *Mediator) changed(i interface{}) {
	switch inst := i.(type) {
	case *CDDriver:
		m.CPU.Process(inst.Data)
	case *CPU:
		m.SoundCard.Play(inst.SoundCard)
		m.VideoCard.Display(inst.VideoCard)
	}
}
