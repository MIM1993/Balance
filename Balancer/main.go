package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var insts []*Instance
	for i := 0; i < 10; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), rand.Intn(255))
		post, _ := strconv.Atoi(fmt.Sprintf("880%d", i))
		inst := NewInstance(host, post)
		insts = append(insts, inst)
	}

	var name = "random"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	for {
		inst, err := DoBalance(name, insts)
		if err != nil {
			log.Fatal("DoBalance err :", err)
			time.Sleep(time.Duration(1))
			continue
		}
		fmt.Printf("%#v\n", inst)
		time.Sleep(time.Duration(1))
	}

}
