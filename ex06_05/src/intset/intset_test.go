package intset

import (
	"fmt"
	"testing"
)

var t0 *testing.T
var testCase string

func TestAddAndRemove(t *testing.T) {
	setup(t, "Add/Remove")

	set := new(IntSet)
	validate(set)

	set.Add(10)
	validate(set, 10)

	set.Add(20)
	validate(set, 10, 20)

	set.Add(10)
	validate(set, 10, 20)

	set.Add(33)
	set.Add(67)
	validate(set, 10, 20, 33, 67)

	set.Remove(10)
	validate(set, 20, 33, 67)

	set.Remove(1)
	validate(set, 20, 33, 67)

	set.Clear()
	validate(set)

	tearDown()
}

func TestAddAll(t *testing.T) {
	setup(t, "AddAll")

	set := new(IntSet)

	set.AddAll(0, 1, 2, 3)
	validate(set, 0, 1, 2, 3)

	set.AddAll(0, 2, 4, 6)
	validate(set, 0, 1, 2, 3, 4 ,6)

	set.AddAll(16, 8, 4, 2, 1)
	validate(set, 0, 1, 2, 3, 4 ,6, 8, 16)

	tearDown()
}

func TestCopy(t *testing.T) {
	setup(t, "Copy")

	originalSet := new(IntSet)

	originalSet.AddAll(0, 1, 2, 3)
	validate(originalSet, 0, 1, 2, 3)

	copySet := originalSet.Copy()
	validate(copySet, 0, 1, 2, 3)

	originalSet.Clear()
	validate(originalSet)
	validate(copySet, 0, 1, 2, 3)

	tearDown()
}

func TestUnion1(t *testing.T) {
	setup(t, "Union1")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create(10, 20, 40, 80, 160)

	set1.UnionWith(set2)

	validate(set1, 10, 20, 30, 40, 50, 60, 80, 160)
	validate(set2, 10, 20, 40, 80, 160)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.UnionWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 10, 20, 30, 40, 50, 60, 80, 160)

	tearDown()
}

func TestUnion2(t *testing.T) {
	setup(t, "Union2")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create()

	set1.UnionWith(set2)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2)

	set2.UnionWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 10, 20, 30, 40, 50, 60)

	tearDown()
}

func TestUnion3(t *testing.T) {
	setup(t, "Union3")

	set1 := create()
	set2 := create()

	set1.UnionWith(set2)

	validate(set1)
	validate(set2)

	tearDown()
}

func TestIntersect1(t *testing.T) {
	setup(t, "Intersect1")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create(10, 20, 40, 80, 160)

	set1.IntersectWith(set2)

	validate(set1, 10, 20, 40)
	validate(set2, 10, 20, 40, 80, 160)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.IntersectWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 10, 20, 40)

	tearDown()
}

func TestIntersect2(t *testing.T) {
	setup(t, "Intersect2")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create()

	set1.IntersectWith(set2)

	validate(set1)
	validate(set2)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.IntersectWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2)

	tearDown()
}

func TestIntersect3(t *testing.T) {
	setup(t, "Intersect3")

	set1 := create()
	set2 := create()

	set1.IntersectWith(set2)

	validate(set1)
	validate(set2)

	tearDown()
}

func TestDifference1(t *testing.T) {
	setup(t, "Difference1")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create(10, 20, 40, 80, 160)

	set1.DifferenceWith(set2)

	validate(set1, 30, 50, 60)
	validate(set2, 10, 20, 40, 80, 160)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.DifferenceWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 80, 160)

	tearDown()
}

func TestDifference2(t *testing.T) {
	setup(t, "Difference2")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create()

	set1.DifferenceWith(set2)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.DifferenceWith(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2)

	tearDown()
}

func TestDifference3(t *testing.T) {
	setup(t, "Difference3")

	set1 := create()
	set2 := create()

	set1.DifferenceWith(set2)

	validate(set1)
	validate(set2)

	tearDown()
}

func TestSymmetrictDifference1(t *testing.T) {
	setup(t, "SymmetricDifference1")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create(10, 20, 40, 80, 160)

	set1.SymmetricDifference(set2)

	validate(set1, 30, 50, 60, 80, 160)
	validate(set2, 10, 20, 40, 80, 160)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.SymmetricDifference(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 30, 50, 60, 80, 160)

	tearDown()
}

func TestSymmetricDifference2(t *testing.T) {
	setup(t, "SymmetricDifference2")

	set1 := create(10, 20, 30, 40, 50, 60)
	set2 := create()

	set1.SymmetricDifference(set2)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2)

	set1 = create(10, 20, 30, 40, 50, 60)
	set2.SymmetricDifference(set1)

	validate(set1, 10, 20, 30, 40, 50, 60)
	validate(set2, 10, 20, 30, 40, 50, 60)

	tearDown()
}

func TestSymmetricDifference3(t *testing.T) {
	setup(t, "SymmetricDifference3")

	set1 := create()
	set2 := create()

	set1.SymmetricDifference(set2)

	validate(set1)
	validate(set2)

	tearDown()
}

//////////////////////////////////////////////////////////////////////

func validate(target *IntSet, expected ...int) {
	assertTrue(target.Len() == len(expected), fmt.Sprintf("Len() is unexpected result : actual = %d, expected = %d", target.Len(), len(expected)))

	// 期待値全てに対してHas()がtrueを返すことを確認
	for _, v := range expected {
		assertTrue(target.Has(v), fmt.Sprintf("IntSet does not contain expected value : %d", v))
	}

	elems := target.Elems()

	// Elems()に重複要素がないことを確認
	for _, e := range elems {
		assertTrue(count(elems, e) == 1, fmt.Sprintf("IntSet contains duplicate elements : %d", e))
	}

	// Elems()に期待値が全て含まれることを確認
	assertTrue(len(elems) == len(expected), fmt.Sprintf("Elems() size is unexpected : actual = %d, expected = %d", target.Len(), len(expected)))
	for _, v := range expected {
		assertTrue(count(elems, v) == 1, fmt.Sprintf("Elems() does not contain expected element : %d\nexpected=%v\nactual  =%v", v, expected, elems))
	}
}

func setup(t *testing.T, testCase0 string) {
	t0 = t
	testCase = testCase0

}

func tearDown() {
	t0 = nil
	testCase = ""
}

func create(values ...int) *IntSet {
	set := new(IntSet)

	for _, v := range values {
		set.Add(v)
	}

	return set
}

func assertTrue(result bool, msg string) {
	if t0 == nil {
		panic(0)
	}

	if !result {
		t0.Errorf("Test Failed [%s] : %s", testCase, msg)
	}
}

func count(values []int, target int) int {
	count := 0
	for _, v := range values {
		if v == target {
			count++
		}
	}
	return count
}
