package main

import (
	"bible-compiler/lexer"
	"bible-compiler/tokenizer"
	"strings"
	"testing"
)

func TestLexerTypes(t *testing.T) {
	t.Run("Test simple int", func(t *testing.T) {
		code := "int aa = 12"
		expectedTypes := []tokenizer.TokenType{"int", "id", "=", "number"}
		expectedResult := []string{"int", "aa", "=", "12"}

		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple char", func(t *testing.T) {
		code := "char a = 'a'"
		expectedResult := []string{"char", "a", "=", "a"}
		expectedTypes := []tokenizer.TokenType{"char", "id", "=", "char"}

		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple string", func(t *testing.T) {
		code := `string aa = "er12"`
		expectedTypes := []tokenizer.TokenType{"string", "id", "=", "string"}
		expectedResult := []string{"string", "aa", "=", `er12`}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple float", func(t *testing.T) {
		code := "float deozuhbfr = 12.1235"
		expectedTypes := []tokenizer.TokenType{"float", "id", "=", "number"}
		expectedResult := []string{"float", "deozuhbfr", "=", "12.1235"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple long", func(t *testing.T) {
		code := "long aa = 124546578756463322453"
		expectedTypes := []tokenizer.TokenType{"long", "id", "=", "number"}
		expectedResult := []string{"long", "aa", "=", "124546578756463322453"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple double", func(t *testing.T) {
		code := "double aa = 3.141592653589793238462643383279502884197169399375105820974944592307816406286"
		expectedTypes := []tokenizer.TokenType{"double", "id", "=", "number"}
		expectedResult := []string{"double", "aa", "=", "3.141592653589793238462643383279502884197169399375105820974944592307816406286"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple bool", func(t *testing.T) {
		code := "bool aa = true"
		expectedTypes := []tokenizer.TokenType{"bool", "id", "=", "bool"}
		expectedResult := []string{"bool", "aa", "=", "true"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test simple array", func(t *testing.T) {
		code := "int[] aa = [1,2,3,4,5]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"intArr", "id", "=", "intArr"}
		expectedResult := []string{"int", "[", "]", "aa", "=", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test multi dim array", func(t *testing.T) {
		code := "int[][] aa = [[1,2,3,4,5],[1,2,3,4,5]]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"intArr", "id", "=", "intArr"}
		expectedResult := []string{"int", "[", "]", "[", "]", "aa", "=", "[", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]", ",", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test array string", func(t *testing.T) {
		code := "string[] aa = [hello,world]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"stringArr", "id", "=", "stringArr"}
		expectedResult := []string{"string", "[", "]", "aa", "=", "[", "hello", ",", "world", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test mult dim array string", func(t *testing.T) {
		code := "string[][] aa = [[hello,world],[hello,world]]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"stringArr", "id", "=", "stringArr"}
		expectedResult := []string{"string", "[", "]", "[", "]", "aa", "=", "[", "[", "hello", ",", "world", "]", ",", "[", "hello", ",", "world", "]", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test array bool", func(t *testing.T) {
		code := "bool[] aa = [true,false]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"boolArr", "id", "=", "boolArr"}
		expectedResult := []string{"bool", "[", "]", "aa", "=", "[", "true", ",", "false", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test mult dim array  bool", func(t *testing.T) {
		code := "bool[][] aa = [[true,false],[false,true]]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"boolArr", "id", "=", "boolArr"}
		expectedResult := []string{"bool", "[", "]", "[", "]", "aa", "=", "[", "[", "true", ",", "false", "]", ",", "[", "false", ",", "true", "]", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test array float", func(t *testing.T) {
		code := "float[] aa = [1.0,2.0]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"floatArr", "id", "=", "floatArr"}
		expectedResult := []string{"float", "[", "]", "aa", "=", "[", "1.0", ",", "2.0", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("Test mult dim array  float", func(t *testing.T) {
		code := "float[][] aa = [[1.0,2.0],[2.0,1.0]]"
		// TODO Define expected types
		expectedTypes := []tokenizer.TokenType{"floatArr", "id", "=", "floatArr"}
		expectedResult := []string{"float", "[", "]", "[", "]", "aa", "=", "[", "[", "1.0", ",", "2.0", "]", ",", "[", "2.0", ",", "1.0", "]", "]"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
	t.Run("failing Case identifier", func(t *testing.T) {
		code := "a int = 1"
		expectedTypes := []tokenizer.TokenType{"id", "int", "=", "int"}
		expectedResult := []string{"ERROR"}
		BodyOfTest(code, expectedTypes, expectedResult, t)
	})
}

func TestLexerFuncs(t *testing.T) {
	t.Run("empty func", func(t *testing.T) {
		// code := "func test() {}"
		// // TODO Define expected types
		// expectedTypes := []tokenizer.TokenType{"func", "id", "(", ")", "{", "}"}
		// expectedResult := []string{"func", "test", "(", ")", "{", "}"}
		// BodyOfTest(code, expectedTypes, expectedResult, t)
	})
}

func BodyOfTest(code string, expectedTypes []tokenizer.TokenType, expectedResult []string, t *testing.T) {
	l := lexer.New(code)

	var result []string
	var resultTypes []tokenizer.TokenType

	for {
		token := l.NextToken()
		result = append(result, token.Literal)
		resultTypes = append(resultTypes, token.Type)
		if l.Ch == 0 {
			break
		}
	}
	if len(expectedResult) != len(result) || len(expectedTypes) != len(resultTypes) {
		t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
		// stop the test
		return
	}
	er := false
	for i, v := range expectedResult {
		if v != result[i] || expectedTypes[i] != resultTypes[i] {
			er = true
		}
	}

	if er {
		t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		var strTypes string
		for _, v := range resultTypes {
			strTypes += string(v)
		}
		var strExpectedTypes string
		for _, v := range expectedTypes {
			strExpectedTypes += string(v)
		}
		t.Errorf("With types %s, got %s", strExpectedTypes, strTypes)
	}
}
