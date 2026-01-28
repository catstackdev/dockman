# Your current setup ✅

```bash
go mod init github.com/catstackdev/dockman
go get -u github.com/spf13/cobra@latest

# Add these:
go get -u github.com/spf13/viper # Config management
go get -u github.com/fatih/color # Colored output go get -u gopkg.in/yaml.v3 #
YAML parsing go get -u github.com/docker/docker/client # Docker SDK (optional)


```

# Verify installation (auto unused remove)

```bash
go mod tidy
```

# Build

```
go build -o dockman .
```

#Install Globally (Optional)

````
# Install to $GOPATH/bin
go install

# Now you can use it anywhere:
dockman up
dockman logs -f
dockman down```
````

# for completion

```
# 1. Create the directory if it doesn't exist
mkdir -p ~/.zsh/completion
# 2. Save the completion script to a file
dockman completion zsh > ~/.zsh/_dockman
# 3. Add these lines to your ~/.zshrc file
echo 'fpath=(~/.zsh $fpath)' >> ~/.zshrc
echo 'autoload -Uz compinit && compinit' >> ~/.zshrc
```
