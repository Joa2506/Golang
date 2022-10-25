package main

import "log"


func main () {
	const ArticleURL string = "https://storage.googleapis.com/aller-structure-task/articles.json"
	const ContentMarketingURL string = "https://storage.googleapis.com/aller-structure-task/contentmarketing.json"
	j := running(ArticleURL, ContentMarketingURL)
	log.Println(string(j))
}

