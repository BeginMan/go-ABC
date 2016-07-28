这个项目是《Go语言编程》命令行计算器程序。
可参看这篇blog: http://siddontang.com/2014/05/13/golang-in-procution-review/
Usage:
	
	$ calc help
	usage: calc command [arguments]...
	
	The commands are:
	sqrt	Square root of a non-negative value.
	add		Addition of two values

工程结构：

	~/goproj/calproj> tree
	.
	├── bin
	├── pkg
	└── src
	    ├── calc
	    ├── calc.go
	    ├── readme.md
	    └── simplemath
	        ├── add.go
	        ├── add_test.go
	        ├── sqrt.go
	        └── sqrt_test.go