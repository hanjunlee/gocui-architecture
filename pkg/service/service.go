package service

import "strings"

var (
	zoo = []string{"alligator", "ant", "bear", "bee"}
	mart = []string{"apple", "banana", "cookie", "donut"}
)

// ZooService is the zoo service.
type ZooService struct {}

// Search search animals in the zoo.
func (z *ZooService) Search(pattern string) []string{
	ret := make([]string, 0)
	for _, a := range zoo {

		if pattern == "" {
			ret = append(ret, a)
			continue
		}

		if i := strings.Index(a, pattern); i != -1 {
			ret = append(ret, a)
		}
	}
	return ret
}

// MartService is the mart service.
type MartService struct {}

// Search search foods in the market.
func (z *MartService) Search(pattern string) []string{
	ret := make([]string, 0)
	for _, f := range mart {

		if pattern == "" {
			ret = append(ret, f)
			continue
		}

		if i := strings.Index(f, pattern); i != -1 {
			ret = append(ret, f)
		}
	}
	return ret
}
