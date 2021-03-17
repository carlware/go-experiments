package main

import (
	"bufio"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"carlware/gcs/fs"

	"cloud.google.com/go/storage"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"

	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
)

//https://storage.cloud.google.com/files-econic-staging-private/proofs/PKFCDPEBBKMJMTR.jpeg
type Proof struct {
	File string `json:"file"`
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

const bucketName = "files-econic-staging-private"
const secretPath = "key.json"
const filePath = "cat2.jpg"

type FileLimiter struct {
	bytes  []byte
	size   int
	cursor int
}

func NewFileLimiter(size int) *FileLimiter {
	return &FileLimiter{
		bytes: make([]byte, size),
		size:  size,
	}
}

func (f *FileLimiter) Write(p []byte) (n int, err error) {
	if f.cursor > f.size {
		n = len(p)
	}
	t := len(p)
	if f.size-f.cursor < t {
		t = f.size - f.cursor
	}
	n = copy(f.bytes[f.cursor:f.cursor+t], p[0:t])
	f.cursor += n
	n = len(p)
	return
}

func (f *FileLimiter) Bytes() []byte {
	return f.bytes
}

func main() {
	data, err := os.Open("cat2.jpg")
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}
	buff := make([]byte, 512)
	data.Read(buff)
	data.Seek(0, 0)

	uploadFile(data, buff)

	// generateURL(bucketName, "proofs/PKFCDPEBBKMJMTR.jpeg",)

	// 	e := echo.New()

	// 	e.Use(middleware.Logger())
	// 	e.Use(middleware.Recover())

	// 	e.POST("/files", upload)

	// 	e.Logger.Fatal(e.Start(":1323"))
}

func generateURL(bucket, object string) (string, error) {
	// bucket := "bucket-name"
	// object := "object-name"
	// serviceAccount := "service_account.json"
	jsonKey, err := ioutil.ReadFile("key.json")
	if err != nil {
		return "", fmt.Errorf("ioutil.ReadFile: %v", err)
	}
	conf, err := google.JWTConfigFromJSON(jsonKey)
	if err != nil {
		return "", fmt.Errorf("google.JWTConfigFromJSON: %v", err)
	}

	opts := &storage.SignedURLOptions{
		Scheme:         storage.SigningSchemeV4,
		Method:         "GET",
		GoogleAccessID: conf.Email,
		PrivateKey:     conf.PrivateKey,
		Expires:        time.Now().Add(1 * time.Minute),
	}
	u, err := storage.SignedURL(bucket, object, opts)
	if err != nil {
		return "", fmt.Errorf("storage.SignedURL: %v", err)
	}

	fmt.Println("Generated GET signed URL:")
	fmt.Printf("%q\n", u)
	fmt.Println("You can use this URL with any user agent, for example:")
	fmt.Printf("curl %q\n", u)
	return u, nil
}

type FilesResponse struct {
	Files []string `json:"files"`
}

func upload(c echo.Context) error {
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	g, err := fs.New("key.json", bucketName)

	files := form.File["files"]
	uris := []string{}

	for _, file := range files {
		src, err := file.Open()
		if err != nil {
			return err
		}
		defer src.Close()

		uri, err := g.Upload(context.TODO(), src)
		if err != nil {
			return err
		}
		uris = append(uris, uri)
	}
	return c.JSON(http.StatusOK, &FilesResponse{uris})
}

func uploadFile(reader io.Reader, b []byte) (string, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	fil, err := os.Create("./temp.jpg")
	if err != nil {
		log.Fatal(err)
	}
	writerFil := bufio.NewWriter(fil)

	// newReader := bufio.NewReader(reader)
	// buf, err := newReader.Peek(512)
	// if err != nil {
	// 	return "", err
	// }
	// contentType := http.DetectContentType(buf)

	// contentType := ".jpg"
	// fmt.Println("contentType: ", contentType)

	// setting sha256 writer
	h := sha256.New()

	fmt.Println("content/type", http.DetectContentType(b))

	name := randomString()
	bucket := client.Bucket(bucketName)
	w := bucket.Object("proofs/" + name).NewWriter(ctx)

	l := NewFileLimiter(512)

	buffer := make([]byte, 32*1024)
	writer := io.MultiWriter(h, w, writerFil, l)
	written, err := io.CopyBuffer(writer, reader, buffer)
	if err != nil {
		fmt.Printf("io.Copy: %v\n", err)
		return "", err
	}
	fmt.Println("written", written)
	if err := w.Close(); err != nil {
		fmt.Printf("Writer.Close: %v\n", err)
		return "", err
	}

	contentType := http.DetectContentType(l.Bytes())

	fmt.Println("content type", contentType)
	fmt.Println("content type", w.Attrs().Name)
	fmt.Println("uploaded", name)
	fmt.Println("sha", hex.EncodeToString(h.Sum(nil)))

	return name, nil
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
