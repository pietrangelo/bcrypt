# bcrypt

## An easy tool to hashing passwords using bcrypt algorithm

bcrypt is a simple library designed to do just one thing: hashing passwords with the bcrypt function.

## Features

- Hash password passed with the `-p` or `--password` argument.
- The cost factor can be changed within `-c` or `--cost` argument.

```bash
Usage of ./bcrypt:
  -c, --cost string       Cost factor (default "12")
  -p, --password string   Password to hash
```

## installation

To install and run this project, follow these steps:

1. Clone the repository:
    ```shell
    git clone https://github.com/pietrangelo/bcrypt.git
    cd bcrypt
    ```

2. Build the project:
    ```shell
    make
    ```

3. Run the project
    ```shell
    ./bcrypt -p [your-strong-password]
    ```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
