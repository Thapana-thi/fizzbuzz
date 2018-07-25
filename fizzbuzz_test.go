package fizzbuzz

import "testing"

func TestFizzBuzzShouldReturnOne(t *testing.T){
	 fb := fizzbuzz("1")

	 if fb !="1"{ 
		 println("value shoud be 1")
		 t.Error("value shoud be 1 but get:",fb)
  
	 }
	 
}

func TestFizzBuzzShouldReturnTwo(t *testing.T){
	fb := fizzbuzz("2")

	if fb !="2"{ 
		println("value shoud be 2")
		t.Error("value shoud be 2 but get:",fb)
 
	}
	
}

func TestFizzBuzzShouldReturnFizz(t *testing.T){
	fb := fizzbuzz("3")

	if fb !="Fizz"{ 
		println("value shoud be Fizz")
		t.Error("value shoud be Fizz but get:",fb)
 
	}
	
}