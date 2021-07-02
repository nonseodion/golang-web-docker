// main_test.go

package main

import (
	"testing"

	"github.com/nonseodion/golang-web-docker/controllers"
)

func TestSum(t *testing.T) {
    if controllers.Sum(2, 5) != 7 {
        t.Fail()
    }
    if controllers.Sum(2, 100) != 102 {
        t.Fail()
    }
    if controllers.Sum(222, 100) != 322 {
        t.Fail()
    }
}

func TestProduct(t *testing.T) {
    if controllers.Product(2, 5) != 10 {
        t.Fail()
    }
    if controllers.Product(2, 100) != 200 {
        t.Fail()
    }
    if controllers.Product(222, 3) != 666 {
        t.Fail()
    }
}