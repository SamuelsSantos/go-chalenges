package main

import (
	"sort"
	"strconv"
	"strings"
)

const listSeparator = ", "

func getIntervals(intervals []string) string {
	return strings.Join(intervals, listSeparator)
}

func madeList(value string) []string {
	return strings.SplitN(value, listSeparator, -1)
}

func getInterval(min int, max int) string {

	start := strconv.Itoa(min)
	end := strconv.Itoa(max)
	interval := "[" + start + "-" + end + "]"
	if start == end {
		interval = "[" + end + "]"
	}

	return interval
}

func groupBy(stack []int, group *[]string) {

	sort.Ints(stack)

	min := stack[0]
	max := stack[len(stack)-1]
	interval := getInterval(min, max)
	*group = append(*group, interval)

}

func agrupar(value string) (string, error) {

	list := madeList(value)
	var stack = []int{}
	var result = []string{}
	for index, element := range list {
		j, _ := strconv.Atoi(element)
		if index == 0 {
			stack = append(stack, j)
		} else {
			old, _ := strconv.Atoi(list[index-1])
			if j-1 == old {
				stack = append(stack, j)
				if index == len(list)-1 {
					groupBy(stack, &result)
				}
				continue
			}

			groupBy(stack, &result)
			stack = nil
			stack = append(stack, j)
			if index == len(list)-1 {
				groupBy(stack, &result)
			}
		}
	}

	return getIntervals(result), nil
}

func main() {

}
