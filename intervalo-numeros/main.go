package main

import (
	"sort"
	"strconv"
	"strings"
)

func agrupar(list string) (string, error) {

	s := strings.SplitN(list, ", ", -1)
	var pilha = []int{}
	var result = []string{}
	for index, element := range s {
		j, _ := strconv.Atoi(element)
		if index == 0 {
			pilha = append(pilha, j)
		} else {
			old, _ := strconv.Atoi(s[index-1])
			if j-1 == old {
				pilha = append(pilha, j)

				if index == len(s)-1 {
					sort.Ints(pilha)
					min := strconv.Itoa(pilha[0])
					max := strconv.Itoa(pilha[len(pilha)-1])
					if min == max {
						result = append(result, "["+max+"]")
					} else {
						result = append(result, "["+min+"-"+max+"]")
					}
				}

				continue
			}

			sort.Ints(pilha)
			min := strconv.Itoa(pilha[0])
			max := strconv.Itoa(pilha[len(pilha)-1])

			if min == max {
				result = append(result, "["+max+"]")
			} else {
				result = append(result, "["+min+"-"+max+"]")
			}
			pilha = nil
			pilha = append(pilha, j)

			if index == len(s)-1 {
				sort.Ints(pilha)
				min := strconv.Itoa(pilha[0])
				max := strconv.Itoa(pilha[len(pilha)-1])
				if min == max {
					result = append(result, "["+max+"]")
				} else {
					result = append(result, "["+min+"-"+max+"]")
				}
			}
		}
	}

	str := strings.Join(result, ", ")
	return str, nil
}

func main() {

}
