# Task-orientierte UIs

- CRUD-UI
  - "Forms over Data"
  - Tabelle mit Datensätzen
    - Neu-Button
    - Bearbeiten-Button
    - Löschen Button
    - Formular mit allen Feldern eines Datensatzes
    - Speichern-Button
  - Probleme
    - Technisch, nicht fachlich
    - Validierung
    - Nichtssagenden Audit-Trail ("Customer Updated")
    - Mobile-Tauglichkeit

- Tasks
  - Fachliche Aktion mit Intention
  - Weg von Daten, hin zu Use-Cases
  - User-zentriert statt Daten-zentriert denken

- DDD
  - Buttons, Menüeinträge, … spiegeln fachliche Intention wieder
  - Ubiquitous Language auch in der UI!

- CQRS
  - Trennung von Commands und Queries bzw. von Schreiben und Lesen sollte sich auch in der UI widerspiegeln
  - Commands
    - Entsprechen Tasks
    - Formulare, Dialoge, Buttons, Menüeinträge, …
  - Queries
    - Entsprechen Projektionen
    - Tabelle, Diagramm, Labels, Read-Only-Textboxem, …
  - Events
    - Aktivitäts-Feed, Notifications, History, …

- Umgang mit Task-Completion
  - Optimistisch
    - UI zeigt Erfolg sofort an, Fehler später korrigieren
  - Pessimistisch
    - UI wartet auf Backend-Bestätigung per Projektion und / oder Event

- Fachliche Anforderungen als User-Stories
  - "Als [Rolle] will ich [Aktion], um [Ziel]."
  - Das "Warum" ist das wichtigste davon
  - Eine Story entspricht einem Task (oder mehreren (Workflow))
  - Für Task definieren
    - Kontext: Informationen für Entscheidungsfindung
    - Eingabe: Was muss der User eingeben?
    - Validierung: Regeln, Situationen, Fehlermeldungen, …
      - Technisch auf Client-Seite
      - Technisch + fachlich auf Server-Seite
    - Ausgabe: Erfolgsmeldung, nächste Schritte, …
  - Von der Idee zum Design
    - Scribble: Pen-and-Paper, Struktur statt Design
    - Wireframes: Digitale Skizzen, Struktur und Navigation
    - Click-Dummy: Simulation der UI für User-Tests

- Best-Practices
  - Konsistenz
  - Berechtigungen (nicht CRUD-basiert)
  - Fehlerbehandlung (synchron, asynchron, …)
  - Stale-Data visualisieren
  - Offline-First
  - Keine God-Forms
  - Keine Hidden-Actions
  - Keine generischen Buttons (kein CRUD)
