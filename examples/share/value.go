package share

import (
	b3 "github.com/toophy/behavior3go"
	//. "github.com/toophy/behavior3go/actions"
	//. "github.com/toophy/behavior3go/composites"
	. "github.com/toophy/behavior3go/config"
	. "github.com/toophy/behavior3go/core"
	//. "github.com/toophy/behavior3go/decorators"
)

// 自定义action节点
type SetValue struct {
	Action
	value int
	key   string
}

func (this *SetValue) Initialize(setting *BTNodeCfg) {
	this.Action.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *SetValue) OnTick(tick Tick) b3.Status {
	tick.GetBlackboard().SetMem(this.key, this.value)
	return b3.SUCCESS
}

// 自定义action节点
type IsValue struct {
	Condition
	value int
	key   string
}

func (this *IsValue) Initialize(setting *BTNodeCfg) {
	this.Condition.Initialize(setting)
	this.value = setting.GetPropertyAsInt("value")
	this.key = setting.GetPropertyAsString("key")
}

func (this *IsValue) OnTick(tick Tick) b3.Status {
	v := tick.GetBlackboard().GetInt(this.key, "", "")
	if v == this.value {
		return b3.SUCCESS
	}
	return b3.FAILURE
}
