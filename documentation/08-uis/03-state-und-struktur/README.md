# State und Struktur

- Wie baut man Listen in Web-Components dynamisch auf?

```javascript
renderTodo(todo) {
  return html`
    <tnw-todo>${todo.title}</tnw-todo>
  `;
}

render() {
  const todos = [
    // ...
  ];

  if (todos.length === 0) {
    return html`
      <div>Keine To-Dos vorhanden</div>
    `;
  }

  return html`
    <tnw-todos>
      ${todos.map(todo => renderTodo(todo))}
    </tnw-todos>
  `;
}
```

- State
  - State = Daten, die sich ändern
  - Änderungen am State erzwingen ein Re-Rendering
  - Unterschied zu Properties
    - Properties werden von Außen in eine Komponente hineingegeben, read-only
    - State wird intern verwaltet, änderbar
  - EVA-Prinzip
    - Props => Eingabe
    - State => Verarbeitung
    - Events => Ausgabe
  - Wo gehört State hin? => State-Management

- Re-Rendering
  - Erfolgt automatisch in Lit, sobald sich eine `@property` ändert
  - `@state` verwenden für interne Variablen statt `@property`

- Wo gehört State hin?
  - "Lift state up" => State so weit oben wie möglich unterbringen
  - State, der möglichst weit "oben" sitzt, ist eine Single-Source of Truth

- Zwei Arten von Komponenten
  - Presentational Component
    - Daten anzeigen
    - Nur Props, kein State
    - Ggf. bei Eingaben Events auslösen
    - Wiederverwendbar und testbar
  - Container Component
    - Hat State
    - Daten laden und ggf. zu speichern, zB über eine API
    - Führt Geschäftslogik aus
    - Verwaltet anwendungsrelevanten Daten
    - Koordiniert Presentational Components
  - Viele Presentational Components, wenig Container Components

- State-Management

```javascript
render() {
  const todos = [];

  return <App todos={todos} />
}

class App {
  render(todos) {
    return <TodoList todos={todos} />
  }
}

class TodoList {
  render(todos) {
    return <Todo todo={...}>
  }
}
```

```html
const todos=...

<tnw-app>
  <tnw-todo-list>
    ${todos.map =>
      <tnw-todo-item text=${todo.title}></tnw-todo-item>
    }
  </tnw-todo-list>
</tnw-app>
```

- Fachliche vs technische Struktur im Frontend

```
models/
  AuthModel.js
  CartModel.js
controllers/
  AuthController.js
  CartController.js
views/
  AuthView.js
  AuthView.html
  AuthView.css
  CartView.js
  CartView.html
  CartView.css
```

vs

```
shared/
  database.js
domain/
  auth/
    Model.js
    Controller.js
    View.js
    View.html
    View.css
  cart/
    Model.js
    Controller.js
    View.js
    View.html
    View.css
```
