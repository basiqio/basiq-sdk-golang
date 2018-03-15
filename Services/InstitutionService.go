package Services

import (
	"github.com/basiqio/basiq-sdk-golang/Utilities"
	"fmt"
	"encoding/json"
)

type Institution struct {
	Id              string                 `json:"id"`
	Name            string                 `json:"name"`
	ShortName       string                 `json:"shortName"`
	Country         string                 `json:"country"`
	ServiceName     string                 `json:"serviceName"`
	ServiceType     string                 `json:"serviceType"`
	LoginIdCaption  string                 `json:"loginIdCaption"`
	PasswordCaption string                 `json:"PasswordCaption"`
	Colors          map[string]interface{} `json:"colors"`
	Logo            map[string]interface{} `json:"logo"`
}

type InstitutionsList struct {
	Count int           `json:"count"`
	Data  []Institution `json:"data"`
}

type InstitutionService struct {
	Session Session
}

func NewInstitutionService(session *Session) *InstitutionService {
	return &InstitutionService{
		Session: *session,
	}
}

func (is *InstitutionService) GetInstitutions() (InstitutionsList, *Utilities.APIError) {
	var data InstitutionsList

	body, statusCode, err := is.Session.api.Send("GET", "institutions", nil)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return data, &Utilities.APIError{Message: err.Error()}
		}

		return data, &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &Utilities.APIError{Message: err.Error()}
	}

	return data, nil
}

func (is *InstitutionService) GetInstitution(institutionId string) (Institution, *Utilities.APIError) {
	var data Institution

	body, statusCode, err := is.Session.api.Send("GET", "institutions/" + institutionId, nil)
	if err != nil {
		return data, &Utilities.APIError{Message: err.Error()}
	}
	if statusCode > 299 {
		response, err := Utilities.ParseError(body)
		if err != nil {
			return data, &Utilities.APIError{Message: err.Error()}
		}

		return data, &Utilities.APIError{
			Response:   response,
			Message:    response.GetMessages(),
			StatusCode: statusCode,
		}
	}

	if err := json.Unmarshal(body, &data); err != nil {
		fmt.Println(string(body))
		return data, &Utilities.APIError{Message: err.Error()}
	}

	return data, nil
}