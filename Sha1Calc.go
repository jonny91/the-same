package main

import (
	"crypto/sha1"
	"encoding/hex"
)

type Sha1Calc struct {
	Calculator
}

func NewSha1Calc() *Sha1Calc {
	c := new(Sha1Calc)
	c.SameList = make(map[string][]string)
	c.results = make(map[string][]string)
	return c
}

func (this *Sha1Calc) Calc(path string, bytes []byte) string {
	r := sha1.Sum(bytes)
	c := hex.EncodeToString(r[:])
	this.hasSame(path, c)
	return c
}

func (this *Sha1Calc) GetSameList() map[string][]string {
	return this.SameList
}
