# Game Recommender

A recommender for what game to play - add your favourites and search through them to find the one that best fits your requirements!

## Getting Started

You will need Go and Docker.

To run the database, in the docker folder, use:

```
docker compose up -d
```

To run the API, in the api folder, use:

```
make api
```

## Functionality

The recommender allows you to add games to a database, update their details and also delete games. You can obtain a list of all games in the database, and search for games that best match your requirements via age rating, number of players, category and genre.
