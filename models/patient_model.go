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
	} `json:"personalInformation" bson:"personalInformation"`
	Location struct {
		Department string `json:"department" bson:"department"`
		Province   string `json:"province" bson:"province"`
		District   string `json:"district" bson:"district"`
		Reference  string `json:"reference" bson:"reference"`
	} `json:"location" bson:"location"`
	Cita struct {
		AppointmentDate primitive.DateTime `json:"appointmentDate" bson:"appointmentDate"`
		Specialties     []struct {
			Specialty string `json:"specialty" bson:"specialty"`
			Value     string `json:"value" bson:"value"`
		} `json:"specialties" bson:"specialties"`
		AppointmentDetail string `json:"appointmentDetail" bson:"appointmentDetail"`
	} `json:"cita" bson:"cita"`
	Question struct {
		QuestionExamRecent bool `json:"questionExamRecent" bson:"questionExamRecent"`
		SpiritualSupport   bool `json:"spiritualSupport" bson:"spiritualSupport"`
		FutureActivities   bool `json:"futureActivities" bson:"futureActivities"`
	} `json:"question" bson:"question"`
	Estate    string             `json:"estate" bson:"estate"`
	CreatedAt primitive.DateTime `json:"createdAt" bson:"createdAt"`
	UpdatedAt primitive.DateTime `json:"updatedAt" bson:"updatedAt"`
	Version   int                `json:"__v" bson:"__v"`
}
