package stats
import (
	"github.com/iamgafurov/bank/v2/pkg/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	for _,payment := range payments{
		if(payment.Status != "Fail"){
			sum= sum + payment.Amount
		}
	}
	return sum / types.Money(len(payments))

}

func TotalInCategory(payments []types.Payment, category types.Category)types.Money{
	total:= types.Money(0)
	for _, payment := range payments{
		if(payment.Category == category && payment.Status != "Fail" ){
			print(payment.Status)
			total = total + payment.Amount
		}
	}
	return total

}