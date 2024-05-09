package main

import (
	"log"
	"sync"
	"testing"
	"unique-codes-generator/generator"
)

func TextMain(t *testing.T) {

	var tests = []struct {
		name  string
		input int
		want  int
	}{
		// the table itself
		{"100", 100, 100},
		{"1000", 1000, 1000},
		{"10000", 10000, 10000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			const number = 100
			codes := make([]string, 0, tt.input)
			var wg sync.WaitGroup
			for i := 0; i <= number-1; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					code, err := generator.GenIdWithSonyFlake()
					if err != nil {
						log.Fatalf("%v", err)
					}
					_ = append(codes, code)
				}()
			}
			wg.Wait()
			if len(codes) != tt.want {
				t.Errorf("got %v, want %v", len(codes), tt.want)
			}
		})
	}

}
