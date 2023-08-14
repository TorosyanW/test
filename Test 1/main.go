package main

import "fmt"

func main() {
	age := GetAge()
	CheckAge(age, 22)

}

func GetAge() int {
	var age int

	fmt.Println("Сколько тебе лет?")

	if _, err :=fmt.Scanf("%d\n", &age); err != nil {
		fmt.Println(err)
		return age

	}
      return age
}

func CheckAge(age, LimitAge int) {
	if age >= LimitAge {
		printAge(age)
		return

	}else if LimitAge == 21 && (age == 10 || age == 12) {
		fmt.Println("Да заходи!")
		return
	}

	  PrintError()
}

  func PrintError(){
	fmt.Println("Уйди малой!")

  }

  func printAge(age int) {
	fmt.Println("Привет твой возраст - %d. \nПроходи!", age)

  }