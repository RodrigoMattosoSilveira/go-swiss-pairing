# Shadows and Forms
**WARNING** This is work in progress; I'll strive to reflect it in the architecture.

Recent developments in software architecture, like the µ-services, rendered concepts like `Entity` obsolete. The complexity of the enterprise systems we build today requires new concepts to help us understand, build, and support them. Our understanding of these systems is constantly growing, as did [Plato’s cave dwellers](https://en.wikipedia.org/wiki/Allegory_of_the_cave), who learned about reality from saw shadows reflected in their walls.

The shadows represent the fragment of reality that they could perceive through their senses, while the objects under the sun represent the true forms that they can only perceive through reason.

For the sake of simplicity, we will use the following names to assist in our definitions:
- application - an enterprise software solution;
- µ-services - an independent element of an application;
- µ-architecture- a means to conceive, design and implement an application using µ-services;

## Plato and Software Architecture and Plato
One of the most ubiquitous modern software architectures,` µ-architecture`, enables us to model, enhance, and integrate our understanding of our applications independently and interactively. Each of its `µ-servicea` has its own representation of the reality required for their operation, orthogonal from other application's `µ-services`; they have the ability to integrate their representation’s state with other µ-services also interested in them, as well as to persist state change required by the application.

As we identify ab application `µ-services` and interact with `domain experts`, `stakeholders`, and `clients` we learn nothing but elements of reality, `Shadows`, each providing us with an increased understanding of our requirements, but only mere incomplete perceptions of the real world; the collection of these `Shadows` offers our best approximation for our domain element's `Form`.

Concepts like the `OO Entity` are too rigid to assist us in learning about our `application` and its `µ-services` piece meal. It would be best to leave `OO Entities` behind and search for a new software architectural model to support our interactive understanding of the `applications` we build.

Thus, borrowing from Plato, we will use `Shadow` to represent a `slice of reality` and `Form` to represent our best understanding of that `integrated reality`.

## Shadow
- Is as a slice of a domain model required by a µ-service;
- µ-services rely on APIs they use to manage their Shadows;
- µ-services own their Shadows APIs;
- µ-services know nothing about other µ-services' Shadows;
- µ-services know nothing about the integrated µ-services' Shadows
- µ-services might have completely different Shadows for the same Form;

## Form:
- Is the best understanding and modeling of a Domain element at a particular point in time;
- Has a well defined and unique persistence structure;
- Its structure supports the APIs of the µ-services dependent on it;
- Its structure is orthogonal to the Shadows dependent on it;
- The only link between a Form and its Shadows are the APIs µ-services use to retrieve their Shadows;