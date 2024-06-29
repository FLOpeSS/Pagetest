package handlers

import (
	"net/http"
	"testing"
)

func sum_test(x, y int) int {
	return x + y
}

func TestStatusJson(t *testing.T) {
	// req, err := http.NewRequest("GET", "http://localhost:3000/api/v1/status", nil)
	// if err != nil {
	// 	t.Error("req didn't work properly")
	// }
	// t.Log(req)
	// t.Log(req.Body)
	resp, err := http.Get("http://localhost:3000/api/v1/status")
	if err != nil {
		t.Error("Error on Get Request, error: ", err)
	}
	t.Logf("response body %d\n", resp.Body)
	t.Log("response only", resp)
}

// func TestSummed(t *testing.T) {
// 	a, b := 1, 2
// 	result := handlers.SumToTest(a, b)
// 	// t.Logf("result using SumToTest : %d", result)
// 	fmt.Printf("result using SumToTest: %d\n", result)
//
// 	expected := 3
// 	if result != expected {
// 		t.Errorf("result sum_test: %d, not equal to expected: %d", result, expected)
// 	}
//
// }
//
// func TestSum(t *testing.T) {
// 	a, b := 1, 2
// 	result := sum_test(a, b)
// 	t.Logf("result using sum_test: %d", result)
//
// 	expected := 3
// 	if result != expected {
// 		t.Errorf("result: %d, not equal to expected: %d", result, expected)
// 	}
// }
//
// func TestStatus(t *testing.T) {
// 	result := sum_test(5, 6)
// 	if result != 11 {
// 		t.Error("sum to test isn't working properly: ", result)
// 	}
// 	fmt.Println("sum_to_test is working properlyjk, result: ", result)
// }
