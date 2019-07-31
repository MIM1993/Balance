package main

type Balance interface {
	/*
	负载均衡算法
	*/
	DoBalance([]*Instance,...string)(*Instance,error)
}
