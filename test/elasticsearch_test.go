package test

import (
	"testing"
	"gopkg.in/olivere/elastic.v5"
	"log"
	"os"
	"context"
)

func TestElasticSearch(t *testing.T){
	file ,_ :=os.OpenFile("elastic.log", os.O_APPEND | os.O_CREATE, 777)
	logger := log.New(file, "elastic:", log.LstdFlags)
	client, err := elastic.NewClient(elastic.SetURL("http://localhost:9200"),elastic.SetErrorLog(logger))
	if err != nil{
		t.Logf("%v", err)
		return
	}
	res,err :=client.Get().Index("megacorp").Type("_all").Id("1").Do(context.Background())
	if err != nil{
		t.Logf("%v", err)
		return
	}

	//删除
	//dres, err := client.Delete().Index("megacorp").Type("employee").Id("2").Do(context.Background())
	//if err != nil{
	//	t.Logf("%v", err)
	//	return
	//}
	//t.Logf("%v", dres)

	c, _ := client.Count("megacorp", "customer").Do(context.Background())
	t.Logf("%d", c)

	s,_ := res.Source.MarshalJSON()
	t.Logf("%s", s)
}