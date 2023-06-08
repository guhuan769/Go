package cache

import (
	"sync"
	"time"
)

// cache的实现类
type memCache struct { //定义一个结构体
	//最大内存
	maxMemorySizeStr string
	//最大内存
	maxMemorySize int64
	// 当前使用内存大小
	currMemorySize int64
	//缓存键值对
	values map[string]*memCacheValue
	//读写锁  通过读写锁 可以让多个携程进行并发读
	locker sync.RWMutex
}

type memCacheValue struct {
	value interface{}
	//插入时间
	inserTime time.Time //时间类型
	//有效时长
	expire time.Duration //时间段类型
	//value 大小
	size int64
	//在golang中 时间类型与时间段类型是2个类型
}

func NewMemCache() *memCache { //返回结构体
	return &memCache{}
}

// size : 1kb 100kb 1MB 2MB 1GB
func (mc *memCache) SetMaxMemory(size string) bool {
	// memCache指针类型
	mc.maxMemorySize, mc.maxMemorySizeStr = ParseSize(size)
	return true
}

// 将value 写入缓存
func (mc *memCache) Set(key string, val interface{}, expire time.Duration) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	valSize := GetValSize(val)
	mcv := &memCacheValue{
		value:     val,
		inserTime: time.Now(),
		expire:    expire,
		size:      valSize,
	}
	mc.del(key)
	mc.add(key, mcv)

	if mc.currMemorySize > mc.maxMemorySize {
		mc.del(key)
		return false
	}
	return true
}

// 根据key值获取value
func (mc *memCache) Get(key string) (interface{}, bool) {
	mc.locker.RLock() //读锁
	defer mc.locker.RUnlock()
	mcv, ok := mc.get(key)
	if ok {
		if mcv.expire != 0 && time.Now().After(mcv.inserTime.Add(mcv.expire)) {
			mc.del(key)
			return nil, false
		}
		return mcv, ok
	}
	return nil, false
}

// 删除key值
func (mc *memCache) Del(key string) bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.del(key)
	return true
}

// 判断key是否存在
func (mc *memCache) Exists(key string) bool {
	mc.locker.RLock()
	defer mc.locker.Unlock()
	_, ok := mc.values[key]
	return ok
	return false

}

// 清空所有key
func (mc *memCache) Flush() bool {
	mc.locker.Lock()
	defer mc.locker.Unlock()
	mc.values = make(map[string]*memCacheValue, 0)
	mc.currMemorySize = 0
	return true
}

// 获取缓存中所有key的数量
func (mc *memCache) Keys() int64 {
	mc.locker.RLock()
	mc.locker.Unlock()
	return int64(len(mc.values))
}

func (mc *memCache) get(key string) (*memCacheValue, bool) {
	val, ok := mc.values[key]
	return val, ok
}

func (mc *memCache) add(key string, val *memCacheValue) {
	mc.values[key] = val
	mc.currMemorySize += val.size

}
func (mc *memCache) del(key string) {
	tmp, ok := mc.get(key)
	delete(mc.values, key) //不存在也不会报错
	if ok && tmp != nil {
		mc.currMemorySize -= tmp.size
	}
}
