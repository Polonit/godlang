package main

import (
	"bible-compiler/lexer"
	"strings"
	"testing"
)

func TestLexerTypes(t *testing.T) {
	t.Run("Test simple int", func(t *testing.T) {
		code := "int aa = 12"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"int", "aa", "=", "12"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple char", func(t *testing.T) {
		code := "char a = 'a'"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"char", "a", "=", "'a'"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}
		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple string", func(t *testing.T) {
		code := `string aa = "er12"`
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"string", "aa", "=", `"er12"`}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple float", func(t *testing.T) {
		code := "float deozuhbfr = 12.1235"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"float", "deozuhbfr", "=", "12.1235"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple long", func(t *testing.T) {
		code := "long aa = 124546578756463322453"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"long", "aa", "=", "124546578756463322453"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple double", func(t *testing.T) {
		code := "double aa = 3.141592653589793238462643383279502884197169399375105820974944592307816406286"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"double", "aa", "=", "3.141592653589793238462643383279502884197169399375105820974944592307816406286"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple bool", func(t *testing.T) {
		code := "bool aa = true"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"bool", "aa", "=", "true"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test simple array", func(t *testing.T) {
		code := "int[] aa = [1,2,3,4,5]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"int", "[", "]", "aa", "=", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test multi dim array", func(t *testing.T) {
		code := "int[][] aa = [[1,2,3,4,5],[1,2,3,4,5]]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"int", "[", "]", "[", "]", "aa", "=", "[", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]", ",", "[", "1", ",", "2", ",", "3", ",", "4", ",", "5", "]", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test array string", func(t *testing.T) {
		code := "string[] aa = [\"hello\",\"world\"]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"string", "[", "]", "aa", "=", "[", "\"hello\"", ",", "\"world\"", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test mult dim array string", func(t *testing.T) {
		code := "string[][] aa = [[\"hello\",\"world\"],[\"hello\",\"world\"]]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"string", "[", "]", "[", "]", "aa", "=", "[", "[", "\"hello\"", ",", "\"world\"", "]", ",", "[", "\"hello\"", ",", "\"world\"", "]", "]"}
		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test array bool", func(t *testing.T) {
		code := "bool[] aa = [true,false]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"bool", "[", "]", "aa", "=", "[", "true", ",", "false", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test mult dim array  bool", func(t *testing.T) {
		code := "bool[][] aa = [[true,false],[false,true]]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"bool", "[", "]", "[", "]", "aa", "=", "[", "[", "true", ",", "false", "]", ",", "[", "false", ",", "true", "]", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test array float", func(t *testing.T) {
		code := "float[] aa = [1.0,2.0]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"float", "[", "]", "aa", "=", "[", "1.0", ",", "2.0", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("Test mult dim array  float", func(t *testing.T) {
		code := "float[][] aa = [[1.0,2.0],[2.0,1.0]]"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"float", "[", "]", "[", "]", "aa", "=", "[", "[", "1.0", ",", "2.0", "]", ",", "[", "2.0", ",", "1.0", "]", "]"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
	t.Run("failing Case identifier", func(t *testing.T) {
		code := "a int = 1"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"ERROR"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 || token.Type == "ERROR" {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
}

func TestLexerFuncs(t *testing.T) {
	t.Run("empty func", func(t *testing.T) {
		code := "func test() {}"
		l := lexer.New(code)

		var result []string
		expectedResult := []string{"func", "test", "(", ")", "{", "}"}

		for {
			token := l.NextToken()
			result = append(result, token.Literal)
			if l.Ch == 0 {
				break
			}
		}

		if len(expectedResult) != len(result) {
			t.Errorf("Expected len %d, got %d", len(expectedResult), len(result))
			// stop the test
			return
		}
		er := false
		for i, v := range expectedResult {
			if v != result[i] {
				er = true
			}
		}
		if er {
			t.Errorf("Expected %s, got %s", strings.Join(expectedResult, " "), strings.Join(result, " "))
		}
	})
}
