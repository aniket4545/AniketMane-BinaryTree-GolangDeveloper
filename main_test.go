package main

import "testing"

var bTree *Tree

func Test_Insert(t *testing.T) {
	bTree = new(Tree)
	bTree.InsertNode(2)
	bTree.InsertNode(4)
	bTree.InsertNode(6)
	bTree.InsertNode(3)
	bTree.InsertNode(1)
	bTree.InsertNode(7)
	bTree.InsertNode(9)
	bTree.InsertNode(5)
	bTree.InsertNode(-1)
	bTree.InsertNode(-2)

	bTree.Print()
}

func Test_InOrder(t *testing.T) {
	var (
		expectedResult = []int{-2, -1, 1, 2, 3, 4, 5, 6, 7, 9}
		obtainedResult = bTree.InOrder()
	)
	if !checkIfSameArray(expectedResult, obtainedResult) {
		t.Errorf("InOrder traversal is wrong: expected: %v\n got:%v\n", expectedResult, obtainedResult)
	}
}

func Test_PreOrder(t *testing.T) {
	var (
		expectedResult = []int{2, 1, -1, -2, 4, 3, 6, 5, 7, 9}
		obtainedResult = bTree.PreOrder()
	)
	if !checkIfSameArray(expectedResult, obtainedResult) {
		t.Errorf("PreOrder traversal is wrong: expected: %v\n got:%v\n", expectedResult, obtainedResult)
	}
}

func Test_PostOrder(t *testing.T) {
	var (
		expectedResult = []int{-2, -1, 1, 3, 5, 9, 7, 6, 4, 2}
		obtainedResult = bTree.PostOrder()
	)
	if !checkIfSameArray(expectedResult, obtainedResult) {
		t.Errorf("PostOrder traversal is wrong: expected: %v\n got:%v\n", expectedResult, obtainedResult)
	}
}

func checkIfSameArray(expectedResult, result []int) bool {

	// if both are nil then return true
	if expectedResult == nil && result == nil {
		return true
	}

	// if any one out of them is nil then return false
	if expectedResult == nil || result == nil {
		return false
	}

	// if lenght vaires then return false
	if len(expectedResult) != len(result) {
		return false
	}

	// if any element is not there at expected index then return false
	for i := 0; i < len(expectedResult); i++ {
		if expectedResult[i] != result[i] {
			return false
		}
	}
	return true
}
