# Containerisieren mit Docker

- Der Begriff "Docker"
  - Software
  - Unternehmen
  - Gattungsbegriff

## VM-basierter Ansatz

```
                                      App 1   App 2
                                        Tooling
                                        Kernel
                                    Virtueller Host
App 1     App 2     ...    App n    Virtualisierung
                   Tooling
                   Kernel
                    Host
```

## Container-basierter Ansatz

```
                               App 3      App 4
            App 1    App 2    Docker (Server + CLI)
                   Tooling
                   Kernel
                    Host
```

## Command vs Entrypoint

docker run esdb

CMD [ "xxx" ]
docker run esdb     => docker run esdb xxx
docker run esdb yyy => docker run esdb yyy

ENTRYPOINT [ "xxx" ]
docker run esdb     => docker run esdb xxx
docker run esdb yyy => docker run esdb xxx yyy


CMD:       docker run esdb --api-token=secret => Fehler
ENTRYPOINT docker run esdb --api-token=secret => OK
