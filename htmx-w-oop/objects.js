class Button {
    constructor ({ text, ...kwargs }) {
        this.text = text;
        let props = Object.entries(kwargs).map(el => {
            let [k, v] = el;
            return `${k}="${v}"`;
        }).join(" ");
        this.html = `<button ${props}>${text}</button>`;
    }
}

module.exports = Button