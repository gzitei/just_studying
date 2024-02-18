class Stack {
    constructor() {
        this.data = [];
        this.top = -1;
    }
    add(value) {
        this.data.push(value);
        this.top = this.top + 1;
    }
    remove() {
        if (this.top < 0) {
            console.log("empty stack");
        } else {
            this.data.pop();
            this.top--;
        }
    }
    peek() {
        if (this.top < 0) {
            console.log("empty stack");
        } else {
            console.log(this.data[this.top]);
        }
    }
}
let s = new Stack();
s.add(5);
s.add(2);
s.add(9);
s.peek(); // 9
console.log(JSON.stringify(s)); // {"data":[5,2,9],"top":2}
s.remove();
s.remove();
s.peek(); // 5
console.log(JSON.stringify(s)); // {"data":[5],"top":0}
