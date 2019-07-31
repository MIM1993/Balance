package main

import "errors"

//注册函数
func init() {
	RegisterBalance("roundorbin", &RoundrobinBalance{})
}

//轮询结构体
type RoundrobinBalance struct {
	curIndex int //下标字段
}

//实现接口
func (p *RoundrobinBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) <= 0 {
		err = errors.New("no Instance")
		return
	}
	//判断重置下表数据
	if p.curIndex > len(insts) {
		p.curIndex = 0
	}
	//返回服务器
	inst = insts[p.curIndex]
	//下表加以
	p.curIndex ++

	return
}
