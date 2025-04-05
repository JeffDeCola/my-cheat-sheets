# CALCULUS CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_The study of continuous change._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#overview)
* [LIMITS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#limits)
* [DIFFERENTIAL CALCULUS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#differential-calculus)
  * [THE DERIVATIVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-derivative)
  * [DEFINITION OF A DERIVATIVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#definition-of-a-derivative)
  * [BASIC DERIVATIVE RULES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#basic-derivative-rules)
* [INTEGRAL CALCULUS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#integral-calculus)
  * [THE INDEFINITE INTEGRAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-indefinite-integral)
  * [THE DEFINITE INTEGRAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-definite-integral)
  * [DEFINITION OF AN INTEGRAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#definition-of-an-integral)
  * [BASIC INTEGRAL RULES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#basic-integral-rules)
* [EXAMPLES - DERIVATIVES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#examples---derivatives)
  * [FINDING VELOCITY (Using the Power Rule)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#finding-velocity-using-the-power-rule)
  * [THE SIGMOID FUNCTION (Using the Quotient Rule)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-sigmoid-function-using-the-quotient-rule)
  * [THE TANH FUNCTION (Using the Quotient Rule)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-tanh-function-using-the-quotient-rule)
  * [EXAMPLE (Using the Chain Rule)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#example-using-the-chain-rule)
* [EXAMPLES - INTEGRALS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#examples---integrals)
  * [FINDING DISTANCE FUNCTION (Using the Power Rule)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#finding-distance-function-using-the-power-rule)

Documentation and Reference

* [differential equations](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet/differential-equations-cheat-sheet)
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

* $dx$ means  infinitesimal step in x
* $dy$ means  infinitesimal step in y
* $\frac{dy}{dx}$ means the rate of change of y with respect to x

So,

$$
dx(t) = y(t)dt
$$

Would mean, the infinitesimal area, dx(t), of this thin slice can be found by multiplying its height, y(t), to its infinitesimal width dt.

There are two branches of Calculus,

* Differential Calculus
  * Rate of change and slopes
  * **Derivatives** - One quantity changes with respect to another
  * Example: If you have a function describing the position of a falling object,
    its derivative gives you the object's velocity, and the derivative of
    velocity gives you acceleration

* Integral Calculus
  * Opposite of differential calculus
  * Finding area under curve
  * **Integrals** - Accumulation of a quantity over an interval
  * Example: If you have a function describing the velocity of a falling object,
    its integral gives you the object's position

## LIMITS

Before we get into calculus, we need to understand limits.
Limits are the foundation of calculus. They describe the value a function
approaches as the input approaches a certain value.

It is written as,

$$
\lim_{{x \to a}} f(x) = L
$$

Where `f(x)` is a function of `x`, `a` is the value `x` is approaching,
and `L` is the limit.

For example, given the function,

$$
f(x) = x² + 2
$$

<p align="center">
    <img src="svgs/f-of-x-equals-x-squared-plus-2.svg"
    align="middle"
</p>

The limit of `x² + 2` as `x` approaches `3` is 11,

$$
\lim_{x \to 3} (x² + 2) = 11
$$

We use limits because some functions are not defined at certain points.

For example, given the function

$$
f(x) = \frac{1}{x}
$$

<p align="center">
    <img src="svgs/f-of-x-equals-1-over-x.svg"
    align="middle"
</p>

This is not defined at `x = 0`.

However, we can still find the limit of `f(x)` as `x` approaches `0`.

$$
\lim_{x \to 0} \frac{1}{x} = \infty
$$

So how does this apply to calculus?  Well, the derivative of a function is
defined as the limit of the average rate of change of the function as the
interval over which the rate of change is calculated approaches zero.
This will make more sense in the next section.

## DIFFERENTIAL CALCULUS

Derivatives are the rate of change of a function.
They describe how one quantity changes with respect to another.

### THE DERIVATIVE

The derivative of a function

$$
f(x) = y
$$

is written as

$$
f'(x) = \frac{dy}{dx}
$$

where `dy` is the change in `y`, `dx` is the change in `x`, and `f'(x)` is the
derivative of `f(x)`.

For example, given the function,

$$
f(x) = y = x² + 2
$$

The derivative of `x²` is `2x` (see next section on how we did this).

$$
f'(x) = \frac{dy}{dx}  = 2x
$$

So if you're at `x = 2`, the slope (change) of the tangent line is `4`.

<p align="center">
    <img src="svgs/f-of-x-equals-x-squared-plus-2-showing-a-tangent-at-x-equals-2.svg"
    align="middle"
</p>

Hence, the derivative of a function is the slope of the tangent line to the curve
at a given point.

You will usually see the following notation to find the derivative of a function,

$$
\frac{d}{dx} f(x)
$$

or

$$
\frac{df(x)}{dx}
$$

### DEFINITION OF A DERIVATIVE

We learned about limits because the derivative of a function is
defined as the limit of the average rate of change of the function as the
interval over which the rate of change is calculated approaches zero.

<p align="center">
    <img src="svgs/limit-definition-of-the-derivative.svg"
    align="middle"
</p>

You can show this as

$$
f'(x) = \lim_{{\Delta x \to 0}} \frac{f(x+\Delta x) - f(x)}{\Delta x}
$$

So giving the example above of $f(x) = x^2 + 2$, the derivative is

$$
\begin{aligned}
f'(x) &= \lim_{{\Delta x \to 0}} \frac{((x+\Delta x)^2 + 2) -
         (x^2 + 2)}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{(x+\Delta x)^2 + 2 - x^2 - 2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{x^2 + 2x\Delta x + \Delta x^2 + 2 -
         x^2 - 2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{2x\Delta x + \Delta x^2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} 2x + \Delta x \\
f'(x) &= 2x
\end{aligned}
$$

### BASIC DERIVATIVE RULES

#### Power Rule

$$
\begin{aligned}
f(x) &= x^n \\
f'(x) &= nx^{n-1}
\end{aligned}
$$

#### Constant Rule

$$
\begin{aligned}
f(x) &= c \\
f'(x) &= 0
\end{aligned}
$$

#### Sum Rule

$$
\begin{aligned}
f(x) &= g(x) + h(x) \\
f'(x) &= g'(x) + h'(x)
\end{aligned}
$$

#### Product Rule

$$
\begin{aligned}
f(x) &= g(x)h(x) \\
f'(x) &= g'(x)h(x) + g(x)h'(x)
\end{aligned}
$$

#### Quotient Rule

$$
\begin{aligned}
f(x) &= \frac{g(x)}{h(x)} \\
f'(x) &= \frac{g'(x)h(x) - g(x)h'(x)}{h(x)^2}
\end{aligned}
$$

#### Chain Rule

$$
f(x) = g(h(x))
$$

$$
\begin{aligned}
\text{\scriptsize h(x) is the inner function} \\
\text{\scriptsize g(u) is the outer function with u = h(x)}
\end{aligned}
$$

$$
f'(x) = g'(h(x)) \cdot h'(x)
$$

## INTEGRAL CALCULUS

Integrals are the area under a curve. They describe the accumulation of a
quantity over an interval.

### THE INDEFINITE INTEGRAL

The integral of a function

$$
f(x)
$$

is written as

$$
\int f(x) dx = F(x) + C
$$

where `F(x)` is the antiderivative of `f(x)` and `C` is the constant of
integration.

```text
The reason we use F(x) is to avoid confusion with the original function f(x)
that is being integrated. F(x) is also written as y(x).
```

For example, given the function,

$$
f(x) = 2x
$$

The integral of `2x` is `x²` (see next section on how we did this).

$$
\int 2x dx = x² + C
$$

So if you're looking for the area under the curve of `2x`, you would find the
area of a triangle with base `x` and height `2x`.

### THE DEFINITE INTEGRAL

The definite integral of a function

$$
f(x)
$$

is written as

$$
\int_{a}^{b} f(x) dx = F(b) - F(a)
$$

where `F(b)` is the antiderivative of `f(x)` evaluated at `b` and
`F(a)` is the antiderivative of `f(x)` evaluated at `a`.

For example, given the function,

$$
f(x) = 2x
$$

The definite integral of `2x` from `0` to `2` is `4`.

$$
\begin{aligned}
\int_{0}^{2} 2x dx &= x² \Big|_{0}^{2} \\
                   &= F(2) - F(0) \\
                   &= 2² - 0² \\
                   &= 4
\end{aligned}
$$

### DEFINITION OF AN INTEGRAL

Because the integral of a function is
defined as the limit of the sum of the areas of
rectangles under the curve as the width of the
rectangles approaches zero.

<p align="center">
    <img src="svgs/limit-definition-of-the-integral.svg"
    align="middle"
</p>

You can show this as

$$
AREA = \int_{a}^{b} f(x) dx = \lim_{{\Delta x \to 0}}
       \sum_{i=1}^{n} f(x_i) \Delta x
$$

where

$$
\Delta x = \frac{b - a}{n}
$$

and

$$
n \to \infty \quad and \quad \Delta x \to 0
$$

So giving the example above of $f'(x) = 2x$, the integral is

$$
\begin{aligned}
\int 2x dx &= \lim_{{\Delta x \to 0}} \sum_{i=1}^{n} 2x_i \Delta x \\
\int 2x dx &= \lim_{{\Delta x \to 0}} \sum_{i=1}^{n} 2x_i \Delta x \\
\int 2x dx &= \lim_{{\Delta x \to 0}} 2 \sum_{i=1}^{n} x_i \Delta x \\
\int 2x dx &= 2 \lim_{{\Delta x \to 0}} \sum_{i=1}^{n} x_i \Delta x \\
\int 2x dx &= 2 \int x dx \\
\int 2x dx &= 2 \frac{x^2}{2} + C \\
\int 2x dx &= x^2 + C
\end{aligned}
$$

### BASIC INTEGRAL RULES

#### The Power Rule

$$
\int x^n dx = \frac{x^{n+1}}{n+1} + C
$$

#### The Constant Rule

$$
\int c dx = cx + C
$$

#### The Sum Rule

$$
\int (f(x) + g(x)) dx = \int f(x) dx + \int g(x) dx
$$

## EXAMPLES - DERIVATIVES

Some derivative examples.

### FINDING VELOCITY (Using the Power Rule)

Lets say a car travels 50 miles every hour.
Instead of $x,y$ we will use $t,s$ for time and distance.
We can write the distance function as

$$
f(t) = s = 50t
$$

where `t` is time in hours and `s` is distance in miles.

Let's find the derivative of the distance function using the Power Rule,

$$
\begin{aligned}
f(x) &= x^n \\
f'(x) &= nx^{n-1}
\end{aligned}
$$

For the distance function,

$$
n = 1
$$

Hence the derivative of the sigmoid function is,

$$
\begin{aligned}
f'(t) &= 1 \cdot 50t^{1-1} \\
f'(t) &= 50t^0 \\
f'(t) &= 50
\end{aligned}
$$

Therefor the velocity of the car is the **derivative** of the distance function.

$$
f'(t) = \frac{ds}{dt} =  50
$$

So the velocity of the car is 50 mph.

### THE SIGMOID FUNCTION (Using the Quotient Rule)

The sigmoid function is a mathematical function having an "S" shaped curve
(also called a sigmoid curve).

<p align="center">
    <img src="svgs/sigmoid-function.svg"
    align="middle"
</p>

The sigmoid function is defined as,

$$
\sigma(x) = \frac{1}{1 + e^{-x}}
$$

The range of the sigmoid function is between 0 and 1.

$$
0 \lt \sigma(x) \lt 1
$$

And the derivative of the sigmoid function is (see solution below),

$$
\sigma'(x) = \sigma(x) \cdot (1 - \sigma(x))
$$

Which is interesting. The derivative of the sigmoid function is the sigmoid
function times 1 minus the sigmoid function.
This is a very important property of the sigmoid function and is used in
neural networks. Because it makes it easy to compute the derivative of
the sigmoid function during backpropagation.

Let's find the derivative of the sigmoid function using the Quotient Rule,

$$
\sigma(x) = \frac{g(x)}{h(x)}
$$

$$
\sigma'(x) = \frac{g'(x)h(x) - g(x)h'(x)}{h(x)^2}
$$

For the sigmoid function,

$$
\begin{aligned}
g(x) &= 1 \\
g'(x) &= 0 \\
h(x) &= 1 + e^{-x} \\
h'(x) &= -e^{-x}
\end{aligned}
$$

Hence the derivative of the sigmoid function is,

$$
\begin{aligned}
\sigma'(x) &= \frac{0 \cdot (1 + e^{-x}) - 1 \cdot (-e^{-x})}{(1 + e^{-x})^2} \\
\sigma'(x) &= \frac{e^{-x}}{(1 + e^{-x})^2}
\end{aligned}
$$

Now, let's simplify this,

$$
\begin{aligned}
\sigma'(x) &= \frac{1}{1 + e^{-x}} \cdot \frac{e^{-x}}{1 + e^{-x}} \\
\sigma'(x) &= \frac{1}{1 + e^{-x}} \cdot \frac{1 + e^{-x} - 1}{1 + e^{-x}} \\
\sigma'(x) &= \frac{1}{1 + e^{-x}} \cdot \left( 1 - \frac{1}{1 + e^{-x}} \right)
\end{aligned}
$$

Hence,

$$
\sigma'(x) = \sigma(x) \cdot (1 - \sigma(x))
$$

### THE TANH FUNCTION (Using the Quotient Rule)

The tanh function is a mathematical function having an "S" shaped curve
(also called a sigmoid curve).

<p align="center">
    <img src="svgs/tanh-function.svg"
    align="middle"
</p>

The tanh function is defined as,

$$
\tanh(x) = \frac{e^x - e^{-x}}{e^x + e^{-x}}
$$

The range of the tanh function is between -1 and 1.

$$
-1 \lt \tanh(x) \lt 1
$$

And the derivative of the tanh function is (see solution below),

$$
\tanh'(x) = 1 - \tanh^2(x)
$$

Which is interesting. The derivative of the tanh function is 1 minus the tanh function squared.
This is a very important property of the tanh function and is used in
neural networks. Because it makes it easy to compute the derivative of
the tanh function during backpropagation.

Let's find the derivative of the tanh function using the Quotient Rule,

$$
\tanh(x) = \frac{g(x)}{h(x)}
$$

$$
\tanh'(x) = \frac{g'(x)h(x) - g(x)h'(x)}{h(x)^2}
$$

For the tanh function,

$$
\begin{aligned}
g(x) &= e^x - e^{-x} \\
g'(x) &= e^x + e^{-x} \\
h(x) &= e^x + e^{-x} \\
h'(x) &= e^x - e^{-x}
\end{aligned}
$$

Hence the derivative of the tanh function is,

$$
\begin{aligned}
\tanh'(x) &= \frac{(e^x + e^{-x})(e^x + e^{-x}) - (e^x - e^{-x})(e^x - e^{-x})}{(e^x + e^{-x})^2} \\
\tanh'(x) &= \frac{(e^x + e^{-x})^2 - (e^x - e^{-x})^2}{(e^x + e^{-x})^2} \\
\end{aligned}
$$

We know that,

$$
\begin{aligned}
(a + b)^2 &= a^2 + 2ab + b^2 \\
(a - b)^2 &= a^2 - 2ab + b^2
\end{aligned}
$$

So

$$
\begin{aligned}
(e^x + e^{-x})^2 &= e^{2x} + 2e^x e^{-x} + e^{-2x} = e^{2x} + 2 + e^{-2x} \\
(e^x - e^{-x})^2 &= e^{2x} - 2e^x e^{-x} + e^{-2x} = e^{2x} - 2 + e^{-2x}
\end{aligned}
$$

Plug this back in, we simplify to,

$$
\begin{aligned}
\tanh'(x) &= \frac{e^{2x} + 2 + e^{-2x} - e^{2x} + 2 - e^{-2x}}{(e^x + e^{-x})^2} \\
\tanh'(x) &= \frac{4}{(e^x + e^{-x})^2} \\
\end{aligned}
$$

Now, with a little leap of faith, 4 can be written as,

$$
\begin{aligned}
4 = (e^x + e^{-x})^2 -(e^x - e^{-x})^2
\end{aligned}
$$

So now the derivative looks like,

$$
\begin{aligned}
\tanh'(x) &= \frac{(e^x + e^{-x})^2 - (e^x - e^{-x})^2}{(e^x + e^{-x})^2} \\
\end{aligned}
$$

We can then split the fraction,

$$
\begin{aligned}
\tanh'(x) &= \frac{(e^x + e^{-x})^2}{(e^x + e^{-x})^2} - \frac{(e^x - e^{-x})^2}{(e^x + e^{-x})^2} \\
\tanh'(x) &= 1 - \frac{(e^x - e^{-x})^2}{(e^x + e^{-x})^2} \\
\end{aligned}
$$

Now if we recognize that the original function,

$$
\begin{aligned}
\tanh(x) &= \frac{e^x - e^{-x}}{e^x + e^{-x}} \\
\end{aligned}
$$

Than squaring the original function would be,

$$
\begin{aligned}
\tanh^2(x) &= \left( \frac{e^x - e^{-x}}{e^x + e^{-x}} \right)^2 \\
\tanh^2(x) &= \frac{(e^x - e^{-x})^2}{(e^x + e^{-x})^2} \\
\end{aligned}
$$

Plug this into the derivative to get,

$$
\begin{aligned}
\tanh'(x) &= 1 - \tanh^2(x)
\end{aligned}
$$

Whew, that was a lot of work.

### EXAMPLE (Using the Chain Rule)

Giving the function,

$$
f(x)=(3x^2 + 5x − 2)^4
$$

Let's find the derivative of the function using the
Chain Rule,

$$
f(x) = g(h(x))
$$

$$
\begin{aligned}
\text{\scriptsize h(x) is the inner function} \\
\text{\scriptsize g(u) is the outer function with u = h(x)}
\end{aligned}
$$

$$
f'(x) = g'(h(x)) \cdot h'(x)
$$

For our function, we take the inner and outer functions as

$$
\begin{aligned}
h(x) &= 3x^2 + 5x − 2 \\
g(u) &= u^4
\end{aligned}
$$

Where those derivatives are,

$$
\begin{aligned}
g'(u) &= 4u^3 \\
h'(x) &= 6x + 5
\end{aligned}
$$

Hence the derivative of the function is,

$$
f'(x) = 4(3x^2 + 5x − 2)^3 \cdot (6x + 5)
$$

## EXAMPLES - INTEGRALS

Some integrals examples.

### FINDING DISTANCE FUNCTION (Using the Power Rule)

Lets say we know the velocity of a car is 50 mph.
We can write the velocity function as

$$
f'(t) = \frac{ds}{dt} = 50
$$

The distance function is the **integral** of the velocity function.

Let's find the distance function using the Power Rule,

$$
\begin{aligned}
\int c dx &= cx + C
\end{aligned}
$$

Hence,

$$
f(t) = \int 50 dt = 50t + C
$$
