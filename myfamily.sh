# ! /bin/bash
# export HERO_ID=70
curl -s https://learn.reboot01.com/assets/superhero/all.json | jq ".[] | select(.id==$HERO_ID)" | jq ".connections.relatives" | tr -d '"'