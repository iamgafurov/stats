package stats
import (
	"github.com/iamgafurov/bank/v2/pkg/types"
)
func Avg(payments []types.Payment) types.Money{
	sum:= types.Money(0)
	kol:= 0
	for _,payment := range payments{
		if(payment.Status != "Fail"){
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
		if(payment.Category == category && payment.Status != "Fail" ){
			total = total + payment.Amount
		}
	}
	return total

}