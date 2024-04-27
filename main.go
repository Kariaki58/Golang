package main
import (
    "fmt"
)

func main()  {
    numbers := []int {1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

    fmt.Printf("numbers = %v\n", numbers)
    fmt.Printf("length = %d\n", len(numbers))
    fmt.Printf("capacity = %d\n", cap(numbers))

    neededNumbers := numbers[:len(numbers) - 10]
    fmt.Println(neededNumbers, "needed number")
    numbersCopy := make([]int, len(neededNumbers))
    fmt.Printf("number ---- %v", numbersCopy)
    copy(numbersCopy, neededNumbers)

    fmt.Printf("numbersCopy = %v\n", numbersCopy)
    fmt.Printf("length = %d\n", len(numbersCopy))
    fmt.Printf("capacity = %d\n", cap(numbersCopy))
}
