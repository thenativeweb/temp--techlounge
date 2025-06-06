# APIs gestalten, versionieren und Co.

- Technologien
  - OSI-Modell (vereinfacht)
    - Hardware-Ebene (Ethernet, WLAN, Token Ring, …)
    - Routing (IP, ICMP, …)
    - Transport (TCP, UDP, …)
    - Anwendung (HTTP, HTTPS, FTP, DNS, SMTP, …)

- Protokolle
  - REST
    - Konzepte
      - Ressourcen (Entitäten, Prozesse, …)
      - Verben (GET, POST, PUT, PATCH, DELETE, …)
      - URLs
    - Alles zusammen
      - Ressource (`/books/42` oder `/todos/23efca`)
      - Aktionen (`POST /books` oder `DELETE /todos/23efca`)
        - Daten üblicherweise im Request-Body als JSON oder XML
        - Bei GET-Anfragen Daten üblicherweise im Querystring (`GET /books?orderBy=title`)
      - Antworten
        - HTTP-Statuscode (`200 OK`, `404 Not Found`, `409 Conflict`, …)
        - Daten üblicherweise im Response-Body als JSON oder XML
    - HATEOAS
      - Hypertext as the Engine of Application State

  - GraphQL
    - Jeder Aufruf erfolgt über POST
    - Jeder Aufruf geht gegen den selben Endpunkt (in der Regel so etwas wie `/graphql`)
    - GraphQL unterscheidet zwei (drei) Arten von Interaktionen mit dem Server
      - Mutations
        ```graphql (Pseudocode)
        mutation UpdateBook {
          id: 42
          title: 'Fräulein Smillas Gespür für Schnee'
        }
        ```

      - Queries
        ```graphql (Pseudocode)
        query {
          profile {
            firstName
            lastName
            profilePictureUrl

            friends {
              firstName
              lastName

              status {
                latestPost
              }
            }
          }
        }
        ```
      - (Subscriptions)

  - gRPC
    - Binärformat, basiert auf Protobuf
    - Sehr performant, sehr klein

- Inhalte
  - CQRS auf eine API abbilden
    - POST für Commands
      - Beispiele
        - `POST /send-email`
        - `POST /memorize-todo`
        - `POST /order-management/submit-order`
        - `POST /order-management/cancel-order`
        - `POST /order-management/delete-order`
        - …

    - GET für Queries
      - Beispiele
        - `GET /all-orders`
        - `GET /order-details/:orderId`

- Versionierung
  - Über die URL
    - `POST /api/v1/send-email`
  - Über das Datum
    - `POST /api/2025-04-25/send-email`
  - Über einen expliziten Versions-Header
    - `X-Online-Shop-API-Version: 1.0.0`
  - Über neue Endpunkte
    - `POST /api/send-email-new`

- Authentifizierung
  - Zugangsdaten in der URL mitschicken
  - Zugangsdaten in einem Cookie mitschicken
  - Zugangsdaten per Header mitschicken
    - `Authorization: Basic jane.doe:secret`
    - `Authorization: Digest encrypted(jane.doe:secret)`
    - `Authorization: Bearer <API_TOKEN>`

## OpenAPI

- `GET /.well-known/openapi`
