package main;

import (
  "fmt"
  "generic/BinaryTree"
)

/*
  Suppose we only have access to the interface.
  If we had access to the underlying struct we could implement this directly
  instead of wrapping.
*/
type ComparableWhatever struct {
  data BinaryTree.WhateverInterface
}

func (w ComparableWhatever) ToInteger() int {
  return w.data.GetIt();
}

func (w ComparableWhatever) Compare(c BinaryTree.Comparable) int {
  return w.ToInteger() - c.ToInteger();
}

func IntoWhatever(w []BinaryTree.Comparable) []ComparableWhatever{
  var ls []ComparableWhatever;
  for _,v := range w {
    ls = append(ls, v.(ComparableWhatever))
  }
  return ls
}

func main() {
  fmt.Println("Testing...")
  one := ComparableWhatever{ BinaryTree.NewWhatever(1) };
  three := ComparableWhatever{ BinaryTree.NewWhatever(3) };
  five := ComparableWhatever{ BinaryTree.NewWhatever(5) };
  seven := ComparableWhatever{ BinaryTree.NewWhatever(7) };
  eleven := ComparableWhatever{ BinaryTree.NewWhatever(11) };
  fmt.Println("Finished creations")
  tree := BinaryTree.New();
  tree.Insert(seven);
  tree.Insert(three);
  tree.Insert(one);
  tree.Insert(five);
  tree.Insert(eleven);
  fmt.Println("Finished insertions")
  inOrder      := IntoWhatever(tree.InOrder())
  preOrder    := IntoWhatever(tree.PreOrder())
  postOrder  := IntoWhatever(tree.PostOrder())
  fmt.Println("Pre order:")
  for _,v := range preOrder {
    fmt.Println(v.data.GetIt())
  }
  fmt.Println("Post order:")
  for _,v := range postOrder {
    fmt.Println(v.data.GetIt())
  }
  fmt.Println("In order:")
  for _,v := range inOrder {
    fmt.Println(v.data.GetIt())
  }
}
