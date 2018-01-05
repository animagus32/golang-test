package routinepool

import (
	"golang.org/x/net/context"
	// "reflect"
)

// var queue = make(chan func(),100)
type Pool struct{
	// Queue chan struct{
	// 	Args []reflect.Value
	// 	Func interface{}
	// } 
	Queue chan func()
	// handler func(interface{}) error
	Limit int
	Ctx context.Context
}

func GetPool(ctx context.Context,limit int) *Pool{
	
	p := Pool{
		Limit : limit,
		Ctx : ctx,
	}
	p.Queue = make(chan func(),limit )

	return &p
}

func (p *Pool) Add(f func()){
	p.Queue <- f
}
//todo 加入日志
func (p *Pool) Run(){
	done := make(chan bool)
	
	for i:=0 ; i < p.Limit ;i++ {
		go func(){
			// defer wg.Done()
			for {
				select{
				case <- done:
					return 
				default:
					{
						//handler
						f := <- p.Queue
						f()
					}
				}
			}
		}()
	}

	<- p.Ctx.Done()
	done <- true
}