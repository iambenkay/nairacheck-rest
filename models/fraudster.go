package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Fraudster struct {
	ID                    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name                  string             `json:"name" bson:"name,omitempty" query:"name"`
	Aliases               string             `json:"aliases" bson:"aliases,omitempty" query:"aliases"`
	KnownLocations        string             `json:"known_locations" bson:"known_locations,omitempty" query:"known_locations"`
	AssociatedPhones      string             `json:"associated_phones" bson:"associated_phones,omitempty" query:"associated_phones"`
	AssociatedAccounts    string             `json:"associated_accounts" bson:"associated_accounts,omitempty" query:"associated_accounts"`
	AssociatedCountries   string             `json:"associated_countries" bson:"associated_countries,omitempty" query:"associated_countries"`
	AssociatedWebsites    string             `json:"associated_websites" bson:"associated_websites,omitempty" query:"associated_websites"`
	AssociatedSocialMedia string             `json:"associated_social_media" bson:"associated_social_media,omitempty" query:"associated_social_media"`
	Classification        string             `json:"classification" bson:"classification,omitempty" query:"classification"`
	ReporterEmail         string             `json:"reporter_email" bson:"reporter_email,omitempty" query:"reporter_email"`
	ReporterContact       string             `json:"reporter_contact" bson:"reporter_contact,omitempty" query:"reporter_contact"`
	DateCreated           string             `json:"date_created" bson:"date_created,omitempty" query:"date_created"`
	Summary               string             `json:"summary" bson:"summary,omitempty" query:"summary"`
}

func FindOneFraudster(filter interface{}) (*Fraudster, error) {
	var result = new(Fraudster)

	err := findOne(filter, result, "fraudsters")

	zero := Fraudster{}

	if *result == zero {
		return nil, err

	}
	return result, err
}

func FindFraudsters(filter interface{}, params *PageSortParams) ([]Fraudster, error) {
	result := make([]Fraudster, 0)

	err := find(filter, params, &result, "fraudsters")
	return result, err
}

func AddFraudster(fraudster Fraudster) (Fraudster, error) {
	id, err := insertOne(fraudster, "fraudsters")
	fraudster.ID = id.(primitive.ObjectID)
	return fraudster, err
}
