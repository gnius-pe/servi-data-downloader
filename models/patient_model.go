package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Patient struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	PersonalInformation struct {
		Name                 string             `json:"name" bson:"name"`
		LastName             string             `json:"lastName" bson:"lastName"`
		NumberIdentification string             `json:"numberIdentification" bson:"numberIdentification"`
		Email                string             `json:"email" bson:"email"`
		FirstNumberPhone     string             `json:"firstNumberPhone" bson:"firtsNumberPhone"`
		SecondNumberPhone    string             `json:"secondNumberPhone" bson:"secondNumberPhone"`
		Sexo                 string             `json:"sexo" bson:"sexo"`
		BirthDate            primitive.DateTime `json:"birthDate" bson:"birthDate"`
		ID                   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	} `json:"personalInformation" bson:"personalInformation"`
	Location struct {
		Department string             `json:"department" bson:"department"`
		Province   string             `json:"province" bson:"province"`
		District   string             `json:"district" bson:"district"`
		Reference  string             `json:"reference" bson:"reference"`
		ID         primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	} `json:"location" bson:"location"`
	Cita struct {
		AppointmentDate   primitive.DateTime `json:"appointmentDate" bson:"appointmentDate"`
		Specialties       []Specialty        `json:"specialties" bson:"specialties"`
		AppointmentDetail string             `json:"appointmentDetail" bson:"appointmentDetail"`
		ID                primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	} `json:"cita" bson:"cita"`
	Question struct {
		QuestionExamRecent bool               `json:"questionExamRecent" bson:"questionExamRecent"`
		SpiritualSupport   bool               `json:"spiritualSupport" bson:"spiritualSupport"`
		FutureActivities   bool               `json:"futureActivities" bson:"futureActivities"`
		ID                 primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	} `json:"question" bson:"question"`
	Estate     string             `json:"estate" bson:"estate"`
	NumberFile int                `json:"numberFile" bson:"numberFile"`
	CreatedAt  primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt  primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	Version    int                `json:"__v" bson:"__v"`
}

type Specialty struct {
	Label string             `json:"label" bson:"label"`
	Value string             `json:"value" bson:"value"`
	ID    primitive.ObjectID `json:"id" bson:"_id,omitempty"`
}
