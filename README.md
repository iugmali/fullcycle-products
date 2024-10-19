# fullcycle-products

Projeto que partiu do estudo de arquitetura hexagonal (ports and adapters) do [curso Fullcycle 3.0](https://fullcycle.com.br).
O objetivo do projeto é implementar uma API Rest e um CLI para gerenciar produtos. 

## Rodando o projeto via Docker

```bash
docker run --name products -e PORT=9001 -p 9001:9001 -d iugmali/products-app
```

O servidor REST ficará disponível em [http://localhost:9001](http://localhost:9001). Os comandos CLI também podem ser executados com o comando abaixo
```bash
docker exec -it products main create --name "Produto 1" --price 10000
```

## Fazendo o Build e rodando por conta própria

### Variáveis de ambiente (inserir em arquivo .env na raiz do projeto)
- PORT: Porta em que o servidor REST irá rodar

### build
Comando utilizado para compilar o projeto
```bash
go build -o bin/products github.com/iugmali/fullcycle-products
```

### Servidor REST
Comando em bash utilizado para rodar o servidor em background na porta 9000 e salvar os logs em server.log
```bash
nohup ./bin/products serve 0<&- &> server.log &
```

### CLI commands
Comando utilizado para criar um produto com nome "Produto 1" e preço de $100
```bash
./bin/products create --name "Produto 1" --price 10000
```
Comando utilizado para buscar o produto de id {id}
```bash
./bin/products get --id {id}
```
Comando utilizado para ativar o produto de id {id}
```bash
./bin/products enable --id {id}
```
Comando utilizado para atualizar o preço do produto de id {id}
```bash
./bin/products setprice --id {id} --price 1000
```
Comando utilizado para desativar o produto de id {id}
```bash
./bin/products disable --id {id}
```

### API Endpoints
|       | Endpoint                     | Ação                                                                                                                  |
|-------|------------------------------|-----------------------------------------------------------------------------------------------------------------------|
| GET   | /product                     | Retorna todos os produtos                                                                                             |
| POST  | /product                     | Cria um novo produto. O corpo deve ser enviado em json e deve ter as propriedades "name" (string) e "price" (integer) |
| GET   | /product/:id                 | Retorna o produto referente a id                                                                                      |
| PATCH | /product/:id/enable          | Habilita o produto caso o price seja maior que zero                                                                   |
| PATCH | /product/:id/disable         | Desabilita o produto caso o price seja igual a zero                                                                   |
| PATCH | /product/:id/setprice/:price | Atualiza o valor do produto baseado no valor numérico informado na url                                                |

## Processo de desenvolvimento

### Iniciando projeto
```bash
go mod init github.com/iugmali/fullcycle-products
```

### Definindo interfaces

- Definição de interfaces para os casos de uso de produto
```go
type ProductInterface interface {
	IsValid() (bool, error)
	Enable() error
	Disable() error
	GetID() string
	GetName() string
	GetStatus() string
	GetPrice() int64
	SetPrice(price int64) error
}

type ProductServiceInterface interface {
	GetAll() ([]ProductInterface, error)
	Get(id string) (ProductInterface, error)
	SetPrice(product ProductInterface, price int64) (ProductInterface, error)
	Create(name string, price int64) (ProductInterface, error)
	Enable(product ProductInterface) (ProductInterface, error)
	Disable(product ProductInterface) (ProductInterface, error)
}

type ProductReaderInterface interface {
	GetAll() ([]ProductInterface, error)
	Get(id string) (ProductInterface, error)
}

type ProductWriterInterface interface {
	Save(product ProductInterface) (ProductInterface, error)
}

type ProductPersistenceInterface interface {
	ProductReaderInterface
	ProductWriterInterface
}
```

### Instalando mockgen para testes

```bash
go install go.uber.org/mock/mockgen@latest
```

- Gerando mocks de interfaces de product.go para testes
```bash
mockgen -destination=application/mocks/product.go -source=application/product.go application
```

### Instalando spf13/cobra-cli para CLI
```bash
go install github.com/spf13/cobra-cli@latest
```

- Comando para gerar boilerplate para CLI
```bash
cobra-cli init
```
