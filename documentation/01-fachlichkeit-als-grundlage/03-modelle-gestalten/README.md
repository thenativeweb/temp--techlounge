# Modelle gestalten

```
                State <---------+
                  |             |
                  v             |
Command -> Command-Handler -> Event(s)
```

- Gemeinsame Sprache
  - Ideen
    - "Shared Language"
    - "Common Language"
    - …

  - Tatsächlicher Begriff
    - "Ubiquitous Language"
      - Allgegenwärtige Sprache

    - Ist das das Gleiche wie eine "Universal Language"?
      - Nein!

- Sprache ist Kontext-abhängig
  - "Bounded Contexts"

- Domain                    Fachthema, Fachlichkeit, …
  - Bounded Contexts        Namespaces für Sprache
    - Ubiquitous Language   Allgegenwärtige Sprache
    - Aggregates            Konsistenz-Grenzen
      - State
      - Commands
      - Command-Handler
      - Events

- Alternativbegriffe
  - Fachlich    =>  Intention (was)
  - Technisch   =>  Implementierung (wie)

- Beispiel für Bounded-Contexts

```
Schach
  Spielen
    Figur
    Bauer
    Pferd (oder Springer?) (im Sinne von "Rolle der Figur")
    Läufer
    …
    Schwarz (im Sinne von "Partei 1")
    Weiß (im Sinne von "Partei 2")
    Felder
    …

  Figurendesigner
    Pferd (im Sinne von "Tier mit vier Beinen")
    Drache
    Hund
    Katze
    …
    Schwarz (im Sinne von `#000000`)
    Weiß (im Sinne von `#ffffff`)
```

## Aggregate-Design

### 1. Command

Command:
  Memorize to-do { title, details, dueDate }

Command-Handler:
  fn(cmd, state) -> []evt
  Regeln:
  - title darf nicht leer sein
  - title muss eindeutig sein in Bezug auf die Liste (?)
    - das werden wir *nicht* umsetzen
  - details dürfen leer sein
  - dueDate darf nicht in der Vergangenheit liegen

Event:
  To-do memorized { toDoId, title, details, dueDate }

### 2. Command

Command:
  Complete to-do { toDoId }

Command-Handler:
  fn(cmd, state) -> []evt
  Regeln:
  - das to-do darf nicht schon completed sein

Event:
  To-do completed { toDoId }

### Was ist der State?

1. Möglichkeit: Ein globaler State

```
State         <-- den gibt es global genau 1x
  []ToDoList
    []ToDo
```

+ Alle Business-Regeln lassen sich damit abdecken
- Schlechte Performance, weil User sich in die Quere kommen

2. Möglichkeit: Ein State pro To-Do-Liste

```State      <-- den gibt es mehrfach, genau 1x pro User
  []ToDo
```

+ Alle Business-Regeln lassen sich damit abdecken
o Gute Performance, weil jeder User für sich (!) arbeitet
  Schlechte Performance, wenn ein User mehrere Commands gleichzeitig sendet

3. Möglichkeit: Ein State pro To-Do

```State      <-- den gibt es mehrfach, genau 1x pro To-Do
  ToDo
```

- Wir können nicht alle Business-Regeln damit abdecken
+ Sehr gute Performance mit kaum Konflikten

## Workshop-Formate

- Event-Storming (aka Model-Storming)
  - Alberto Brandolini

  - Wer?
    - Interdisziplinäres Team, nicht zu groß (< 10)

  - Wo?
    - Große Wand, beschreib- und beklebbar
    - Keine Tische, keine Stühle

  - Was?
    - Viel Papierrollen ;-)
    - Post-Its
    - Stifte

  - Ablauf?
    - 0. Scope festlegen (klein!)
    - 1. Events sammeln (jede:r für sich)
    - 2. Events in eine grobe (!) zeitliche Reihenfolge bringen
    - 3. Events clustern nach fachlicher Ähnlichkeit
    - 4. Events konsolidieren
    - 5. Commands sammeln
    - 6. Parameter und State festlegen
    - 7. Aggregates definieren
    - 8. Iterieren ab Schritt 1

- Domain-Storytelling
  - C1 WPS
  - https://domainstorytelling.org/
