package popcount

import (

)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i / 2] + byte(i & 1)
	}
}

// テーブルを利用してポピュレーションカウントを取得します
func PopCountWithTable(x uint64) int {
	return int(pc[byte(x >> (0 * 8))] +
		pc[byte(x >> (1 * 8))] +
		pc[byte(x >> (2 * 8))] +
		pc[byte(x >> (3 * 8))] +
		pc[byte(x >> (4 * 8))] +
		pc[byte(x >> (5 * 8))] +
		pc[byte(x >> (6 * 8))] +
		pc[byte(x >> (7 * 8))])
}

// テーブル+ループを利用してポピュレーションカウントを取得します
func PopCountWithTableAndLoop(x uint64) int {
	count := 0

	for i := 0; i < 8; i++ {
		count += int(pc[byte(x >> (uint(i) * 8))])
	}

	return count
}

// ビットシフトと最下位ビットの確認によりポピュレーションカウントを取得します
func PopCountWithLowestBitLoop(x uint64) int {
	count := 0

	for i := 0; i < 64; i++ {
		count += int(x & 1)
		x >>= 1
	}

	return count
}

// 最下位の1のビットを削除していくことでポピュレーションカウントを取得します
func PopCountWithBitRemoval(x uint64) int {
	count := 0

	for x != 0 {
		count++
		x &= (x - 1)
	}

	return count
}
