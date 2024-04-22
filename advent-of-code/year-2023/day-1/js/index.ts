let init = Date.now();
import fs from 'fs';

const content = fs.readFileSync('../input.txt');
let sum = 0;
let text;
let start = 0;
let pos = 0;
let end;

do {

    end = content.indexOf(10, start);
    
    text = content.subarray(start, end);
    
    for (let i = 0; i < text.length; i++) {
        let curr = text[i] - 48;
        if (curr >= 0 && curr < 10) {
            sum = sum + 10*curr;
            pos = i;
            break;
        }
    }
    
    for (let i = text.length - 1; i >= pos; i--) {
        let curr = text[i] - 48;
        if (curr >= 0 && curr < 10) {
            sum = sum + curr;
            break;
        }
    }
    
    start = end+1;

} while (text.length > 0)

const finish = Date.now();

console.log(sum, finish - init, 'ms');
