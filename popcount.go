package popcount

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount は x のポピュレーションカウント(1が設定されているビット数)
// を数える
func PopCount(x uint64) int {
	var sum int
	var i uint64
	for i = 0; i < 8; i++ {
		sum += int(pc[byte(x>>(i*8))])
	}
	return sum
}

// 愚直な実装
func PopCount2(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 最下位ビットの検査を64回繰り返す方法
func PopCount3(x uint64) int {
	var sum int
	var i uint64
	for i = 0; i < 64; i++ {
		sum += int(byte((x >> i) & 1))
	}
	return sum
}
