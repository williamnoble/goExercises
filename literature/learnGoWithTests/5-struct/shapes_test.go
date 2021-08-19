package shapes

import "testing"

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{
			name: "Rectangle",
			shape: Rectangle{
				width:  12,
				length: 6,
			},
			hasArea: 72.0,
		},
		{
			name:    "Circle",
			shape:   Circle{Radius: 10},
			hasArea: 314.1592653589793,
		},
		{
			name: "Triangle",
			shape: Triangle{
				Base:   12,
				Height: 6,
			},
			hasArea: 36.0, // 26.0 for error ex.
		},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}

//func TestPerimeter(t *testing.T) {
//	rectangle := Rectangle{
//		width:  10.0,
//		length: 10.0,
//	}
//	got := rectangle.Perimeter()
//	want := 40.0
//
//	if got != want {
//		// %.2f -> 2 decimmal places
//		t.Errorf("got %.2f want %.2f", got, want)
//	}
//}
//
//func TestArea(t *testing.T) {
//
//	checkArea := func(t testing.TB, shape Shape, want float64) {
//		t.Helper()
//		got := shape.Area()
//		if got != want {
//			t.Errorf("got %g want %g", got, want)
//		}
//	}
//
//
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{
//			width:  12,
//			length: 6,
//		}
//		want := 72.0
//		checkArea(t, rectangle, want)
//	})
//
// 	t.Run("circles", func(t *testing.T) {
//		circle := Circle{Radius: 10.0}
//		want :=  314.1592653589793
//		checkArea(t, circle, want)
//	})
//
//}
