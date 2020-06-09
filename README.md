### Proxy API Covid Saúde

Este proxy serve a api do portal da saude - https://covid.saude.gov.br/

Este pequeno proxy foi criado para resolver o problema de cors que alguns serviços estão usando - https://twitter.com/breakzplatform/status/1270107471252721669

Ele sobrescreve alguns headers de requisição e resposta
### Requerimentos
- Golang ```1.14.x``` - https://golang.org/dl/

### Compilando
```bash
make build
```

### Executando com docker
```bash
docker run --rm -p 8080:8080 jhonatamenezes/covid-saude-proxy
```
Acesse - http://localhost:8080/prod/PortalMunicipio
