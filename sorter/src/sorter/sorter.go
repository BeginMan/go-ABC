//sorter.go:主程序：实现快速排序
//bubblesort.go: 实现冒泡排序

package main

import "flag"
import "fmt"
import "bufio"
import "io"
import "time"
import "os"
import "strconv"
import "algorithm/bubblesort"
import "algorithm/qsort"

var infile *string = flag.String("i", "unsorted.dat", "File contains values for sorting")
var outfile *string = flag.String("o", "sorted.dat", "File to receive sorted values")
var algorithm *string = flag.String("a", "qsort", "Sort algorithm")

func readValues(infile string) (values []int, err error) {
	file, err := os.Open(infile)
	if err != nil {
		fmt.Println("Failed to open the input file", infile)
		return
	}

	defer file.Close()

	br := bufio.NewReader(file)
	//数组切片
	values = make([]int, 0)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			break
		}

		if isPrefix {
			fmt.Println("A too long line, same unexpected.")
			return
		}

		str := string(line)              // 转换字符数组为字符串
		value, err1 := strconv.Atoi(str) // 注意，只能转换数字，如果有空行或非数字字符则报错
		if err1 != nil {
			err = err1
			return
		}

		values = append(values, value)
	}
	return

}

func writeValues(result []int, path string) error {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Create file error")
		return err
	}
	defer file.Close() // 别忘记close了
	for _, value := range result {
		file.WriteString(strconv.Itoa(value) + "\n")
	}

	return nil
}

func main() {
	flag.Parse()
	if infile != nil {
		fmt.Println("infile=", *infile, "outfile=", *outfile, "algorithm=", *algorithm)
	}
	values, err := readValues(*infile)
	if err == nil {
		t1 := time.Now()
		switch *algorithm {
		case "qsort":
			qsort.QuickSort(values)
		case "bubblesort":
			bubblesort.BubbleSort(values)
		default:
			fmt.Println("Sorting algorithm", *algorithm, "is either unknown or unsupported.")
		}
		t2 := time.Now()
		fmt.Println("The sorting process costs", t2.Sub(t1), "to complete.")
		writeValues(values, *outfile)
	} else {
		fmt.Println(err)
	}

}
