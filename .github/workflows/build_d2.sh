for file_dir in "${@}"; do
    case "$file_dir" in 
        *.d2)
        d2 --layout=elk "$file_dir";;
    esac

done
