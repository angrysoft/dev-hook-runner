set -l hooks (/usr/bin/ls .hooks.d | sed 's/\.hook\.yml//')

complete -c dhr -n "not __fish_seen_subcommand_from $hooks" -a "$hooks" -f
complete -c dhr -f -r -n "not __fish_seen_subcommand_from $hooks; and not __fish_contains_opt -s a" -s a -d "Argument to pass to the hook"
complete -c dhr -r -n "not __fish_seen_subcommand_from $hooks" -s d -a "(__fish_complete_directories)" -F -d "Directory containing hooks" 
complete -c dhr -f -n "not __fish_seen_subcommand_from $hooks; and not __fish_contains_opt -s a" -s h -d "Print help message"
complete -c dhr -f -n "not __fish_seen_subcommand_from $hooks; and not __fish_contains_opt -s a" -s l -d "List all hooks"

complete -c dhr -f -n "__fish_seen_subcommand_from $hooks"
