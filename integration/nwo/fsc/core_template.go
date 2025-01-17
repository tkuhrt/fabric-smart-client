/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fsc

const CoreTemplate = `---
logging:
 spec: {{ Topology.Logging.Spec }}
 format: {{ Topology.Logging.Format }}
fsc:
  id: {{ Peer.ID }}
  networkId: {{ Registry.NetworkID }}
  address: 127.0.0.1:{{ .NodePort Peer "Listen" }}
  addressAutoDetect: true
  listenAddress: 127.0.0.1:{{ .NodePort Peer "Listen" }}
  identity:
    cert:
      file: {{ .NodeLocalCertPath Peer }}
    key:
      file: {{ .NodeLocalPrivateKeyPath Peer }}
  tls:
    enabled:  true
    clientAuthRequired: {{ .ClientAuthRequired }}
    cert:
      file: {{ .NodeLocalTLSDir Peer }}/server.crt
    key:
      file: {{ .NodeLocalTLSDir Peer }}/server.key
    clientCert:
      file: {{ .NodeLocalTLSDir Peer }}/server.crt
    clientKey:
      file: {{ .NodeLocalTLSDir Peer }}/server.key
    rootcert:
      file: {{ .NodeLocalTLSDir Peer }}/ca.crt
    clientRootCAs:
      files:
      - {{ .NodeLocalTLSDir Peer }}/ca.crt
    rootCertFile: {{ .CACertsBundlePath }}
  keepalive:
    minInterval: 60s
    interval: 300s
    timeout: 600s
  p2p:
    listenAddress: /ip4/127.0.0.1/tcp/{{ .NodePort Peer "P2P" }}
    bootstrapNode: {{ .BootstrapNode Peer }}
  kvs:
    persistence:
      type: badger
      opts:
        path: {{ NodeKVSPath }}

{{ range Extensions }}
{{.}}
{{- end }}
`
