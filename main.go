package main

import (
	"fmt"
	rss "github.com/jteeuwen/go-pkg-rss"
	"os"
	"./lib"
)

/*
http://www.infobae.com/rss
http://www.pagina12.com.ar/usuarios/rss.php
http://stackoverflow.com/questions/6002619/unmarshal-an-iso-8859-1-xml-input-in-go
*/

var pendingItems chan *rss.Item

func main() {

	fmt.Println("Start")

	feeds := []string{
		//"http://cdn01.ib.infobae.com/adjuntos/162/rss/politica.xml",
		"http://cdn01.ib.infobae.com/adjuntos/162/rss/Infobae.xml",
		//"http://cdn01.ib.infobae.com/adjuntos/162/rss/economia.xml",
		//"http://cdn01.ib.infobae.com/adjuntos/162/rss/sociedad.xml",
		//"http://cdn01.ib.infobae.com/adjuntos/162/rss/finanzas.xml",
		//"http://cdn01.ib.infobae.com/adjuntos/162/rss/policiales.xml",
		"http://www.pagina12.com.ar/diario/rss/principal.xml",
		//"http://www.pagina12.com.ar/diario/rss/ultimas_noticias.xml",
	}
	
	pendingItems = make(chan *rss.Item, 50)
	
	PullFeeds(feeds)
	
	for item := range pendingItems {
        
        fmt.Println("\t",item.Title)
		fmt.Println("\t","Links")
		
		for j := range item.Links {
			link := item.Links[j]
			fmt.Println("\t\t", link.Href)
		}
		
		fmt.Println("\t","Categories")
		
		for k := range item.Categories {
			category := item.Categories[k]
			fmt.Println("\t\t", category.Domain, category.Text)
		}
			
			
    }

}

func PullFeeds(feeds []string) {
	feed := rss.New(5, true, chanHandler, itemHandler)
	
	for f := range feeds {
		
		uri := feeds[f]
		
		go Pull(feed, uri)
	}
}

func Pull(feed *rss.Feed, uri string) {
	if err := feed.Fetch(uri, lib.CharsetReader); err != nil {
		fmt.Fprintf(os.Stderr, "[e] %s: %s", uri, err)
		fmt.Println(err)
		return
	}
}

func chanHandler(feed *rss.Feed, newchannels []*rss.Channel) {
	fmt.Printf("%d new channel(s) in %s\n", len(newchannels), feed.Url)
}

func itemHandler(feed *rss.Feed, ch *rss.Channel, newitems []*rss.Item) {
	fmt.Printf("%d new item(s) in %s\n", len(newitems), feed.Url)
	
	for i := range newitems {
		
		 pendingItems <- newitems[i]
			
	}
}





