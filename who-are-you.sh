#! /bin/bash
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ".[52].name"