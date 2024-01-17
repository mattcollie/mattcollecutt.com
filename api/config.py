import toml

config: dict = toml.load("config.toml")


if __name__ == '__main__':
    print(config)
