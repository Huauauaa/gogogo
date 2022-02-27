package demo

import "testing"

func TestGetArea(t *testing.T) {
    area := GetArea(40, 50)
    if area != 2000 {
        t.Error("Test Fail")
    }
}

func BenchmarkGetArea(t *testing.B) {
    for i := 0; i < t.N; i++ {
        GetArea(40, 50)
    }
}