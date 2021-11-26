#!/bin/bash
#Script to generate/restore before.rules for UFW
#It reads CIDR adresslist from (.zone) files and inserts them as rules
#Exitcodes:     0 Ok, nothing special
#               1 Error
#               2 New rules applied
#               3 Rules restored

#Koen Kumps February 2019

#Environment specific parameters
#before.rules file location
PATH_TO_RULES="/etc/ufw"
#CIDR list(s) location. Store files with .zone extension
#Aggregated zone files are automatically downloaded from www.ipdeny.com
#Files should only contain a list of CIDR formatted IP addresses
#Comments in zone files should be made with "#"
PATH_TO_LISTS="/root/ban_ip/ip_zones"
#.ex.zone file is an exclusion file with a list of CIDR which should be excluded
#Countries to block
COUNTRIES="af ar \
           ba bb bd bo br bw \
           ca cm cn co \
           do \
           ec eg et \
           fr \
           gm \
           id il in it \
           jp \
           ke kn kp kr kz \
           lb lt \
           ma mn mw mx my \
           na no np \
           pe ph pk pl \
           ro ru \
           sc sg sk \
           th tk tn tw \
           ua us \
           ve vn \
           za"

########################################
#Don't change anything below these lines
#Version:
VERSION="1.0b - IPV4 Only"

#Functions
#Show usage
show_usage()
{
    echo "Usage:"
    echo "UFW_CreateRules [-c|-d|-g|-r|-v|-h]"
    echo "UFW_action -h shows help message"
    exit 0
}

#Show extensive usage
show_usage_long()
{
    echo "Usage:"
    echo "UFW_CreateRules [-c|-d|-g|-r|-v|-h]"
    echo "-c Create new before.rules file and apply to UFW, more or less silently"
    echo "-d Dry-run/Debug, creates test.out file in directory where script is run, gives more feedback"
    echo "-g Get new zone files from www.ipdeny.com"
    echo "-r Revert back to previous Firewall rules"
    echo "-v Version"
    echo "-h Shows this message"
    exit 0
}

#Show version
show_version()
{
    echo "UFW Create Rules v: $VERSION"
    echo "Created by Koen Kumps"
    exit 0
}

#Prepare environment
prepare_env()
{
    mkdir -p $PATH_TO_LISTS >/dev/null 2>&1
    if [ -n "$GET_FILES" ]; then
        #Download aggregated zone files and extract
        wget -O $PATH_TO_LISTS/all-zones.tar.gz http://www.ipdeny.com/ipblocks/data/countries/all-zones.tar.gz >/dev/null 2>&1
        tar -zxf $PATH_TO_LISTS/all-zones.tar.gz -C $PATH_TO_LISTS
        #Download Github IP adresses used for webhooks
        curl -H "Accept: application/vnd.github.v3+json" https://api.github.com/meta | jq .hooks[] | sed 's/"//g' > $PATH_TO_LISTS/.ex.zone
    fi
}

#Revert to previous rules
restore_rules()
{
    cp $PATH_TO_RULES/before.rules.1 $PATH_TO_RULES/before.rules
    echo "Rules restored to previous version,"
}

#Echo to shell upon error
echo_upon_error()
{
    echo $1
    exit 1
}

#Exit script with predefined exitcode
exit_script()
{
    exit $1
}

