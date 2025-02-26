# DIFFERENTIAL EQUATIONS CHEAT SHEET

_The study of continuous change._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#overview)
* [NOMENCLATURE OF DERIVATIVES, INTEGRALS AND DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#nomenclature-of-derivatives-integrals-and-differential-equations)
  * [f(x) in CALCULUS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-calculus)
  * [f(x) in DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#fx-in-differential-equations)
  * [SUMMARY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#summary)
* [FIRST-ORDER AND SECOND-ORDER DIFFERENTIAL EQUATIONS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet#first-order-and-second-order-differential-equations)

Documentation and Reference

* [calculus](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet)
  cheat sheet

## OVERVIEW

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

## UNDERSTANDING f(x)

Understanding the role of $f(x)$ is important because:

* In **calculus**, you typically **differentiate or integrate** $f(x)$
  to get new functions.
* In **differential equations**, you **start with $f(x) = \frac{dy}{dx}$
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
\frac{dy}{dx} = f(x)
$$

To solve for y(x), we integrate both sides of the equation,

$$
\int \frac{dy}{dx} dx = \int f(x) dx
$$

$$
y(x) = \int f(x) dx
$$

As an example,

$$
\frac{dy}{dx} = f(x) = 2x + 3
$$

Solve for y(x),

$$
\int \frac{dy}{dx} dx = \int (2x + 3) dx
$$

$$
y(x) = x² + 3x + C
$$

### SUMMARY

| CONTEXT                | WHAT f(x) REPRESENTS               | WHAT WE WANT TO FIND  |
|------------------------|------------------------------------|-----------------------|
| CALCULUS (Derivatives) | $f(x)$ is the original function    | $f'(x)=\frac{dy}{dx}$ |
| CALCULUS (Integrals)   | $f(x)$ is function to integrate    | $F(x)=\int f(x)dx$    |
| DIFFERENTIAL EQUATIONS | $f(x)$ is the derivative of $y(x)$ | Solve for $y(x)$      |

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



