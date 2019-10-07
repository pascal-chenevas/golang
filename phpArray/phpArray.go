package main

import (
	"fmt"
	s "strings"
	"strconv"
	"regexp"
	"encoding/json"
)

type KeyValue struct {
	Key, Value string
}

type PhpArrays struct {
	KeysValues []KeyValue
}
	
	
func main() {
	array := `Array
(
    [A (column1)] => PHP code te
ster Sandbox Online
    [B (column
2)] => bar
    [C (column 
              e 
              [f%])] => 1
              2
              3
              4 
    [D (column4 d 
                Version)] => 7.
3.5
[A (column1)] => PHP code te
ster Sandbox Online
    [B (column
2)] => bar
    [C (column 
              e 
              [f%])] => 1
              2
              3
              4 
    [D (column4 d 
                Version)] => 7.
3.5
 [A (1)] => PHP code tester Sandbox Online
    [B (foo)] => bar
    [C (2)] => 5
    [D (5)] => 89009
    [E (case)] => Random Stuff: 873
    [F (PHP Version)] => 7.3.5
)`

	str := s.Fields(array)
	for i,a := range str {
		fmt.Println("i:" + strconv.Itoa(i) + string(a))
	}
	
	str2 := s.ReplaceAll(array, "\n", "")
	re := regexp.MustCompile(`\s+`)
	str3 := re.ReplaceAllString(str2, " ") 
	fmt.Println(str2)
	
	fmt.Println(str3)
	
	newStr := str3[len("Array ("):len(str3)-1]

	fmt.Println("---")	
	fmt.Println(newStr)
	
		
	fmt.Println("====")
	
	columnNames := []string{}
	u := regexp.MustCompile(`\[[A-Z]+ \((.*?)\)\]`)
	for _, submatches := range u.FindAllStringSubmatch(newStr, -1) {
		columnNames = append(columnNames, submatches[0])
	}
	
	fmt.Println("====")
	
	for _,col := range columnNames {
		newStr = s.ReplaceAll(newStr, col, "")
	}
	
	fmt.Println(newStr)
	
	ar :=  s.Split(newStr, " => ")
	columnValues := []string{}
	for _, fff := range ar {
		if len(fff) > 0 {
			columnValues = append(columnValues, fff)
		}
	}	
	
	
	fmt.Println("==[result]==")
	fmt.Println(columnNames)
	
	php := PhpArrays{}
	
	data := make(map[int][]map[string]string)
	nMap := make(map[string]string)
	count := 0
	for idx, k := range columnNames {
		pos := s.Index(k, "(")
		pos++
		name := s.Trim(k[pos:len(k)-2], " ")
		v := s.Trim(columnValues[idx], " ")
		
		php.KeysValues = append(php.KeysValues, KeyValue{name, v} )
		
		nMap[name] = v
		if idx == len(columnNames) -1 {
			data[count] = append(data[count], nMap)
			count++
			fmt.Println(nMap)
		}
		
	}
	
	fmt.Println(php.KeysValues)
	
	b, err := json.MarshalIndent(php.KeysValues, " ", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	
	fmt.Println(string(b))
	fmt.Println("====================================")
	
	b, err = json.MarshalIndent(data, " ", " ")
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(b))
	
}

