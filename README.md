# envar
A simple wrapper of github.com/joho/godotenv, for methods to get typed environment variables.
E.g. envar.Int("key"), envar.Bool("key"), or even envar.MilliSeconds("key").

# Changelog
2021-06-01 Add GetDefs method so I can log it. (Why not just log it? Because user might want to fmt.Print it, or logger.Print it, to keep it simple, I choose to just return default key/value pairs.)
