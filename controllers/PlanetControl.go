package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"planeta/models"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/julienschmidt/httprouter"
)

type PlanetController struct{}

func (u PlanetController) GetPlanet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	erro := models.Erro{}
	// os parametros de conexao com o DYNAMODB estao nas variaveis de ambiente.
	awsConfig := &aws.Config{}
	awsConfig.Region = aws.String(os.Getenv("AWS_DEFAULT_REGION"))
	awsConfig.Credentials = credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	awsConfig.CredentialsChainVerboseErrors = aws.Bool(true)
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}
	ddbClient := dynamodb.New(sess)
	// fmt.Println("teste1")
	planeta := p.ByName("planeta")
	if planeta == "" {

		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}
	// fmt.Println("teste2")
	// fmt.Println(planeta + "------")
	output, err := ddbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("teste"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(planeta),
			},
		},
	})
	// fmt.Println("teste3")

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	b2wPlanet := models.B2WPlanet{}
	err = dynamodbattribute.UnmarshalMap(output.Item, &b2wPlanet)
	// fmt.Println("teste4")
	// fmt.Println(output.Item)
	// fmt.Println(b2wPlanet.Nome + "teste4")

	if b2wPlanet.Nome == "" {
		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}

	url := "https://swapi.co/api/planets/"

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Cache-Control", "no-cache")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	listPlant := models.SwapiListPlanet{}

	json.Unmarshal(body, &listPlant)

	i := 0
	for i == 0 {
		for index := 0; index < len(listPlant.Results); index++ {
			if strings.ToUpper(planeta) == strings.ToUpper(listPlant.Results[index].Name) {
				b2wPlanet.QtdFilmes = len(listPlant.Results[index].Films)
				i = 1
			}
		}
		if i == 0 {
			if listPlant.Next != "" {
				url = listPlant.Next

				req, _ = http.NewRequest("GET", url, nil)
				req.Header.Add("Cache-Control", "no-cache")
				res, _ = http.DefaultClient.Do(req)

				defer res.Body.Close()
				body, _ = ioutil.ReadAll(res.Body)
				listPlant = models.SwapiListPlanet{}

				json.Unmarshal(body, &listPlant)
			} else {
				i = 1
			}
		}
	}

	// fmt.Println(res)
	// fmt.Println(string(body))

	// Marshal provided interface into JSON structure
	rj, _ := json.Marshal(b2wPlanet)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", rj)
}

func (u PlanetController) PostPlanet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

	// os parametros de conexao com o DYNAMODB estao nas variaveis de ambiente.
	awsConfig := &aws.Config{}
	awsConfig.Region = aws.String(os.Getenv("AWS_DEFAULT_REGION"))
	awsConfig.Credentials = credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	awsConfig.CredentialsChainVerboseErrors = aws.Bool(true)
	sess, err := session.NewSession(awsConfig)
	erro := models.Erro{}

	if err != nil {
		w.WriteHeader(404)
		return
	}

	ddbClient := dynamodb.New(sess)

	b2wPlanetPost := models.B2WPlanetPost{}
	b2wPlanetGet := models.B2WPlanet{}

	err = json.NewDecoder(r.Body).Decode(&b2wPlanetPost)
	fmt.Println(b2wPlanetPost)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	param := b2wPlanetPost.Nome
	if param == "" {
		erro.Detail = "Nome do planeta não informado"
		rj, _ := json.Marshal(erro)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}
	fmt.Println("1")

	// if recPut.Prods[0].Codigo != "" {

	output, err := ddbClient.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("teste"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(param),
			},
		},
	})
	fmt.Println("2")

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	err = dynamodbattribute.UnmarshalMap(output.Item, &b2wPlanetGet)
	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	av, err := dynamodbattribute.MarshalMap(b2wPlanetPost)
	if err != nil {
		w.WriteHeader(404)
		return
	}

	if b2wPlanetGet.Nome != "" {
		erro.Detail = "Nome do planeta já existente"
		rj, _ := json.Marshal(erro)
		// Write content-type, statuscode, payload
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}
	fmt.Println("3")

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("teste"),
	}
	_, err = ddbClient.PutItem(input)

	if err != nil {
		w.WriteHeader(404)
		return
	}
	w.Header().Set("Content-Type", "Text")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", "Transação OK")
}

