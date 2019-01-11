package array

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {

	var result []string
	if len(list1) == 0 || len(list2) == 0 {
		return result
	}

	dorisMem := configMap(list2)

	minIndexSum := int(^uint(0) >> 1)
	for index1, str1 := range list1 {
		if index2, ok := dorisMem[str1]; ok {
			indexSum := index1 + index2
			if indexSum < minIndexSum {
				minIndexSum = indexSum
				result = result[:0]
				result = append(result, str1)
			} else if indexSum == minIndexSum {
				result = append(result, str1)
			}
		}
	}

	return result
}

func configMap(list []string) map[string]int {
	result := make(map[string]int)
	for index, str := range list {
		result[str] = index
	}

	return result
}

func array() {
	list1 := []string{"Shogun", "Tapioca Express", "Burger King", "KFC"}
	list2 := []string{"Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"}
	list3 := []string{"KFC", "Shogun", "Burger King"}
	fmt.Println(findRestaurant(list1, list2))
	fmt.Println(findRestaurant(list1, list3))
}
