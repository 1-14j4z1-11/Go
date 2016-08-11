package intset

import (

)

type baseType uint64

type IntSet struct {
	words []baseType
}

const (
	unit = 32 << (^baseType(0) >> 63)
)

// 値xが含まれているか判定します
func (s *IntSet) Has(x int) bool {
	word, bit := offset(x)
	return word < len(s.words) && s.words[word] & (1 << bit) != 0
}

// 値xを追加します
func (s *IntSet) Add(x int) {
	word, bit := offset(x)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// 別のIntSetとの和集合を取得します
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}
	}
}

// 要素数を取得します (6-1)
func (s *IntSet) Len() int {
	count := 0
	for _, word := range s.words {
		count += popcount(word)
	}
	return count
}

// 値xが含まれている場合は削除します (6-1)
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := offset(x)
		s.words[word] &= baseType(^(1 << bit))
	}
}

// 全ての値を削除します (6-1)
func (s *IntSet) Clear() {
	s.words = nil
}


// コピーを作成します (6-1)
func (s *IntSet) Copy() *IntSet {
	t := new(IntSet)
	for _, word := range s.words {
		t.words = append(t.words, word)
	}
	return t
}

// 値を追加します (6-2)
func (s *IntSet) AddAll(vals ...int) {
	for _, v := range vals {
		s.Add(v)
	}
}

// 別のIntSetとの積集合を取得します(6-3)
func (s *IntSet) IntersectWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= tword
		}
	}

	for i := len(t.words); i < len(s.words); i++ {
		s.words[i] = 0;
	}
}

// 別のIntSetとの差集合を取得します(6-3)
func (s *IntSet) DifferenceWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] &= ^tword
		}
	}
}

// 別のIntSetとの対称差を取得します(6-3)
func (s *IntSet) SymmetricDifference(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] ^= tword
		}
	}

	for i := len(s.words); i < len(t.words); i++ {
		s.words = append(s.words, t.words[i])
	}
}

// 現在含まれている要素のスライスを取得します (6-4)
func (s *IntSet) Elems() []int {
	var elems []int
	for i, word := range s.words {
		offset := 0
		for word != 0 {
			if word & 1 != 0 {
				elems = append(elems, i * unit + offset)
			}
			word >>= 1
			offset++
		}
	}
	return elems
}

func popcount(x baseType) int {
	count := 0

	for x != 0 {
		count++
		x &= (x - 1)
	}

	return count
}

func offset(x int) (words int, bit uint) {
	return x/unit, uint(x%unit)
}
