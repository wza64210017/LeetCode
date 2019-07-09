package main

import "fmt"

type MyLinkedList struct {
	Size int
	Head *ListNode
	Tail *ListNode
}

type ListNode struct {
	Val  int
	Next *ListNode
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	tail := &ListNode{}
	head := &ListNode{Next: tail}

	list := MyLinkedList{
		Head: head,
		Tail: tail,
	}

	return list
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || this.Size <= index {
		return -1
	}

	q := this.Head.Next
	for index > 0 {
		q = q.Next
		index--
	}

	return q.Val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	node := &ListNode{
		Val:  val,
		Next: this.Head.Next,
	}

	this.Head.Next = node
	this.Size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	this.Tail.Val = val
	node := &ListNode{}
	this.Tail.Next = node
	this.Tail = node
	this.Size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.Size {
		return
	}

	if index == 0 {
		this.AddAtHead(val)
		return
	}

	if index == this.Size {
		this.AddAtTail(val)
		return
	}

	p := this.Head
	for index > 0 {
		p = p.Next
		index--
	}

	node := &ListNode{
		Val:  val,
		Next: p.Next,
	}

	p.Next = node
	this.Size++
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.Size || index < 0 {
		return
	}

	p := this.Head
	for index > 0 {
		p = p.Next
		index--
	}

	p.Next = p.Next.Next
	this.Size--
}

func main() {
	linkedList := Constructor()
	//linkedList.AddAtHead(1)
	//linkedList.AddAtTail(3)
	//linkedList.AddAtIndex(1,2)    //链表变为1-> 2-> 3
	//linkedList.Get(1)            //返回2
	//linkedList.DeleteAtIndex(1)  //现在链表是1-> 3
	//ppp(linkedList)
	//linkedList.Get(1)            //返回3

	linkedList.AddAtIndex(-1, 0)
	ppp(linkedList)
	fmt.Println(linkedList.Get(0))
	linkedList.DeleteAtIndex(-1)
}

func ppp(list MyLinkedList) {
	fmt.Println("size", list.Size)
	for list.Head != nil {
		fmt.Print(fmt.Sprintf("%d%s", list.Head.Val, "--------"))
		list.Head = list.Head.Next
	}
	fmt.Println()

}
