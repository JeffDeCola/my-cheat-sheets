# DIFFERENTIAL EQUATIONS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_The study of continuous change._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#overview)
* [HOW TO SOLVE A DIFFERENTIAL EQUATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#how-to-solve-a-differential-equation)
* [UNDERSTANDING f(x)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#understanding-fx)
  * [f(x) in CALCULUS (DERIVATIVE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-calculus-derivative)
  * [f(x) in CALCULUS (INTEGRAL)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-calculus-integral)
  * [f(x) in DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-differential-equations)
* [CLASSIFICATION OF DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#classification-of-differential-equations)
  * [BY TYPE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#by-type)
  * [BY ORDER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#by-order)
* [EXAMPLES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#examples)
  * [POPULATION GROWTH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#population-growth)

Documentation and Reference

* [calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet)
* [my-latex-renders](https://github.com/JeffDeCola/my-latex-renders)
* [latex math mode](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/languages/latex-cheat-sheet/latex-math-mode.md)

## OVERVIEW

* **Calculus** is a broad field of mathematics that includes
differentiation (finding derivatives) and
integration (finding integrals).
It focuses on rates of change and accumulation.

* **Differential Equations** (DiffEQ) are a specific branch of mathematics
that deals with equations involving derivatives.
A differential equation expresses a relationship between a function and its derivatives.
Differential equations are used to
model real-world phenomena involving rates of
change and accumulation.

It's important to understand the role of $f(x)$ in both calculus and differential equations.

| CONTEXT                | WHAT f(x) REPRESENTS               | WHAT WE WANT TO FIND                 |
|------------------------|------------------------------------|--------------------------------------|
| CALCULUS (Derivatives) | $f(x)$ is the original function    | The derivative $f'(x)=\frac{dy}{dx}$ |
| CALCULUS (Integrals)   | $f(x)$ is function to integrate    | Function $F(x)=\int f(x)dx + C$      |
| DIFFERENTIAL EQUATIONS | $f(x)$ is the derivative of $y(x)$ | Solve for $y(x)$ by integrating      |

It's also important to understand the notation,

* $dx$ means a small change in x
* $dy$ means a small change in y
* $\frac{dy}{dx}$ means the rate of change of y with respect to x

## HOW TO SOLVE A DIFFERENTIAL EQUATION

```text
The goal of diffEQ is to find the function y(x) that satisfies the
differential equation.
```

Given this first-order ordinary differential equation
let's walk through the steps to solve it.

$$
\frac{dy}{dx} = 2x
$$

Separate the variables,

$$
dy = 2x dx
$$

Integrate both sides.

$$
\int dy = \int 2x dx
$$

$$
y = x² + C
$$

Solve for the constant C using initial conditions.

$$
y(0) = 0² + C = 0
$$

Hence

$$
C = 0
$$

Therefore, the solution to the differential equation is

$$
y = x²
$$

## UNDERSTANDING f(x)

Understanding the role of $f(x)$ is important because:

* In **calculus**, you typically **differentiate or integrate** $f(x)$
  to get new functions.
* In **differential equations**, you **start with $y'(x) = f(x) = \frac{dy}{dx}$
  and integrate** to recover $y(x)$.
* The notation can be tricky, but knowing whether **$f(x)$ is the function
  or its derivative** helps avoid confusion.

### f(x) in CALCULUS (DERIVATIVE)

Given a function, find the rate of change.

Original function,

$$
f(x) = y
$$

Derivative of that function (find rate of change),

$$
f'(x) = \frac{dy}{dx}
$$

As an example,

$$
f(x) = y = x² + 3x + 5
$$

The derivative of this function (rate of change) is,

$$
f'(x) = \frac{dy}{dx} = 2x + 3
$$

### f(x) in CALCULUS (INTEGRAL)

Given the rate of change, find the function.

Original function,

$$
f(x) = \frac{dy}{dx}
$$

Integral of that function (find function),

$$
F(x) = \int f(x) dx
$$

As an example,

$$
f(x) = \frac{dy}{dx} = 2x + 3
$$

The integral of this function is (find function),

$$
F(x) = \int (2x +3 ) dx  = x² + 3x + C
$$

You may also see the integral written as,

$$
F(x) = y(x)
$$

### f(x) in DIFFERENTIAL EQUATIONS

In differential equations, f(x) is often used to represent
the derivative of another function y(x).
Like integrals, start with the rate of change and find the function.

$$
f(x) = \frac{dy}{dx}
$$

Separate the variables

$$
f(x){dx} = {dy}
$$

integrate both sides of the equation,

$$
 \int f(x) dx = \int {dy}
$$

$$
\int f(x) dx = y(x)
$$

As an example,

$$
f(x) = \frac{dy}{dx} = 2x + 3
$$

or

$$
\frac{dy}{dx} = 2x + 3
$$

Separate the variables,

$$
dy = (2x + 3) dx
$$

Integrate both sides,

$$
\int dy = \int (2x + 3) dx
$$

$$
y(x) = x² + 3x + C
$$

## CLASSIFICATION OF DIFFERENTIAL EQUATIONS

Differential equations can be classified in many ways, we will
classify them by type and order.

### BY TYPE

* **Ordinary Differential Equations** (ODEs) involve only one independent variable. For example,

$$
\frac{dy}{dx} = 2x
$$

* **Partial Differential Equations** (PDEs) involve more than one independent variable.
For example,

$$
\frac{\partial u}{\partial t} = \frac{\partial² u}{\partial x²}
$$

where $u$ is a function of $x$ and $t$.

### BY ORDER

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

## EXAMPLES

Since differential equations are used to model real-world phenomena,
let's consider a simple example. Remember, the goal of differential
equations is to find the function $y(x)$ that satisfies the equation.

### POPULATION GROWTH

Consider a population of bacteria that grows at a rate proportional to the
current population. The differential equation that models this growth is,

$$
P'(t)= \frac{dP}{dt} = kP
$$

where $P$ is the population and $k$ is the growth rate constant.
We want to find P(t), the population at a particular time.

To solve this differential equation, we integrate both sides,

$$
\int \frac{dP}{dt} dt = \int kP dt
$$

$$
\int \frac{dP}{P} = \int k dt
$$

$$
\ln P = kt + C
$$

Solving for natural log we get

$$
P(t) = e^{kt + C}
$$

$$
P(t) = P_0e^{kt}
$$

where $P_0$ is the initial population at time $t=0$ and k is

$$
k = \frac{1}{t} ln \frac{P(t)}{P_0}
$$

As an example, if we have an initial popular of 100 bacteria
thats doubles every 3 hours,

$$ P_0 = 100 \quad and \quad P(3) = 2P_0 = 200 $$

the constant k would be

$$ k = \frac{1}{3} ln \frac{200}{100} = \frac{1}{3} ln 2 =0.231$$

Therefore, the population at time t would be,

$$ P(t) = 100e^{0.231t} $$
