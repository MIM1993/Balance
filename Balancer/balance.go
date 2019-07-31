package main

/* 服务器信息*/
type Instance struct {
	host string
	port int
}

//初始化服务器
func NewInstance(host string, port int) *Instance {
	return &Instance{
		host: host,
		port: port,
	}
}

//获取host
func (p *Instance) GetHost() string {
	return p.host
}

//获取port
func (p *Instance) GetPort() int {
	return p.port
}


