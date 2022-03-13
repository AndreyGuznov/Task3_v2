package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var command = map[int]string{}

func main() {
	var result string
	handleComand()
	for i := 2; i < len(command); i++ {
		result += fmt.Sprintf("\n%d\n", findRanges(command, i))
	}
	err := ioutil.WriteFile(string(command[0]), []byte(result), 0777)
	if err != nil {
		fmt.Printf("Error of creating file")
	}
}

func handleComand() {
	fmt.Println("Greeting You. This program is for searching simple numbers. Example of correct command: find_primes--file testfile.txt--timeout 10--range 1:10--range 400:500..(more ranges).")
	fmt.Print("find_primes--file ")
	inp := bufio.NewScanner(os.Stdin)
	inp.Scan()
	comand := inp.Text()
	if !(strings.Contains(comand, ".txt")) {
		fmt.Println("Incorrect command. Were file.txt?")
		return
	} else if strings.Count(comand, "--") < 2 {
		fmt.Println("Invalid command!")
		return
	}
	for i := 0; i < strings.Count(comand, "--")+1; i++ {
		command[i] = (strings.Split(comand, "--"))[i]
	}
	err := ioutil.WriteFile(command[0], []byte(comand), 0777)
	if err != nil {
		fmt.Println("Error of writing UserCommand to file")
		return
	}
	// args := os.Args
	// for i := 1; i < len(args); i++ {
	// 	_, err := os.ReadFile(args[i])
	// 	if err != nil {
	// 		fmt.Println("Error", err.Error())
	// 		return
	// 	}
	// }
}
func findRanges(command map[int]string, i int) []int {
	str := command[i]
	var integ [2]int
	r, err := regexp.Compile(`[0-9]+`)
	if err != nil {
		fmt.Println("Uncorrect range")
		return nil
	}
	matches := r.FindAllString(string(str), -1)
	for i, m := range matches {
		num, err := strconv.Atoi(m)
		if err != nil {
			fmt.Println(err)
		}
		integ[i] = num
	}
	slice := make([]int, 0)
	fmt.Println(integ)
	if integ[0] > integ[1] {
		fmt.Println("Uncorrect range")
	}

	arr := make([]bool, integ[1])
	for i := 2; i <= int(math.Sqrt(float64(integ[1])+1)); i++ {
		if arr[i] == false {
			for j := i * i; j < integ[1]; j += i {
				arr[j] = true
			}
		}
	}
	var primes []int

	for i, isComposite := range arr {
		if i > 1 && !isComposite {
			primes = append(primes, i)
		}
	}
	for _, val := range primes {
		if val >= integ[0] {
			slice = append(slice, val)
		}
	}

	return slice
}

//testfile.txt--timeout 10--range 1:10--range 400:500
