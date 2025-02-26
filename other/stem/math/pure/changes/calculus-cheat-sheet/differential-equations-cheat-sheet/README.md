# DIFFERENTIAL EQUATIONS CHEAT SHEET

_The study of continuous change._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#overview)
* [HOW TO SOLVE A DIFFERENTIAL EQUATION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#how-to-solve-a-differential-equation)
* [UNDERSTANDING f(x)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#understanding-fx)
  * [f(x) in CALCULUS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-calculus)
  * [f(x) in DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-differential-equations)
* [CLASSIFICATION OF DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#classification-of-differential-equations)
  * [BY TYPE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#by-type)
  * [BY ORDER](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#by-order)
* [A REAL WORLD EXAMPLE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#a-real-world-example)
  * [POPULATION GROWTH](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#population-growth)

Documentation and Reference

* [calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet)
  cheat sheet

## OVERVIEW

| CONTEXT                | WHAT f(x) REPRESENTS               | WHAT WE WANT TO FIND                 |
|------------------------|------------------------------------|--------------------------------------|
| CALCULUS (Derivatives) | $f(x)$ is the original function    | The derivative $f'(x)=\frac{dy}{dx}$ |
| CALCULUS (Integrals)   | $f(x)$ is function to integrate    | Function $F(x)=\int f(x)dx$          |
| DIFFERENTIAL EQUATIONS | $f(x)$ is the derivative of $y(x)$ | Solve for $y(x)$ by integrating      |

* **Calculus** is a broad field of mathematics that includes
differentiation (finding derivatives) and
integration (finding integrals).
It focuses on rates of change and accumulation.

* **Differential Equations** (DiffEQ) are a specific branch of mathematics
that deals with equations involving derivatives.
A differential equation expresses a relationship between a function and its derivatives.
As we will see, differential equations are used to
model real-world phenomena involving rates of
change and accumulation.

```text
The goal of diffEQ is to find the function y(x) that satisfies the equation.
```

## HOW TO SOLVE A DIFFERENTIAL EQUATION

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

### f(x) in CALCULUS

In calculus, we define f(x) as the original function, and we
find the derivative or integral on that function.

**Derivative**

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

**Integral**

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
y'(x) = \frac{dy}{dx} = f(x)
$$

Separate the variables

$$
{dy} = f(x){dx}
$$

integrate both sides of the equation,

$$
\int {dy} = \int f(x) dx
$$

$$
y(x) = \int f(x) dx
$$

As an example,

$$
y'(x) = \frac{dy}{dx} = f(x) = 2x + 3
$$

or,

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

**Ordinary Differential Equations** (ODEs) involve only one independent variable. For example,

  $$
  \frac{dy}{dx} = 2x
  $$

**Partial Differential Equations** (PDEs) involve more than one independent variable.
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

## A REAL WORLD EXAMPLE

Since differential equations are used to model real-world phenomena,
let's consider a simple example. Remember, the goal of differential
equations is to find the function $y(x)$ that satisfies the equation.

### POPULATION GROWTH

Consider a population of bacteria that grows at a rate proportional to the
current population. The differential equation that models this growth is,

$$
\frac{dP}{dt} = kP
$$

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

$$ P_0 = 100 \; and \; P(3) = 2P_0 = 200 $$

the constant k would be

$$ k = \frac{1}{3} ln \frac{200}{100} = \frac{1}{3} ln 2 =0.231$$

Therefore, the population at time t would be,

$$ P(t) = 100e^{0.231t} $$
