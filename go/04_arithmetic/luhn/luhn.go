package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	
	idInt, err := strconv.Atoi(id)
	if err != nil{
		return false
	}


	panic("Please implement the Valid function")
}
