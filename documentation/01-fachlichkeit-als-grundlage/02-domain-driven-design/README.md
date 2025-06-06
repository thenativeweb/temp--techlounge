# Domain-Driven Design (DDD)

- TDD: Test-Driven Development
- BDD: Behavior-Driven Development
- MDD: Model-Driven Design

- Domain-Driven Design
  - Fachlichkeit in den Vordergrund rücken
  - Um den Entwurf von Software zu verbessern

- Seit 2003
  - "Domain-Driven Design: Tackling Complexity in the Heart of Software"
  - Eric Evans

## Building-Blocks von Domain-Driven Design

- Domain-Events
  - ein neues To-Do wurde notiert
  - ein To-Do wurde angepasst
  - ein To-Do wurde nach hinten geschoben
  - …

- Commands
  - Notiere ein neues To-Do!
  - Passe das To-Do an!
  - Schiebe das To-Do nach hinten!
  - …

- Beispiel für 1:n von Commands zu Events
  - Open group   -->  Group opened + Group joined
  - Join group   -->  Group joined
  - Leave group  -->  Group left (+ Group closed)
  - Close group  -->  Group left + … + Group closed

- EVA-Prinzip
  - Commands = Eingabe
  - ?        = Verarbeitung
  - Events   = Ausgabe

- Command -> Command-Handler -> Event(s)

        func: Command -> Event[]

```golang
func NoteToDo (command NoteToDoCommand): ([]DomainEvent, error) {
  title := strings.TrimSpace(command.Title)
  description := strings.TrimSpace(command.Description)

  if title == "" {
    // Fachlich relevante Fehler:
    // return []DomainEvent{
    //   NewNoteTodoRejected("title must not be empyt")
    // }, nil

    // Technische Fehler:
    return nil, errors.New("title must not be empty")
  }

  return []DomainEvent{
    NewToDoNoted(title, description)
  }, nil
}
```

                  State <---------+
                    |             |
                    v             |
- Command -> Command-Handler -> Event(s)

    func: (Command, State) -> []Event

- Aggregates
  - Garantieren die Konsistenz ihres inneliegenden States
  - Indem sie Commands sequenzialisieren
  - Das heißt, innerhalb eines Aggregates wird Konsistenz garantiert; Aggregate-übergreifend gilt das jedoch nicht
  - Ein Aggregate (d.h., der darin enthaltene State) soll so klein wie möglich sein (wegen Parallelisierung von Commands), aber so groß wie nötig (um die Konsistenz zu wahren)
