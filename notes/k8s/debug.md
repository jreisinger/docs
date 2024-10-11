```
k get pods                                  # READY, STATUS, RESTARTS
k describe pod $POD                         # Events
k logs $POD
k exec -it nginx -- /bin/sh                 # ps aux ; netstat ; ...
k debug -it $POD --image=busybox \          # See Ephemeral Containers
  --target $CONTAINER                       # in k describe $POD.
k debug -it $POD --image=alpine \
  --share-processes --copy-to=$POD-debug    # now install your stuff
```
