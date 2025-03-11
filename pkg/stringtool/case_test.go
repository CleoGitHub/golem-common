package stringtool

import "testing"

func TestSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "foo"},
		{"two words", "fooBar", "foo_bar"},
		{"three words", "fooBarBaz", "foo_bar_baz"},
		{"four words", "fooBarBazQux", "foo_bar_baz_qux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeCase(tt.in); got != tt.want {
				t.Errorf("SnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCamelCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "foo"},
		{"two words", "foo_bar", "fooBar"},
		{"three words", "foo_bar_baz", "fooBarBaz"},
		{"three words", "foo bar baz", "fooBarBaz"},
		{"four words", "foo_bar_baz_qux", "fooBarBazQux"},
		{"four words", "foo_bar bAz_qux", "fooBarBAzQux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelCase(tt.in); got != tt.want {
				t.Errorf("CamelCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPascalCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "Foo"},
		{"two words", "foo_bar", "FooBar"},
		{"three words", "foo_bar_baz", "FooBarBaz"},
		{"three words", "foo bar baz", "FooBarBaz"},
		{"four words", "foo_bar_baz_qux", "FooBarBazQux"},
		{"four words", "foo_bar bAz_qux", "FooBarBAzQux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PascalCase(tt.in); got != tt.want {
				t.Errorf("PascalCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDashCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "foo"},
		{"two words", "foo_bar", "foo-bar"},
		{"three words", "foo_bar_baz", "foo-bar-baz"},
		{"three words", "foo bar baz", "foo-bar-baz"},
		{"four words", "foo_bar_baz_qux", "foo-bar-baz-qux"},
		{"four words", "foo_bar bAz_qux", "foo-bar-b-az-qux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DashCase(tt.in); got != tt.want {
				t.Errorf("DashCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperSnakeCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "FOO"},
		{"two words", "foo_bar", "FOO_BAR"},
		{"three words", "foo_bar_baz", "FOO_BAR_BAZ"},
		{"three words", "foo bar baz", "FOO_BAR_BAZ"},
		{"four words", "foo_bar_baz_qux", "FOO_BAR_BAZ_QUX"},
		{"four words", "foo_bar bAz_qux", "FOO_BAR_B_AZ_QUX"},
		{"four words", "foBbar bAz_qux", "FO_BBAR_B_AZ_QUX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperSnakeCase(tt.in); got != tt.want {
				t.Errorf("UpperSnakeCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUpperDashCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "FOO"},
		{"two words", "foo_bar", "FOO-BAR"},
		{"three words", "foo_bar_baz", "FOO-BAR-BAZ"},
		{"three words", "foo bar baz", "FOO-BAR-BAZ"},
		{"four words", "foo_bar_baz_qux", "FOO-BAR-BAZ-QUX"},
		{"four words", "foBbar bAz_qux", "FO-BBAR-B-AZ-QUX"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UpperDashCase(tt.in); got != tt.want {
				t.Errorf("UpperDashCase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSpaceCase(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"one word", "foo", "foo"},
		{"two words", "foo_bar", "foo bar"},
		{"two words", "foo-bar", "foo bar"},
		{"two words", "fooBar", "foo bar"},
		{"three words", "foo_bar_baz", "foo bar baz"},
		{"three words", "foo bar baz", "foo bar baz"},
		{"four words", "foo_bar_baz_qux", "foo bar baz qux"},
		{"four words", "foo_bar bAz_qux", "foo bar b az qux"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SpaceCase(tt.in); got != tt.want {
				t.Errorf("SpaceCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