func (u PlanetController) DeletePlanet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	erro := models.Erro{}
	// os parametros de conexao com o DYNAMODB estao nas variaveis de ambiente.
	awsConfig := &aws.Config{}
	awsConfig.Region = aws.String(os.Getenv("AWS_DEFAULT_REGION"))
	awsConfig.Credentials = credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	awsConfig.CredentialsChainVerboseErrors = aws.Bool(true)
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", rj)
		return
	}
	ddbClient := dynamodb.New(sess)
	planeta := p.ByName("planeta")
	if planeta == "" {

		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", rj)
		return
	}

	_, err = ddbClient.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String("teste"),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(planeta),
			},
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}
	w.Header().Set("Content-Type", "Text")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", "Exclusão com sucesso")
}

func (u PlanetController) ListPlanet(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	erro := models.Erro{}
	// os parametros de conexao com o DYNAMODB estao nas variaveis de ambiente.
	awsConfig := &aws.Config{}
	awsConfig.Region = aws.String(os.Getenv("AWS_DEFAULT_REGION"))
	awsConfig.Credentials = credentials.NewStaticCredentials(os.Getenv("AWS_ACCESS_KEY_ID"), os.Getenv("AWS_SECRET_ACCESS_KEY"), "")
	awsConfig.CredentialsChainVerboseErrors = aws.Bool(true)
	sess, err := session.NewSession(awsConfig)
	if err != nil {
		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "%s", rj)
		return
	}
	ddbClient := dynamodb.New(sess)

	output, err := ddbClient.Scan(&dynamodb.ScanInput{
		TableName: aws.String("teste"),
	})

	if err != nil {
		panic(fmt.Sprintf("Failed to unmarshal Record, %v", err))
	}

	b2wListPlanet := []models.B2WPlanet{}
	err = dynamodbattribute.UnmarshalListOfMaps(output.Items, &b2wListPlanet)

	if len(b2wListPlanet) == 0 {
		erro.Detail = "Not found"
		rj, _ := json.Marshal(erro)
		w.Header().Set("Content-Type", "application/json")
		//w.WriteHeader(404)
		fmt.Fprintf(w, "%s", rj)
		return
	}

	j := 0
	for i := 0; i < len(b2wListPlanet); i++ {
		j = 0

		url := "https://swapi.co/api/planets/"

		req, _ := http.NewRequest("GET", url, nil)
		req.Header.Add("Cache-Control", "no-cache")
		res, _ := http.DefaultClient.Do(req)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		listPlant := models.SwapiListPlanet{}
		json.Unmarshal(body, &listPlant)
		for j == 0 {
			for index := 0; index < len(listPlant.Results); index++ {
				fmt.Println(b2wListPlanet[i].Nome)
				fmt.Println(listPlant.Results[index].Name)
				if strings.ToUpper(b2wListPlanet[i].Nome) == strings.ToUpper(listPlant.Results[index].Name) {
					b2wListPlanet[i].QtdFilmes = len(listPlant.Results[index].Films)
					j = 1
				}
			}
			if listPlant.Next != "" {
				url = listPlant.Next

				req, _ = http.NewRequest("GET", url, nil)
				req.Header.Add("Cache-Control", "no-cache")
				res, _ = http.DefaultClient.Do(req)

				defer res.Body.Close()
				body, _ = ioutil.ReadAll(res.Body)
				listPlant = models.SwapiListPlanet{}

				json.Unmarshal(body, &listPlant)
			} else {
				j = 1
			}
		}
	}

	rj, _ := json.Marshal(b2wListPlanet)

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)

	fmt.Fprintf(w, "%s", rj)
}
