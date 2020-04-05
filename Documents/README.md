# event-dispatcher
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Semantic Versioning 2.0.0](https://img.shields.io/badge/semver-2.0.0-brightgreen?style=flat-square)](https://semver.org/spec/v2.0.0.html)
[![License](https://img.shields.io/github/license/Anadian/event-dispatcher)](https://github.com/Anadian/event-dispatcher/Documents/LICENSE)

> A simple, machine-local, thread-safe pub/sub event queue and dispatcher for golang.
# Table of Contents
- [Background](#Background)
- [Install](#Install)
- [Usage](#Usage)
- [API](#API)
- [Contributing](#Contributing)
- [License](#License)
# Background
I created this project because, quite bluntly, there wasn't any other project, among the [42 repos listed under the "Messaging" section of the awesome-go README](https://github.com/avelino/awesome-go#messaging), that suited my needs. First of all, most of the projects were designed to interface with some third-party API to exchange messages between clusters of machines running on cloud servers: not traditional single-process event queues. As for the few projects listed that were local event handlers, all of _them_ were either abandoned, poorly-documented, needlessly obtuse in their design, or just too limited to do anything really. So created this project with explicite design goals of it being optionally synchronous or asynchronous, optionally supporting wildcard/RegEx event matching, be completely safe regardless of whatever options are used, and to be as simple as possible while meeting all previous design goals.
# Install
# Usage
# API
# Contributing
Changes are tracked in [Documents/CHANGES.md](./CHANGES.md).
# License
MIT ©2019 Anadian

SEE LICENSE IN [Documents/LICENSE](./LICENSE)
