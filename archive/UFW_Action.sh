#!/bin/bash
#Script to block IP ranges read from a txt file with UFW

#Koen Kumps 13/02/2019

#UFW action to apply deny|reject
UFW_ACTION_TYPE=deny

#Show usage
show_usage()
{
    echo "Usage:"
    echo "UFW_action add|remove <filename>"
    echo "UFW_action -h shows this message"
}

echo_upon_error()
{
    echo $1
    exit 1
}

#If file with IP's exist continue, else quit with error
check_file()
{
    if [ ! -f $IP_FILE ]; then
        echo_upon_error "IP list-file not found!!"
    fi
}

#Apply UFW action for each line in file
apply_ufw_action()
{
    COUNT=1
    while IFS= read -r ip_address; do
        ufw insert 1 $UFW_ACTION_TYPE from $ip_address to any &> /dev/null
        echo -ne  "$COUNT rule(s) inserted."\\r
        COUNT=$[COUNT + 1]
    done < "$IP_FILE"
    echo "$COUNT rule(s) inserted."
    echo "Done"
}

#Remove UFW action for each line in file
remove_ufw_action()
{
    COUNT=1
    while IFS= read -r ip_address; do
        ufw delete $UFW_ACTION_TYPE from $ip_address &> /dev/null
        echo -ne "$COUNT rules(s) removed."\\r
        COUNT=$[COUNT + 1]
    done < "$IP_FILE"
    echo "$COUNT rule(s) removed."
    echo "Done"
}

#Main script doing everything it needs to do
alter_ufw_actions()
{
    OPTIND=1;
    while getopts ":h" opt; do
        case $opt in
            h) show_usage;;
            *) show_usage;;
        esac
    done
    shift $(($OPTIND - 1))
    if [ $ADD_REMOVE == "add" ]; then
        apply_ufw_action
    elif [ $ADD_REMOVE == "remove" ]; then
        remove_ufw_action
    else echo_upon_error "Invalid action provided, try add|remove"
    fi
}

#Eval
if [ $# == 0 ] ; then
        show_usage
else
    ADD_REMOVE=$1
    IP_FILE=$2
    #check_file $IP_FILE
    check_file $2
    alter_ufw_actions $@
fi
exit 0
