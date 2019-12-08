package test

import (
	"testing"
	"rsc.io/pdf"
	"fmt"
	"github.com/google/go-tika/tika"
	"context"
	"os"
	"net/http"
	pdfs "github.com/unidoc/unipdf/v3/model"
	"github.com/unidoc/unipdf/v3/extractor"
	"image/jpeg"
	"archive/zip"
)

func TestOpenPdf(t *testing.T)  {
	//file, err := pdf.Open("E:/go/src/ReadRpc/pages.pdf")
	file, err := pdf.Open("C:/Users/msbox/Desktop/[深入理解计算机系统(原书第2版)]_清晰带标签.pdf")
	if err != nil {
		panic(err)
	}
	for i:=2;i <= 50; i++  {
		fmt.Println(file.Page(i).V.Name())
	}

}

func TestTikaPdf(t *testing.T)  {
	//server,_ := tika.NewServer(" E:/go/src/ReadRpc/tika-server.jar", "7010")
	server,_ := tika.NewServer("E:/go/src/ReadRpc/tika-server-1.16.jar", "7010")
	fmt.Println(server.URL())
	server.Start(context.Background())
}

func TestTikaParse(t *testing.T)  {
	client := tika.NewClient(http.DefaultClient, "http://localhost:7010")
	file, _ := os.Open("C:/Users/msbox/Desktop/程序员教程（第4版）_13652157.pdf")
	defer file.Close()
	str,err :=client.Parse(context.Background(), file)
	if err!=nil {
		panic(err)
	}
	fmt.Println(str)
}

func TestDownloadTika(t *testing.T)  {
	err := tika.DownloadServer(context.Background(), "1.20", "tika-server-1.16.jar")
	if err != nil {
		t.Error(err)
	}
}

func TestUnidocImages(t *testing.T)  {
	f, err := os.Open("C:/Users/msbox/Desktop/程序员教程（第4版）_13652157.pdf")
	if err != nil {
		panic(err)
	}

	pdfReader, err := pdfs.NewPdfReader(f)

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		panic(err)
	}
	fmt.Printf("PDF Num Pages: %d\n", numPages)


	// Prepare output archive.
	zipf, err := os.Create("image.zip")
	if err != nil {
		panic(err)
	}

	defer zipf.Close()
	zipw := zip.NewWriter(zipf)

	totalImages := 0
	for i := 0; i < numPages; i++ {
		fmt.Printf("-----\nPage %d:\n", i+1)

		page, err := pdfReader.GetPage(i + 1)
		if err != nil {
			panic(err)
		}


		pextract, err := extractor.New(page)
		if err != nil {
			panic(err)
		}

		pimages, err := pextract.ExtractPageImages(nil)
		if err != nil {
			panic(err)
		}

		fmt.Printf("%d Images\n", len(pimages.Images))
		for idx, img := range pimages.Images {
			fmt.Printf("Image %d - X: %.2f Y: %.2f, Width: %.2f, Height: %.2f\n",
				totalImages+idx+1, img.X, img.Y, img.Width, img.Height)
			fname := fmt.Sprintf("p%d_%d.jpg", i+1, idx)

			gimg, err := img.Image.ToGoImage()
			if err != nil {
				panic(err)
			}

			imgf, err := zipw.Create(fname)
			if err != nil {
				panic(err)
			}
			opt := jpeg.Options{Quality: 100}
			err = jpeg.Encode(imgf, gimg, &opt)
			if err != nil {
				panic(err)
			}
		}
		totalImages += len(pimages.Images)
	}
	fmt.Printf("Total: %d images\n", totalImages)

	// Make sure to check the error on Close.
	err = zipw.Close()
	if err != nil {
		panic(err)
	}

}

func TestUnidocText(t *testing.T)  {
	f, err := os.Open("C:/Users/msbox/Desktop/程序员教程（第4版）_13652157.pdf")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	pdfReader, err := pdfs.NewPdfReader(f)
	if err != nil {
		panic(err)
	}

	numPages, err := pdfReader.GetNumPages()
	if err != nil {
		panic(err)
	}

	fmt.Printf("--------------------\n")
	fmt.Printf("PDF to text extraction:\n")
	fmt.Printf("--------------------\n")
	for i := 0; i < numPages; i++ {
		pageNum := i + 1

		page, err := pdfReader.GetPage(pageNum)
		if err != nil {
			panic(err)
		}

		ex, err := extractor.New(page)
		if err != nil {
			panic(err)
		}

		text, err := ex.ExtractText()
		if err != nil {
			panic(err)
		}

		fmt.Println("------------------------------")
		fmt.Printf("Page %d:\n", pageNum)
		fmt.Printf("\"%s\"\n", text)
		fmt.Println("------------------------------")
	}
}