package tests

import (
	"fmt"
	"testing"

	"github.com/szwtdl/aqscwlxy/utils"
)

func TestUtils(t *testing.T) {
	image, err := utils.ImageToBase64("/Users/mac/go/aqscwlxy/demo/output_0005.jpg")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(image)
}

// 计算数量

func TestCalculateNum(t *testing.T) {
	m := []int{1, 2, 3, 4, 5, 6}
	oValue := 8
	num := utils.CalculateNum(m, oValue)
	fmt.Println(num)
}
