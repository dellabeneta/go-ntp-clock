## NTP Clock

![Pipeline](https://github.com/dellabeneta/go-ntp-clock/actions/workflows/main.yaml/badge.svg)

Um serviço web para consulta de hora via NTP (Network Time Protocol), desenvolvido em Go. Consulte rapidamente a hora atual usando uma interface web simples e responsiva.

### Aviso Importante

**Este serviço é apenas para fins educacionais e de teste.** Não utilize para fins ilegais. O projeto foi criado para aprendizado e demonstração de tecnologias.

### Funcionalidades

- Consulta da hora atual via NTP (pool.ntp.br)
- Interface web responsiva para desktop e mobile
- Exibição da hora atual em formato legível
- Backend em Go com endpoint `/time`
- Pronto para rodar em Docker e Kubernetes
- CI/CD com GitHub Actions

### Começando

**Pré-requisitos**
- Go 1.24.3 ou superior
- Docker (opcional)
- Kubernetes/k3s (opcional)

**Instalação Local**
```bash
# Clone o repositório
git clone https://github.com/yourusername/go-ntp-clock.git
cd go-ntp-clock

# Execute a aplicação
go run main.go
```

A aplicação estará disponível em `http://localhost:8080`

**Usando Docker**
```bash
# Construa a imagem
docker build -t ntp-clock .

# Execute o container
docker run -p 8080:8080 ntp-clock
```

**Deploy no Kubernetes**
```bash
# Crie o namespace
kubectl apply -f k3s/namespace.yaml

# Deploy da aplicação
kubectl apply -f k3s/deployment.yaml
kubectl apply -f k3s/service.yaml
```

A aplicação estará disponível na porta `30084` do cluster.

### Tecnologias

- **Backend**: Go 1.24.3
- **Frontend**: HTML5, CSS3
- **Container**: Docker Alpine
- **Orquestração**: Kubernetes/k3s
- **Protocolo**: NTP

### Como Funciona

O serviço implementa a seguinte lógica:

1. **Recebe a solicitação**: O usuário acessa a interface web.
2. **Consulta NTP**: O backend consulta um servidor NTP para obter a hora atual.
3. **Resposta**: A hora atual é retornada em JSON e exibida na interface.

### Configuração

O serviço pode ser configurado através das seguintes variáveis de ambiente:

| Variável | Descrição         | Padrão |
|----------|-------------------|--------|
| PORT     | Porta do servidor | 8080   |

### Estrutura do Projeto
```
della@ubuntu:~/projetos/go-ntp-clock$ tree
.
├── Dockerfile
├── go.mod
├── go.sum
├── icon.png
├── k3s
│   ├── deployment.yaml
│   ├── namespace.yaml
│   └── service.yaml
├── main.go
├── nuke.sh
├── README.md
└── static
    ├── favicon.ico
    └── index.html

3 directories, 13 files
```

### Scripts Úteis

**nuke.sh**: Script para limpeza completa do Docker (containers, imagens, volumes e redes)

```bash
chmod +x nuke.sh
./nuke.sh
```

### Licença

Este projeto está sob a licença MIT. Veja o arquivo [LICENSE](LICENSE)
