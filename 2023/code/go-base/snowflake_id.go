package main

import "time"

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2023/12/25
 * @Desc:
 * @Project: 2023
 */

var (
	machineID     int64 // 机器Id 占10位 [0, 1023]
	sn            int64 // 序列号 占12位 [0, 4095]
	lastTimeStamp int64 // 上次的时间戳(毫秒级)
)

func init() {
	lastTimeStamp = time.Now().UnixNano() / 1000000
}

func SetMachineId(mid int64) {
	// 把机器 id 左移 12 位移 12 位 让出 12 位空间给序列号使用
	machineID = mid << 12
}

func GetSnowflakeId() int64 {
	curTimeStamp := time.Now().UnixNano() / 1000000
	// 同一毫秒
	if curTimeStamp == lastTimeStamp {
		sn++
		// 序列号占12位 十进制范围是[0, 4095]
		if sn > 4095 {
			time.Sleep(time.Millisecond)
			curTimeStamp = time.Now().UnixNano() / 1000000
			lastTimeStamp = curTimeStamp
			sn = 0
		}

	}
}
