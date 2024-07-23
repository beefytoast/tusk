# tusk
Task manager for project file written in GO utilising git as sync tool.

## Features

- Initialize a new project
- Add and list issues
- Add and list tasks in issues
- Log time for tasks
- Switch active issue

## Usage

```sh
# Initialize a new project
cd /path/to/newProject
tusk init

# Add a new issue
tusk issue --add --desc "Issue description"

# List all issues
tusk issue --list

# Switch active issue
tusk issue --switch <hash_issue>

# Add a new task to the active issue
tusk track --add --desc "Task description"

# Start time logging for the last task in the active issue
tusk track --start

# Stop time logging for the last task in the active issue
tusk track --stop
```

## Instalation
**Clone repository**
```bash
git clone https://github.com/yourusername/tusk.git
```
**Navigate to the project directory:**
```bash
cd tusk
```
**Build the application:**
```bash
go build -o tusk ./cmd/tusk
```
**Move the binary to a directory in your PATH:**
```bash
mv tusk /usr/local/bin/
```

## Contributing
If you have any suggestions or improvements, feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.

