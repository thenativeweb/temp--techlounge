# Event-Sourcing und CQRS

## Event-Sourcing

Ein Event ist ein Ereignis, das geschehen ist, und das eine gewisse Relevanz aus fachlicher Sicht hat.

- To-do memorized
- To-do adjusted
- To-do postponed
- To-do completed
- To-do discarded

CRUD = Create, Read, Update, Delete

Konten (klassisch: Status Quo)
-------------------------
KNr    | Saldo | SldVnVr1
-------------------------
23     | 3.000 |    2.500
42     |   750 |   10.000
17     |  -327 |     -417

Konten (mit Event-Sourcing: Veränderungen, die zum Status Quo geführt haben)
--------------------------------------------------------------------------
Order | KNr    | Typ            | Zeit       | Details
--------------------------------------------------------------------------
1     | 23     | Konto eröffnet | 17.02.2025 | { inhaber: "…" }
2     | 23     | Eingezahlt     | 18.02.2025 | { betrag: 1000 }
3     | 42     | Konto eröffnet | 21.02.2025 | { inhaber: "…" }
4     | 23     | Abgebucht      | 22.02.2025 | { betrag: 400, text: "…" }
5     | 42     | Eingezahlt     | 24.02.2025 | { betrag: 50 }
6     | 42     | Abgebucht      | 25.02.2025 | { betrag: 17, text: "lotto" }
7     | 23     | Überwiesen     | 27.02.2025 | { betrag: 300, ziel: "…" }
8     | 23     | Gutgeschrieben | 01.03.2025 | { betrag: 300, grund: "…" }
...

Wir beschränken uns bewusst (!) auf Create und Read. Update und Delete kommen nicht vor.

### Benefits

- Historische Daten
- Analysierbarkeit und Reports
- Interpretierbarkeit / Projektionen
- Fehlersuche und Debugging
- …

### Herausforderungen

- Stetig anwachsenden Speicherbedarf
  - Speicherbedarf pro Event
    - 1 KByte pro Event
    - 1.000 Events => 1 MByte
    - 1.000.000 Events => 1 GByte
    - 1.000.000.000 Events => 1 TByte (mein MacBook Pro von 2019)
  - Anzahl der Events
    - 5 Finanztransaktionen pro Tag
    - 5 * 365 = 1825
    - 100.000 Kunden => 182.500.000
    - => 5,5 Jahre bis 1 Mrd Events erreicht sind

- Replay dauert immer länger
  - Snapshots, um Zwischenstände zu speichern
  - Versuchen, Event-Ströme kurz zu halten
  - Snapshots evtl als Events denken

- DSGVO Artikel 17 "Recht auf Vergessenwerden"
  - DSGVO ignorieren (nicht empfohlen)
  - Event "DSGVO Art 17 ist eingetreten" und Projektionen bereinigen
  - Crypto-Thrashing / Crypto-Shredding (Daten verschlüsseln und ggf Schlüssel wegwerfen)
  - PII in eine extra Datenbank auslagern und im Event nur einen Pointer speichern
  - Pseudonymisieren / Anonymisieren
  - Migration in einen neuen Event-Store, wobei während der Migration gefiltert wird (ETL = Extract, Transform, Load)

- Semantische Änderung an Events im Lauf der Zeit
  - `überweisung wurde ausgeführt` (Typ des Events)
  - Schritt 1: Versions-Suffix einführen (im Event-Store)
    - `überweisung wurde ausgeführt V2`
    - V1 und V2 gesondert behandeln
  - Schritt 2: Upcasting (in der Anwendung / Projektion)
    - `fn (eventV1) -> eventV2`
    - Für die Anwendung sieht das so aus, als gäbe es nur noch V2-Events
  - Schritt 3: ETL-Szenario + Migration

- Unique-Constraint-Problem
  - Alle relevanten (!) Events auf mögliche Duplikate auswerten, bevor man Events speichert
  - Auf die Projektion zugreifen, aber evtl Eventual-Consistency-Problem
  - Reservation-Pattern nutzen ("Konto beantragt"), und Freischaltung durch House-Keeping-Task
  - Gesonderte Lock-Tabelle, aber hier kann es passieren, dass zwar das Lock acquired wird, aber dann das System abstürzt => Lock-Tabelle enthält Locks, die nicht genutzt werden

## CQRS

CQRS = Command Query Responsibility Segregation

```
                         +------ Replay ---------+
                         |                       |
                         v                       |
       --> Command ---> API ---> Event(s) ---> Event-Store
      /                                          |
     /                                           v
    /                                       Message-Queue ---> SMS-Service
User                                             |
    \                                            v
     \                                       Projector
      \                                          v
       --> Query -----> API -----------> Projektionen / Views
```
