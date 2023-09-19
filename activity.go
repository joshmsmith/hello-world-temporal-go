package app

import (
    "errors"
    "context"
    "log"
    "math/rand"
    "fmt"
)

func ComposeGreeting(ctx context.Context, name string) (string, error) {
    
	//throw an error sometime 
	i := rand.Intn(2)

   info := fmt.Sprintf("the number that should not be 1 is %v:  ", i)

    if name == "" { 
	    return "", errors.New("empty name oops")
    } 

    if i  == 1 {
	    log.Print("I broke at random because ", info, ", darnit")
	    return info, errors.New("I broke at random")
    }

    greeting := fmt.Sprintf("Hello %s!", name)
    return greeting, nil
}

