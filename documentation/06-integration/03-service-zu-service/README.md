# Service-zu-Service-Kommunikation

- Verteilte Systeme basieren auf Netzwerken
  - Dadurch werden sie inhärent komplex
  - Netzwerke sind unzuverlässig
  - Fehler und Latenz sind normal

- Spektrum an Kommunikation
  - Direkte Aufrufe bis zu entkoppelten Nachrichten
  - Es gibt keinen "richtigen" oder "falschen" Ansatz
    - Eher: Es gibt "passende" und "unpassende" Ansätze
  - Jeder Ansatz zieht Trade-Offs nach sich

- Synchrone Kommunikation
  - Request-Response
  - Direkter Aufruf, direktes Ergebnis
  - Enge Kopplung bezüglich Zeit und Verfügbarkeit
  - Vorteile
    - Einfach
    - Ergebnis liegt "sofort" vor
    - Einfache Fehlerbehandlung
  - Nachteile
    - Zeitliche Kopplung und Verfügbarkeit
    - Latenzen summieren sich

- Asynchrone Kommunikation
  - "Fire and forget" / "Eventual Response"
  - Kein direkter Aufruf, kein direktes Ergebnis
  - Lose Kopplung bezüglich Zeit und Verfügbarkeit
  - Message-Queues als Vermittler
  - Vorteile
    - Zeitliche Entkopplung und Möglichkeit der Nicht-Verfügbarkeit
    - Bessere Resilienz
  - Nachteile
    - Höhere Komplexität
    - Ergebnis liegt nicht "sofort" vor
    - Aufwändigere Fehlerbehandlung

- CAP-Theorem
  - Consistency vs Availability im Falle einer Partition
  - Synchron ist CP, asynchron ist AP
  - Die Frage nach CP vs AP ist *keine* technische Frage!
    - Sondern eine Business-Frage (Risiko und Wahrscheinlichkeit)

- Latenz berücksichtigen
  - Latenz ist unvermeidbar
  - Netzwerk-Technologie, geografische Distanz, Netzwerk-Hops, Verarbeitungszeit

- Timeouts (bei synchroner Kommunikation), um Latenz zu kontrollieren
  - Darf nicht zu kurz sein, ansonsten erhält man (vermeintliche) Fehler
  - Darf nicht zu lang sein, ansonsten verschwendet man Ressourcen und Zeit
  - Vorsicht
    - A -> B (mit Timeout von 30s)
    - A -> B -> C (jeweils mit Timeout von 30s)
      - A-B braucht einen längeren (!) Timeout als B-C

- Fehler sind die Regel, nicht die Ausnahme
  - Denn in einem verteilten System können jederzeit Komponenten / Services ausfallen
  - Partielle Ausfälle vs Totalausfall

- Protokolle für synchrone Service-zu-Service-Kommunikation
  - HTTP als Standard-Protokoll für synchrone Kommunikation
  - HTTP/2 und HTTP/3 als Alternativen
  - gRPC als High-Performance-Alternative

```
          Entkoppeln von API und Fachlichkeit
                      |
                      v
Client -> API (HTTP) --> Fachlogik
       -> API (gRPC) -->
```

- Protokolle für asynchrone S2S-Kommunikation
  - Message-Queues, zB AMQP, MQTT, …
  - At-least-once vs at-most-once
  - Push vs Pull

- Push vs Pull
  - Service A braucht Informationen von Service B
  - Option 1: Service A *fragt* Service B
    - Pull, synchron
    - Fördert Autorität
    - Favorisiert CP
  - Option 2: Service B *schickt proaktiv* Daten zu Service A
    - Push, asynchron
    - Fördert Autonomie
    - Favorisiert AP

```
Service X
  POST /register-webhook { url: 'http://...', ... }

1. Kommunikation: Registrieren -> synchron
2. Kommunikation: Callback (per HTTP) -> synchron

oder

1. Kommunikation: Registrieren -> Callback -> asynchron
2. Kommunikation: Callback (per HTTP) -> synchron
```

- Was macht man, wenn ein Aufruf schiefgeht?
  - Retry-Pattern
    - Fehler durch Wiederholen versuchen zu beheben
    - Idempotenz als (weiche) Voraussetzung
    - Kurze Retry-Zeiten, aber zufällig gestreut, um Spitzen zu vermeiden
    - Retry-Zeiten mit exponenziellem Backoff, also zB 100ms, dann 200ms, dann 400ms, dann 800ms, …
      - Aber gecappt (Obergrenze)!

  - Circuit-Breaker
    - Automatische Abschaltung bei wiederholten Fehlern
    - Drei Zustände
      - Closed (Normal)
      - Open (Blockiert)
      - Half-Open (Testet, ob Verbindung wieder möglich ist)

  - Bulkhead-Pattern
    - Isolation von Ressource-Nutzung
    - Durch Namespaces, durch Limits, durch Pools, …
    - Anschauliches Beispiel
      - Prozess-Isolation von Tabs im Browser

- Idempotenz
  - Mehrfache Ausführung führt zum gleichen Ergebnis
  - Unter Umständen kritisch für Retry-Logik
    - Eigentlich will man Exactly-Once
  - Da Exactly-Once nicht möglich ist
    - Weicht man in der Regel auf At-Least-Once aus
    - Das heißt, Nachrichten kommen eventuell doppelt
    - Das Ziel ist, Nachrichten zu deduplizieren
      - Eine Möglichkeit ist, mit idempotenten Aktionen zu arbeiten
      - Eine andere Möglichkeit ist, mit Message-IDs zu arbeiten

- USE und RED
  - USE
    - Utilization
      - Misst die Auslastung von CPU, Memory, Disk, …
      - ZB "CPU-Last liegt bei 80%"
    - Saturation
      - Die Überlastung, wie viel Arbeit gerade nicht erledigt werden kann
      - ZB "50 Threads warten auf CPU-Zeit"
    - Errors
  - RED
    - Rate
      - Der Durchsatz des Services (wird der Service genutzt)
      - ZB "1000 Requests/s"
    - Errors
    - Duration
      - Ist mein Service schnell genug?
      - ZB "700ms pro Request"
  - Vergleich
    - USE misst interne Faktoren des Services ("Warum?")
    - RED misst externe Faktoren des Services ("Was?")

- Chaos Engineering
  - Netflix Chaos Monkey
