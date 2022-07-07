package args

import "testing"

func TestValidateArgsPositive(t *testing.T) {
	args := []string{
		"arg1",
		"arg2",
	}

	filename, err := validateArgs(args)

	if err != nil {
		t.Errorf("Expected %v found %s", nil, err.Error())
	}

	if filename != args[1] {
		t.Errorf("Expected %s found %s", args[1], filename)
	}
}

func TestValidateArgsNegative(t *testing.T) {
	args := []string{
		"arg1",
	}

	_, err := validateArgs(args)

	if err == nil {
		t.Errorf("Expected error found %s", err)
	}
}
