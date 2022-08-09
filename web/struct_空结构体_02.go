package main

import "fmt"

//空结构体实现集合

type Set map[string]struct{}

func (s Set) Has(key string) bool {
	_, ok := s[key]
	return ok
}
func (s Set) Add(key string) {
	s[key] = struct{}{} //空结构体不占用内存
}
func (s Set) Delete(key string) {
	delete(s, key)
}

func main() {
	s := make(Set)
	s.Add("Name")
	fmt.Println(s.Has("qq"))
}
