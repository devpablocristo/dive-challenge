#!/bin/sh

# Load environment variables from the .env file
if [ -f .env ]; then
  export $(cat .env | grep -v ^# | xargs)
fi

# Function to log messages
log() {
  echo "ENTRYPOINT: $1"
}

# Function to build the server binary
buildServer() {
  log "Building server binary"
  go build -gcflags "all=-N -l" -buildvcs=false -o "/app/bin/$APP_NAME" "/app/cmd/main.go"
  # Verificar si el archivo binario se ha creado y es ejecutable
  if [ -f "/app/bin/$APP_NAME" ]; then
    log "Binary file created successfully"
    chmod +x "/app/bin/$APP_NAME"
  else
    log "Failed to create binary file"
    exit 1
  fi
}

# Function to run the server in normal mode
runServer() {
  log "Run server in normal mode"

  log "Killing old server"
  pkill -f dlv || true
  pkill -f "/app/bin/$APP_NAME" || true

  log "Starting server in normal mode"
  "/app/bin/$APP_NAME" &
}

# Function to run the server in debug mode
runServerDebug() {
  log "Run server in debug mode"

  log "Killing old server"
  pkill -f dlv || true
  pkill -f "/app/bin/$APP_NAME" || true

  log "Starting server with delve"
  dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec "/app/bin/$APP_NAME"
}

# Function to rebuild and rerun the server
rerunServer() {
  log "Rerun server"
  buildServer
  runServer
}

# Function to rebuild and rerun the server in debug mode
rerunServerDebug() {
  log "Rerun server in debug mode"
  buildServer
  runServerDebug
}

# Function to monitor file changes and trigger server restart
liveReloading() {
  log "Run liveReloading"
  inotifywait -e modify,delete,move -m -r --format '%w%f' --exclude '.*(\.tmp|\.swp)$' /app | (
    while read file; do
      if [[ "$file" == *.go ]]; then
        log "File $file changed. Reloading..."
        if [ "$DEBUG" = "true" ]; then
          rerunServerDebug
        else
          rerunServer
        fi
      fi
    done
  )
}

# Function to initialize the file change logger
initializeFileChangeLogger() {
  echo "" > /tmp/filechanges.log
  tail -f /tmp/filechanges.log &
}

# Main function to orchestrate the process
main() {
  initializeFileChangeLogger
  buildServer
  if [ "$DEBUG" = "true" ]; then
    runServerDebug
  else
    runServer
  fi
  liveReloading
}

# Call the main function to start the process
main
