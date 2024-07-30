# prar

`prar` is a program to request review in GitHub pull requests using users defined in a file.

## Usage

To use you should create a file with the users you want to request review.
The file should be located in ~/.config/prar and have the following format:

```json
{
  "repository": [
    "user1"
  ],
  "another-repository": [
    "user4",
    "user"
  ]
}
```

Then you can run the program with the following command:

```bash
prar <repository>
```

Note: you should be in the <repository> directory.

## Install

```bash
sudo GOBIN=/usr/local/bin/ go install
```
