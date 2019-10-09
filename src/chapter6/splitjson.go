package main

import (
	"encoding/json"
	"fmt"
)

// 使用匿名结构体分离json数据
type Screen struct {
	Size float32
	ResX, ResY int
}

type Battery struct {
	Capacity int
}

func getJsonData() []byte  {
	raw := &struct {
		Screen
		Battery
		HasTouchID bool
	}{
		Screen:Screen{
			Size: 5.5,
			ResX: 1920,
			ResY: 1080,
		},
		Battery:Battery{Capacity:2910},
		HasTouchID: true,
	}

	jsonData,_ := json.Marshal(raw)

	return jsonData
}

func main() {
	jsonData := getJsonData()

	fmt.Println(string(jsonData))

	screenAndTouch := struct {
		Screen
		HasTouchID bool
	}{

	}
	json.Unmarshal(jsonData, screenAndTouch)

	fmt.Printf("%+v \n", screenAndTouch)

	BatteryAndTouch := struct {
		Battery
		HasTouchID bool
	}{}

	json.Unmarshal(jsonData, BatteryAndTouch)
	fmt.Printf("%+v\n", BatteryAndTouch)
}
