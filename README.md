# Clear Badge Cache CLI 🧹✨

[![Go Version](https://img.shields.io/github/go-mod/go-version/judahpaul16/clear-badge-cache)](https://go.dev/dl/)
[![Go Report Card](https://goreportcard.com/badge/github.com/judahpaul16/clear-badge-cache)](https://goreportcard.com/report/github.com/judahpaul16/clear-badge-cache)
[![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/judahpaul16/clear-badge-cache/go.yml)](https://github.com/judahpaul16/clear-badge-cache/actions)
[![GitHub](https://img.shields.io/github/license/judahpaul16/clear-badge-cache)](LICENSE)
[![Website](https://img.shields.io/badge/website-https://clear--badge--cache.com-blue)](https://clear-badge-cache.com/)

Clear Badge Cache is a tool designed to purge cached GitHub badge images. Based on the repository URL and badge file, it will delete the cached image from the GitHub CDN.

Inspired by [github-badge-cache-buster](https://github.com/sbts/github-badge-cache-buster) 

There is a large list of bugs open (and incorrectly closed) relating to this.
The main issue is
- [#224](https://github.com/github/markup/issues/224) - Aggressive image caching breaks image badges

Some of the others are.....
- [#9](https://github.com/rvagg/nodei.co/issues/9) - GitHub Image Proxy breaking Nodei.co images in markdown
- [#15](https://github.com/badges/pypipins/issues/15) - Badges do not update with GitHubs caching
- [#111]() - can't find the URL for this one
- [#116]() - can't find the URL for this one
- [#134](https://github.com/badges/shields/issues/134) - Serve cached images upon request timeout
- [#137](https://github.com/badges/shields/issues/137) - Travis badges seem to be broken
- [#218](https://github.com/codecov/support/issues/218) - GitHub caching affecting my badge
- [#220](https://github.com/lemurheavy/coveralls-public/issues/220) - Set `Cache-Control` or `Expires` on S3 assets
- [#257](https://github.com/drone/drone/pull/257) - Badge caching
- [#414](https://github.com/lemurheavy/coveralls-public/issues/414) - Incorrect coverage badge
- [#1970](https://github.com/travis-ci/travis-ci/issues/1970) - Build status image cached by GitHub
- [#3122](https://github.com/PowerShell/PowerShell/issues/3122) - Code Coverage Status Badge is cached
- [#3150](https://github.com/PowerShell/PowerShell/issues/3150) - Unix nightly badge is inaccurate
- [#6040](https://github.com/pouchdb/pouchdb/pull/6040) - fix coveralls badge issue in README
- [#17057](https://gitlab.com/gitlab-org/gitlab-ce/issues/17057) - Aggressive GitHub caching breaks Gitlab badge images.

More information [here](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/about-anonymized-urls).

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
  sh ./clear-badge-cache.sh
  ```
- Windows:

  ```cmd
  .\clear-badge-cache.exe
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
