package main

import "fmt"

const LoginToken string = "logintokenn78" //We created this const with first letter capital. So this is a public variable.

/*Public variable :->

1. Case Sensitivity - Start with capital letter
2. Scope - Public variables are declared at a package level outside any function.
3. Access - Can be accessed from any file within the same package using their full name. * Public variables can be accessed from other packages by importing the package.
*/

func main() {
	//Var can be changed.
	var username string = "kapil" //String data type
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true //Boolean data type
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 256 //Integer data type
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float64 = 255.455845148451895 //Floating point data type
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	//Default values and some Aliases

	var anotherVar int
	fmt.Println(anotherVar)
	fmt.Printf("Variable is of type: %T \n", anotherVar)
	//We can declare a variable without specifying types for it.
	var website = "kapil-kumar.netlify.app"
	fmt.Println(website)
	fmt.Printf("Variable is of type: %T \n", website)

	//No var style
	//:= can only be used inside functions, not at the package level.
	//:= cannot be used for re-assigning values to existing variables. Use = for that.
	//:= simplifies code by avoiding explicit type declarations and makes it more concise and readable.

	withoutVar := 5849
	fmt.Println(withoutVar)

	//We can also define multiple variables in a single line using := like below.

	a, b, c := 5, 4, 3
	fmt.Println(a, b, c)

	fmt.Println(LoginToken)
	anotherFunc()

}
func anotherFunc() {
	fmt.Println("Hello How are you")
}
