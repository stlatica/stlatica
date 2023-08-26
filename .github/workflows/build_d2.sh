for file_dir in "${@}"; do
    case "$file_dir" in 
        *.d2)
        d2 --theme=300 --dark-theme=200 -l elk --pad 0 "$file_dir";;
    esac

done
