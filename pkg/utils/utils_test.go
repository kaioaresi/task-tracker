package utils

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_Help(t *testing.T) {
	oldStdout := os.Stdout
	defer func() { os.Stdout = oldStdout }()

	r, w, _ := os.Pipe()

	os.Stdout = w

	Help()

	w.Close()

	out, _ := io.ReadAll(r)

	output := string(out)

	expectedMsg := []string{
		"add: add new task",
		"list: list all tasks",
		"update: update a task description",
		"delete: delete a task",
		"mark-in-progress: update task status to 'in-progress'",
		"mark-done: update task status to 'done'",
	}

	for _, msg := range expectedMsg {
		if !strings.Contains(output, msg) {
			t.Errorf("expected output %s but got %v", msg, output)
		}
	}
}

func Test_CheckInput(t *testing.T) {

	// para evitar que os demais tests quebrem, vamos salvar o valor original, para que quando finalizar esse bloco de tests,
	//  os demais test não sejam impactados por essa alteração.
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	testCases := []struct {
		name      string
		inputTest []string
		expectErr bool
	}{
		{
			name:      "3 inputs",
			inputTest: []string{"cli", "add", "description qualquer"},
			expectErr: false,
		},
		{
			name:      "empty inputs",
			inputTest: []string{"cli"},
			expectErr: true,
		},
		{
			name:      "2 inputs",
			inputTest: []string{"cli", "add"},
			expectErr: true,
		},
		{
			name:      "empty description",
			inputTest: []string{"cli", "add", ""},
			expectErr: true,
		},
		{
			name:      "more than 3 inputs",
			inputTest: []string{"cli", "add", "description", "mora than 3"},
			expectErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			os.Args = tc.inputTest

			result := CheckInput(tc.inputTest)

			if tc.expectErr {
				if result == nil {
					t.Errorf("%s: expected a error but got a %v", tc.name, result)
				}
			}

			if !tc.expectErr {
				if result != nil {
					t.Errorf("%s: not expected a error but got %v", tc.name, result)
				}
			}
		})
	}
}

func Test_InputToInt(t *testing.T) {
	testCases := []struct {
		name           string
		input          string
		expectedResult int
		expectErr      bool
	}{
		{
			name:           "Valide input",
			input:          "1",
			expectedResult: 1,
			expectErr:      false,
		},
		{
			name:           "Invalide input",
			input:          "abc",
			expectedResult: 0,
			expectErr:      true,
		},
		{
			name:           "Empty input",
			input:          "",
			expectedResult: 0,
			expectErr:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			result, err := InputToInt(tc.input)

			if tc.expectErr {
				if err == nil {
					t.Errorf("%s: expected error but got %v with error %v", tc.name, result, err)
				}
			}

			if !tc.expectErr {
				if err != nil {
					t.Errorf("%s: error not expected but got %v with error %v", tc.name, result, err)
				}
			}

		})
	}
}

func Test_GetTaskID(t *testing.T) {
	testCases := []struct {
		name           string
		input          []string
		expectedResult int
		expectErr      bool
	}{
		{
			name:           "Valide input",
			input:          []string{"cli", "add", "1"},
			expectedResult: 1,
			expectErr:      false,
		},
		{
			name:           "Invalide input",
			input:          []string{"cli", "add", "aaa"},
			expectedResult: 0,
			expectErr:      true,
		},
		{
			name:           "Invalide input",
			input:          []string{"cli", "add"},
			expectedResult: 0,
			expectErr:      true,
		},
		{
			name:           "Empty input",
			input:          []string{"cli", ""},
			expectedResult: 0,
			expectErr:      true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			os.Args = tc.input
			result, err := GetTaskID(tc.input)

			if tc.expectErr {
				if err == nil {
					t.Errorf("%s: expected error but got %v, result %v", tc.name, err, result)
				}
			}

			if !tc.expectErr {
				if err != nil {
					t.Errorf("%s: not expected erro but got %v, result %v", tc.name, err, result)
				}
			}
		})
	}
}
