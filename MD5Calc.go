package main

import (
	"crypto/md5"
	"encoding/hex"
)

type Md5Calc struct {
	Calculator
}

func NewMd5Calc() *Md5Calc {
	c := new(Md5Calc)
	c.SameList = make(map[string][]string)
	c.results = make(map[string][]string)
	return c
}

func (this *Md5Calc) Calc(path string, bytes []byte) string {
	md5Hash := md5.Sum(bytes)
	c := hex.EncodeToString(md5Hash[:])
	this.hasSame(path, c)
	return c
}

func (this *Md5Calc) GetSameList() map[string][]string {
	return this.SameList
}
