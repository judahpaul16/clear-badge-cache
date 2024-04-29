# Clear Badge Cache 🧹✨

[![Go Version](https://img.shields.io/github/go-mod/go-version/judahpaul16/clear-badge-cache)](https://go.dev/dl/)
[![GitHub](https://img.shields.io/github/license/judahpaul16/clear-badge-cache)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/judahpaul16/clear-badge-cache)](https://goreportcard.com/report/github.com/judahpaul16/clear-badge-cache)

Clear Badge Cache is a tool designed to purge cached GitHub badge images. Based on the repository URL and badge file, it will delete the cached image from the GitHub CDN.

Read this article to learn more about the [GitHub badge cache problem](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/about-anonymized-urls).

## Quick Start 🚀

### Prerequisites 📋

- Go (1.22.2 or later)

*Don't want to install Go? No worries; you can use the pre-built binaries in the `dist` directory using the same arguments as the Go program.*

### Installation 🛠

1. Clone the repository:

   ```bash
   git clone
    ```
2. Navigate to the project directory:

   ```bash
   cd clear-badge-cache
   ```
3. Run the go file directly with or without arguments:

   ```bash
   go run main.go [repoURL] [badgeFile]
   ```

   If no arguments are provided, the program will prompt you to enter the repository URL.
️

### Building 🔨

Run the appropriate build script from the `build-scripts` directory for your operating system:

- For Linux or MacOS, open your terminal and run:

  ```bash
  ./build-scripts/build.sh
  ```

- For Windows, run the following in Command Prompt or PowerShell:

  ```cmd
  .\build-scripts\build.bat
  ```

This will compile the source code and produce an executable in the `dist` directory.

### Running 🏃

To run Clear Badge Cache, navigate to the `dist` directory and execute the binary:

- Linux/MacOS:

  ```bash
  ./clear-badge-cache-linux.sh
  ```

- MacOS (alternative):

  ```bash
  ./clear-badge-cache-mac.sh
  ```

- Windows:

  ```cmd
  .\clear-badge-cache-win.exe
  ```

#### Command Line Arguments 🖊️

You can provide the repository URL and optionally a badge file as command-line arguments:

```bash
clear-badge-cache [repoURL] [badgeFile]
```

If no arguments are provided, the program will prompt you to enter the repository URL.

#### Help

To display the help message, use the `-h` or `--help` flag:

```bash
clear-badge-cache -h
```

The help message is as follows:

```
usage: clear-badge-cache [ -h | --help | [ repoURL [ badgeFile ] ] ]

       repoURL   required unless provided on command line
       badgeFile defaults to ''
```

## Contributing 🤝

Contributions are what make the open-source community great. Any contributions you make are **much appreciated**.

## License 📝

Distributed under the GNU GPL-3 License. See `LICENSE` for more information.