#Rotate before.rules (backup previous version)
rotate_rules()
{
    if ([ "$DRYRUN" -eq 1 ] && [ "$CREATE" -eq 0 ]); then
        echo "Doing dry-run!!"
        echo "Rules will not be applied, firewall will not be reloaded."
    else
        if [ -f $PATH_TO_RULES/before.rules.5.gz ]; then
            rm $PATH_TO_RULES/before.rules.5.gz
        fi
        if [ -f $PATH_TO_RULES/before.rules.4.gz ]; then
            mv $PATH_TO_RULES/before.rules.4.gz $PATH_TO_RULES/before.rules.5.gz
        fi
        if [ -f $PATH_TO_RULES/before.rules.3.gz ]; then
            mv $PATH_TO_RULES/before.rules.3.gz $PATH_TO_RULES/before.rules.4.gz
        fi
        if [ -f $PATH_TO_RULES/before.rules.2.gz ]; then
            mv $PATH_TO_RULES/before.rules.2.gz $PATH_TO_RULES/before.rules.3.gz
        fi
        if [ -f $PATH_TO_RULES/before.rules.1 ]; then
            gzip $PATH_TO_RULES/before.rules.1
            mv $PATH_TO_RULES/before.rules.1.gz $PATH_TO_RULES/before.rules.2.gz
        fi
        cp $PATH_TO_RULES/before.rules $PATH_TO_RULES/before.rules.1
    fi
}
#Create new before.rules file
create_rules()
{
    #Check if there are files at specified location
    if [ -n "$(ls -A $PATH_TO_LISTS)" ]; then
        if ([ "$DRYRUN" -eq 1 ] && [ "$CREATE" -eq 0 ]); then
            echo "Zone files found in $PATH_TO_LISTS"
            LINES=$(sed -n '0,/# End required lines/p' $PATH_TO_RULES/before.rules | wc -l)
            head -n $LINES $PATH_TO_RULES/before.rules > test.out
            echo "" >> test.out
            echo "################" >> test.out
            echo "# Country blocking list" >> test.out
            echo "################" >> test.out
            COUNT_FILES=0
            for country in $COUNTRIES; do
                ZONE_FILE="$PATH_TO_LISTS/$country.zone"
                echo "################" >> test.out
                echo "#" >> test.out
                echo "# $ZONE_FILE" >> test.out
                echo "#" >> test.out
                echo "################" >> test.out
                while IFS= read -r ip_address; do
                    #skip lines in *.zone files starting  with "#"
                    #continue breaks out iteration
                    case $ip_address in
                            \#*) continue ;;
                    esac
                    #skip lines where ip is in .ex.zone file
                    #continue breaks out iteration
                    if ! grep -Fxq $ip_address $PATH_TO_LISTS/.ex.zone
                        echo "-A ufw-before-input -s $ip_address -j DROP" >> test.out
                    fi
                done < $ZONE_FILE
                COUNT_FILES=$[COUNT_FILES + 1]
                echo -ne "$COUNT_FILES CIDR file(s) processed"\\r
            done
            echo -ne "$COUNT_FILES CIDR file(s) processed"\\r
            echo "" >> test.out
            sed -n -e '/# allow all on loopback/,$p' $PATH_TO_RULES/before.rules >> test.out
            echo ""
            echo "test.out created"
        else
            LINES=$(sed -n '0,/# End required lines/p' $PATH_TO_RULES/before.rules.1 | wc -l)
            head -n $LINES $PATH_TO_RULES/before.rules.1 > $PATH_TO_RULES/before.rules
            echo "" >> $PATH_TO_RULES/before.rules
            echo "################" >> $PATH_TO_RULES/before.rules
            echo "# Country blocking list" >> $PATH_TO_RULES/before.rules
            echo "################" >> $PATH_TO_RULES/before.rules
            echo "" >> $PATH_TO_RULES/before.rules
            for country in $COUNTRIES; do
                ZONE_FILE="$PATH_TO_LISTS/$country.zone"
                echo "################" >> $PATH_TO_RULES/before.rules
                echo "#" >> $PATH_TO_RULES/before.rules
                echo "# $ZONE_FILE" >> $PATH_TO_RULES/before.rules
                echo "#" >> $PATH_TO_RULES/before.rules
                echo "################" >> $PATH_TO_RULES/before.rules
                while IFS= read -r ip_address; do
                    #skip lines in *.zone files starting with "#"
                    #continue breaks out iteration
                    case $ip_address in
                            \#*) continue ;;
                    esac
                    #skip lines where ip is in .ex.zone file
                    #continue breaks out iteration
                    if ! grep -Fxq $ip_address $PATH_TO_LISTS/.ex.zone
                    echo "-A ufw-before-input -s $ip_address -j DROP" >> $PATH_TO_RULES/before.rules
                    fi
                done < $ZONE_FILE
            done
            echo "" >> $PATH_TO_RULES/before.rules
            sed -n -e '/# allow all on loopback/,$p' $PATH_TO_RULES/before.rules.1 >> $PATH_TO_RULES/before.rules
            echo "New rules applied,"
        fi
    else
            echo_upon_error "No .zone files found at $PATH_TO_LISTS!!"
    fi
}

#Reload UFW to apply new rules
reload_ufw()
{
    if ([ "$DRYRUN" -eq 1 ] && [ "$CREATE" -eq 0 ]); then
        echo "Dry-run finished, validate test.out before running:"
        echo "UFW_CreateRules -c"
    else
        ufw reload
    fi
}

#Main script doing everything it needs to do
main()
{
    OPTIND=1;
    while getopts ":cdgrvh" opt; do
        case $opt in
            c) CREATE=1; DRYRUN=0;;
            d) DRYRUN=1; CREATE=0;;
            g) GET_FILES=1;;
            r) RESTORE=1;;
            v) show_version;;
            h) show_usage_long;;
            *) show_usage;;
        esac
    done
    shift $(($OPTIND - 1))
    prepare_env
    if [ -n "$RESTORE" ]; then
        restore_rules
        EXIT_CODE=3
    else
        rotate_rules
        create_rules
        EXIT_CODE=2
    fi
    reload_ufw
    exit_script $EXIT_CODE
}

#Eval
if [ $# == 0 ] ; then
    show_usage
else
    main $@
fi
exit 0
