package main
import (
	"fmt"
)
type node struct {
	val int
	next *node
}
type LinkedList struct {
	head *node
	len int
}
func (l *LinkedList) addLast(val int)  {
	newNode := node{val: val}
	if l.head == nil{
		l.head = &newNode
	}else {
		current := l.head
		for current.next != nil{
			current = current.next
		}
		current.next = &newNode
	}
}
//   first -> 2
//   head = &first
//   temp = &first
//   second= 4
//   head = &second
//   head.next = temp

func (l *LinkedList) addFirst(val int)  {
	newNode := &node{val: val}
	if l.head == nil{
		l.head = newNode
	}else{
		temp := l.head
		l.head = newNode
		l.head.next = temp
	}
}
/*3 4 5     ->     3 4 7 5
insert(2, 7)
count = 0
head = 3
current = head = 3
current = current.next = 4
count++ = 1
*/
// 3 4 5, 2
func (l *LinkedList) insert(position, val int)  {
	if l.head == nil{
		panic("empty list")
	}else{
		count := 0
		newNode := &node{val: val}
		current := l.head
		for current != nil{
			if count == position-1{
				temp := current.next
				current.next = newNode
				newNode.next = temp
				return
			}
			count++
			current=current.next
		}
		panic(fmt.Sprintf("position %d out of bound", position))
	}
}
func (l *LinkedList) print()  {
	if l.head == nil{
		fmt.Println("[]")
	}else{
		current:= l.head
		for current != nil{
			fmt.Print(current.val, " ")
			current = current.next
		}
	}
}
func main() {
	var list1 = &LinkedList{}
	var list2 = &LinkedList{}
	list2.addFirst(3)
	list1.addFirst(2)
	list1.addLast(5)
	list1.addFirst(4)
	list1.insert(3,6)
	list1.print()
	list2.print()
}
/*l.head -> &node(3)
l.add(5)
current = 3
current.next = nil
current.next = &node(5)
l.add(7)
head =3
current = 3
current = current.next = 5
current.next = nil
current.next = 7
*/
/*func (l *LinkedList) addFirst(val int)  {
    newNode := &node{val: 2}
    node2 := *newNode
    if l.head == nil{
        l.head = newNode
        l.head.val = 3
        fmt.Println(l.head, "head")
        fmt.Println(newNode, "newNode")
        fmt.Println(node2, "Node2")
    }else{
        return
    }
}*/