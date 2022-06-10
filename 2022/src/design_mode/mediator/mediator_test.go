package mediator

import "testing"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/6/10 17:02
 * @Desc:
 */

func TestGetMediator(t *testing.T) {
	mediator := GetMediator()
	mediator.CDDriver = &CDDriver{}
	mediator.CPU = &CPU{}
	mediator.VideoCard = &VideoCard{}
	mediator.SoundCard = &SoundCard{}

	// Tiggle
	mediator.CDDriver.ReadData()
	if mediator.CDDriver.Data != "music,image,video" {
		t.Error("CDDriver data is not music,image,video")
	}
	if mediator.CPU.SoundCard != "music" {
		t.Error("CPU sound card is not music")
	}
}
