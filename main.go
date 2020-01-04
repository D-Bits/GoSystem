
package main

import (
	"fmt"
)


func main()  {

	// Get the directory from the user
	fmt.Print("Enter the full path of the parent directory where you want to create your new directories: ")
	var userDir string
	fmt.Scanf("%s", &userDir)

	// Prompt the user to enter how many dirs they want to make
	fmt.Print("How many directories do you want to create?: ")
	var uDirNum int
	fmt.Scanf("%e", uDirNum)
	createDirs(uDirNum, userDir)

	/* fmt.Print("Enter a string: ")
	var uInput string
	fmt.Scanf("%s", &uInput)
	fmt.Println("Your string is:", uInput)

	fmt.Print("Press enter to end program: ")
	var endProgram string
	fmt.Scanf("%s", endProgram) */

}