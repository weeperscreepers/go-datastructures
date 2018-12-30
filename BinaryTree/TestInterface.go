package BinaryTree;

/*
  Just a dumb interface for examples/proof of concept
*/

type WhateverInterface interface {
  GetIt() int
}

type intWrapper struct {
  interior int
}

func (w *intWrapper) GetIt() int {
  return (*w).interior;
}

func NewWhatever(i int) WhateverInterface {
  return &intWrapper{ i };
}
