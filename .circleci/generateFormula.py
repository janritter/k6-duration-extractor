#!/usr/bin/env/python
from jinja2 import Environment, FileSystemLoader
import os

envVersion = os.environ['VERSION']
envFileHashMac = os.environ['FILE_HASH_MAC']
envFileHashLinux = os.environ['FILE_HASH_LINUX']

file_loader = FileSystemLoader('templates')
env = Environment(loader=file_loader)

template = env.get_template('k6-duration-extractor.rb.j2')

output = template.stream(
    version=envVersion, sha256Mac=envFileHashMac, sha256Linux=envFileHashLinux).dump('../HomebrewFormula/k6-duration-extractor.rb')
