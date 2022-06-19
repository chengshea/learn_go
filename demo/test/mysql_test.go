package test

import (
	"demo/com/cs/learn/model"
	"testing"
)

func Test_crud(t *testing.T) {
	type User struct {
		user model.User
	}
	tests := []struct {
		id   string
		name string
		age  int
		sex  bool
	}{
		// TODO: Add test cases.
		{
			// |   12 | 第三      |   11 |    1 |
			// |  113 | 的人      |   23 |    0 |
			// |   15 | 推广      |   19 |    1 |
			// |   21 | 发广告    |   21 |    1 |
			// |   51 | 若非      |   35 |    0 |
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//create()

			//query()

			//update()

		})
	}
}
