# event-dispatcher
[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg?style=flat-square)](https://github.com/RichardLitt/standard-readme)
[![Semantic Versioning 2.0.0](https://img.shields.io/badge/semver-2.0.0-brightgreen?style=flat-square)](https://semver.org/spec/v2.0.0.html)
[![License](https://img.shields.io/github/license/Anadian/event_dispatcher)](https://github.com/Anadian/event_dispatcher/LICENSE)
[![GoDoc](https://godoc.org/github.com/Anadian/event_dispatcher/source?status.svg)](https://godoc.org/github.com/Anadian/event_dispatcher/source)
[![Build Status](https://travis-ci.org/Anadian/event_dispatcher.svg?branch=master)](https://travis-ci.org/Anadian/event_dispatcher)
[![Coverage Status](https://coveralls.io/repos/github/Anadian/event_dispatcher/badge.svg?branch=master)](https://coveralls.io/github/Anadian/event_dispatcher?branch=master)

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
Changes are tracked in [CHANGES.md](./CHANGES.md).
# License
MIT ©2019-2020 Anadian

SEE LICENSE IN [LICENSE](./LICENSE)
