package commands

import "testing"

func TestLinuxPackageNameApt(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"go", "golang-go"},
		{"git", "git"},
		{"python", "python3"},
	}

	for _, c := range cases {
		if got := linuxPackageName("apt-get", c.name); got != c.want {
			t.Fatalf("linuxPackageName(apt-get, %q) = %q, want %q", c.name, got, c.want)
		}
	}
}

func TestLinuxPackageNameOther(t *testing.T) {
	cases := []struct {
		name string
		want string
	}{
		{"go", "go"},
		{"git", "git"},
		{"python", "python3"},
	}

	for _, c := range cases {
		if got := linuxPackageName("dnf", c.name); got != c.want {
			t.Fatalf("linuxPackageName(dnf, %q) = %q, want %q", c.name, got, c.want)
		}
	}
}
