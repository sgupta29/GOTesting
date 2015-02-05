package newsletter

import (
        "reflect"
        "testing"

        "appengine/mail"
)

func TestComposeNewsletter(t *testing.T) {
        want := &mail.Message{
                Sender:  "newsletter@appspot.com",
                To:      []string{"User <user@example.com>"},
                Subject: "Weekly App Engine Update",
                Body:    "Don't forget to test your code!",
        }
        if msg := composeNewsletter(); !reflect.DeepEqual(msg, want) {
                t.Errorf("composeMessage() = %+v, want %+v", msg, want)
        }
}