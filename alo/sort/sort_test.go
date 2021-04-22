package sort

import (
	"fmt"
	"math/rand"
	"testing"
)
var arr []int
func init(){
	for i:=0;i<100;i++{
		arr=append(arr,rand.Int()%100)
	}
}
func TestBubleSort(t *testing.T) {
	fmt.Println(arr)
	BubleSort(arr)
	fmt.Println(arr)

}
func TestSelectSort(t *testing.T){
	fmt.Println(arr)
	SelectSort(arr)
	fmt.Println(arr)
}
func TestInsertSort(t *testing.T){
	fmt.Println(arr)
	InsertSort(arr)
	fmt.Println(arr)
}
