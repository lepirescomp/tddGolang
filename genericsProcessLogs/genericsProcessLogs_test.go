package main

import (
	"reflect"
	"testing"
)

func TestGetLogKeys(t *testing.T) {

	t.Run("Test get log keys output", func(t *testing.T) {
		logRecord1 := map[string]string{
			"1": "User1",
			"2": "paymentAmount1",
			"3": "data1",
		}
		expectedRecords1Keys := []string{"1", "2", "3"}

		logRecord2 := map[int]string{
			1: "User2",
			2: "paymentAmount2",
			3: "data2",
		}
		expectedRecords2Keys := []int{1, 2, 3}

		logKeys1 := getLogKeys(logRecord1)
		if !reflect.DeepEqual(expectedRecords1Keys, logKeys1) {
			t.Errorf("There is an error in processing logs with key as strings")
		}

		logKeys2 := getLogKeys(logRecord2)
		if !reflect.DeepEqual(expectedRecords2Keys, logKeys2) {
			t.Errorf("There is an error in processing logs with key as ints")
		}
	})
}
