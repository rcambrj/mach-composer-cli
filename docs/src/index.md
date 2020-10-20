# ![MACH composer](./_img/logo.png)

This documentation describes the workings of the MACH composer. Intended to setup and manage a **M**icroservices based, **A**PI-first, **C**loud-native SaaS and **H**eadless platform.


## What is it?

MACH composer is a framework that you use to orchestrate and extend modern digital commerce & experience platforms, based on MACH technologies and cloud native services. It provides a standards-based, future-proof tool-set and methodology to hand to your teams when building these types of platforms.

It includes:

- A configuration framework for managing MACH-services configuration, using infrastructure-as-code underneath (powered by Terraform)
- A microservices architecture based on modern serverless technology (AWS Lambda and Azure Functions), including (alpha) support for building your microservices with the Serverless Framework
- Multi-tenancy support for managing many instances of your platform, that share the same library of micro services
- CI/CD tools for automating the delivery of your MACH ecosystem
- Tight integratoin with AWS an Azure, including an (opinionated) setup of these cloud environments

The framework is intended as the 'center piece' of your MACH architecture and incorporates industry best practises such as the 12 Factor Methodology, Infrastrucure-as-code, DevOps, immutable deployments, FAAS, etc.

With combining (and requiring) these practises, using the framework has significant impact on your engineering methodology and organisation. On the other hand, by combining those practises we believe it offers an accelerated 'way in' in terms of embracing modern engineering practises in your organisation.


## How does it work?

The MACH composer takes a [YAML configuration](./syntax.md) as input, and will translate this into a Terraform configuration. It will then execute the terraform configuration, which will deploy all resources for the site architecture.

[![MACH diagram](./_img/mach.png)](./_img/mach.png)

The MACH composer is intended for managing multiple instances of the architecture.