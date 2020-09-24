package stats
import (
	"github.com/iamgafurov/bank/pkg/bank/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	for _,payment := range payments{
		sum= sum + payment.Amount
	}
	return sum / types.Money(len(payments))

}

func TotalInCategory(payments []types.Payment, category types.Category)types.Money{
	total:= types.Money(0)
	for _, payment := range payments{
		if(payment.Category == category){
			total = total + payment.Amount
		}
	}
	return total

}