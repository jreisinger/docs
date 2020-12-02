## Install on Linux

```
# Remove system/package npm
sudo apt-get remove npm
sudo apt autoremove

# Install node.js (includes npm)
curl -sL https://deb.nodesource.com/setup_15.x | sudo -E bash -
sudo apt-get install -y nodejs

# Install Angular CLI tool
sudo npm install -g @angular/cli

# Start new Angular app
ng new my-app
cd my-app
ng serve --open
```

[source](https://angular.io/guide/setup-local)
