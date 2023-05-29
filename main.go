package main

import (
	"fmt"
	"panel/lib"
	"strconv"
)

func main() {
	r := lib.SetupApi()

	r.Run(fmt.Sprintf(":%s", strconv.Itoa(int(lib.ApiConfig.Port))))
}
