class Queue {
    constructor (head = null) {
        this.head = head;
        this.tail = head;
    }
    enqueue(node) {
        if (this.head == null) {
            this.head = node;
            this.tail = this.head;
        } else {
            this.tail.next = node;
            this.tail = node;
        }
    }
    dequeue() {
        if (this.head == null) {
            console.log("empty queue");
        } else if (this.head.next == null) {
            this.head = null;
            this.tail = null;
        } else {
            this.head = this.head.next;
        }
    }
    peek() {
        if (this.head == null) {
            console.log("empty queue");
        } else {
            console.log(this.head.data);
        }
    }
}
class node {
    constructor (data=null, next=null) {
        this.data = data;
        this.next = next;
    }
}
let q = new Queue();
let node1 = new node(1);
q.enqueue(node1);
let node2 = new node(4);
q.enqueue(node2);
let node3 = new node(7);
q.enqueue(node3);
q.peek(); // 1
console.log(JSON.stringify(q)); // {"head":{"data":1,"next":{"data":4,"next":{"data":7,"next":null}}},"tail":{"data":7,"next":null}}
q.dequeue();
q.peek(); // 4
console.log(JSON.stringify(q)); // {"head":{"data":4,"next":{"data":7,"next":null}},"tail":{"data":7,"next":null}}
q.dequeue();
q.peek(); // 7
console.log(JSON.stringify(q)); // {"head":{"data":7,"next":null},"tail":{"data":7,"next":null}}
q.dequeue(); 
q.dequeue(); // empty queue
q.enqueue(new node(8)); 
console.log(JSON.stringify(q)); // {"head":{"data":8,"next":null},"tail":{"data":8,"next":null}}