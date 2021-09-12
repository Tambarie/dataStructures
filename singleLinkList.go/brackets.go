package main

func main()  {

}

func isValid(s string) bool {
	myMap := map[rune]rune{'(' : ')', '[' : ']', '{' : '}'}
	var stack [] rune
	for _, v:= range s{
		if _, ok := myMap[v]; ok{
			stack = append(stack, v)
		}else{
			if len(stack) > 0{
				topItem := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if myMap[topItem] != v {
					return false
				}
			}else{
				return false
			}
		}
	}
	if len(stack) != 0{
		return false
	}
	return true
}