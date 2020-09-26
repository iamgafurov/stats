package stats
import (
	"github.com/iamgafurov/bank/v2/pkg/types"
	"fmt"
)
func AvgTest(){
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
		Status: "Fail",
	},
	types.Payment{
		ID: 2,
		Amount: 1232,
		Status: "Ok",
	},
	types.Payment{
		ID: 3,
		Amount: 1232,
		Status: "Ok",
	},
	types.Payment{
		ID: 4,
		Amount: 1232,
		Status: "Ok",
	},
	}
fmt.Println(Avg(payments))
//Output: 
//1232
}

func TotalInCategoryTest(){
	payments := []types.Payment{
	types.Payment{
		ID: 1,
		Amount: 1232,
		Category:"b",
		Status: "Fail",
	},
	types.Payment{
		ID: 2,
		Category: "ss",
		Status: "Fail",
		Amount: 1232,
	},
	types.Payment{
		ID: 3,
		Amount:1232,
		Category:"b",
		Status: "Fail",
	},
	types.Payment{
		ID: 4,
		Category:"g",
		Amount: 1232,
	},
	}
fmt.Println(TotalInCategory(payments,"b"))

//Output: 
//2464
}