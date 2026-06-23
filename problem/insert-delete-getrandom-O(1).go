package problem

import (
	"math/rand"
	"time"
)

type RandomizedSet struct {
	randomAccess []int
	another      map[int]int
	randmizer    *rand.Rand
}

func Constructor() RandomizedSet {
	randmizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	return RandomizedSet{
		randomAccess: []int{},
		another:      map[int]int{},
		randmizer:    randmizer,
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.another[val]; ok {
		return false
	}
	this.randomAccess = append(this.randomAccess, val)
	this.another[val] = len(this.randomAccess) - 1
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.another[val]
	if !ok {
		return false
	}

	lastIndex := len(this.randomAccess) - 1
	lastValue := this.randomAccess[lastIndex]

	if index != lastIndex {
		this.randomAccess[index] = lastValue
		this.another[lastValue] = index
	}
	this.randomAccess = this.randomAccess[:lastIndex]
	delete(this.another, val)
	return true
}

func (this *RandomizedSet) GetRandom() int {
	return this.randomAccess[this.randmizer.Intn(len(this.randomAccess))]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
