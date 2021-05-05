# Fix Python CA certificates

```
export ALL_CA_CERTIFICATES="/usr/local/share/ca-certificates/cacert.pem"

# NOTE: we are appending so don't run these commands multiple times
sudo bash -c "cat $(python3 -m certifi) >> $ALL_CA_CERTIFICATES"
sudo bash -c "cat your-ca.crt >> $ALL_CA_CERTIFICATES"

# Put this into ~/.bashrc
export REQUESTS_CA_BUNDLE=$ALL_CA_CERTIFICATES
```

# Take screenshots (printscreens)

## Whole screen

1. Command (⌘) + Shift + 3

## Part of screen

1. Command (⌘) + Shift + 4
1. hold down Control and make your selection
1. Command (⌘) +  V

# Enable key repeats (like in VSCodium)

Run this in terminal and then restart the given application:

```
defaults write NSGlobalDomain ApplePressAndHoldEnabled -bool false
```

# Mount NFS

```
sudo mount -o nolocks -t nfs 192.168.100.100:/srv/nfs/public ~/nfs
```

# Fix Home and End keyboard keys

```
sudo -i 
mkdir -p ~/Library/KeyBindings ; cd ~/Library/KeyBindings
vim DefaultKeyBinding.dict
```

```
{
/* Remap Home / End keys */
/* Home Button*/
"\UF729" = "moveToBeginningOfLine:"; 
/* End Button */
"\UF72B" = "moveToEndOfLine:"; 
/* Shift + Home Button */
"$\UF729" = "moveToBeginningOfLineAndModifySelection:"; 
/* Shift + End Button */
"$\UF72B" = "moveToEndOfLineAndModifySelection:"; 
/* Ctrl + Home Button */
"^\UF729" = "moveToBeginningOfDocument:"; 
/* Ctrl + End Button */
"^\UF72B" = "moveToEndOfDocument:"; 
 /* Shift + Ctrl + Home Button */
"$^\UF729" = "moveToBeginningOfDocumentAndModifySelection:";
/* Shift + Ctrl + End Button*/
"$^\UF72B" = "moveToEndOfDocumentAndModifySelection:"; 
}
```

Restart MacBook.

Source: https://medium.com/@elhayefrat/how-to-fix-the-home-and-end-buttons-for-an-external-keyboard-in-mac-4da773a0d3a2
