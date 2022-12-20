package utils
import (
	"fmt"
)

type Map map[string]string;

var Pairs_Map = make(Map);

// IsEmpty: check if map of key-value pairs is empty
func (m *Map) IsEmpty() bool {
	return len(*m) == 0
}

// Print: print stack elements
func (m *Map) Print() {
	fmt.Println("Environment Variables ::")
	for k, v := range *m {
		fmt.Printf("%v : %v\n", k, v);
	}
	fmt.Println();
}

func (m *Map) containsKey(key string) bool {
	if _, exists := (*m)[key]; exists {
		return true;
	} else {
		return false;
	}
}

// ACCESS --> Pairs_Map[key]
// ADD --> Pairs_Map[key] = value
