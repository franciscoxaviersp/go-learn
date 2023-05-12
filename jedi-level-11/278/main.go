package main

import (
	"fmt"
	"log"
)

type sqrtError struct {
	value float64
	err   error
}

func (se sqrtError) Error() string {
	return fmt.Sprintf("math error: %v %v", se.value, se.err)
}

func main() {
	_, err := sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		e := fmt.Errorf("value must positive. value given: %v", f)
		return 0, sqrtError{
			value: f,
			err:   e,
		}
	}
	return 42, nil
}
