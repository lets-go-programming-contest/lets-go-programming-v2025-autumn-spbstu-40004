package main
func main(){
  var firstNum, secondNum, result int
	var operator string
	_, err := fmt.Scanln(&firstNum)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}

	_, err = fmt.Scanln(&secondNum)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}

	_, err = fmt.Scanln(&operator)
	if err != nil {
		fmt.Println("Invalid operation")
		return
}
