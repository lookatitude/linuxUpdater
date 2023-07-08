#!/bin/bash

# Define color codes
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

# Function to log user activity
log_user_activity() {
    if [[ $2 == "error" ]]
    then
        echo -e "[`date`] - ${RED}$1${NC}" | tee -a user_activity.log
    else
        echo -e "[`date`] - ${GREEN}$1${NC}" | tee -a user_activity.log
    fi
}

# Function to update apt
update_apt() {
    echo "Updating apt..."
    log_user_activity "Started updating apt."

    sudo apt-get update -y

    if [ $? -eq 0 ]
    then
        log_user_activity "Updated apt successfully."
    else
        log_user_activity "Failed to update apt." "error"
        exit 1
    fi
}

# Function to upgrade apt
upgrade_apt() {
    echo "Upgrading apt..."
    log_user_activity "Started upgrading apt."

    sudo apt-get upgrade -y

    if [ $? -eq 0 ]
    then
        log_user_activity "Upgraded apt successfully."
    else
        log_user_activity "Failed to upgrade apt." "error"
        exit 1
    fi
}

# Function to check and upgrade distribution
dist_upgrade_apt() {
    echo "Checking for distribution upgrade..."
    log_user_activity "Started checking for distribution upgrade."

    sudo apt-get dist-upgrade --dry-run | grep '^Inst' 

    if [ $? -eq 0 ]
    then
        log_user_activity "New distribution version available."
        echo "New distribution version available. Do you want to upgrade? (y/n)"
        read answer
        if echo "$answer" | grep -iq "^y" ;then
            sudo apt-get dist-upgrade -y
            if [ $? -eq 0 ]
            then
                log_user_activity "Distribution upgraded successfully."
            else
                log_user_activity "Failed to upgrade distribution." "error"
                exit 1
            fi
        else
            log_user_activity "Distribution upgrade skipped by user."
        fi
    else
        log_user_activity "No new distribution version available."
    fi
}

# Function to clean apt
clean_apt() {
    echo "Cleaning apt..."
    log_user_activity "Started cleaning apt."

    sudo apt-get autoremove -y
    sudo apt-get autoclean -y

    if [ $? -eq 0 ]
    then
        log_user_activity "Cleaned apt successfully."
    else
        log_user_activity "Failed to clean apt." "error"
        exit 1
    fi
}

# Function to install brew
install_brew() {
    echo "Installing brew..."
    log_user_activity "Started installing brew."

    /bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"

    if [ $? -eq 0 ]
    then
        log_user_activity "Brew installed successfully."
    else
        log_user_activity "Failed to install brew." "error"
        exit 1
    fi
}

# Function to check brew
check_brew() {
    echo "Checking if brew is installed..."
    log_user_activity "Started checking if brew is installed."

    brew --version 2> /dev/null

    if [ $? -ne 0 ]
    then
        log_user_activity "Brew not found."
        echo "Brew not found. Do you want to install it? (y/n)"
        read answer
        if echo "$answer" | grep -iq "^y" ;then
            install_brew
        else
            log_user_activity "Brew installation skipped by user."
        fi
    else
        log_user_activity "Brew is already installed."
    fi
}

# Function to update brew
update_brew() {
    echo "Updating brew..."
    log_user_activity "Started updating brew."

    brew update

    if [ $? -eq 0 ]
    then
        log_user_activity "Updated brew successfully."
    else
        log_user_activity "Failed to update brew." "error"
        exit 1
    fi
}

# Function to upgrade brew
upgrade_brew() {
    echo "Upgrading brew..."
    log_user_activity "Started upgrading brew."

    brew upgrade

    if [ $? -eq 0 ]
    then
        log_user_activity "Upgraded brew successfully."
    else
        log_user_activity "Failed to upgrade brew." "error"
        exit 1
    fi
}

# Function to clean brew
clean_brew() {
    echo "Cleaning brew..."
    log_user_activity "Started cleaning brew."

    brew cleanup

    if [ $? -eq 0 ]
    then
        log_user_activity "Cleaned brew successfully."
    else
        log_user_activity "Failed to clean brew." "error"
        exit 1
    fi
}

# Function to diagnose brew
diagnose_brew() {
    echo "Running brew doctor..."
    log_user_activity "Started running brew doctor."

    brew doctor

    if [ $? -eq 0 ]
    then
        log_user_activity "Brew Doctor ran successfully."
    else
        log_user_activity "Failed to run brew doctor." "error"
        exit 1
    fi
}

# Main function
main() {
    update_apt
    upgrade_apt

    if [ "$1" != "--skip-distro-upgrade" ]; then
        dist_upgrade_apt
    fi

    clean_apt

    if [ "$1" != "--skip-brew-install" ]; then
        check_brew
    fi

    if [ "$1" != "--skip-brew-update" ]; then
        update_brew
    fi

    upgrade_brew
    clean_brew
    diagnose_brew
}

main $1
