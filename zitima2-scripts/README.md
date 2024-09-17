# Ζήτημα 2 - Δοχειοποίηση RESTful Υπηρεσίας

Αυτό το έργο αποτελεί την υλοποίηση του Ζητήματος 2 της εργασίας για το μάθημα "Τεχνολογίες Νέφους". Περιλαμβάνει τη δοχειοποίηση (containerization) της RESTful υπηρεσίας και της βάσης δεδομένων PostgreSQL που αναπτύχθηκαν στο Ζήτημα 1.

## Περιεχόμενα

- [Προαπαιτούμενα](#προαπαιτούμενα)
- [Δομή Έργου](#δομή-έργου)
- [Εγκατάσταση](#εγκατάσταση)
- [Χρήση](#χρήση)
- [Διαχείριση Containers](#διαχείριση-containers)
- [Αντιμετώπιση Προβλημάτων](#αντιμετώπιση-προβλημάτων)

## Προαπαιτούμενα

- Docker
- Docker Compose (προαιρετικό, για μελλοντική χρήση)
- Git

## Δομή Έργου

```
icsd-cloud-2024/
├── app/
│   └── ... (κώδικας εφαρμογής από το Ζήτημα 1)
├── zitima2-scripts/
│   ├── setup.sh
│   ├── start.sh
│   └── stop.sh
├── Dockerfile
└── README.md
```

## Εγκατάσταση

1. Κλωνοποιήστε το αποθετήριο:
   ```
   git clone <URL_του_αποθετηρίου>
   cd icsd-cloud-2024
   ```

2. Μεταβείτε στο branch του Ζητήματος 2:
   ```
   git checkout zitima-2
   ```

3. Βεβαιωθείτε ότι το Docker είναι εγκατεστημένο και εκτελείται στο σύστημά σας.

## Χρήση

Όλα τα scripts βρίσκονται στον φάκελο `zitima2-scripts`. Για να τα χρησιμοποιήσετε:

1. Μεταβείτε στον φάκελο των scripts:
   ```
   cd zitima2-scripts
   ```

2. Δώστε δικαιώματα εκτέλεσης στα scripts:
   ```
   chmod +x setup.sh start.sh stop.sh
   ```

3. Εκτελέστε τα scripts με την ακόλουθη σειρά:

   - Για να δημιουργήσετε την εικόνα Docker και τον volume της βάσης δεδομένων:
     ```
     ./setup.sh
     ```

   - Για να ξεκινήσετε τα containers:
     ```
     ./start.sh
     ```

   - Για να σταματήσετε τα containers:
     ```
     ./stop.sh
     ```
## Διαχείριση Containers

- Το container της βάσης δεδομένων ονομάζεται `postgres`.
- Το container της RESTful υπηρεσίας ονομάζεται `restful-service`.
- Και τα δύο containers συνδέονται στο δίκτυο `myapp-network`.

## Αντιμετώπιση Προβλημάτων

1. Αν αντιμετωπίζετε προβλήματα σύνδεσης μεταξύ των containers:
   - Ελέγξτε ότι και τα δύο containers είναι συνδεδεμένα στο `myapp-network`:
     ```
     docker network inspect myapp-network
     ```

2. Για να ελέγξετε τη συνδεσιμότητα μεταξύ των containers:
   ```
   docker exec restful-service ping -c 4 postgres
   ```

3. Για να ελέγξετε τη σύνδεση με τη βάση δεδομένων:
   ```
   docker exec restful-service sh -c "nc -zv postgres 5432"
   ```

4. Αν χρειάζεστε περισσότερες πληροφορίες, ελέγξτε τα logs των containers:
   ```
   docker logs postgres
   docker logs restful-service
   ```
