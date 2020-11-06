package model

import (
	"fmt"
	"log"
	"strings"

	"firebase.google.com/go/db"
)

type Antrian struct {
	Id     string `json:"id"`
	Status bool   `json:"status"`
}

func AddAntrian() (bool, error) {
	_, _, dataAntrian := GetAntrian()
	var Id string
	var antrianRef *db.Ref
	ref := client.NewRef("antrian")

	if dataAntrian == nil {
		Id = fmt.Sprintf("B-0")
		antrianRef = ref.Child("0")
	} else {
		Id = fmt.Sprintf("B-%d", len(dataAntrian))
		antrianRef = ref.Child(fmt.Sprintf("%d", len(dataAntrian)))
	}
	antrian := Antrian{
		Id:     Id,
		Status: false,
	}
	err := antrianRef.Set(ctx, antrian)
	if err != nil {
		log.Fatalln(err)
		return false, err
	}

	return true, nil
}

func GetAntrian() (bool, error, []map[string]interface{}) {
	var data []map[string]interface{}
	ref := client.NewRef("antrian")
	err := ref.Get(ctx, &data)
	if err != nil {
		return false, err, nil
	}
	return true, nil, data
}

func UpdateAntrian(idAntrian string) error {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	antrian := Antrian{
		Id:     idAntrian,
		Status: true,
	}
	err := childRef.Set(ctx, antrian)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	return nil
}

func DeleteAntrian(idAntrian string) error {
	ref := client.NewRef("antrian")
	id := strings.Split(idAntrian, "-")
	childRef := ref.Child(id[1])
	err := childRef.Delete(ctx)
	if err != nil {
		return err
	}

	return nil
}
