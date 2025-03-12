#!/usr/bin/env bash

abuild clear
abuild clean
abuild checksum
abuild build
abuild -r