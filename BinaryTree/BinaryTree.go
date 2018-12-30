package BinaryTree;

import(
  // "fmt"
)

/*
  Things that go in a binary tree must, at a minimum, be "comparable" to each other.
  We ask to provide ToInteger so that we can easily have a nonhomogeneous BinaryTree
  ( ie we can't expect everyone implementing this interface to check for every type )
*/
type Comparable interface {
  ToInteger() int
  Compare( Comparable ) int
}

type BinaryTree interface {
  Insert(Comparable)
  PreOrder() []Comparable
  InOrder() []Comparable
  PostOrder() []Comparable
}

type BinaryNode struct {
  leftChild *BinaryNode
  rightChild  *BinaryNode
  data Comparable
}

func defaultEmpty(b *BinaryNode) []Comparable{
  if b== nil || b.data ==nil {
      return []Comparable{};
  }
  return []Comparable{ b.data }
}

func (b *BinaryNode) Insert(c Comparable) {
  if b.data == nil {
    b.data = c;
    return;
  }
  comparison := c.Compare(b.data);
  if comparison < 0 {
    if b.leftChild == nil {
      b.leftChild = &BinaryNode{ nil, nil, nil };
    }
    b.leftChild.Insert(c);
  }else{
    if b.rightChild == nil {
      b.rightChild = &BinaryNode{ nil, nil, nil };
    }
    b.rightChild.Insert(c);
  }
}

/*
  Pre-order and post-order are depth-first traversals
  so they can be done fairly easily with recursion.
*/

func (b *BinaryNode) PreOrder() []Comparable {
  ls := []Comparable{};
  if b == nil || b.data == nil { return ls; }
  ls = append(ls, b.leftChild.PreOrder() ... )
  ls = append(ls, b.data)
  ls = append(ls, b.rightChild.PreOrder() ... )
  return ls;
}

func (b *BinaryNode) PostOrder() []Comparable {
  ls := []Comparable{};
  if b == nil || b.data == nil { return ls; }
  ls = append(ls, b.leftChild.PostOrder() ... )
  ls = append(ls, b.rightChild.PostOrder() ... )
  ls = append(ls, b.data)
  return ls;
}

/*
  In-order is a breadth-first traversal and so is more complex
*/
func (b *BinaryNode) InOrder() []Comparable {

  selfLs := defaultEmpty(b);
  if len(selfLs) <= 0 { return selfLs };

  var helper func (b *BinaryNode) []Comparable;
  helper = func (b *BinaryNode) []Comparable {
    ls := []Comparable{};
    if b == nil || b.data == nil { return ls; }
    if b.leftChild != nil && b.leftChild.data != nil {
      ls = append( ls, b.leftChild.data )
    }
    if b.rightChild != nil && b.rightChild.data != nil {
      ls = append( ls, b.rightChild.data )
    }
    ls = append(ls, helper(b.leftChild) ... );
    ls = append(ls, helper(b.rightChild) ... );
    return ls;
  }
  return append(selfLs, helper(b) ... );

}

func New () BinaryTree {
  return &BinaryNode{ nil, nil, nil };
}
