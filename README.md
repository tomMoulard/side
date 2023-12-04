# Side

Here is the little help to understand this project !

## Run it!

To start this project, run:
```bash
$ make
```

After a while, your should have an OpenFaaS infrastructure ready to be used.
The infrastructure is composed of:
 - a kubenetes cluster (k3d)
 - a Traefik gateway
 - a Openfaas gateway ([dashboard](http://localhost/ui)))
 - a Nats queue
 - a Prometheus server
 - a MongoDB database
