package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fraudster struct {
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                  string `json:"name" bson:"name"`
	Aliases               string `json:"aliases" bson:"aliases"`
	KnownLocations        string `json:"known_locations" bson:"known_locations"`
	AssociatedPhones      string `json:"associated_phones" bson:"associated_phones"`
	AssociatedAccounts    string `json:"associated_accounts" bson:"associated_accounts"`
	AssociatedCountries   string `json:"associated_countries" bson:"associated_countries"`
	AssociatedWebsites    string `json:"associated_websites" bson:"associated_websites"`
	AssociatedSocialMedia string `json:"associated_social_media" bson:"associated_social_media"`
	Classification        string `json:"classification" bson:"classification"`
	ReporterEmail         string `json:"reporter_email" bson:"reporter_email"`
	ReporterContact       string `json:"reporter_contact" bson:"reporter_contact"`
	DateCreated           string `json:"date_created" bson:"date_created"`
	Summary               string `json:"summary" bson:"summary"`
}

func FindFraudsters(filter interface{}, params *PageSortParams) ([]Fraudster, error) {
	result := make([]Fraudster, 0)

	err := find(filter, params, &result, "fraudsters")

	return result, err
}