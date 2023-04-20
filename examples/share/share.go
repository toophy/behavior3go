package share

import (
	"fmt"

	b3 "github.com/toophy/behavior3go"

	//. "github.com/toophy/behavior3go/actions"
	//. "github.com/toophy/behavior3go/composites"
	. "github.com/toophy/behavior3go/config"
	. "github.com/toophy/behavior3go/core"
	//. "github.com/toophy/behavior3go/decorators"
)

// 自定义action节点
type LogTest struct {
	Action
	info string
}

func (this *LogTest) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.info = setting.GetPropertyAsString("info")
}

func (this *LogTest) OnTick(tick Tick) b3.Status {
	fmt.Println("logtest:", tick.GetLastSubTree(), this.info)
	return b3.SUCCESS
}
