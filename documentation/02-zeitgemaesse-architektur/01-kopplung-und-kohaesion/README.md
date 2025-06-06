# Kopplung und Kohäsion

## Kopplung

- Verbundenheit, Zusammenhang
- Beispiel für funktionale Kopplung

    ```golang
    func Foo() {
      // Kopplung
      bar()
    }

    func bar() {
      // ...
    }
    ```

- Beispiel für Datenkopplung

    ```golang
    const magicNumber = 19

    func Brutto(netto float64) float64 {
      // Kopplung an Konstante
      brutto := netto * (1 + magicNumber / 100)
      return brutto
    }

    func Steuersatz() int {
      // Kopplung an Konstante
      return magicNumber
    }

    func PferdeAufKoppel() int {
      return magicNumber
    }
    ```

## Kohäsion

- Inhaltliche Zusammengehörigkeit

    ```
    /
      cart/
        controller.go
        controller_test.go
        model.go
        list.go
        view.go
      search/
        controller.go
        model.go
        view.go
    ```

- Keine technische Gleichartigkeit

    ```
    /
      controllers/
        cart.go
        search.go
      models/
        cart.go
        search.go
      tests/
        controllers/
          cart_test.go
      utils/
        list.go
      views/
        cart.go
        search.go
    ```

## Idealfall

- Niedrige Kopplung
  - Weil eine Änderung nicht zig andere Änderung nach sich zieht
  - Änderungen sind lokal möglich und überschaubar

- Hohe Kohäsion
  - Alle erforderlichen Änderungen sind nah beinander und nicht weit verstreut

## DRY-Prinzip

- "Don't Repeat Yourself"
  - Wird oft interpretiert als "kein Copy / Paste"

- Beispiel "Versandhandel"
  - Eine Bestellung darf nicht mehr als 3 Artikel enthalten

    ```golang
    if numberOfArticles > 3 {
      fmt.Println("maximal 3 Artikel")
      return
    }
    ```

  - Es dürfen nicht mehr als 3 Artikel in einem Paket versendet werden

    ```golang
    if numberOfArticles > 3 {
      fmt.Println("maximal 3 Artikel")
      return
    }
    ```

  - Was wird gemacht?
    - Man lagert den (vermeintlich) gleichen Code in eine Funktion aus

    ```golang
    func checkArticles(context string, count int) string {
      if context == "bestellen" {
        if count > 3 {
          return "maximal 3 Artikel"
        }
      } else if context == "versand" {
        if count > 10 {
          return "maximal 10 Artikel"
        }
      }

      return ""
    }
    ```

## Entkopplungsmechanismen

- Asynchrone Kommunikation statt synchroner Kommunikation
- Fehler bei der Kommunikation (Latenz, Timeouts, Ausfall, …) einplanen
- Indirekte Kommunikation statt direkter Kommunikation

  ```
  Direkt:

    A ----> B

  Indirekt:

               /----> B
    A ----> MQ -----> C
               \----> D
  ```

- Events
  - Verwenden, um ein Bedürfnis zu signalisieren
  - Oder verwenden, um ein Ergebnis zu signalisieren
  - "Es ist etwas passiert …"
  - Events sind nicht zielgerichtet, sie gehen an die Allgemeinheit
  - Events erwarten keine synchrone Antwort

  ```
                Bestellung       /-------------> Versand
                aufgegeben      /
  Bestellwesen ------------> MQ ---------------> Buchhaltung
                              | \
                              |  \-------------> Reklamation
                              |
                              +-> Adapter -----> Kommunikation
                                (Evt -> Cmd)
  ```

- Abhängigkeiten
  - Message-Queue an sich
  - Message-Queue-Protokoll (`amqp://`, …)
  - Event-Format (wobei man das über Adapter lösen kann)
    - Solche Adapter nennt man auch "Context Mapper"

```golang
func HandleEvent(event Event) {
  if event.Type == "bestellung-aufgegeben" {
    kommunikationSvc.SendCommand("send-whatsapp", {
      // ...
    })
  }
}
```
