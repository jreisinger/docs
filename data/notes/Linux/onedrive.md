```
# Install and enable
sudo add-apt-repository ppa:yann1ck/onedrive
sudo apt-get install onedrive
systemctl --user enable onedrive
systemctl --user start onedrive

# Check status and logs
systemctl --user status onedrive
journalctl --user -u onedrive
```
