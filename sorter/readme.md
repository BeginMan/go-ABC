# 项目需求

我们准备开发一个排序算法的比较程序
从命令行指定输入的数据文件和输出的数据文件,并指定对应的排序算法。该程序的用法如下所示:

    USAGE: sorter –i <in> –o <out> –a <qsort|bubblesort>

一个具体的执行过程如下:

    $ ./sorter –I in.dat –o out.dat –a qsort
    The sorting process costs 10us to complete.
	
当然,如果输入不合法,应该给出对应的提示。

# 项目结构

```
.
├── bin
├── pkg
├── readme.md
├── src
│   ├── algorithm
│   │   ├── bubblesort
│   │   │   ├── bubblesort.go
│   │   │   └── bubblesort_test.go
│   │   └── qsort
│   │       ├── qsort.go
│   │       └── qsort_test.go
│   └── sorter
│       └── sorter.go
└── unsorted.dat
```

# 构建

在确认已经设置好GOPATH后,我们可以直接运行以下命令来构建和测试程序:

	$ echo $GOPATH
	~/goyard/sorter
	
	$ go build algorithm/qsort
	$ go build algorithm/bubblesort
	
	$ go test algorithm/qsort
	$ go test algorithm/bubblesort
	
	$ go install algorithm/qsort
    $ go install algorithm/bubblesort
	
    $ go build sorter
    $ go install sorter
	
