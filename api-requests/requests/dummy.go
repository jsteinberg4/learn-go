package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Product struct {
	Id          int
	Title       string
	Description string
	Price       int
	Discount    float32
	Rating      float32
	Stock       int
	Brand       string
	Category    string
	Thumbnail   string
	Images      []string
}

func Dummy() {
	fmt.Println("Making get request")
	resp, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("err")
	} else {
		defer resp.Body.Close()
		resp_bytes, _ := ioutil.ReadAll(resp.Body)

		var resp_json interface{}
		json.Unmarshal(resp_bytes, &resp_json)
		fmt.Printf("Resp: %+v", resp_json)
	}
}
