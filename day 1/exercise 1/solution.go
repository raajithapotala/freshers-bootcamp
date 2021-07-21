package main

import (
	"encoding/json"
	"fmt"
)

type Matrix struct {
	rows int
	columns int
	elements [][]int
}

func getrow (m Matrix) int {
	return m.rows
}

func getcol (m Matrix) int {
	return m.columns
}

func setelement (m Matrix , i int, j int, element int) {
	m.elements[i][j] =  element
}

func (matrix Matrix) AddMatrix(addMatrix Matrix) {
	if (len(matrix.elements) != len(addMatrix.elements)) || (len(matrix.elements[0]) != len(addMatrix.elements[0])) {
		panic("Unsupported matrix addition")
	}
	for i:=0 ; i<len(matrix.elements) ; i++ {
		for j:=0 ; j<len(matrix.elements[0]) ; j++ {
			matrix.elements[i][j] = matrix.elements[i][j] + addMatrix.elements[i][j]
		}
	}
}

func printmatrix (m Matrix){
	data , mat := json.MarshalIndent(m,"","  ")
	if(mat!=nil) {
		fmt.Println(mat)
	}
	fmt.Println(data)
}
func main() {
	m1 := [][] int {{1,1,1},{2,2,2},{3,3,3}}
	n1 := [][] int {{0,0,0},{1,1,1},{2,2,2}}
	m := Matrix{3,3,m1}
	n := Matrix{2,3,n1}
	fmt.Println("No of rows = ",getrow(m))
	fmt.Println("No of columns = ",getcol(m))
	setelement(m,1,1,5)
	fmt.Println(m.elements)
	printmatrix(m)
	m.AddMatrix(n)
	fmt.Println(m.elements)
	printmatrix(m)
}
