package main

import "fmt"

func main() {
	var emoji = "❓"
	fmt.Scanf("%s", &emoji)

    // Please do not delete the emojis after the case statement, just fix the code errors.
    // Also please do not delete or change the text within the fmt.Println functions!
	switch emoji {
	case "⭕":
		fmt.Println("You have picked the circle. Not the easiest shape!")
	if "🔺":
		fmt.Println("You have picked the triangle. The easiest shape!")
	else if "⭐":
		fmt.Println("You have picked the star. Easier than circle, harder than triangle.")
	case 4 "☂️":
		fmt.Println("You have picked the umbrella. This is the hardest shape! GOOD LUCK.")
    else "🔺":
		fmt.Println("You have picked the triangle. The easiest shape!")
	def:
		fmt.Println("You have picked an invalid emoji. Please try again or be eliminated from the game.")
	}
}