class node {
    data: number;
    next: node | null;
    constructor(data: number, next = null) {
        this.data = data;
        this.next = next;
    }
}
class LinkedList {
    head: node | null;
    constructor(head = null) {
        this.head = head;
    }
    inserir(node:node) {
        if (this.head == null) {
            this.head = node;
        } else {
            let current: node | null | undefined;
            current = this.head;
            while(Boolean(current?.next)) {
                current = current?.next;
            }
            current!.next = node;
        }
    }
}
let head = new LinkedList();
for (let i:number = 0; i < 4; i++) {
    let new_node: node;
    new_node = new node(i);
    head.inserir(new_node);
}
console.log(JSON.stringify(head)); // {"head":{"data":0,"next":{"data":1,"next":{"data":2,"next":{"data":3,"next":null}}}}}
