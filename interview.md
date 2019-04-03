# JS technical test

## JS- Scope
Consider the following code:
```javascript
var a = b = 5;
console.log(b);
```
What will be printed on the console?

## JS- Hoisting
Describe what the following will return and why
```javascript
function bar() {
   console.log(a);
   console.log(foo());
   var a = 1;
   function foo() {
      return 2;
   }
}
bar();
```

## JS- Flow
Describe the predicted output of, and why
```javascript
console.log("john");
setTimeout(function() {
    console.log("jack");
}, 0);
console.log("jake");
```

## JS- Extend basic prototypes
Write javascript to make the following:
```javascript
console.log('helloworld'.repeatit(3));
```
output: ```helloworldhelloworldhelloworld```

## JS- Classes
- How do you declare a class in Javascript?
- How do you extend a class in Javascript?
- How do you mixin a class in Javascript?
- How do you create a new instance of a class?

## JS- Quickfire
 - Describe the difference between `==` and `===`
 - Assuming ```function myfunc() {}``` describe the difference between ```myfunc()``` ```myfunc.call()``` and ```myfunc.apply()``` 
 - Describe the difference between ```null``` and ```undefined```
 - What is an error-first callback?
 - What is and how can you avoid 'callback hell'?
 - How can you listen on port 80 with Node?
 - What's the event loop?
 - Why npm shrinkwrap is useful?

--- 
# NOT JS -

## Git
Assume you have commits pushed to origin:```ABCDEF```
- You discover another developer put a password in ```C```, discuss how you could go about editing that commit
- You find it would make more semantic sense to reorder the commits to ```AEBCF```
- You find it would make more semantic sense to combine commits ```A & B```
Assume have another branch with commits: ```WXYZ```
- You want to take commit ```Y``` and add it after ```E```

## Testing
- Describe a test pyramid
- How can you implement it when talking about HTTP APIs?
- Describe mocking, stubbing
- Describe the difference between BDD and TDD
- Why would you use BDD/TDD, are they mutually exclusive?

## Docker
 - Describe the difference between `RUN`, `CMD`, and `ENTRYPOINT`
 - Describe the difference between `ADD` and `COPY`
 - What is a multi-stage build and why would you use one?
 - Figure how to start up and return a salted hash with the docker image `chrisns/docker-devtest`

## Kubernetes
 - What is the difference between `livenessProbe` and `readinessProbe`
 - What is the sidecar pattern
 - What is a a custom operator
 - [whiteboard] Draw the relationship between describing what they do
   - `cluster`
   - `configMap`
   - `container`
   - `cronjob`
   - `daemonset`
   - `deployment`
   - `headless`
   - `ingress`
   - `job`
   - `kubectl`
   - `kubelet`
   - `kubeproxy`
   - `namespace`
   - `persistentVolume`
   - `persistentVolumeClaim`
   - `pod`
   - `replicaSet`
   - `replicationController`
   - `secret`
   - `service`
   - `statefulSet`
 - What is an `initContainer` and how would it fit in?
 - How would you manage the resources your container needs
 - Can you split a pod over multiple nodes?
 - What is a node `taint` or `tolerate`
 - Why would you set a `podAntiAffinity` or `podAffinity`?
 - Affinity, what is the difference between `requiredDuringSchedulingIgnoredDuringExecution` and `preferredDuringSchedulingIgnoredDuringExecution`
 - What do each of these pod phases mean, and why would it get there?
   - `Pending`
   - `Running`
   - `Succeeded`
   - `Failed`
   - `Unknown`
   - `Completed`
   - `CrashLoopBackOff`
 - What is the difference between a service type and why would you use each
   - `ClusterIP`
   - `NodePort`
   - `LoadBalancer`
   - `ExternalName`
   - `Headless`
 - What is `minikube`?
 - What protocol underpins `kubectl`
 - How do you do a canary deployment?
 - Your metrics appear to show that when the pod schedudles it bounces queries for a bit, why might this be? [1]
 <details>
  <summary>[1]</summary>

```yaml
---
kind: Deployment
apiVersion: apps/v1
metadata:
  name: mydeployment
spec:
  replicas: 30
  selector:
    matchLabels:
      name: mydeployment
  template:
    metadata:
      labels:
        name: mydeployment
    spec:
      containers:      
      - image: chrisns/mywonderful-notreal-app
        name: mydeployment
        livenessProbe:
          httpGet:
            path: /health/liveness
            port: http
          initialDelaySeconds: 90
          periodSeconds: 3
          timeoutSeconds: 1

        readinessProbe:
          httpGet:
            path: /health/readiness
            port: http
          initialDelaySeconds: 60
          periodSeconds: 3
          timeoutSeconds: 1
```
</details>

## Coding tests
 - Pair with one of our juniors/apprentices on solving [a kata](https://www.codewars.com/kata/multiples-of-3-or-5/go)

 ## Discussion points
 - What's the difference between operational and programmer errors?
 - AWS/rackspace/similar cloud
 - NoSQL stores
 - Technology you're excited about
 - Mysql when would you use INNODB vs MYSIAM
 - Talk about what HTTP methods there are and what they are used for.
 - Talk about what RESTful means to you
 - Describe what HATEOAS means to you
 - What tools can be used to assure consistent style?
 - What is CI/CD/CD
 - Working in an agile team, what does that mean

## What don't you like about:
 - JS/angular/backbone/ember/other framework/es6
 - Node.js
 - Mysql/postgres
 - NoSQL store
 - Docker
 - Kubernetes
 - other technology