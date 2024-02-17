class LinkedList {
    constructor(head = null) {
        this.head = head;
    }
    inserir(node) {
        if (this.head == null) {
            this.head = node;
        } else {
            let current = this.head;
            while(Boolean(current.next)) {
                current = current.next;
            }
            current.next = node;
        }
    }
}
class Node {
    constructor(data = null, next = null) {
        this.data = data;
        this.next = next;
    }
}
let head = new LinkedList();
for (let i = 0; i < 10; i++) {
    let new_node = new Node(i);
    head.inserir(new_node);
}
console.log(JSON.stringify(head));
