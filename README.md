## Solid Principle Implementation

### Single Responsibility Principle
- A class should have only one reason to change.
- A class should have only one job.
- A class should have only one responsibility.

### Open Closed Principle
- Software entities (classes, modules, functions, etc.) should be open for extension, but closed for modification.
- Open for extension means that we should be able to add new features or components to the application without breaking existing code.
- Closed for modification means that we should not introduce breaking changes to existing functionality, because that would force you to refactor a lot of existing code.

### Liskov Substitution Principle
- Subtypes must be substitutable for their base types.
- Objects in a program should be replaceable with instances of their subtypes without altering the correctness of that program.
- If S is a subtype of T, then objects of type T may be replaced with objects of type S without altering any of the desirable properties of that program.
- A type S is a subtype of type T if objects of type S may be replaced with objects of type T without altering any of the desirable properties of that program.
- A type S is a subtype of type T if it can be used in any place where T is expected.

### Interface Segregation Principle
- A client should never be forced to implement an interface that it doesn’t use or clients shouldn’t be forced to depend on methods they do not use.
- Many client-specific interfaces are better than one general-purpose interface.
- No client should be forced to depend on methods it does not use.
- ISP splits interfaces that are very large into smaller and more specific ones so that clients will only have to know about the methods that are of interest to them. Such shrunken interfaces are also called role interfaces. ISP is intended to keep a system decoupled and thus easier to refactor, change, and redeploy.
- ISP is intended to keep a system decoupled and thus easier to refactor, change, and redeploy.

### Dependency Inversion Principle
- High-level modules should not depend on low-level modules. Both should depend on abstractions.
- Abstractions should not depend on details. Details should depend on abstractions.
- The principle states that high-level modules should not depend on low-level modules, both should depend on abstractions. In other words, the principle states that the dependency between high-level and low-level modules should be abstracted using an interface.
- The principle states that abstractions should not depend on details, details should depend on abstractions. In other words, the principle states that the dependency between abstractions and details should be abstracted using an interface.

## Run

### Command to run the project

#### Migrate & Seed
```bash
  go run main.go db:migrate or make migrate
  go run main.go db:seed or make seed
```

#### Run CRUD
```bash
  go run main.go {command}:{argument} or make {command}-{argument}
```

Or using **_Makefile_**