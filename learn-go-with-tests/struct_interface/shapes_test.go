package main

import "testing"

// func TestPerimeter(t *testing.T) {
// 	rectangle := Rectangle{10.0, 10.0}
// 	got := Perimeter(rectangle)

// 	want := 40.0

// 	if got != want {
// 		t.Errorf("got %.2f, wnat %2.f", got, want)
// 	}
// }

func TestArea(t *testing.T) {
	// TDD 测试代码
	// checkArea := func(t *testing.T, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()

	// 	if got != want {
	// 		t.Errorf("got %.2f want %.2f", got, want)
	// 	}
	// }

	// t.Run("rectangle", func(t *testing.T) {
	// 	rectangle := Rectangle{12, 6}
	// 	// got := rectangle.Area()
	// 	want := 72.0
	// 	checkArea(t, rectangle, want)
	// })

	// t.Run("circles", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	// got := circle.Area()
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })

	// 表格测试代码
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12.0, 6.0}, 72.0},
		{Circle{10}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %.2f want %.2f %#v", tt.shape, got, tt.want, tt.shape)
		}
	}
}
