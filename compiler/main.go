package main

// read all lines of file
// dictionary [(identifier , (type , value))]
//
//line by line, parse and run code by either modifying , creating or reading the dictionary

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	identifiers map[string]variable
	regexes     map[string]*regexp.Regexp
)

type variable struct {
	gtype string
	value string
}

func main() {
	identifiers = make(map[string]variable)
	regexes = make(map[string]*regexp.Regexp)
	regexCompiling()
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
	identifiers[tokens[1]] = variable{gtype: tokens[0], value: evaluateExpression(tokens[3:])}
}

func evaluateExpression(exp []string) string {
	if len(exp) == 1 {
		uni := exp[0]
		if identifiers[uni] == (variable{}) {
			return uni
		}
		return identifiers[uni].value
	}
	str := ""

	leftHand := evaluateExpression([]string{exp[0]})
	rightHand := evaluateExpression(exp[2:])
	// tuple leftHand rightHand
	tuple := func(leftHand, rightHand string) string {
		return leftHand + " " + rightHand
	}

	switch tuple(evaluateType(leftHand), evaluateType(rightHand)) {
	case "bool bool":
		evaluateBoolOperator(leftHand, rightHand)
	case "char char":
		evaluateStringOperator(leftHand, rightHand)
	case "int int":
		evaluateIntOperator(leftHand, rightHand)
	case "float float":
		evaluateFloatOperator(leftHand, rightHand)
	case "string string":
		evaluateFloatOperator(leftHand, rightHand)

	}

	str += operator(leftHand, rightHand)

	return str
}

func evaluateBoolOperator(op string) func(string, string) string {
	switch op {
	case "&&":
		return func(leftHand, rightHand string) string { return leftHand && rightHand }
	case "||":
		return func(leftHand, rightHand string) string { return leftHand || rightHand }
	case "==":
		return func(leftHand, rightHand string) string { return leftHand == rightHand }
	case "!=":
		return func(leftHand, rightHand string) string { return leftHand != rightHand }
	case "^":
		return func(leftHand, rightHand string) string { return (leftHand && !rightHand) || (!leftHand && rightHand) }
	}
	panic("Bool operator evaluation failed")
}

func evaluateStringOperator(op string) func(string, string) string {
	switch op {
	case "+":
		return func(leftHand, rightHand string) string { return leftHand + rightHand }
	}
	panic("String operator evaluation failed")
}

func evaluateIntOperator(op string) func(int, int) string {
	switch op {
	case "+":
		return func(leftHand, rightHand int) string { return itoa(leftHand + rightHand) }
	case "-":
		return func(leftHand, rightHand int) string { return itoa(leftHand - rightHand) }
	case "*":
		return func(leftHand, rightHand int) string { return itoa(leftHand * rightHand) }
	case "/":
		return func(leftHand, rightHand int) string { return itoa(leftHand / rightHand) }
	case "%":
		return func(leftHand, rightHand int) string { return itoa(leftHand % rightHand) }
	}
	panic("Int operator evaluation failed")
}

func evaluateFloatOperator(op string) func(float64, float64) string {
	switch op {
	case "+":
		return func(leftHand, rightHand float64) string { return strconv.FormatFloat(leftHand+rightHand, 'f', 6, 64) }
	case "-":
		return func(leftHand, rightHand float64) string { return strconv.FormatFloat(leftHand-rightHand, 'f', 6, 64) }
	case "*":
		return func(leftHand, rightHand float64) string { return strconv.FormatFloat(leftHand*rightHand, 'f', 6, 64) }
	case "/":
		return func(leftHand, rightHand float64) string { return strconv.FormatFloat(leftHand/rightHand, 'f', 6, 64) }
	case "%":
		return func(leftHand, rightHand float64) string {
			return strconv.FormatFloat(leftHand/100*rightHand, 'f', 6, 64)
		}
	}
	panic("Float operator evaluation failed")
}

func evaluateType(str string) string {
	if identifiers[str] == (variable{}) {
		for key, value := range regexes {
			if value.MatchString(str) {
				return key
			}
		}
		panic("No valid type found")
	}
	return identifiers[str].gtype
}

func regexCompiling() {
	regexes["int"] = regexp.MustCompile("[0-9]+")
	// char
	regexes["char"] = regexp.MustCompile("\"[a-zA-Z0-9]\"")
	regexes["str"] = regexp.MustCompile("\"[a-zA-Z0-9]*\"")
	regexes["bool"] = regexp.MustCompile("true|false")
	// float
	regexes["float"] = regexp.MustCompile("[0-9]+.[0-9]+")
}

func atoi(str string) int {
	retval, _ := strconv.Atoi(str)
	return retval
}

func itoa(str int) string {
	return strconv.Itoa(str)
}
