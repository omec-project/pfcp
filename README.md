<!--
SPDX-FileCopyrightText: 2021 Open Networking Foundation <info@opennetworking.org>
Copyright 2019 free5GC.org

SPDX-License-Identifier: Apache-2.0

-->

# pfcp

The PFCP protocol is the interface between the Control and User Plane. 
The SMF (Session Management Function) uses the PFCP module to communicate with the UPF (User Plane Function). 
The PFCP protocol is defined in 3GPP specification 29.244.

## Supported Features

PFCP module provides following functionality,
1. The PFCP module provides APIs to send/receive PFCP messages to/from UPF.
2. The PFCP module maintains the transaction framework per UPF to map a transaction between SMF and UPF.
3. The PFCP module provides PFCP message transaction timeout and callback function in Application.
4. The PFCP module also manages the Retransmission of PFCP messages with 3 retries (not configurable).
5. The Application layer(here SMF) needs to form and send different types of PFCP messages to the PFCP module.

## Planned Features

Following features/enhancements are planned for PFCP module,
1. UE IP-Address allocation via UPF
2. Parsing new IEs based on future 3gpp releases

Note : PFCP module is used only by SMF.  ONF’s UPF does not use this pfcp module.
