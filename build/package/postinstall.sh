#!/bin/bash

chmod 711 /usr/bin/systemservice 
chmod 640 /lib/systemd/system/dm-system.service
systemctl daemon-reload
systemctl enable dm-system.service
systemctl restart dm-system.service
