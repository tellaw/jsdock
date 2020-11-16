package prompt

/*
This file should contain all func usefull for the prompt command
*/

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

/*
InList is a func used to ask in pultiple choice from command prompt
*/
func InList(values []string, title string) string {

	prompt := promptui.Select{
		Label: title,
		Items: values,
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		panic(err)
	}

	return result

}
