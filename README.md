# Graph Validation framework

Validation is often flat, we validate an entity 
against a set of rules and check for errors.

In some cases our validation scheme gets a little
complex and validations have dependencies on other
validations. 

We usually get around this by doing validation 
in phases. This gets quite cumbersome when our 
validation relationship starts getting more complex.

## DAAAAG

One way to get around this issue of phases is to
represent validations as a **directed acyclic graph**.
A simple depth first will traverse validation nodes
and stop when a validation fails. 

## vlad

Enter **vlad**, a simple dag validation framework.
vlad provides an intuitive way to construct validation
relationships and run the validation. vlad returns a
list of error messages that honor dependencies.

A successfull validation will return no errors.

A validation is represented as a simple graph node that 
consists of a  `validation function` and an `error message`.
The validation function takes in an entity and returns a
boolean indicating if the validation was honored or not.

```
validation_node {
    validation : e -> bool,
    msg : string,
}
```

```
dependency {
    from: validation_node,
    to: validation_node,
}
```
