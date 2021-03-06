# Abraxas

A BreedTV media archive bot. It looks for videos posted to Slack and then saves them to Firestore.

## Building

```docker build -t abraxas:latest .```

### Running

```docker run --env-file .env -p 3000:3000 abraxas:latest```
