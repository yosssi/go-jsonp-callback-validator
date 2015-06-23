package validator

import "testing"

var cases = []struct {
	in   string
	want bool
}{
	{`foo`, true},
	{`foo123`, true},
	{`fos.Router.data`, true},
	{`$.callback`, true},
	{`_.callback`, true},
	{`hello`, true},
	{`foo23`, true},
	{`$210`, true},
	{`_bar`, true},
	{`some_var`, true},
	{`$`, true},
	{`somevar`, true},
	{`$.ajaxHandler`, true},
	{`array_of_functions[42]`, true},
	{`array_of_functions[42][1]`, true},
	{`$.ajaxHandler[42][1].foo`, true},
	{`array_of_functions["key"]`, true},
	{`_function`, true},
	{`petersCallback1412331422[12]`, true},
	{`(function xss(x){evil()})`, false},
	{``, false},
	{`alert()`, false},
	{`test()`, false},
	{`a-b`, false},
	{`23foo`, false},
	{`function`, false},
	{` somevar`, false},
	{`$.23`, false},
	{`array_of_functions[42]foo[1]`, false},
	{`array_of_functions[]`, false},
	{`myFunction[123].false`, false},
	{`myFunction .tester`, false},
	{`:myFunction`, false},
	{`array_of_functions["k"ey"]`, false},
	{`array_of_functions["k\"ey"]`, true},
	{`array_of_functions["k""y"]`, false},
	{`array_of_functions["""y"]`, false},
	{`array_of_functions[""""]`, false},
	{`array_of_functions["\""]`, true},
	{`array_of_functions["k\"e\""]`, true},
	{`array_of_functions["k\'ey"]`, true},
	{`array_of_functions['k'ey']`, false},
	{`array_of_functions['k\"ey']`, true},
	{`array_of_functions['k\'ey']`, true},
	{`array_of_functions['\'key']`, true},
	{`array_of_functions['key\'']`, true},
	{`array_of_functions['k'ey'']`, false},
	{`array_of_functions[''']`, false},
	{`array_of_functions['\'']`, true},
}

func TestValidate(t *testing.T) {
	for _, c := range cases {
		if got := Validate(c.in); got != c.want {
			t.Errorf("Validate(%q) => %t, want %t", c.in, got, c.want)
		}
	}
}
