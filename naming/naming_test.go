package naming

import "testing"

func TestDog_Bark(t *testing.T) {

}

func TestDog_Bark_muzzeled(t *testing.T) {

}

func TestDog_Bark_unmuzzeled(t *testing.T) {

}

func TestSpeak_spanish(t *testing.T) {

}

func TestColour(t *testing.T) {
	arg := "blue"
	want := "#0000FF"
	got := Colour(arg)
	if got != want {
		t.Errorf("Colour(%s) = %s; want %s", arg, got, want)
	}
}
