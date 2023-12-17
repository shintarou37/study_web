import (
  "log"

  firebase "firebase.google.com/go"
  "google.golang.org/api/option"
)

// Use the application default credentials
ctx := context.Background()
conf := &firebase.Config{ProjectID: projectID}
app, err := firebase.NewApp(ctx, conf)
if err != nil {
  log.Fatalln(err)
}

client, err := app.Firestore(ctx)
if err != nil {
  log.Fatalln(err)
}
defer client.Close()

_, _, err := client.Collection("users").Add(ctx, map[string]interface{}{
        "first": "Ada",
        "last":  "Lovelace",
        "born":  1815,
})
if err != nil {
        log.Fatalf("Failed adding alovelace: %v", err)
}

_, _, err = client.Collection("users").Add(ctx, map[string]interface{}{
        "first":  "Alan",
        "middle": "Mathison",
        "last":   "Turing",
        "born":   1912,
})
if err != nil {
        log.Fatalf("Failed adding aturing: %v", err)
}
