# REST
Requisitos:

- A API deve ser REST
- Para cada planeta, os seguintes dados devem ser obtidos do banco de dados da aplicação, sendo inserido manualmente:
- Nome
- Clima
- Terreno

- Para cada planeta também devemos ter a quantidade de aparições em filmes, que podem ser obtidas pela API pública do Star Wars:  https://swapi.co/

Funcionalidades desejadas: 

- Adicionar um planeta (com nome, clima e terreno)
- Listar planetas
- Buscar por nome
- Remover planeta

## Build Setup

``` bash
# build for production
go build

# Start project
nohup ./planeta &

```

## Local ##

- xpto

## ENV_VARIABLES ##

Conexão com DynamoDB

- export AWS_ACCESS_KEY_ID=xpto
- export AWS_SECRET_ACCESS_KEY=xpto
- export AWS_DEFAULT_REGION=us-east-1

## Serviços ##

GET

- R.GET("/teste/v1/planets/:planeta", p.GetPlanet)

```
/teste/v1/planets/{name}
```

POST

- R.POST("/teste/v1/planets/", p.PostPlanet)

```
/teste/v1/planets/
```

- Exemplo Request : Body

```
{
    "ID": "Alderaan",
    "clima": "Tropical",
    "terreno": "plano"
}
```

DELETE

- R.DELETE("/teste/v1/planets/:planeta", p.DeletePlanet)

```
/teste/v1/planets/{name}
```


GET (LIST)

- R.GET("/teste/v1/planets/:planeta", p.GetPlanet)

```
/teste/v1/planets/
```

 
