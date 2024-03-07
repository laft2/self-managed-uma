#!/usr/bin/bash
set -euxo pipefail

HOST_IPADDR=$(hostname -I | awk '{print $1}')

{
    echo "C_IPADDR=$HOST_IPADDR"
    echo "C_PORT=10001"
    echo "RS_IPADDR=$HOST_IPADDR"
    echo "RS_PORT=10002"
    echo "QS_IPADDR=$HOST_IPADDR"
    echo "QS_PORT=10003"
} >.env_eg
