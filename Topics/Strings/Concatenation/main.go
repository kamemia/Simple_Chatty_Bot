package main

import "fmt"

func reconstructSpeech(words []string) (speech string) {
	// From the very beginning the speech is empty:
	speech = ""

	for _, word := range words {
		speech += word // Concatenate the 'speech' and 'word' variables here!
	}

	// DO NOT remove the 'return speech' statement!
	return speech
}

// DO NOT MODIFY the main function below! It contains the parts of the speech:
func main() {
	var words = []string{
		"Akatsuki wa tsudota!\n",
		"Konan, Itachi, Kisame, Sasori, Deidara, Kakuzu, Hidan, Orochimaru, Zetsu.\n",
		"Each of you has your own agenda and your own dreams...\n",
		"We have finally come this far...\n",
		"And now at last we will move into full-scale action.",
	}

	fmt.Println(reconstructSpeech(words))
}
