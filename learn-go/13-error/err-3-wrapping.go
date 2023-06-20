package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	err := openSomething("test.txt")
	if err != nil {
		//errors.Is reports whether any error in err's tree matches target.
		//it is used when we want to check if any specific error is present in the chain and the basis of it we would
		//take some actions
		if errors.Is(err, ErrFileNotFound) {
			//if this is true then we can take actions that might solve the problem otherwise will stop the func flow
			fmt.Println("err matched in the chain")
			return
		}
		fmt.Println("err is not in the chain, quit the func")
		return
	}
	fmt.Println("no err reported")

}

var ErrFileNotFound = errors.New("not found in root directory")

func openSomething(fileName string) error {

	_, err := os.OpenFile(fileName, 0, 0)
	if err != nil {
		//wrapping two values together and trying to return err value to the caller func
		//%w wrap the values
		return fmt.Errorf(" %w", err)
	}
	return nil
}
