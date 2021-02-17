package dinner

import (
	"math/rand"
	"time"
)

func Choose() string {
	dinners := []string{
		"tacos",
		"pizza",
		"ramen",
	}
	rand.Seed(time.Now().Unix())

	return dinners[rand.Intn(len(dinners))]
}
