package cache

import (
	"fmt"
	"sync"
	"time"
)

// cache的实现类
type memCache struct { //定义一个结构体
	//最大内存
	maxMemorySizeStr string
	//最大内存
	maxMemorySize int64
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

	return false
}

// 根据key值获取value
func (mc memCache) Get(key string) (interface{}, bool) {
	fmt.Println("called func get")
	return nil, false
}

// 删除key值
func (mc *memCache) Del(key string) bool {

	return false
}

// 判断key是否存在
func (mc *memCache) Exists(key string) bool {
	return false

} //清空所有key
func (mc *memCache) Flush() bool {
	return false
}

// 获取缓存中所有key的数量
func (mc *memCache) Keys() int64 {
	return 0
}
