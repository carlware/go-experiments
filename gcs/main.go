package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"google.golang.org/api/option"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const bucketName = "proofs-staging"
const folderName = "gs://proofs-staging/proofs"
const secretPath = "key.json"
const filePath = "cat2.jpg"

func main() {
	uploadFile(filePath)
}

func uploadFile(path string) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	name := randomString()
	bucket := client.Bucket(bucketName)
	w := bucket.Object("proofs/" + name).NewWriter(ctx)
	if _, err = io.Copy(w, f); err != nil {
		fmt.Printf("io.Copy: %v\n", err)
	}
	if err := w.Close(); err != nil {
		fmt.Printf("Writer.Close: %v\n", err)
	}
	fmt.Println("uploaded", name)
}

func randomString() string {
	bytes := make([]byte, 15)
	for i := 0; i < 15; i++ {
		bytes[i] = byte(randInt(65, 90))
	}
	return string(bytes)
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

// func testAFS() {
// 	jwtConfig, err := gs.NewJwtConfig(option.NewLocation(secretPath))
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	JSON, err := json.Marshal(jwtConfig)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	jsonAuth := goption.WithCredentialsJSON(JSON)
// 	opt1 := gs.NewClientOptions(jsonAuth)

// 	buf, err := os.Open(filePath)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	r := bufio.NewReader(buf)

// 	ctx := context.Background()
// 	service := afs.New()
// 	err = service.Upload(ctx, folderName+"/cats.jpeg", 0644, r, opt1)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	// reader, err := service.DownloadWithURL(ctx, folderName+"/cats.jpeg", opt1)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// data, err := ioutil.ReadAll(reader)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// fmt.Printf("data: %s\n", data)

// 	// service := afs.New()
// 	// ctx := context.Background()
// 	// objects, err := service.List(ctx, folderName, opt1)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// for _, object := range objects {
// 	// 	url := object.URL()
// 	// 	fmt.Printf("%v %v\n", object.Name(), strings.Replace(url, "gs://", "https://storage.googleapis.com/", 1))
// 	// 	if object.IsDir() {
// 	// 		continue
// 	// 	}
// 	// }
// 	// err = service.Copy(ctx, folderName, "tmp")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// }
