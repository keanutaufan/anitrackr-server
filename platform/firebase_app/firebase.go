package firebase_app

import (
	"context"
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func NewFirebaseClient(ctx context.Context, credentialFile string) (*FirebaseClient, error) {
	opt := option.WithCredentialsFile(credentialFile)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	auth, err := app.Auth(ctx)
	if err != nil {
		return nil, err
	}

	return &FirebaseClient{
		App:  app,
		Auth: auth,
	}, nil
}
