# DIFFERENTIAL EQUATIONS CHEAT SHEET

_The study of continuous change._

Table of Contents

Documentation and Reference

* [calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet)
  cheat sheet

## OVERVIEW

* **Calculus** is a broad field of mathematics that includes 
differentiation (finding derivatives) and
integration (finding integrals).
It focuses on rates of change and accumulation.

* **Differential** equations (DiffEQ) are a specific branch of mathematics
that deals with equations involving derivatives.
A differential equation expresses a relationship between a function and its derivatives.
As we will see, differential equations are used to
model real-world phenomena involving rates of
change and accumulation.

### DERIVATIVE

```text
Start with the function and find the rate of change at any point.

"How quickly is this function changing at this particular point?"
```

For example, giving the function,

$$
f(x) = y = x² + 3x + 5
$$

The derivative of this function is,

$$
f'(x) = \frac{dy}{dx} = 2x + 3
$$

This tells us that at any point x, the rate of change is 2x + 3.

This is useful in physics, where we could use the derivative to find the
velocity of an object.

### INTEGRAL

```text
Start with the rate of change and find the function.
```

For example, giving the rate of change,

$$
f(x) = \frac{dy}{dx} = 2x + 3
$$

```text
NOTE: f(x) is defined as the derivative of F(x). We
want to find F(x) or also written as y(x).
```

The integral of this function is,

$$
\int f(x) dx = F(x) = y(x) = x² + 3x + C
$$

This is useful in physics, where we could use the integral to find the
position of an object.

### DIFFERENTIAL EQUATION

A differential equation is an equation that contains a derivative.

```text

Start with the rate of change and find the function.

I know how the function changes, but what is the actual function?"
```

For example, for the differential equation,

$$
f(x) = \frac{dy}{dx} = 2x + 3
$$

The function is,

$$
F(x) = y = x² + 3x + C
$$

This is useful in physics, where we could use the differential equation
to find the position of an object.

```text
But this looks like an integral, right?  Yes, but the difference is
that a differential equation is a general solution that can be used
to solve many problems. Not every differential equation can be solved by
integration. And not every integral is a differential equation.
```

For example, the differential equation,




## FIRST-ORDER AND SECOND-ORDER DIFFERENTIAL EQUATIONS

* **First-order** differential equations involve only the first derivative.
  For example,

  $$
  y'(x) = \frac{dy}{dx} = 2x
  $$
* **Second-order** differential equations involve the second derivative.
  For example,

  $$
  y''(x) = \frac{d²y}{dx²} = 2x
  $$



