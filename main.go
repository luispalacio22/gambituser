package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/luispalacio22/gambituser/awsgo"
	"github.com/luispalacio22/gambituser/bd"
	"github.com/luispalacio22/gambituser/models"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicioAWS()
	if !ValidoParametros() {
		fmt.Println("Error en los parametors. Debe enviar 'SecretName'")
		err := errors.New("error en los parametros debe enviar secretName")
		return event, err
	}
	var datos models.SignUp

	for row, att := range event.Request.UserAttributes {
		switch row {
		case "email":
			datos.UserEmail = att
			fmt.Println("Email = " + datos.UserEmail)
		case "sub":
			datos.UserUUID = att
			fmt.Println("Sub = " + datos.UserUUID)
		}

	}
	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("error al leer el secret " + err.Error())
		return event, err
	}
	return event, nil
	err = bd.SignUp(datos)
	return event, err
}

func ValidoParametros() bool {
	var traeParametro bool
	_, traeParametro = os.LookupEnv("SecretName")
	return traeParametro
}
