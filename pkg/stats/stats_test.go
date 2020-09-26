package stats
import (
	"github.com/iamgafurov/bank/v2/pkg/types"
	"reflect"
	"testing"
)
func TestCategoriesTotal(t *testing.T){
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount:1_000_00},
		{ID: 1, Category: "food", Amount:2_000_00},
		{ID: 1, Category: "auto", Amount:3_000_00},
		{ID: 1, Category: "auto", Amount:4_000_00},
		{ID: 1, Category: "fun", Amount:5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto":8_000_00,
		"food": 2_000_00,
		"fun":5_000_00,
	}
	result := CategoriesTotal(payments)
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}


func TestCategoriesAvg(t *testing.T){
	payments := []types.Payment{
		{ID: 1, Category: "auto", Amount:2_000_00},
		{ID: 1, Category: "food", Amount:2_000_00},
		{ID: 1, Category: "auto", Amount:3_000_00},
		{ID: 1, Category: "auto", Amount:4_000_00},
		{ID: 1, Category: "fun", Amount:5_000_00},
	}
	expected := map[types.Category]types.Money{
		"auto":3_000_00,
		"food": 2_000_00,
		"fun":5_000_00,
	}
	result := CategoriesAvg(payments)
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}

func TestPeriodsDynamic(t *testing.T){
	first := map [types.Category]types.Money{
		"auto": 10,
		"food": 20,
	}
	second := map [types.Category]types.Money{
		"auto":20,
		"food":5,
		 "fun":5,
	}
	expected := map[types.Category]types.Money{
		"auto":10,
		"food": -15,
		"fun":5,
	}
	result := PeriodsDynamic(first,second)
	if !reflect.DeepEqual(expected, result){
		t.Errorf("invalid result, expected: %v, actual: %v",expected,result)
	}
}