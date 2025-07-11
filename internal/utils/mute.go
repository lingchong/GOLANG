package utils

import "sync"

type SafeCount struct {
	mu sync.Mutex
	AY []int
}

func AddEle(safeCount *SafeCount) {
	safeCount.mu.Lock()
	safeCount.AY = append(safeCount.AY, 1)
	safeCount.mu.Unlock()
}
