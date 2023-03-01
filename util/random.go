package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomTasks() string {
	tasks := []string{"Washing", "Cooking", "Gym", "Running", "Shopping", "Bathing"}
	n := len(tasks)

	return tasks[rand.Intn(n)]
}

func RandomBool() bool {
	return rand.Intn(2) == 0
}
