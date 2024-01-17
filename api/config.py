import toml

config: dict = toml.load("api/config.toml")


if __name__ == '__main__':
    print(config)
