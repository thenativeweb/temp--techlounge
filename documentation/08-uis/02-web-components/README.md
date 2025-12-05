# Web-Components

- Komponente
  - Gekapselte Einheit: HTML, CSS und JavaScript
  - Wiederverwendbarkeit und Isolation
  - Atomic Design
    - Atom: Button, Input, Combobox, …
    - Molekül: Kombinationen aus Atomen, zB Formular-Feld
    - Organismus: Login-Form

- Web-Components
  - Native Komponenten im Browser, ohne Frameworks
  - Drei APIs
    - Custom-Elements
    - Shadow DOM
    - HTML-Templates

- Custom-Elements
  - <tnw-calendar>...</tnw-calendar>
  - JavaScript-Klasse als Implementierung
    - Enthält Lifecycle-Methoden
    - ZB `connectedCallback` und `disconnectedCallback`
  - Wert-Übergabe / Properties
    - Als Attribut im HTML: <tnw-calender year="2026">
    - Als Property im JavaScript: tnwCalendar.year = 2026;
  - Custom-Events
    - Komponente kann eigene Events definieren
- Shadow-DOM
  - Separates DOM, das ins globale DOM eingeklinkt wird
  - Entspricht einem isolierten Namespace für CSS
  - Slots: Platzhalter für Content
- HTML-Templates
  - Das `<template>...</template>`-Element
  - Erfordert das Klonen von Knoten im DOM, fühlt sich an wie 1995
