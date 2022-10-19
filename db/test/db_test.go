package test

import (
	"context"
	"fmt"
	"io"
	"os"
	"plugin-demo-go/db"
	"plugin-demo-go/db/gen/item"
	"testing"
	"time"
)

func TestPut(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client := db.Connection()

	dir := "/home/amelia/Desktop/conf.json"
	f, err := os.Open(dir)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		panic("open file failed")
	}
	fileBytes, err := io.ReadAll(f)
	if err != nil {
		panic("read file failed")
	}

	client.Item.Update().
		Where(item.Name("conf")).
		SetFile(fileBytes).
		SaveX(ctx)

	client.Close()
}
func TestGet(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1*time.Second)
	client := db.Connection()

	fileBytes := client.Item.Query().
		Where(item.NameEQ("sms")).
		OnlyX(ctx)
	f, err := os.OpenFile("./sms_impl.zip",os.O_CREATE|os.O_RDWR,0766)
	if err != nil {
		panic("open file failed")
	}
	n, err := f.Write(fileBytes.File)
	if err != nil {
		panic("write file error")
	}
	println(n)
	f.Close()
	client.Close()
}