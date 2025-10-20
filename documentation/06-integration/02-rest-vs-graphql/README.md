# REST vs GraphQL

- These: Fachlichkeit vor Technologie
  - Commands und Queries als fachliche Ankerpunkte
  - HTTP + JSON reichen (erst einmal) völlig aus
  - KISS: Technologie dient der Fachlichkeit, nicht umgekehrt
    - Verben sind wichtig, sie drücken Intention aus
- Commands vs Queries
  - Commands ändern den Zustand, aber geben nichts zurück
  - Queries ändern den Zustand nicht, geben aber etwas zurück

```
POST /submit-order   (fachlich)
       vs
POST /order          (technisch)

GET /get-pending-todos
       vs
GET /todos?filterField=status&filterValue=pending
```

- REST (Representational State Transfer)
  - Roy Fielding, im Jahr 2000
  - URL: "Alles ist eine Ressource" => Substantiv
  - HTTP-Verben: GET, POST, PUT, DELETE, PATCH, …
  - HATEOAS
    - Hypermedia as the Engine of Application State
- GraphQL (Graph Query Language)
  - Facebook, im Jahr 2012
  - Client (!) definiert Abfragen auf Graph, die vom Server ausgeführt werden
  - Statisches Typsystem
  - Drei Arten von Interaktionsmöglichkeiten mit dem Graphen
    - Queries: Entsprechen Queries
    - Mutations: Entsprechen Commands
    - Subscriptions: Entsprechen Events
  - GraphQL hat seinen Preis
- Caching
  - Funktioniert bei REST so wie bei HTTP
  - GraphQL verwendet ausschließlich `POST /graphql`
    - Deshalb ist Caching bei GraphQL sehr viel komplexer
    - Letztlich muss sich der Client von Hand darum kümmern
- "Parse, Don't Validate"
  - JSON Schema in Verbindung mit OpenAPI

```json
{
  "type": "object",
  "properties": {
    "name": {
      "type": "string",
      "minLength": 3
    },
    "xp": {
      "type": "number",
      "min": 0,
      "max": 100
    }
  },
  "required": [ "name", "xp" ],
  "additionalProperties": false
}
```

- Cross-Cutting Concerns
  - Versionierung
    - HTTP-Route
      - `/api/v1/submit-order`
      - `/api/1.0.0/submit-order`
      - `/api/2025-10-17/submit-order`
        - Client verwendet entweder das *aktuelle* Datum für die Latest-Route oder ein hart-codiertes Datum, wo man dann weiß, wann das zum letzten Mal "angefasst" wurde
    - HTTP-Header
      - `X-Version: v1`
      - `X-Version: 1.0.0`
  - Sicherheit
    - Authentifizierung und Autorisierung
    - Bei GraphQL ist Query-Analyse erforderlich
  - Monitoring
    - Tracing für fachliche (evtl. verteilte) Prozesse

- gRPC für Service-zu-Service-Kommunikation
  - Stark typisiert
  - Sehr kompakt
  - Sehr performant
  - Aber: Es hat seinen Preis
- tRPC auf Basis von TypeScript
  - Typsicherheit, aber "ohne Schema"
  - Ende-zu-Ende-Verwendung von TypeScript

- Entscheidungsmatrix
  - Mit HTTP + JSON beginnen (mit Commands und Queries)
  - Alles Weitere erst bei Bedarf
  - Typische Fehler
    - GraphQL, gRPC, tRPC, … als magische Silver Bullet
    - Aufwand und Abhängigkeiten für GraphQL, gRPC, tRPC, … unterschätzen
    - Technische Patterns statt fachlicher Ausrichtung
    - Over-Engineering
