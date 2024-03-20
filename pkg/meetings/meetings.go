package meetings

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"example-app/pkg/google/utils"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

var calService *calendar.Service
var defaultCalendar string = "primary"
var defaultNumCalendarEvents int64 = 10

func GetCalendarEvents(token string) error {

	if calService == nil {
		tok, err := utils.GetTokenFromString(token)
		if err != nil {
			return err
		}
		ctx := context.Background()
		oauth2.NewClient(ctx, TokenSource(ctx, &tok))

		b, err := os.ReadFile("credentials.json")
		if err != nil {
			log.Fatalf("Unable to read client secret file: %v", err)
		}

		// If modifying these scopes, delete your previously saved token.json.
		config, err := google.ConfigFromJSON(b, calendar.CalendarReadonlyScope)
		if err != nil {
			log.Fatalf("Unable to parse client secret file to config: %v", err)
		}

		client := utils.GetClient(config)

		calService, err = calendar.NewService(ctx, option.WithHTTPClient(client))
		if err != nil {
			log.Fatalf("Unable to retrieve Calendar client: %v", err)
		}

	}

	t := time.Now().Format(time.RFC3339)
	events, err := calService.Events.List(defaultCalendar).ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(defaultNumCalendarEvents).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	fmt.Println("Upcoming events:")
	if len(events.Items) == 0 {
		fmt.Println("No upcoming events found.")
	} else {
		for _, item := range events.Items {
			date := item.Start.DateTime
			if date == "" {
				date = item.Start.Date
			}
			fmt.Printf("%v (%v)\n", item.Summary, date)
			if item.ConferenceData != nil {
				fmt.Printf("Conference: %+v\n", *item.ConferenceData)
			}
		}
	}
}
