package utils

import (
	"io"
	"log"
	"os"
	"strconv"
)

func ReadFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	buffer, err := io.ReadAll(file)
	return string(buffer)
}

func ConvertStringArrayToInt(values []string) []int {
	var result = []int{}
	for _, val := range values {
		newVal, err := strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, newVal)
	}
	return result
}

func Count[T any](slice []T, f func(T) bool) int {
	count := 0
	for _, s := range slice {
		if f(s) {
			count++
		}
	}
	return count
}

func Reduce[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, val := range s {
		acc = f(acc, val)
	}
	return acc
}

func Where[T any](slice []T, f func(T) bool) []T {
	result := []T{}
	for _, s := range slice {
		if f(s) {
			result = append(result, s)
		}
	}
	return result
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
