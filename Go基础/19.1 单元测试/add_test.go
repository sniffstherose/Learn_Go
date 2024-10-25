package main

import "testing"

/*
表驱动法（Table-Driven Method）是一种基于数据表的程序设计方法。

Golang采用了Table-Driven组织测试。

Table Driven主要分成三个部分：

测试用例的定义：即每一个测试用例需要有什么。
具体的测试用例：你设计的每一个测试用例都在这里。
执行测试用例：这里面还包括了对测试结果进行断言。
*/
func Test_add(t *testing.T) {
	// 测试用例的定义
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// 具体测试用例
		{
			name: "Add positive numbers",
			args: args{a: 1, b: 2},
			want: 3,
		},
		{
			name: "Add negative numbers",
			args: args{a: -1, b: -2},
			want: -3,
		},
		{
			name: "Add mixed numbers",
			args: args{a: 1, b: -2},
			want: -1,
		},
		{
			name: "Add zeros",
			args: args{a: 0, b: 0},
			want: 0,
		},
		{
			name: "Add large numbers",
			args: args{a: 1000000, b: 2000000},
			want: 3000000,
		},
	}
	// 执行测试用例
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("add() = %v, want = %v", got, tt.want)
			}
		})
	}
}
