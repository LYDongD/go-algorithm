package main

import "fmt"
import "strings"

type Stack struct {
	strs []string
}

func (self *Stack) isEmpty() bool {
	return len(self.strs) == 0
}

func (self *Stack) push(str string) {
	self.strs = append(self.strs, str)
}

func (self *Stack) pop() string {
	if self.isEmpty() {
		panic("stack is empty")
	}

	top := self.strs[len(self.strs)-1]
	self.strs = self.strs[:len(self.strs)-1]
	return top
}

func simplifyPath(path string) string {
	if len(path) == 0 {
		return ""
	}

	stack := &Stack{strs: []string{}}
	dirList := strings.Split(path, "/")
	for _, dir := range dirList {
		if dir == ".." {
			if !stack.isEmpty() {
				stack.pop()
			}
			continue
		}

		if dir != ".." && dir != "" && dir != "." {
			stack.push(dir)
		}
	}

	result := ""
	for !stack.isEmpty() {
		result = stack.pop() + result
		result = "/" + result
	}

	if result == "" {
		result = result + "/"
	}

	return result
}

func main() {
	fmt.Println(simplifyPath("/home/"))
	fmt.Println(simplifyPath("/../"))
	fmt.Println(simplifyPath("/home//foo/"))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
	fmt.Println(simplifyPath("/a/../../b/../c//.//"))
	fmt.Println(simplifyPath("/a//b////c/d//././/.."))
	fmt.Println(simplifyPath("/."))
	fmt.Println(simplifyPath("/..."))
	fmt.Println(simplifyPath("/a/./b/../../c/"))
}
