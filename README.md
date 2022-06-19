# MonoPoly
A CLI tool for creating and managing mono-poly repositories, a hybrid that provides the advantages of both Monorepos and Polyrepos

## About
When working with multiple projects/services, we are faced with 2 choices for source code management or version control,
we either use the monorepo architecture, in which one repository is used with multiple directories for each project, or 
the Polyrepo architecture, in which multiple repositories are used for each project. 

The Monorepo architecture provides is most commonly used as it eases the development, encourages collaboration and code 
sharing, and it allows the developers to have a better, higher view of the whole application. However, it falls back when
it comes to scaling. Imagine having tens of directories for each project when you are only responsible for working on one
or two.

The Polyrepo architecture on the other hand makes the projects loosely coupled, and allows for a better view over the 
history of each project as commits are added for each repository separately, while having better support for CI/CD as
each repository is built and tested separately, without having to deal with subdirectories triggers that some CI vendors
offer, or rely on third-party tools.

Monopoly is a tool that facilitates using a hybrid architecture where one central repository, which contains all the
shared resources and code, is used in conjunction with the projects/services that are only needed by the developer at a 
time locally, keeping all projects in separate repositories, while freely "fetching" and "discarding" repositories locally
as per need.

![EHRouteLogoImage](https://i.imgur.com/jnHmXfr.png)
