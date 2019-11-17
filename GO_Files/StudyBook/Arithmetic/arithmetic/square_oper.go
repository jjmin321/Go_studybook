//29-1

package arithmetic

//x, y 제곱의 합을 계산해서 반환

func (o *Numbers) SquarePlus() int {
	return (o.X * o.X) + (o.Y * o.Y)
}

//x, y 제곱의 차를 계산해서 반환

func (o *Numbers) SquareMinus() int {
	return (o.X * o.X) - (o.Y * o.Y)
}
