package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

//DogListing the json representation of /api/breeds/list/all
type DogListing struct {
	Message map[string][]string `json:"message"`
}

//DogImage the json representation of /api/breeds/image/random and
// /api/breeds/{breed name}/images/random
type DogImage struct {
	Message string `json:"message"`
}

//GetDogBreeds gets the list of dog breeds and sub breeds
func GetDogBreeds() (DogListing, error) {
	var dogs DogListing
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://dog.ceo/api/breeds/list/all", nil)
	if err != nil {
		return dogs, err
	}
	req.Header.Set("User-Agent", "GoDogCeo/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return dogs, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return dogs, err
	}
	json.Unmarshal(body, &dogs)
	return dogs, nil
}

//GetRandomImage gets a random image of any dog breed
func GetRandomImage() (string, error) {
	var image DogImage
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://dog.ceo/api/breeds/image/random", nil)
	if err != nil {
		return "", err
	}
	req.Header.Set("User-Agent", "GoDogCeo/0.1")
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	json.Unmarshal(body, &image)
	return image.Message, nil
}

func main() {
	breeds, err := GetRandomImage()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(breeds)
}
