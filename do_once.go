package com

import (
	"sync"
)

func NewDoOnce() *DoOnce {
	return &DoOnce{
		data: map[interface{}]*sync.WaitGroup{},
	}
}

type DoOnce struct {
	lock sync.RWMutex
	data map[interface{}]*sync.WaitGroup
}

// CanSet 同一时刻只有一个请求能获取执行权限，获得执行权限的线程接下来需要执行具体的业务逻辑，完成后调用release方法通知其他线程，操作完成，获取资源即可，其他请求接下来需要调用wait方法
// reqTag 请求标识 用于标识同一个资源
func (u *DoOnce) CanSet(reqTag interface{}) bool {
	u.lock.Lock()

	if u.data == nil {
		u.data = map[interface{}]*sync.WaitGroup{}
	} else {
		_, ok := u.data[reqTag]
		if ok {
			u.lock.Unlock()
			return false
		}
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)
	u.data[reqTag] = wg

	u.lock.Unlock()
	return true
}

// Wait 调用wait方法将处于阻塞状态，直到获得执行权限的线程处理完具体的业务逻辑，调用release方法来通知其他线程资源ok了
func (u *DoOnce) Wait(reqTag interface{}) {
	u.lock.RLock()
	w, ok := u.data[reqTag]
	u.lock.RUnlock()
	if !ok {
		return
	}

	w.Wait()
}

// Release 获得执行权限的线程需要在执行完业务逻辑后调用该方法通知其他处于阻塞状态的线程
func (u *DoOnce) Release(reqTag interface{}) {
	u.lock.Lock()

	if _, ok := u.data[reqTag]; !ok {
		u.lock.Unlock()
		return
	}
	u.data[reqTag].Done()
	delete(u.data, reqTag)
	u.lock.Unlock()
}
