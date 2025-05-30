package firebase_app

import (
	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
)

type FirebaseClient struct {
	App  *firebase.App
	Auth *auth.Client
}
