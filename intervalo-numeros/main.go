package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Input list of itens
type Input []string

const listSeparator = ", "

func getIntervals(intervals []string) string {
	return strings.Join(intervals, listSeparator)
}

func makeList(value string) Input {
	return strings.SplitN(value, listSeparator, -1)
}

func (b *Input) makeOrderedIntList() ([]int, error) {

	stack := make([]int, len((*b)))

	for index, element := range *b {
		var value int
		var err error
		if value, err = strconv.Atoi(element); err != nil {
			return nil, err
		}

		stack[index] = value
	}

	return stack, nil
}

func getInterval(min int, max int) string {
	interval := fmt.Sprintf("[%d-%d]", min, max)
	if min == max {
		interval = fmt.Sprintf("[%d]", max)
	}
	return interval
}

func sortGroup(group []int) (min int, max int) {
	sort.Ints(group)
	return group[0], group[len(group)-1]
}

func makeInterval(stack []int, group *[]string) {
	interval := getInterval(sortGroup(stack))
	*group = append(*group, interval)
}

func doProcessar(value string) (string, error) {

	input := makeList(value)
	var list []int
	var err error
	if list, err = input.makeOrderedIntList(); err != nil {
		return "", err
	}

	stack := []int{}
	var result = []string{}
	for index, element := range list {
		if index == 0 {
			stack = append(stack, element)
		} else {
			old := list[index-1]
			if element-1 == old {
				stack = append(stack, element)
				if index == len(list)-1 {
					makeInterval(stack, &result)
				}
				continue
			}

			makeInterval(stack, &result)
			stack = nil
			stack = append(stack, element)
			if index == len(list)-1 {
				makeInterval(stack, &result)
			}
		}
	}

	return getIntervals(result), nil
}

func main() {

}
