package track

import (
	"fmt"
	"sort"
	"testing"
)

var t0 *testing.T
var testCase string

// 実行中の変更禁止
var baseList TrackList = TrackList{
	Track{Title: "0", Artist: "0", Album: "0", Year: 1900, Length: Length("0s")},
	Track{Title: "1", Artist: "3", Album: "4", Year: 1904, Length: Length("2s")},
	Track{Title: "2", Artist: "1", Album: "3", Year: 1903, Length: Length("1s")},
	Track{Title: "5", Artist: "3", Album: "3", Year: 1903, Length: Length("1s")},
	Track{Title: "1", Artist: "2", Album: "2", Year: 1902, Length: Length("0s")},
	Track{Title: "0", Artist: "5", Album: "1", Year: 1901, Length: Length("5s")},
}

//////////////////////////////////////////////////////////////

func TestChangeAndResetOrderKeys(t *testing.T) {
	setup(t, "ChangeAndResetOrderKeys")

	assertStrings([]string{"Title", "Artist", "Album", "Year", "Length"}, lessOrder)

	SetFirstOrder("Artist")
	assertStrings([]string{"Artist", "Title", "Album", "Year", "Length"}, lessOrder)

	SetFirstOrder("Album")
	assertStrings([]string{"Album", "Artist", "Title", "Year", "Length"}, lessOrder)

	SetFirstOrder("Length")
	assertStrings([]string{"Length", "Album", "Artist", "Title", "Year"}, lessOrder)

	ResetOrder()
	assertStrings([]string{"Title", "Artist", "Album", "Year", "Length"}, lessOrder)

	tearDown()
}

func TestSortDefaultOrder(t *testing.T) {
	setup(t, "SortDefaultOrder")

	ResetOrder()
	list := createList()

	sort.Sort(&list)
	assertPoints(list, 0, 5, 4, 1, 2, 3)

	tearDown()
}

func TestSortChangedOrder1(t *testing.T) {
	setup(t, "SortChangedOrder1")

	ResetOrder()
	SetFirstOrder(SortKeyArtist)
	list := createList()

	sort.Sort(&list)
	assertPoints(list, 0, 2, 4, 1, 3, 5)

	tearDown()
}

func TestSortChangedOrder2(t *testing.T) {
	setup(t, "SortChangedOrder2")

	ResetOrder()
	SetFirstOrder(SortKeyYear)
	SetFirstOrder(SortKeyArtist)
	list := createList()

	sort.Sort(&list)
	assertPoints(list, 0, 2, 4, 3, 1, 5)

	tearDown()
}

//////////////////////////////////////////////////////////////

func createList() TrackList {
	cp := make([]Track, len(baseList))
	copy(cp, baseList)
	return TrackList(cp)
}

func setup(t *testing.T, testCase0 string) {
	t0 = t
	testCase = testCase0
}

func tearDown() {
	t0 = nil
	testCase = ""
}

func assertPoints(actual TrackList, eIndices ...int) {
	assertTrue(len(eIndices) == len(actual), fmt.Sprintf("Mismatch length %d != %d", len(eIndices), len(actual)))
	result := true

	for i := 0; i < len(eIndices); i++ {
		result = result && assertTrue(actual[i] == baseList[eIndices[i]], fmt.Sprintf("Mismatch element at %d", i))
	}

	if !result {
		fmt.Printf("Actual : %v\n", actual)
	}
}

func assertStrings(expected, actual []string) {
	assertTrue(len(expected) == len(actual), fmt.Sprintf("Mismatch length %d != %d", len(expected), len(actual)))

	for i := 0; i < len(expected); i++ {
		assertTrue(expected[i] == actual[i], fmt.Sprintf("Mismatch element at %d", i))
	}
}

func assertInt(expected, actual int) {
	assertTrue(expected == actual, fmt.Sprintf("Unexpected value : %d != %d", expected, actual))
}

func assertTrue(result bool, msg string) bool {
	if t0 == nil {
		panic(0)
	}

	if !result {
		t0.Errorf("Test Failed [%s] : %s", testCase, msg)
	}

	return result
}
