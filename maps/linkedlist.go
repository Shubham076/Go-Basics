package main
import "fmt"

type Node struct {
	data int;
	next *Node;
}

type List struct {
	size int;
	head *Node
	tail *Node
}

func (l *List) add(data int) {
	n := Node{data, nil}
	if l.head == nil && l.tail == nil {
		l.head = &n;
		l.tail = &n; 
	} else {
		l.tail.next = &n;
		l.tail = &n;
	}
	l.size += 1;
}

func (l List) print() {
	for temp := l.head; temp != nil; temp = temp.next {
		fmt.Println(temp.data);
	}
}

func main(){
	ll := List{}
	arr := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(arr); i++ {
		ll.add(arr[i]);
	}
	ll.print()
	fmt.Println(ll); 
}