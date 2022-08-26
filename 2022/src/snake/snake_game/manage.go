package game

/**
 * @Author: PFinal南丞
 * @Author: lampxiezi@163.com
 * @Date: 2022/7/25 14:07
 * @Desc:
 */
import . "2022/src/snake/snake_game/component"

func Start()  {
	Run(func() {
		NewGameService().Start(NewGameData())
	})
}