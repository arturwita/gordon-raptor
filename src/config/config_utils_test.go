package config_test

import (
	"fmt"
	config "gordon-raptor/src/config"
	"os"
	"testing"
)

func TestGetEnvUtils(t *testing.T) {
	VARIABLE_NAME := "EXAMPLE_ENV_VAR"

	t.Run("GetStringEnv", func(t *testing.T) {
		t.Run("returns the env var if it's not empty", func(t *testing.T) {
			// given
			defaultValue := "hello"
			expected := "example"
			os.Setenv(VARIABLE_NAME, expected)

			// when
			result := config.GetStringEnv(VARIABLE_NAME, defaultValue)

			// then
			if result != expected {
				t.Errorf("Test failed. Expected: = %s; Actual: %s", expected, result)
			}
		})

		t.Run("returns the default value if env var is empty", func(t *testing.T) {
			// given
			expected := "example"
			os.Unsetenv(VARIABLE_NAME)

			// when
			result := config.GetStringEnv(VARIABLE_NAME, expected)

			// then
			if result != expected {
				t.Errorf("Test failed. Expected: = %s; Actual: %s", expected, result)
			}
		})
	})

	t.Run("GetIntEnv", func(t *testing.T) {
		t.Run("returns the int env var if it's not empty", func(t *testing.T) {
			// given
			defaultValue := 1
			expected := 2
			os.Setenv(VARIABLE_NAME, "2")

			// when
			result := config.GetIntEnv(VARIABLE_NAME, defaultValue)

			// then
			if result != expected {
				t.Errorf("Test failed. Expected: = %d; Actual: %d", expected, result)
			}
		})

		t.Run("returns the default value if env var is empty", func(t *testing.T) {
			// given
			expected := 2
			os.Unsetenv(VARIABLE_NAME)

			// when
			result := config.GetIntEnv(VARIABLE_NAME, expected)

			// then
			if result != expected {
				t.Errorf("Test failed. Expected: = %d; Actual: %d", expected, result)
			}
		})

		t.Run("returns the default value when tried to parse to an integer", func(t *testing.T) {
			// given
			testCases := []struct {
				name  string
				value any
			}{
				{"a float", 2.5},
				{"a boolean", true},
				{"a string", "hello"},
			}
			expected := 1

			for _, testCase := range testCases {
				t.Run(testCase.name, func(t *testing.T) {
					os.Setenv(VARIABLE_NAME, fmt.Sprintf("%v", testCase.value))

					// when
					result := config.GetIntEnv(VARIABLE_NAME, expected)

					// then
					if result != expected {
						t.Errorf("Test case '%s' failed. Expected: = %d; Actual: %d", testCase.name, expected, result)
					}
				})
			}
		})
	})
}
