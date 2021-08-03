# k-diff
Kubernetes Resource Diff Tool using helm-diff


# Build
```
go build -o kdiff main.go
```

# Use
```
kdiff --old=old.yaml --new=new.yaml
```

```
default, my-service, Service (v1) has changed:
  ---
  apiVersion: v1
  kind: Service
  metadata:
    name: my-service
  spec:
    selector:
-     app: MyApp
+     app: 222
    ports:
      - protocol: TCP
        port: 80
        targetPort: 9376
```
