all: stop start setup-faas traefik setup-mongo stat deploy
.PHONY: all

TIMEOUT := "8760h" # 365 * 24
FAAS_CLI=./faas-cli --gateway http://localhost

clean:
	$(RM) -f faas-cli arkade
.PHONY: clean

faas-cli:
	curl -sL https://cli.openfaas.com | sh

arkade:
	curl -SLsf https://get.arkade.dev/ | sh

start:
	k3d cluster create mycluster \
		--api-port 6550 \
		--timestamps \
		-p 80:80@loadbalancer \
		-p 8080:8080@loadbalancer \
		-p 443:443@loadbalancer \
		--k3s-arg '--no-deploy=traefik@server:0'
		-i rancher/k3s:v1.21.4-k3s1
	k3d image import traefik:v3.0 -c mycluster
.PHONY: start

setup-faas: arkade
	./arkade install openfaas \
		--basic-auth=false \
		--set gateway.upstreamTimeout=${TIMEOUT} \
		--set gateway.writeTimeout=${TIMEOUT} \
		--set gateway.readTimeout=${TIMEOUT} \
		--set faasnetes.writeTimeout=${TIMEOUT} \
		--set faasnetes.readTimeout=${TIMEOUT} \
		--set queueWorker.ackWait=${TIMEOUT} \
		--wait
.PHONY: setup-faas

setup-mongo: arkade
	./arkade install mongodb \
		--wait
	${FAAS_CLI} secret create mongo-db-username --from-literal admin
	kubectl get secret --namespace default mongodb -o jsonpath="{.data.mongodb-root-password}" | \
		base64 --decode | \
		${FAAS_CLI} secret create mongo-db-password
.PHONY: setup-faas

traefik:
	kubectl apply -f stack/
.PHONY: traefik

traefik-logs:
	kubectl logs -l app=traefik
.PHONY: traefik-logs

stat:
	kubectl get all --all-namespaces
.PHONY: stat

deploy: faas-cli
	${FAAS_CLI} up -f side.yml
.PHONY: deploy

logs: faas-cli
	${FAAS_CLI} logs webui
.PHONY: logs

test:
	go test -race ./...

stop:
	k3d cluster delete mycluster
.PHONY: stop
