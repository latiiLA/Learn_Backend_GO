package main

import "fmt"

func average_computer(grades []float64) float64{
	/*
	This function recieves list of floating numbers
	The function calculates the average value from the list of numbers
	The average will then returned which most likey be floating number
	*/
	if len(grades) == 0{
		return 0.0
	}
	sum := 0.0
	for _, grade := range(grades){
		sum += grade
	}

	return sum / float64(len(grades))
}

func main(){
	var name string
	fmt.Println("Please enter your name: ")
	fmt.Scan(&name)

	var number_of_subjects int
	fmt.Println("Please enter number subects you have taken")
	fmt.Scan(&number_of_subjects)

	grades := make([]float64, number_of_subjects)
	subjects := make([]string, number_of_subjects)
	start := 0

	for start < number_of_subjects{
		fmt.Println("Please enter subject name.")
		var subject string
		fmt.Scan(&subject)
		fmt.Println("Enter course grade achieved.")
		var grade float64
		fmt.Scan(&grade)
		if grade > 100.0{
			fmt.Println("Course score cannot be greater than 100. Please reenter the correct course score again.")
			continue
		}else if grade < 0{
			fmt.Println("Course score cannot be less than 0. Please reenter the correct course score again.")
			continue
		}else{
			grades[start] = grade
			subjects[start] = subject
			start++
		}
	}

	average_grade := average_computer(grades)
	fmt.Printf("Student %s enrolled in %d courses has achieved an average grade of %f.", name, number_of_subjects, average_grade)
}