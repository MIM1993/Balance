package main

import (
	"errors"
	"fmt"
)

//定义负载管理器
type BalanceMgr struct {
	allbalance map[string]Balance
}

//生成一个管理器
var BM = BalanceMgr{
	allbalance: make(map[string]Balance),
}

//注册函数（私有）
func (p *BalanceMgr) registerBalance(name string, b Balance) {
	p.allbalance[name] = b
}

//注册函数（开放）
func RegisterBalance(name string, b Balance) {
	BM.registerBalance(name, b)
}

//执行函数
func DoBalance(name string, insts []*Instance) (inst *Instance, err error) {
	v, ok := BM.allbalance[name]
	if !ok {
		err = errors.New("no balance name")
		return
	}

	inst, err = v.DoBalance(insts)
	if err != nil {
		err = fmt.Errorf("DoBalance err", err)
		return
	}

	return
}
