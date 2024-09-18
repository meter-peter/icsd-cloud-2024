# Ζήτημα 4: Μοντελοποίηση σε Kubernetes

## Εισαγωγή

Σε αυτό το ζήτημα, μοντελοποιήσαμε το backend σε μια σειρά από YAML αρχεία με βάση το πλαίσιο Kubernetes, χρησιμοποιώντας το εργαλείο kind για τοπική ανάπτυξη και δοκιμή.

## Προαπαιτούμενα

- kubectl
- kind (Kubernetes IN Docker)
- Docker

## Περιεχόμενα

- `app-deployment.yaml`: Deployment για τη RESTful υπηρεσία
- `postgres-statefulset.yaml`: StatefulSet για τη βάση δεδομένων PostgreSQL
- `app-service.yaml`: LoadBalancer Service για την εφαρμογή
- `postgres-headless-service.yaml`: Headless Service για τη βάση δεδομένων
- `postgres-pv-pvc.yaml`: PersistentVolume και PersistentVolumeClaim για τη βάση δεδομένων
- `db-secrets.yaml`: Secret για τα διαπιστευτήρια της βάσης δεδομένων
- `postgres-init-script.yaml`: ConfigMap για το script αρχικοποίησης της βάσης δεδομένων

## Λεπτομέρειες Υλοποίησης

### Deployment (app-deployment.yaml)
- Αριθμός αντιγράφων: 3
- Περιορισμοί πόρων: CPU και μνήμη
- Έλεγχοι υγείας: readiness και liveness probes
- Πολιτική ανανέωσης: RollingUpdate

### StatefulSet (postgres-statefulset.yaml)
- Αριθμός αντιγράφων: 1
- Χρήση PersistentVolumeClaim για μόνιμη αποθήκευση
- Έλεγχος υγείας

### Services
- LoadBalancer για την εφαρμογή
- Headless Service για τη βάση δεδομένων

### Διαχείριση Μυστικών
- Χρήση Kubernetes Secrets για τα διαπιστευτήρια της βάσης δεδομένων

## Χρήση του kind

1. Δημιουργία cluster:
   ```
   kind create cluster
   ```

2. Φόρτωση της εικόνας Docker στο cluster:
   ```
   kind load docker-image restful-service
   ```

3. Εφαρμογή των YAML αρχείων:
   ```
   kubectl apply -f .
   ```

4. Έλεγχος της κατάστασης:
   ```
   kubectl get pods
   kubectl get services
   ```

## Πρόσβαση στην Εφαρμογή

Για να αποκτήσετε πρόσβαση στην εφαρμογή:

```
kubectl port-forward service/restful-service 8080:8080
```

Η εφαρμογή θα είναι διαθέσιμη στο `http://localhost:8080`.

## Προκλήσεις και Λύσεις

1. **Διαχείριση Εικόνων Docker**: 
   - Πρόκληση: Το kind cluster δεν έχει αυτόματη πρόσβαση σε τοπικές εικόνες Docker.
   - Λύση: Χρησιμοποιήσαμε την εντολή `kind load docker-image` για να φορτώσουμε την εικόνα στο cluster.

## Συμπεράσματα

Η μοντελοποίηση του backend σε Kubernetes μας επέτρεψε να δημιουργήσουμε μια ευέλικτη και κλιμακώσιμη αρχιτεκτονική. Μέσω αυτής της διαδικασίας, αποκτήσαμε βαθύτερη κατανόηση των εννοιών του Kubernetes όπως Deployments, StatefulSets, Services, και PersistentVolumes, καθώς και πώς αυτά συνεργάζονται για να δημιουργήσουν μια ολοκληρωμένη εφαρμογή.

[Προηγούμενο: Ζήτημα 3 - Μοντελοποίηση με Docker Compose](../README-zitima3.md)
