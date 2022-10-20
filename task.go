package main

import "runtime"

// Pool 协程池
type Pool struct {
	TaskChannel chan func() // fuc类型任务队列
	GoNum       int         // 任务数量
}

// NewPool 创建一个协程池
func NewPool(cap ...int) *Pool {
	// 获取 worker 数量
	var n int
	if len(cap) > 0 {
		n = cap[0]
	}
	if n == 0 {
		n = runtime.NumCPU() // 默认等于CPU线程数
	}
	// p 是 Pool的引用
	p := &Pool{
		TaskChannel: make(chan func()),
		GoNum:       n,
	}
	return p
}

// StartPool 启动协程池
func StartPool(p *Pool) {
	// 创建指定数量 worker 从任务队列取出任务执行
	for i := 0; i < p.GoNum; i++ {
		go func() {
			for task := range p.TaskChannel {
				task()
			}
		}()
	}
}

// Submit 提交任务
func (p *Pool) Submit(f func()) {
	p.TaskChannel <- f
}
