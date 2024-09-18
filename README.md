# Εργασία Τεχνολογίες Νέφους - RESTful Υπηρεσία Διαχείρισης Χρηστών

Καλώς ήρθατε στο repository της εργασίας για το μάθημα "Τεχνολογίες Νέφους". Αυτό το έργο περιλαμβάνει την ανάπτυξη, δοχειοποίηση και ορχήστρωση μιας RESTful υπηρεσίας διαχείρισης χρηστών.

## Δομή Εργασίας

Η εργασία χωρίζεται σε τέσσερα βασικά ζητήματα, καθένα από τα οποία βρίσκεται σε ξεχωριστό branch:

1. [Ζήτημα 1: Ανάπτυξη RESTful Υπηρεσίας](app/README.md)
2. [Ζήτημα 2: Δοχειοποίηση με Docker](zitima2-scripts/README.md)
3. [Ζήτημα 3: Μοντελοποίηση με Docker Compose](README-zitima3.md)
4. [Ζήτημα 4: Μοντελοποίηση σε Kubernetes](yaml/README.md)

## Περιήγηση στην Εργασία

Αγαπητέ καθηγητή,

Για να περιηγηθείτε στην εργασία, προτείνουμε να ακολουθήσετε την εξής σειρά:

1. Ξεκινήστε με το [README του Ζητήματος 1](app/README.md) για να δείτε την υλοποίηση της βασικής RESTful υπηρεσίας.
2. Συνεχίστε με το [README του Ζητήματος 2](zitima2-scripts/README.md) για τη δοχειοποίηση με Docker.
3. Προχωρήστε στο [README του Ζητήματος 3](README-zitima3.md) για τη μοντελοποίηση με Docker Compose.
4. Τέλος, εξερευνήστε το [README του Ζητήματος 4](yaml/README.md) για τη μοντελοποίηση σε Kubernetes.

## Branches

Όλα τα ζητήματα έχουν αναπτυχθεί σε ξεχωριστά branches (zitima-1, zitima-2, zitima-3, zitima-4) και έχουν συγχωνευθεί στο main branch χωρίς συγκρούσεις.

## Δομή Repository

```
.
├── app
│   ├── cmd
│   │   ├── cmd
│   │   └── main.go
│   ├── docs
│   │   ├── docs.go
│   │   ├── swagger.json
│   │   └── swagger.yaml
│   ├── go.mod
│   ├── go.sum
│   ├── internal
│   │   ├── handlers
│   │   │   └── user_handlers.go
│   │   ├── models
│   │   │   └── user.go
│   │   ├── routes
│   │   │   └── routes.go
│   │   └── validators
│   │       └── user_validators.go
│   └── README.md
├── docker-compose.yaml
├── Dockerfile
├── migrations
│   └── init.sql
├── README.md
├── yaml
│   ├── app-deployment.yaml
│   ├── app-service.yaml
│   ├── db-secrets.yaml
│   ├── postgres-headless-service.yaml
│   ├── postgres-init-script.yaml
│   ├── postgres-pv-pvc.yaml
│   ├── postgres-statefulset.yaml
│   └── README.md
└── zitima2-scripts
    ├── README.md
    ├── setup.sh
    ├── start.sh
    └── stop.sh
```

Ελπίζω να βρείτε την εργασία πλήρη. Είμαι στη διάθεσή σας για οποιεσδήποτε ερωτήσεις ή διευκρινίσεις.