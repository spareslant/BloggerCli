package main

/*
https://pkg.go.dev/google.golang.org/api/blogger/v3?tab=doc
https://godoc.org/google.golang.org/api/option
https://pkg.go.dev/golang.org/x/oauth2/google?tab=doc
https://pkg.go.dev/golang.org/x/oauth2?tab=doc#Config.AuthCodeURL
*/

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"

	"google.golang.org/api/blogger/v3"
	"google.golang.org/api/option"
)

func checkErrorAndExit(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func printNoOfPosts(postsService *blogger.PostsService, BlogID string) {
	postListCall := postsService.List(BlogID).MaxResults(500)
	postList, err := postListCall.Do()
	checkErrorAndExit(err)
	fmt.Printf("Total posts in this blog: %d\n", len(postList.Items))
}

func getBloggerClient() *blogger.Service {
	ctx := context.Background()

	contents, err := ioutil.ReadFile("OAuth2.0_secret.json")
	checkErrorAndExit(err)

	conf, err := google.ConfigFromJSON(contents, blogger.BloggerScope)
	checkErrorAndExit(err)

	url := conf.AuthCodeURL("state")
	fmt.Println("Visit the following URL for the auth dialog:")
	fmt.Printf("\n%v\n", url)

	var code string
	fmt.Printf("\nEnter the code obtained from above here: ")
	fmt.Scanln(&code)

	token, err := conf.Exchange(oauth2.NoContext, code)
	checkErrorAndExit(err)

	bloggerService, err := blogger.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
	checkErrorAndExit(err)

	return bloggerService
}

func checkMandatoryArguments(args []*string) {

	allArgsExists := true
	for _, val := range args {
		if len(*val) == 0 {
			allArgsExists = allArgsExists && false
		} else {
			allArgsExists = allArgsExists && true
		}
	}

	if !allArgsExists {
		flag.PrintDefaults()
		os.Exit(2)
	}
}

func createNewPost(title *string, labels *string, fileToUpload *string) *blogger.Post {

	labelsSlice := strings.Split(*labels, ",")
	contents, err := ioutil.ReadFile(*fileToUpload)
	checkErrorAndExit(err)

	newPost := blogger.Post{}

	if len(*labels) == 0 {
		newPost = blogger.Post{
			Content: string(contents),
			Title:   *title,
		}
	} else {
		newPost = blogger.Post{
			Content: string(contents),
			Labels:  labelsSlice,
			Title:   *title,
		}
	}
	return &newPost
}

func insertAPost(title *string, labels *string, fileToUpload *string, BlogID string, postsService *blogger.PostsService) {
	fmt.Println("============= Inserting a new Post ===============")

	newPost := createNewPost(title, labels, fileToUpload)

	postInsertCall := postsService.Insert(BlogID, newPost)
	post, err := postInsertCall.Do()
	checkErrorAndExit(err)
	fmt.Println("A new post added to Blog with following details")
	fmt.Printf("Title: %s\nPostID: %s\n", post.Title, post.Id)
	fmt.Println("=================================================")
	printNoOfPosts(postsService, BlogID)
}

func updateAPost(title *string, labels *string, fileToUpload *string, BlogID string, postsService *blogger.PostsService, postID string) {
	fmt.Println("============= updating a Post ===============")

	// Search post from Title if postID not specified.
	if len(postID) == 0 {
		postsSearchCall := postsService.Search(BlogID, *title)
		postsList, err := postsSearchCall.Do()
		checkErrorAndExit(err)
		if len(postsList.Items) > 1 {
			fmt.Printf("Following multiple posts found for title = %s\n", *title)
			for _, post := range postsList.Items {
				fmt.Println(post.Title)
			}
			log.Fatal("\n\nERROR: Not updating post. Title must identify a single post only. Try specifying PostID")
		} else {
			postID = postsList.Items[0].Id
		}
	}

	newPost := createNewPost(title, labels, fileToUpload)
	postsPatchCall := postsService.Patch(BlogID, postID, newPost)
	post, err := postsPatchCall.Do()
	checkErrorAndExit(err)
	fmt.Println("Following post has been updated in Blog with following details")
	fmt.Printf("Title: %s\nPostID: %s\n", post.Title, post.Id)
	fmt.Println("=================================================")
	printNoOfPosts(postsService, BlogID)
}

func main() {

	fileToUpload := flag.String("f", "", "html file to upload (mandatory)")
	title := flag.String("t", "", "Post title (mandatory)")
	blogid := flag.String("b", "", "Blod ID (mandatory)")
	labels := flag.String("l", "", "comma separated list of labels (optional)")
	modify := flag.Bool("m", false, "Modify a post contents (optional: modifies post in mentioned in -t )")
	postID := flag.String("p", "", "post ID (optional: To be specified when updating a post)")
	flag.Parse()

	if flag.NFlag() == 0 {
		flag.PrintDefaults()
		os.Exit(1)
	}

	mandatoryArgs := []*string{fileToUpload, title, blogid}
	checkMandatoryArguments(mandatoryArgs)

	BlogID := *blogid
	PostID := *postID

	bloggerService := getBloggerClient()
	postsService := bloggerService.Posts

	printNoOfPosts(postsService, BlogID)

	if !*modify {
		insertAPost(title, labels, fileToUpload, BlogID, postsService)
	} else {
		updateAPost(title, labels, fileToUpload, BlogID, postsService, PostID)
	}
}
