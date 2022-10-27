# cachos_prometheus_exporter

> Exporta metricas do varnishtop para o Prometheus

## Como funciona?

Executa o comando `varnishtop -1` e fornece algumas dessas métricas pelo
endpoint `/metrics`

## Instalação e Execução

```sh
cd cachos_prometheus_exporter
make build
./cachos_prometheus_exporter -port 9132 -timeout 10
```

## Desenvolvimento

```sh
make install
make test
make lint
make run
make build
```
