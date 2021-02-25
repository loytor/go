package main

type Rect struct {
	width, height float64
}

func main() {

}

func (r *Rect) size() float64 {
	return r.width * r.height
}
