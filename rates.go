package main

import (
	"fmt"
	"sync"
)

type Rates struct {
	rates map[string]float64
	mutex sync.RWMutex
}

func NewRates() *Rates {
	return &Rates{
		rates: map[string]float64{
			"USD": 90.0,
			"EUR": 98.5,
			"YEN": 0.5,
		},
	}
}

func (r *Rates) GetRate(currency string) (float64, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	rate, exists := r.rates[currency]
	if !exists {
		return 0, fmt.Errorf("currency not supported: %s", currency)
	}

	return rate, nil
}

func (r *Rates) SetRate(currency string, rate float64) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	r.rates[currency] = rate
}

func (r *Rates) GetAllRates() map[string]float64 {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	result := make(map[string]float64)
	for currency, rate := range r.rates {
		result[currency] = rate
	}
	return result
}
