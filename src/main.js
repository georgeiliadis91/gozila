import { html, css, LitElement } from "lit";

export class SimpleGreeting extends LitElement {
  static styles = css`
    p {
      color: blue;
    }
  `;

  static properties = {
    count: { type: Number },
  };

  constructor() {
    super();
    this.count = 0;
  }

  _increment() {
    this.count++;
  }

  _decrement() {
    this.count--;
  }

  render() {
    return html`
      <h2>Hello from the lit component</h2>
      <button @click="${this._increment}">+</button>
      <button @click="${this._decrement}">-</button>
      <p>Count:, ${this.count}!</p>
    `;
  }
}

customElements.define("simple-greeting", SimpleGreeting);
