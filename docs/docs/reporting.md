---
sidebar_position: 7
title: 📝 Reporting
---

# 📝 Vet Reporting

`vet` default scan uses an opinionated [Summary Reporter](https://github.com/safedep/vet/blob/main/pkg/reporter/markdown.template.md) which presents a consolidated summary of findings. Thats NOT about it. Read more for expression based filtering and policy evaluation.

## Summary Report

- The default output format for the `vet` is console with summary. You can run the following command to get the summary report `--report-summary` flag. Which is basically a summary report with actionable advice.

```bash
vet scan --report-summary -D demo-client-java
```

![vet summary report](/img/vet/vet-report-summary.png)

## Console

- You can run the console output format for the `vet` using `--report-console` flag. Which returns the whole report of the scan in detailed format in the terminal or console.

```bash
vet scan --report-console -D demo-client-java
```

```bash
 ❯ vet scan --report-console -D demo-client-java
Scanning packages    ... 95.7% [######################.] [110 in 6.401304s]
Scanning manifests   ...  0.0% [.......................] [0 in 6.401344s]
Manifest: /Users/madhuakula/vet/demo-client-java/gradle.lockfile
┌──────────────────────────────────────────────────────────────────────────────────┬────────────────┬───────────────────┐
│ PACKAGE                                                                          │ ATTRIBUTE      │ SUMMARY           │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ antlr:antlr/2.7.7                                                                │                │                   │
│                                                                                  │ Version Drift  │ 2.7.7 > 20030911  │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ com.fasterxml.jackson.core:jackson-databind/2.13.4                               │                │                   │
│                                                                                  │ Vulnerability  │ Critical:0 High:1 │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ com.sun.activation:jakarta.activation/1.2.2                                      │                │                   │
│                                                                                  │ Version Drift  │ 1.2.2 > 2.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ com.sun.istack:istack-commons-runtime/3.0.12                                     │                │                   │
│                                                                                  │ Low Popularity │ Stars:9 Issues:2  │
│                                                                                  │ Version Drift  │ 3.0.12 > 4.1.1    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ com.zaxxer:HikariCP/4.0.3                                                        │                │                   │
│                                                                                  │ Version Drift  │ 4.0.3 > 5.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ commons-fileupload:commons-fileupload/1.4                                        │                │                   │
│                                                                                  │ Vulnerability  │ Critical:0 High:1 │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.openfeign:feign-core/11.8                                              │                │                   │
│                                                                                  │ Version Drift  │ 11.8 > 12.2       │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.openfeign:feign-slf4j/11.8                                             │                │                   │
│                                                                                  │ Version Drift  │ 11.8 > 12.2       │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-annotations/1.7.0                            │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-circuitbreaker/1.7.0                         │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-circularbuffer/1.7.0                         │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-consumer/1.7.0                               │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-core/1.7.0                                   │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-framework-common/1.7.0                       │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-micrometer/1.7.0                             │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-ratelimiter/1.7.0                            │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-retry/1.7.0                                  │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-spring-boot2/1.7.0                           │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-spring/1.7.0                                 │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ io.github.resilience4j:resilience4j-timelimiter/1.7.0                            │                │                   │
│                                                                                  │ Version Drift  │ 1.7.0 > 2.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ jakarta.annotation:jakarta.annotation-api/1.3.5                                  │                │                   │
│                                                                                  │ Version Drift  │ 1.3.5 > 2.1.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ jakarta.persistence:jakarta.persistence-api/2.2.3                                │                │                   │
│                                                                                  │ Version Drift  │ 2.2.3 > 3.1.0     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ jakarta.transaction:jakarta.transaction-api/1.3.3                                │                │                   │
│                                                                                  │ Version Drift  │ 1.3.3 > 2.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ jakarta.validation:jakarta.validation-api/2.0.2                                  │                │                   │
│                                                                                  │ Version Drift  │ 2.0.2 > 3.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ jakarta.xml.bind:jakarta.xml.bind-api/2.3.3                                      │                │                   │
│                                                                                  │ Version Drift  │ 2.3.3 > 4.0.0     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ net.minidev:json-smart/2.4.8                                                     │                │                   │
│                                                                                  │ Vulnerability  │ Critical:0 High:1 │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.apache.tomcat.embed:tomcat-embed-core/9.0.65                                 │                │                   │
│                                                                                  │ Version Drift  │ 9.0.65 > 10.1.7   │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.apache.tomcat.embed:tomcat-embed-el/9.0.65                                   │                │                   │
│                                                                                  │ Version Drift  │ 9.0.65 > 10.1.7   │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.apache.tomcat.embed:tomcat-embed-websocket/9.0.65                            │                │                   │
│                                                                                  │ Version Drift  │ 9.0.65 > 10.1.7   │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.glassfish.jaxb:jaxb-runtime/2.3.6                                            │                │                   │
│                                                                                  │ Version Drift  │ 2.3.6 > 4.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.glassfish.jaxb:txw2/2.3.6                                                    │                │                   │
│                                                                                  │ Version Drift  │ 2.3.6 > 4.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.slf4j:jul-to-slf4j/1.7.36                                                    │                │                   │
│                                                                                  │ Version Drift  │ 1.7.36 > 2.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.slf4j:slf4j-api/1.7.36                                                       │                │                   │
│                                                                                  │ Version Drift  │ 1.7.36 > 2.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-autoconfigure/2.7.4                         │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-configuration-processor/2.7.4               │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-devtools/2.7.4                              │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-aop/2.7.4                           │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-data-jpa/2.7.4                      │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-data-rest/2.7.4                     │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-jdbc/2.7.4                          │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-json/2.7.4                          │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-logging/2.7.4                       │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-oauth2-resource-server/2.7.4        │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-security/2.7.4                      │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-tomcat/2.7.4                        │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-validation/2.7.4                    │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter-web/2.7.4                           │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot-starter/2.7.4                               │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.boot:spring-boot/2.7.4                                       │                │                   │
│                                                                                  │ Version Drift  │ 2.7.4 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-circuitbreaker-resilience4j/2.1.4         │                │                   │
│                                                                                  │ Version Drift  │ 2.1.4 > 3.0.0     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-commons/3.1.4                             │                │                   │
│                                                                                  │ Version Drift  │ 3.1.4 > 4.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-context/3.1.4                             │                │                   │
│                                                                                  │ Version Drift  │ 3.1.4 > 4.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-openfeign-core/3.1.4                      │                │                   │
│                                                                                  │ Version Drift  │ 3.1.4 > 4.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-starter-circuitbreaker-resilience4j/2.1.4 │                │                   │
│                                                                                  │ Version Drift  │ 2.1.4 > 3.0.0     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-starter-openfeign/3.1.4                   │                │                   │
│                                                                                  │ Version Drift  │ 3.1.4 > 4.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.cloud:spring-cloud-starter/3.1.4                             │                │                   │
│                                                                                  │ Version Drift  │ 3.1.4 > 4.0.1     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.data:spring-data-commons/2.7.3                               │                │                   │
│                                                                                  │ Version Drift  │ 2.7.3 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.data:spring-data-jpa/2.7.3                                   │                │                   │
│                                                                                  │ Version Drift  │ 2.7.3 > 3.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.data:spring-data-rest-core/3.7.3                             │                │                   │
│                                                                                  │ Version Drift  │ 3.7.3 > 4.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.data:spring-data-rest-webmvc/3.7.3                           │                │                   │
│                                                                                  │ Version Drift  │ 3.7.3 > 4.0.4     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.hateoas:spring-hateoas/1.5.2                                 │                │                   │
│                                                                                  │ Version Drift  │ 1.5.2 > 2.0.3     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-config/5.7.3                        │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-core/5.7.3                          │                │                   │
│                                                                                  │ Vulnerability  │ Critical:1 High:0 │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-crypto/5.7.3                        │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-oauth2-core/5.7.3                   │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-oauth2-jose/5.7.3                   │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-oauth2-resource-server/5.7.3        │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework.security:spring-security-web/5.7.3                           │                │                   │
│                                                                                  │ Version Drift  │ 5.7.3 > 6.0.2     │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-aop/5.3.23                                            │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-aspects/5.3.23                                        │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-beans/5.3.23                                          │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-context/5.3.23                                        │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-core/5.3.23                                           │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-expression/5.3.23                                     │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-jcl/5.3.23                                            │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-jdbc/5.3.23                                           │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-orm/5.3.23                                            │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-tx/5.3.23                                             │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-web/5.3.23                                            │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.springframework:spring-webmvc/5.3.23                                         │                │                   │
│                                                                                  │ Version Drift  │ 5.3.23 > 6.0.7    │
├──────────────────────────────────────────────────────────────────────────────────┼────────────────┼───────────────────┤
│ org.yaml:snakeyaml/1.30                                                          │                │                   │
│                                                                                  │ Vulnerability  │ Critical:0 High:2 │
Scanning packages    ... done! [115 in 6.422s]
Scanning manifests   ... done! [1 in 6.422s]
  **   Summary of Findings

  ** 1 critical, 5 high and 5 other vulnerabilities were identified

  ** 1 potentially unpopular library identified as direct dependency

  ** 78 libraries are out of date with major version drift in direct dependencies

  ** across 115 libraries in 1 manifest(s)

Consider upgrading the following libraries for maximum impact:

┌─────────────────────────────────────────────────────────┬───────────┬────────┐
│ PACKAGE                                                 │ UPDATE TO │ IMPACT │
├─────────────────────────────────────────────────────────┼───────────┼────────┤
│ org.yaml:snakeyaml@1.30                                 │ 2.0       │ 13     │
│  vulnerability   drift                                  │           │        │
├─────────────────────────────────────────────────────────┼───────────┼────────┤
│ org.springframework.security:spring-security-core@5.7.3 │ 6.0.2     │ 7      │
│  vulnerability   drift                                  │           │        │
├─────────────────────────────────────────────────────────┼───────────┼────────┤
│ com.sun.istack:istack-commons-runtime@3.0.12            │ 4.1.1     │ 4      │
│  low popularity   drift                                 │           │        │
├─────────────────────────────────────────────────────────┼───────────┼────────┤
│ commons-fileupload:commons-fileupload@1.4               │ 1.5       │ 3      │
│  vulnerability                                          │           │        │
├─────────────────────────────────────────────────────────┼───────────┼────────┤
│ net.minidev:json-smart@2.4.8                            │ 2.4.10    │ 3      │
│  vulnerability                                          │           │        │
└─────────────────────────────────────────────────────────┴───────────┴────────┘

There are 77 more libraries that should be upgraded to reduce risk
Run vet with `--report-markdown=/path/to/report.md` for details

Run with `vet --filter="..."` for custom filters to identify risky libraries
For more details https://github.com/safedep/vet
```

## Markdown

- You can run the Markdown output format for the `vet` using `--report-markdown` flag. Which it generates consolidated markdown report to file.

```bash
vet scan --report-markdown=vet-markdown-report.md -D demo-client-java
```

![vet markdown report file](/img/vet/vet-markdown-report-file.png)

## SARIF

🚧 Work-in-Progress (WIP)

## JSON

🚧 Work-in-Progress (WIP)

## CSV

🚧 Work-in-Progress (WIP)
