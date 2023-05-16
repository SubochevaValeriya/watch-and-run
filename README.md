# watch-and-run
Console application that allows monitor changes in various directories and execute a standard set of console commands

- In configs/config.yaml you can set the directories to watch, the ability to enable / disable the file by regexp mask and log file, where will the logs from executing commands be sent.
- If one of the commands fails, the following will not be executed.
- The history of file changes and launches is stored in the database in tables Event and Launch:
- When the application stops, all commands stop (Graceful Shutdown) and the text appears: "finished".

### To run an app:

```
make build && make run
```

If you are running the application for the first time, you must apply the migrations to the database:

```
make migrate
```
