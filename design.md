encoding/csv for writing out as a csv file
strconv for turning types into strings and visa versa
text/tabwriter for writing out tab aligned output
os for opening and reading files
github.com/spf13/cobra for the command line interface
github.com/mergestat/timediff for displaying relative friendly time differences (1 hour ago, 10 minutes ago, etc)


1. cobra for commands
2. mopple add "description"  -> creates a row in csv file  -- cobra, os, csv
3. mopple list -> list all rows in csv file -- tabwriter, timediff, csv, strconv
4. mopple tick <task id> -> completes the task
5. mopple delete <task id>

> any files in a single directory is a single package and must have same package name -- random russian guy
