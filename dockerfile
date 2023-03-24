# docker file is a file, which contains instructions to run docker image
FROM golang:1.20-alpine
WORKDIR /ds-stats-bot

COPY . .

RUN go mod download
#COPY .env

# ENV DISCORD_BOT_TOKEN=MTA2NDIwMDEwMDAyNjUxOTY3NQ.Gcow87.9qnU9JIVEUvAPLDj6-7W56IOPuWHdkgLIiBwKM
# ENV DATABASE_DSN="user=postgres dbname=ds-stats password=Matwyenko1_ host=localhost port=5432 sslmode=disable"
# docker run \
# -e DISCORD_BOT_TOKEN=MTA2NDIwMDEwMDAyNjUxOTY3NQ.Gcow87.9qnU9JIVEUvAPLDj6-7W56IOPuWHdkgLIiBwKM \
# -e DATABASE_DSN="user=postgres dbname=ds-stats password=Matwyenko1_ host=localhost port=5432 sslmode=disable"

EXPOSE 8080

CMD ["go", "run", "main.go"]
