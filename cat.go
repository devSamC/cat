package cat

import "github.com/devSamC/kitten"

func Meow() func() string {
	meow := kitten.Talk
	return meow
}

func Meows() func() string {
	return kitten.Talks
}
