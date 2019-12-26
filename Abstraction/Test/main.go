package main

import (
	"../../Abstraction/Classes"
	"../../Abstraction"
	"../../Interface"
	"fmt"
)

func main() {

	t := Classes.Triangle{Base: 15, Height: 25}
	r := Classes.Rectangle{Length: 5, Width: 10}
	c := Classes.Circle{Radius: 5}

	fmt.Println(t.Area(), "triangle")
	fmt.Println(r.Area(), "Rectangle")
	fmt.Println(c.Area(), "Circle")

	//Create an Instance of Mobile Type and Call Methods
	mobilePrd := Classes.Mobile{
		Abstraction.Product{
			"Apple iPhone 6s (Rose Gold, 32 GB)  (2 GB RAM)",
			"Handset, Apple EarPods with Remote and Mic, Lightning to USB Cable",
			40999, 143, 10,
		},
		[]string{"2 GB RAM", "4.7 inch Retina HD Display", "12MP Primary Camera"},
		[]string{"32 GB ROM", "4.7 inch Retina HD Display", "5MP Front"},
	}

	//Create an Instance of Shirts Type and Call Methods
	shirtPrd := Classes.Shirts{
		Abstraction.Product{
			"Reebok Striped Men's Round Neck Grey T-Shirt",
			"Machine Wash as per Tag, Do not Dry Clean, Do not Bleach",
			1124, 400, 25,
		},
		"XL", "Striped", "Half Sleeve", "Cotton",
	}

	//Create an Instance of Catalog Type
	category := Classes.Catalog{
		"Amazon",
		[]Interface.Information{mobilePrd, shirtPrd},
	}
	category.CatalogDetails()

	var num =7
	fmt.Print("Enter any positive integer: ")
	//fmt.Scanln(&num)
	mul := multiply(num)
	fmt.Println("Ten times of a given number is: ", mul.tentimes())
}

type multiply int

func (m multiply) tentimes() int {
	return int(m * 10)
}
