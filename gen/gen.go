package main

import (
	"fmt"
)

func permute(m *[16]uint32) {
	*m = [16]uint32{
		m[2], m[6], m[3], m[10],
		m[7], m[0], m[4], m[13],
		m[1], m[11], m[12], m[5],
		m[9], m[14], m[15], m[8],
	}
}

func main() {

	var m [16]uint32

	for i := range m {
		m[i] = uint32(i)
	}
	for x := 1; x < 8; x++ {
		fmt.Printf(`// round%d
		
	// Mix the columns.
	gx(&state, 0, 4, 8, 12, n.block[%d])
	gy(&state, 0, 4, 8, 12, n.block[%d])
	gx(&state, 1, 5, 9, 13, n.block[%d])
	gy(&state, 1, 5, 9, 13, n.block[%d])
	gx(&state, 2, 6, 10, 14, n.block[%d])
	gy(&state, 2, 6, 10, 14, n.block[%d])
	gx(&state, 3, 7, 11, 15, n.block[%d])
	gy(&state, 3, 7, 11, 15, n.block[%d])

	// Mix the diagonals.
	gx(&state, 0, 5, 10, 15, n.block[%d])
	gy(&state, 0, 5, 10, 15, n.block[%d])
	gx(&state, 1, 6, 11, 12, n.block[%d])
	gy(&state, 1, 6, 11, 12, n.block[%d])
	gx(&state, 2, 7, 8, 13, n.block[%d])
	gy(&state, 2, 7, 8, 13, n.block[%d])
	gx(&state, 3, 4, 9, 14, n.block[%d])
	gy(&state, 3, 4, 9, 14, n.block[%d])

`, x, m[0], m[1], m[2], m[3], m[4], m[5], m[6], m[7], m[8], m[9], m[10], m[11], m[12], m[13], m[14], m[15])
		permute(&m)
	}
}