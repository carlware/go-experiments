package main

import (
	"bufio"
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/viant/afs"
	"github.com/viant/afs/option"
	"github.com/viant/afsc/gs"
	_ "github.com/viant/afsc/gs"
	goption "google.golang.org/api/option"
)

const folderName = "gs://proofs-staging/proofs"
const secretPath = "key.json"
const filePath = "cats.jpeg"

func main() {
	jwtConfig, err := gs.NewJwtConfig(option.NewLocation(secretPath))
	if err != nil {
		log.Fatal(err)
	}
	JSON, err := json.Marshal(jwtConfig)
	if err != nil {
		log.Fatal(err)
	}
	jsonAuth := goption.WithCredentialsJSON(JSON)
	opt1 := gs.NewClientOptions(jsonAuth)

	buf, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	r := bufio.NewReader(buf)

	ctx := context.Background()
	service := afs.New()
	err = service.Upload(ctx, folderName+"/cats.jpeg", 0644, r, opt1)
	if err != nil {
		log.Fatal(err)
	}
	// reader, err := service.DownloadWithURL(ctx, folderName+"/cats.jpeg", opt1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// data, err := ioutil.ReadAll(reader)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("data: %s\n", data)

	// service := afs.New()
	// ctx := context.Background()
	// objects, err := service.List(ctx, folderName, opt1)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, object := range objects {
	// 	url := object.URL()
	// 	fmt.Printf("%v %v\n", object.Name(), strings.Replace(url, "gs://", "https://storage.googleapis.com/", 1))
	// 	if object.IsDir() {
	// 		continue
	// 	}
	// }
	// err = service.Copy(ctx, folderName, "tmp")
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
