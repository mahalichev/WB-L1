package main

import (
	"fmt"
	"sync"
)

// Использование sync.Mutex (в момент времени может происходить либо 1 запись, либо 1 чтение)
// Использование дженериков: ключ - сравниваемый тип данных (к которому применима операция ==), значение - любой тип данных
type MutexMap[K comparable, V any] struct {
	base  map[K]V
	mutex sync.Mutex
}

func (mMap *MutexMap[K, V]) Write(key K, value V) {
	// Блокировка доступа к общему ресурсу
	mMap.mutex.Lock()
	mMap.base[key] = value
	// Разблокировка доступа к общему ресурсу
	mMap.mutex.Unlock()
}

func (mMap *MutexMap[K, V]) Read(key K) (V, bool) {
	// Блокировка доступа к общему ресурсу
	mMap.mutex.Lock()
	result, ok := mMap.base[key]
	// Разблокировка доступа к общему ресурсу
	mMap.mutex.Unlock()
	return result, ok
}

func NewMutexMap[K comparable, V any]() *MutexMap[K, V] {
	return &MutexMap[K, V]{base: make(map[K]V), mutex: sync.Mutex{}}
}

// Использование sync.RWMutex (в момент времени может происходить либо 1 запись, либо N чтений)
type RWMutexMap[K comparable, V any] struct {
	base  map[K]V
	mutex sync.RWMutex
}

func (mMap *RWMutexMap[K, V]) Read(key K) (V, bool) {
	// Блокировка доступа к общему ресурсу (доступно только чтение)
	mMap.mutex.RLock()
	result, ok := mMap.base[key]
	// Разблокировка доступа к общему ресурсу (запись будет доступна, если при N RWMutex.RLock() было выполнено N RWMutex.RUnlock())
	mMap.mutex.RUnlock()
	return result, ok
}

func (mMap *RWMutexMap[K, V]) Write(key K, value V) {
	// Блокировка доступа к общему ресурсу
	mMap.mutex.Lock()
	mMap.base[key] = value
	// Разблокировка доступа к общему ресурсу
	mMap.mutex.Unlock()
}

func NewRWMutexMap[K comparable, V any]() *RWMutexMap[K, V] {
	return &RWMutexMap[K, V]{base: make(map[K]V), mutex: sync.RWMutex{}}
}

func RunMutexMap() {
	safeMap := NewMutexMap[int, int]()
	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				safeMap.Write(keyVal, keyVal)
			}
		}(i)

		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				val, ok := safeMap.Read(keyVal)
				for ; !ok; val, ok = safeMap.Read(keyVal) {
				}
				fmt.Printf("%d:%d; ", keyVal, val)
			}
		}(i)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()
}

func RunRWMutexMap() {
	safeMap := NewRWMutexMap[int, int]()
	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				safeMap.Write(keyVal, keyVal)
			}
		}(i)

		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				val, ok := safeMap.Read(keyVal)
				for ; !ok; val, ok = safeMap.Read(keyVal) {
				}
				fmt.Printf("%d:%d; ", keyVal, val)
			}
		}(i)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()
}

// Использование sync.Map
func RunSyncMap() {
	var syncMap sync.Map
	// Создание WaitGroup для ожидания выполнения набора горутин
	wg := new(sync.WaitGroup)

	for i := 0; i < 5; i++ {
		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				syncMap.Store(keyVal, keyVal)
			}
		}(i)

		// Увеличение счётчика горутин на 1
		wg.Add(1)
		go func(tens int) {
			// Оповещение о завершении работы горутины (уменьшение счётчика горутин на 1)
			defer wg.Done()
			for j := 0; j < 3; j++ {
				keyVal := tens*10 + j
				val, ok := syncMap.Load(keyVal)
				for ; !ok; val, ok = syncMap.Load(keyVal) {
				}
				fmt.Printf("%d:%d; ", keyVal, val.(int))
			}
		}(i)
	}
	// Ожидание выполнения всех горутин
	wg.Wait()
}

func main() {
	fmt.Println("Using sync.Mutex:")
	RunMutexMap()

	fmt.Println("\nUsing sync.RWMutex:")
	RunRWMutexMap()

	fmt.Println("\nUsing sync.Map:")
	RunSyncMap()
}
