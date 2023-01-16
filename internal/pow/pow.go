package pow

func Pow(x, y uint32) uint32 {
	res := uint32(1)
	for y > 0 {
		if y%2 == 0 {
			x *= x
			y >>= 1
		} else {
			res *= x
			y--
		}
	}
	return res
}
