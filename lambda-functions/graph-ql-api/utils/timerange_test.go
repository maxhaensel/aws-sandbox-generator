package utils_test

import (
	"fmt"
	"lambda/aws-sandbox/graph-ql-api/utils"
	"testing"
)

func TestTimeRange(t *testing.T) {
	a, b := utils.TimeRange(2022, 02, 02)
	fmt.Println(*a)
	fmt.Println(*b)
}
