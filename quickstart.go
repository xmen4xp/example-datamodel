package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

/* // Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
} */

func GetCalendarHandle(body io.ReadCloser) (*calendar.Service, error) {

	// Get access token.
	tok := &oauth2.Token{}
	err := json.NewDecoder(body).Decode(tok)
	if err != nil {
		return nil, err
	}

	// Create a calendar service handle.
	calServer, err := calendar.NewService(context.Background(), option.WithHTTPClient(config.Client(context.Background(), tok)))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}
	return calServer, nil

}

func calendarEventGetTitle(event *calendar.Event) string {
	if len(event.Summary) > 0 {
		return event.Summary
	}

	return "[ Title Unknown ]"
}

func calendarEventGetMeetingLink(event *calendar.Event) string {
	if len(event.Location) > 0 {
		return event.Location
	}

	// Get meeting link from description
	return "[ Meeting Link Unknown ]"
}

func GetCalendarEvents(w http.ResponseWriter, r *http.Request) {

	calService, err := GetCalendarHandle(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t := time.Now().Format(time.RFC3339)
	events, err := calService.Events.List("primary").ShowDeleted(false).SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	type eventsResponse struct {
		Title       string `json:"event_title"`
		MeetingLink string `json:"event_meeting_link"`
	}
	var response []eventsResponse

	for _, item := range events.Items {
		response = append(response, eventsResponse{
			Title:       calendarEventGetTitle(item),
			MeetingLink: calendarEventGetMeetingLink(item),
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			fmt.Printf("\nItem: %+v\n", *item)
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("Summary: %v (%v)\n", item.Summary, date)
			if item.ConferenceData != nil {
				fmt.Printf("Conference data: %+v\n", *item.ConferenceData)
			}
		}
	}

}

var config *oauth2.Config

func initCalendarService() {
	b, err := os.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err = google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
}

func main() {

	initCalendarService()

	http.HandleFunc("/events", GetCalendarEvents)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
