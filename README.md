# Go Kafka Produce Consumer Docker

A sample for networking kafka in docker using targeted builds

## Local

The only dependency is docker and a `.env` with a EULA for lenses box

Use the make file to run the commands.

## Deploy

To deploy with heroku run:

```bash
heroku create your-app-name --manifest

git push heroku master
```

That's it! Kafka is included in the heroku.yml and everything is configured to work out of the box.
