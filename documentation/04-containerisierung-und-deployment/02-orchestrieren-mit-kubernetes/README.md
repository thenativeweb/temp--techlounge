# Orchestrieren mit Kubernetes

- Ziele eines Clusters
  - Load-Balancing (Last auf mehrere Server verteilen)
  - Failover

- Zwei Arten von Servern
  - Control-Plane-Server (1, 3 oder 5)
  - Worker-Server (1 bis n)

- https://kubernetes.io/

- Leichtgewichtige(re) Alternativen zu K8s
  - K0s (https://k0sproject.io/)
  - K3s (https://k3s.io/)
  - K9s (https://k9scli.io/)

- Konfigurationsdatei
  - `~/.kube/config`
  - `kubectl --kubeconfig <datei>`

- kubectl
  - `<verb> <resource>`
  - zB: `kubectl get nodes`

```
LoadBalancer ----+
                 |
                 v
              Ingress (Nginx)
                 |
                 v
              Service (App, ClusterIP)
                 |
                 v
                Pod <------ Deployment
                 |
                 v
            Container (App)
                 |
                 v
              Service (Datenbank, ClusterIP)
                 |
                 v
                Pod <------ Deployment
                 |
                 v
            Container (Datenbank)
```
