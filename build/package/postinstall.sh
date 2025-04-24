#!/bin/bash
# Copyright © Siemens 2020 - 2025. ALL RIGHTS RESERVED.
# Licensed under the MIT license
# See LICENSE file in the top-level directory

chmod 711 /usr/bin/systemservice 
chmod 640 /lib/systemd/system/dm-system.service
systemctl daemon-reload
systemctl enable dm-system.service
systemctl restart dm-system.service
