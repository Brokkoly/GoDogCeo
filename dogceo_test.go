package godogceo

import "testing"

func TestBreedList(t *testing.T) {
	breedList, err := GetDogBreeds()
	if err != nil {
		t.Error(err)
	}
	if breedList.Message == nil {
		t.Error("Recieved an empty map of breeds")
	}
}

func TestRandomImage(t *testing.T) {
	image, err := GetRandomImage()
	if err != nil {
		t.Error(err)
	}
	if image == "" {
		t.Error("Received an empty image string")
	}
}

func TestRandomBreedImage(t *testing.T) {
	image, err := GetRandomBreedImage("shiba")
	if err != nil {
		t.Error(err)
	}
	if image == "" {
		t.Error("Recieved an empty image string")
	}
}

func TestRandomSubBreedImage(t *testing.T) {
	image, err := GetRandomSubBreedImage("bulldog", "french")
	if err != nil {
		t.Error(err)
	}
	if image == "" {
		t.Error("Recieved an empty image string")
	}
}
