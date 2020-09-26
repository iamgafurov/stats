package stats
import (
	"fmt"
	"github.com/iamgafurov/bank/v2/pkg/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	kol:= 0
	for _,payment := range payments{
		if(payment.Status != types.StatusFail ){
			sum= sum + payment.Amount
			kol = kol + 1
		}
	}
	if(kol == 0  ){
		return 0
	}else{
	return sum / types.Money(kol)
	}
}

func TotalInCategory(payments []types.Payment, category types.Category)types.Money{
	total:= types.Money(0)
	for _, payment := range payments{
		if(payment.Category == category && (payment.Status != types.StatusFail) ){
			total = total + payment.Amount
		}
	}
	return total

}

func CategoriesTotal(payments []types.Payment)map[types.Category]types.Money{
	categories := map[types.Category]types.Money{}
	for _,payment := range payments{
		categories[payment.Category] += payment.Amount
	}
	return categories
}

func CategoriesAvg(payments []types.Payment)map[types.Category]types.Money{
	categories := make(map[types.Category]types.Money)
	counts := make(map[types.Category]types.Money)
	for _,payment := range payments{
		categories[payment.Category] += payment.Amount
		counts[payment.Category] += 1
	}
	for key,category := range categories{
		categories[key] = category / counts[key]
		fmt.Println(counts)
	}
	return categories
}