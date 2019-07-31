package main

import (
	"errors"
	"math/rand"
)

//注册函数
func init() {
	RegisterBalance("random", &RandomBalance{})
}

//生成随机结构体，实现接口
type RandomBalance struct {
}

func (p *RandomBalance) DoBalance(insts []*Instance, key ...string) (inst *Instance, err error) {
	if len(insts) == 0 {
		err = errors.New("no Instance")
		return
	}
	//生成随机数
	index := rand.Intn(len(insts))
	//返回调用的服务器
	inst = insts[index]
	return
}
