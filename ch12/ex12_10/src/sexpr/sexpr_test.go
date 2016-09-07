package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	type Movie struct {
		Title, Subtitle string
		Year            int
		Actor           map[string]string
		Oscars          []string
		Sequel          *string
	}
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		},
		Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}

	data, err := Marshal(strangelove)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	t.Logf("Marshal() = %s\n", data)

	reader := bytes.NewBuffer(data)
	decoder, _ := NewDecoder(reader)
	var movie Movie
	if err := Unmarshal(decoder, &movie); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}
	t.Logf("Unmarshal() = %+v\n", movie)

	if !reflect.DeepEqual(movie, strangelove) {
		t.Fatal("not equal")
	}

	data, err = MarshalIndent(strangelove)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("MarshalIdent() = %s\n", data)
}

////////////////////////////////////////////////////////

type inner struct {
	X int
	Y int
}

type sample struct {
	B bool
	F float64
	I interface{}
}

func TestEx12_10(t *testing.T) {
	RegisterTypeMapping("inner", reflect.TypeOf(inner{}))

	testcaseRedecode(t, sample{B: false, F: 1.2, I: inner{X: 1, Y: 2}})
	testcaseRedecode(t, sample{B: true, F: 0.0, I: inner{X: 0, Y: 0}})
}

func testcaseRedecode(t *testing.T, original sample) {
	data, err := Marshal(original)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}
	fmt.Println(string(data))

	reader := bytes.NewBuffer(data)
	decoder, _ := NewDecoder(reader)

	var redecoded sample
	if err := Unmarshal(decoder, &redecoded); err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if original.B != redecoded.B || original.F != redecoded.F ||
		reflect.ValueOf(original.I) == reflect.ValueOf(redecoded.I) {
		t.Fatalf("Decode failed")
	}
}
