# watch-and-run
Console application that allows monitor changes in various directories and execute a standard set of console commands

- In configs/config.yaml you can set the directories to watch, regexp masks for files and log file, where will the logs from executing commands be sent.
- If one of the commands fails, the following will not be executed.
- The history of file changes and launches is stored in the database in tables Event and Launch:
![Image alt](https://github.com/SubochevaValeriya/watch-and-run/blob/main/internal/files/event_table.png)
![Image alt](https://github.com/SubochevaValeriya/watch-and-run/blob/main/internal/files/launch_table.png)

- When the application stops, all commands stop (Graceful Shutdown) and the text appears: "finished".

### To run an app:

Fill in the part with Path_and_commands in configs/config.yaml.

```
make build && make run
```

If you are running the application for the first time, you must apply the migrations to the database:

```
make migrate
```

Also when db is up, you can run the application by executing the command: go run cmd/watch-and-run/main.go.
