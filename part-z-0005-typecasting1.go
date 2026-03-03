package main
import "fmt";

func main() {

	sales := 50_000

	var fixedCost uint16 = 38_500 

	salesPremiumRate := 0.08
	
	var netProfit = float64(sales - int(fixedCost)) * (1 - salesPremiumRate)

	fmt.Println(netProfit) // 10580
}
