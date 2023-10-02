# twitchGO
Helpers for Twitch API integration in Golang
To use: create .env file in same directory

```
CLIENT_ID= <your_client_id>
CLIENT_SECRET=<your_client_secret>
```

# Description
Helpers for twitch API connections
- env.go takes CLIENT_ID and CLIENT_SECRET from a .env file in the same directory
- oauth.go returns an Oauth access token based on .env file
- user.go returns functions related to the user

Main currently for mainly:
1. Getting the profile picture of the username of your choice
2. Getting login of username of choice
3. Getting the created date of username of choice

