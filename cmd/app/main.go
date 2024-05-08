package main

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/axelgeorg/ImersaoGo_03022023/internal/infra/akafka"
	"github.com/axelgeorg/ImersaoGo_03022023/internal/infra/repository"
	"github.com/axelgeorg/ImersaoGo_03022023/internal/infra/web"
	"github.com/axelgeorg/ImersaoGo_03022023/internal/usecase"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/go-chi/chi/v5"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(host.docker.internal:3306)/products")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repository.NewProductRepositoryMysql(db)
	createProductUseCase := usecase.NewCreateProductUseCase(repository)
	listProductsUseCase := usecase.NewListProductsUseCase(repository)

	productHandlers := web.NewProductHandlers(createProductUseCase, listProductsUseCase)

	r := chi.NewRouter()
	r.Post("/products", productHandlers.CreateProductHandler)
	r.Get("/products", productHandlers.ListProductsHandler)

	go http.ListenAndServe(":8000", r)

	msgChan := make(chan *kafka.Message)
	go akafka.Consume([]string{"product"}, "host.docker.internal:9092", msgChan)

	for msg := range msgChan {
		dtoInput := usecase.CreateProductInputDto{}
		err := json.Unmarshal(msg.Value, &dtoInput)
		if err != nil {
			continue
			//logar erro
		}

		_, err = createProductUseCase.Execute(dtoInput)
	}
}
