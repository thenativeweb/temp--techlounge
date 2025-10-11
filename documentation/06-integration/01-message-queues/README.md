# Message-Queues und Integrationsmuster

- Message-Queues
  - Asynchrone Zustellung von Nachrichten
  - Entkopplung von Sender und Empfänger
    - Zeit
    - Bekanntheit
    - Erreichbarkeit
    - …
  - Fire-and-forget
  - Vorteile
    - Robustheit
    - Fehlertoleranz
    - Skalierbarkeit
    - Verlässlichkeit
  - Herausforderungen
    - Fehler
    - Reihenfolgerichtigkeit
    - Nachvollziehbarkeit
  - CAP-Theorem
    - Consistency, Availability, Partition Tolerance

```
Publisher -> Message -> MQ

                        MQ -> Message -> Consumer
```

- Push vs Pull
  - Push (zB RabbitMQ): MQ sendet Nachrichten an Consumer
  - Pull (zB Kafka): Consumer holen Nachrichten bei der MQ ab
- At least once, at most once, exactly once
  - Consumer sendet ACK (verbunden mit Timeout) oder NACK
  - At least once: Nachricht erst verarbeitet, dann bestätigt
  - At most once: Nachricht erst bestätigt, dann verarbeitet
  - Exactly once: Unmöglich
    - Deduplizieren von Nachrichten
      - Durch Duplikatserkennung
      - Idempotente Nachrichten
- Dead-Letter-Queues
  - Fangen Nachrichten auf, die nicht zugestellt werden können

```
# SEDA (Staged Event-Driven Architecture)

Video-Upload -> MQ -> SD -> MQ -> HD -> MQ -> 4K -> Publish
```

- Integrationsmuster
  - Pipes and Filters
    ```
    --|--|--|--
    ```
  - Request / Response
    ```
    A ---> B
    A <--- B
    ```
  - Publish / Subscribe
    ```
           B1
    A ---> B2
           B3
    ```

    - An welches B wird die Nachricht zugestellt?
      - Routing-Strategien
        - Fanout
        - Round-Robin
        - Topic-basiertes Routing (statisch, per Topic-Key)
        - Dynamisches Routing
          - "Modulo"-Verteilung
          - Consistent Hashing

```
M = Message
C = Consumer
Farbe = Topic-/Routing-Key

M1(rot)  -> hash(rot)  -> 7fba03e4
M2(blau) -> hash(blau) -> 6a4f120a
M3(rot)  -> hash(rot)  -> 7fba03e4
M4(grün) -> hash(grün) -> bad5fe74
M5(blau) -> hash(blau) -> 6a4f120a

C1(127.0.0.1:3000) -> hash(127.0.0.1:3000) -> 2a9fb47a
C2(127.0.0.1:4000) -> hash(127.0.0.1:4000) -> 6ad14b7f
C3(127.0.0.1:5000) -> hash(127.0.0.1:5000) -> a7bdfa82

                00000000
                  |
               __---__  ___ 2a9fb47a (C1)
              /       \/
             /         \
            |           |
             \         /
 a7bdfa82 ___/\__   __/\___ 6ad14b7f (C2)
   (C3)          --- |      Zuständig: 6ad14b7f - a7bdfa81
                     |
                  7fba03e4 (rot)

=> rote Nachrichten werden von C2 verarbeitet
   blaue Nachrichten werden von C1 verarbeitet
   grüne Nachrichten werden von C3 verarbeitet
```

- Bücher
  - Enterprise Integration Patterns: https://www.amazon.de/dp/0321200683
  - Understanding Distributed Systems: https://www.amazon.de/dp/1838430210
  - Designing Data-Intensive Applications: https://www.amazon.de/dp/1449373321
