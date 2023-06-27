# MongoDB Container Custom Liveness Probe

---

### Purpose


With the release of MongoDB 6.0 the `mongo` cli tool was deprecated in favor
of the new `mongosh` tool.  Unfortunately, the `bitnami/mongodb` container 
that we use changed to use the `mongosh` tool for liveness which when executed
continuously causes high CPU load just for the probes.  The MongoDB thread
is here: https://jira.mongodb.org/browse/MONGOSH-1240

To work around this, we've created a small golang connector which should be
able to execute very quickly and with low overhead.  This is an adaptation
of sorts to the https://github.com/syndikat7/mongodb-rust-ping tool that was
developed for the same reason but we decided to remake it because we prefer
to always statically link all of these small tools to prevent issues down
the line but the rust Mongo client used in that tool had a dependency on a
crypto library named "ring" that could not be statically linked.

The method of deployment (pulling something from github on every container
start) is also not consistent so we decided on pulling from a static image.

### Deployment
```yaml
customLivenessProbe:
  failureThreshold: 6
  initialDelaySeconds: 30
  periodSeconds: 20
  successThreshold: 1
  timeoutSeconds: 10
  exec:
    command:
      - /custom-scripts/mongodb-liveness-probe

initContainers:
  - name: install-mongodb-liveness-probe
    image: us-docker.pkg.dev/vonix-io/public/mongodb-liveness-probe:1.0
    imagePullPolicy: IfNotPresent
    command:
      - sh
      - -c
      - |
        #!/usr/bin/env bash -e
        cp /mongodb-liveness-probe /custom-scripts/
        chmod +x /custom-scripts/mongodb-liveness-probe
    volumeMounts:
      - mountPath: "/custom-scripts"
        name: mongodb-ping-volume

extraVolumeMounts:
  - name: mongodb-ping-volume
    mountPath: /custom-scripts

extraVolumes:
  - name: mongodb-ping-volume
    emptyDir:
      sizeLimit: 100Mi
```

