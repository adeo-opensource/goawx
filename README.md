# goawx

[![All Contributors](https://img.shields.io/github/all-contributors/adeo-opensource/goawx?style=flat&label=Contributors&color=informational)](#contributors)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_goawx&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_goawx)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_goawx&metric=coverage)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_goawx)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_goawx&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_goawx)
[![Bugs](https://sonarcloud.io/api/project_badges/measure?project=adeo-opensource_goawx&metric=bugs)](https://sonarcloud.io/summary/new_code?id=adeo-opensource_goawx)

## Description

Go library for AWX (Ansible Tower/Ansible Automation Platform) manager. 
Used in particular by the [terraform-provider-awx][terraform-provider-awx] project.

## Roadmap

[Resources managed by this library](ROADMAP.md)

## AWX authentication configuration options

You must have these values in your possession to use this library:

* url
* username/password or token

## Usage

```go
go get -u github.com/adeo-opensource/goawx
```

In the directory [examples](./examples) you can find examples of use

## How to contribute?

[Learn about how to contribute](CONTRIBUTING.md)

## Where to ask Questions?

Questions can be asked in form of issues in this repository:
[Open an issue][open-issue]

## Changelog

[Learn about the latest improvements](CHANGELOG.md)

## License

Project is under Apache 2.0 license. See [License](LICENSE) file for more information.

## Contributors ✨

Thanks goes to these wonderful people ([emoji key][all-contributors-emoji-url]):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors][all-contributors-url] specification.  
Contributions of any kind welcome!


[all-contributors-url]: https://github.com/all-contributors/all-contributors
[all-contributors-emoji-url]: https://allcontributors.org/docs/en/emoji-key
[open-issue]: https://github.com/adeo-opensource/terraform-provider-awx/issues/new/choose
[terraform-provider-awx]: https://github.com/adeo-opensource/terraform-provider-awx
