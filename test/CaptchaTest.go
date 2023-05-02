package main

import (
	"errors"
	"fmt"
	"github.com/dchest/captcha"
	"time"
)

const (
	// Default number of digits in captcha solution.
	DefaultLen = 4
	// The number of captchas created that triggers garbage collection used
	// by default store.
	CollectNum = 100
	// Expiration time of captchas used by default store.
	Expiration = 10 * time.Minute
)

const (
	// Standard width and height of a captcha image.
	StdWidth  = 240
	StdHeight = 80
)

var (
	ErrNotFound = errors.New("captcha: id not found")
)

func main() {
	id := captcha.New()

	fmt.Println(id)
}
