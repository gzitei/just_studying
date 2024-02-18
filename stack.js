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
            throw new Error("empty stack.");
        } else {
            this.data.pop();
            this.top--;
        }
    }
}
let s = new Stack();
s.add(1);
s.add(2);
s.add(10); 
console.log(JSON.stringify(s)); // {"data":[1,2,10],"top":2}
s.remove();
s.remove();
console.log(JSON.stringify(s)); // {"data":[1],"top":0}
