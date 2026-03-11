package main

import (
	"fmt"
)

func main() {

	getLastYears(2026) // 2026 2025 2024 2023 2022 2021 2020 

	getLastYears(2000) 
}

func getLastYears(year uint16) {

	if year < 2020 { return }

	fmt.Printf("%d ", year)

	getLastYears(year - 1) 
}



