package tic_tac_toe_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestTicTacToe(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "TicTacToe Suite")
}
