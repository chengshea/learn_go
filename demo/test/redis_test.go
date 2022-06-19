package test

import (
	"log"
	"testing"
)

func Test_redisGet(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test1", //key
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Println("tt:", tt.name) //tt: test1
			redisGet(tt.name)
		})
	}
}
