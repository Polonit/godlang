package main

// read all lines of file
// dictionary [(identifier , (type , value))]
//
//line by line, parse and run code by either modifying , creating or reading the dictionary

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var (
	identifiers map[string]string
)

func main() {
	identifiers = make(map[string]string)
	filepath := "src.god"
	file, err := os.Open(filepath)
	if err != nil {
		print(err)
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		processLine(scanner.Text())
	}
}

func processLine(line string) {
	tokens := strings.Split(line, " ")
	if len(tokens) < 3 {
		return
	}
	if tokens[0] == "print" { //strings.Contains(tokens[0], "print") {
		// strings.Index(tokens, "(")
		println(evaluateExpression(tokens[2 : len(tokens)-1]))
	} else {
		evaluateVarAssignment(tokens)
	}
}

func evaluateVarAssignment(tokens []string) {
	// if int str etc...
	identifiers[tokens[1]] = evaluateExpression(tokens[3:])
}

func evaluateExpression(exp []string) string {
	if len(exp) == 1 {
		uni := exp[0]
		if identifiers[uni] == "" {
			return uni
		}
		return identifiers[uni]
	}
	str := ""
	leftHand := atoi(evaluateExpression([]string{exp[0]}))
	rightHand := atoi(evaluateExpression(exp[2:]))
	switch exp[1] {
	case "+":
		str += itoa(leftHand + rightHand)
	case "-":
		str += itoa(leftHand - rightHand)
	case "*":
		str += itoa(leftHand * rightHand)
	case "/":
		str += itoa(leftHand / rightHand)
	case "%":
		str += itoa(leftHand % rightHand)
	}
	return str
}

func atoi(str string) int {
	retval, _ := strconv.Atoi(str)
	return retval
}

func itoa(str int) string {
	return strconv.Itoa(str)
}
