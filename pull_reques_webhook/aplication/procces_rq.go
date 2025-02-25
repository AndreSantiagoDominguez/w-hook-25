package aplication

import (
	"encoding/json"
	"log"
	"web-hook/pull_reques_webhook/domain"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload domain.PullRequestEventPayload

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403
	}

	if (eventPayload.Action == "closed") {
		log.Printf("Se hace el pull hacia la rama: %s \n", eventPayload.PullRequest.Base.Ref)
		log.Printf("El pull viene de la rama:  %s \n", eventPayload.PullRequest.Head.Ref)
		log.Printf("Usuario: %s \n", eventPayload.PullRequest.User.Login)
		log.Printf("Nombre del repositorio: %s \n", eventPayload.Repository.FullName)
	}

	return 200
}