package shape

func Perimeter(widt float64, height float64) float64 {
	perimeter := (widt + height) * 2
	return perimeter
}

func Area(width float64, height float64) float64 {
	area := width * height
	return area
}
