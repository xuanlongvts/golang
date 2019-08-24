package  main

import "fmt"
import "validator"

//import helperComm "happy-new-year/helper" // alias other name
import "happy-new-year/helper"

func main() {
	fmt.Println("This is main method")

	//helperComm.HelperConvert()
	helperMain.HelperConvert()

	validator.ValidatorEmail()
	validator.ValidatorIsPhone()
}