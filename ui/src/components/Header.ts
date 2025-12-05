import { css, html, LitElement } from 'lit';
import { customElement, property } from 'lit/decorators.js';

@customElement('tl-header')
class Header extends LitElement {
  static styles = css`
    header {
      padding: 16px;
      text-align: center;
    }

    h1 {
      font-family: sans-serif;
      font-size: 36px;
      font-weight: 400;
      color: var(--highlight-color);
    }

    @media (max-width: 600px) {
      h1 {
        font-size: 24px;
      }
    }
  `;

  @property({ type: String })
  year = '';

  render() {
    return html`
      <header>
        <h1>tech:lounge Masterclass ${this.year}</h1>
      </header>
    `;
  }
}
