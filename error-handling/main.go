package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func downloadImage(url string) ([]int, error) {
	if !strings.Contains(url, "http") {
		return nil, errors.New("Must provide a valid url")
	}

	return []int{1, 2, 3}, nil
}

// these two are a simple example
func getProfileImage() {
	var imageBuffer, downloadError = downloadImage("http://zzzlololol.com")
	if downloadError != nil {
		fmt.Printf("Failed to download image: %v", downloadError)
		os.Exit(-1)
	}

	fmt.Printf("Successfully downloaded image: %v\n", imageBuffer)
}

// custom error handling examples below
type EnrichmentError struct {
	message   string
	faultyInt int
}

func (e *EnrichmentError) Error() string {
	return fmt.Sprintf("Enrichment Error: %v | Faulty argument: %v", e.message, e.faultyInt)
}

func enrichNumbers(numbers []int) ([]int, error) {
	enrichedNumbers := make([]int, len(numbers), len(numbers))
	for idx, number := range numbers {
		if number == 8 {
			return nil, &EnrichmentError{
				message:   "You can't use this number because I FREAKIN' hate it!!!",
				faultyInt: number,
			}
		}
		enrichedNumbers[idx] = number + 1
	}

	return enrichedNumbers, nil
}

// example with panicking
func coolFunctionThatTotallyWontBrickYourComputer(input int) []int {
	// deferred functions are called at the end of wrapping function's execution
	defer func() {
		// however, if the wrapping function panics, the deferred function will not be called
		// UNLESS it also calls recover() inside itself
		// this recover function returns an interface type
		if recoveryResult := recover(); recoveryResult != nil {
			fmt.Println("RECOVERED FROM:", recoveryResult)
			fmt.Println("End of the world scenario prevented!")
		}
	}()

	if input > 2 {
		// force an unhandled error with a string as its content
		panic("Oh my GAWD, the world is ending!")
	}

	return []int{9, 9, 9}
}

func main() {
	fmt.Println("-- Go Error Handling ---")
	numbers := []int{1, 2, 3}
	firstEnrichmentNumbers, firstEnrichmentError := enrichNumbers(numbers)
	if firstEnrichmentError != nil {
		fmt.Println("First enrichment failed:", firstEnrichmentError)
	} else {
		fmt.Println("First enrichment successful", firstEnrichmentNumbers)
	}

	otherNumbers := []int{6, 7, 8}
	// shorthand syntax
	if secondRes, secondErr := enrichNumbers(otherNumbers); secondErr != nil {
		fmt.Println("Second enrichment failed:", secondErr)
	} else {
		fmt.Println("Second enrichment successful", secondRes)
	}

	yetAnotherSlice := []int{8, 9, 10}
	if thirdRes, thirdErr := enrichNumbers(yetAnotherSlice); thirdErr != nil {
		fmt.Println("Third enrichment failed:", thirdErr)

		// let's say we want to extract EnrichmentError's props for some reason
		// we first need to try to assert ("cast") the error as an EnrichmentError
		if enrichmentError, ok := thirdErr.(*EnrichmentError); ok {
			// assertion successful
			// if I just print it like this, it will call it's Error method
			fmt.Println(enrichmentError)
			// I can also extract props
			fmt.Println("Error message extracted:", enrichmentError.message)
			fmt.Println("Faulty integer extracted:", enrichmentError.faultyInt)
		}

	} else {
		fmt.Println("Third enrichment successful", thirdRes)
	}

	someMoreNumbers := coolFunctionThatTotallyWontBrickYourComputer(3)
	fmt.Println("Here are some more numbers:", someMoreNumbers)
	fmt.Println("-- End of Go Error Handling --")
}
