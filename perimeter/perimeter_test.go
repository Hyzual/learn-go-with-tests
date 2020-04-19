package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36},
	}

	for _, testData := range areaTests {
		t.Run(testData.name, func(t *testing.T) {
			got := testData.shape.Area()
			if got != testData.hasArea {
				t.Errorf("%#v got %.2f hasArea %.2f", testData, got, testData.hasArea)
			}
		})
	}
}
