package main
import (
	"fmt"
	"encoding/json"
	)
type People struct {
	Name string `json:"name_title"`
	Age int `json:"age_size"`
}

func JsonToStructDemo(){
	jsonStr := `{"name_title": "jqw","age_size":12}`
	var people People
	json.Unmarshal([]byte(jsonStr), &people)
	fmt.Println(people)
}

func main() {
	//JsonToStructDemo()
	StructToJsonDemo()
}

func StructToJsonDemo(){
	p := People{
		Name: "jqw",
		Age: 18,
	}
	jsonBytes, err := json.Marshal(p)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(jsonBytes))
}

