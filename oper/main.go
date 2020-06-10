package main

import (
	"fmt"
	"strings"
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack) Length() int {
	return len(*s)
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Dequeue() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		element := (*s)[0]
		*s = (*s)[1:]
		return element, true
	}
}

func (s *Stack) Enqueue(str string) {
	newSlice := []string{str}
	*s = append(newSlice, *s...)
}

func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) SimulatePop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		return element, true
	}
}

func isOperator(input string) bool {
	m := map[string]bool{">": true, "has": true, "<": true, "=": true}
	_, ok := m[input]
	return ok
}

func isLogic(input string) bool {
	m := map[string]bool{"&&": true, "or": true}
	_, ok := m[input]
	return ok
}

func main() {
	var values Stack
	var operators Stack
	var logic Stack
	var result Stack

	// add '(' and ')'
	//inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] )"
	//inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] or id = 10 )"
	//inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] or id = 10 or b < 1 )"
	//inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] or id = 10 or ( b < 1 && x = 4 ) )"
	//inputString := "( ( tags has ['a','b'] ) or ( ( b = 1 ) && ( c = 0 ) ) )"
	//inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] or id = 10 or ( b < 1 && ( x < 1 or y = 0 ) ) )"
	// inputString := "( ( age > 34 && created > 'date' ) or tags has ['all'] )"
	//inputString := "( ( ( ( tags has ['calm'] ) && ( tags has ['small'] ) ) or ( ( ( age > 10) && ( age < 30 ) ) or ( id = 10 ) ) ) && ( name = Pongo ) )"
	inputString := "( tags has ['calm','small'] ) or ( ( age > 10 ) && ( age < 30 ) )"
	// inputString := "( ( tags has ['calm','small'] ) or ( age > 10 ) ) && ( age < 30 )"
	// inputString := "( tags has ['calm','small'] ) && ( ( age > 10 ) or ( age < 30 ) )"
	splitString := strings.Split(inputString, " ")
	for _, item := range splitString {
		if item == "(" {
			continue
		} else if isOperator(item) {
			operators.Enqueue(item)
		} else if isLogic(item) {
			logic.Push(item)
		} else if item == ")" {
			fmt.Println(values)
			fmt.Println(result)
			fmt.Println(operators)
			fmt.Println(logic)

			if !operators.IsEmpty() {
				op, ok := operators.Pop()
				fmt.Println("operator", op)
				for ok {
					l, rok := values.Pop()
					r, lok := values.Pop()
					if !rok && !lok {
						panic("invalid format")
					}
					result.Push(fmt.Sprintf("( %s %s %s )", l, op, r))
					op, ok = operators.Pop()
				}

			} else if !logic.IsEmpty() {

				lo, _ := logic.Pop()
				fmt.Println("logic", lo)

				fmt.Println(result)
				switch lo {
				case "&&":
					r, rok := result.Pop()
					l, lok := result.Pop()
					if !rok && !lok {
						panic("invalid format")
					}
					result.Push(fmt.Sprintf("( %s %s %s )", l, lo, r))
				case "or":
					r, rok := result.Pop()
					l, lok := result.Pop()
					if !rok && !lok {
						panic("invalid format")
					}
					result.Push(fmt.Sprintf("( %s %s %s )", l, lo, r))
				}
			}

		} else {
			values.Enqueue(item)
		}
	}

	fmt.Println("stacks")
	fmt.Println("values:", values)
	fmt.Println("operators:", operators)
	fmt.Println("logic:", logic)
	fmt.Println("result:", result)
	fmt.Println()
	fmt.Println("result")
	fmt.Println("input", inputString)
	fmt.Println("result", "( "+result[0]+" )")

}
