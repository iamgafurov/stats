package stats
import (
	"github.com/iamgafurov/bank/pkg/types"
	"fmt"
)
func AvgTest(){
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
	},
	types.Payment{
		ID: 2,
		Amount: 1232,
	},
	types.Payment{
		ID: 3,
		Amount: 1232,
	},
	types.Payment{
		ID: 4,
		Amount: 1232,
	},
	}
fmt.Println(Avg(payments))
//Output: 1232
}

func TotalInCategoryTest(){
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
		Category:"b",
	},
	types.Payment{
		ID: 2,
		Category: "ss",
		Amount: 1232,
	},
	types.Payment{
		ID: 3,
		Amount:1232,
		Category:"b",
	},
	types.Payment{
		ID: 4,
		Category:"g",
		Amount: 1232,
	},
	}
fmt.Println(TotalInCategory(payments,"b"))
//Output: 2464
}