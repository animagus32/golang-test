package utils

import "container/list"

type RedpacketItem struct {

	Id uint
	Timeout	int64
}


type RedpacketRetreiver struct {
	List list.List
}

func (r *RedpacketRetreiver) Insert(id uint,timeout int64)  {
	timeout += 30 //todo 测试用 30秒后回收
	r.List.PushBack(RedpacketItem{id,timeout})
}

