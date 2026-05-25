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
Differential equations are used to model real-world phenomena involving rates of
change and accumulation.

In short: calculus gives you the tools (derivatives and integrals), and differential equations puts those tools to work modeling real systems.

It's important to understand the role of $f(x)$ in both calculus and differential equations.

| CONTEXT                | WHAT f(x) REPRESENTS               | WHAT WE WANT TO FIND                 |
|------------------------|------------------------------------|--------------------------------------|
| CALCULUS (Derivatives) | $f(x)$ is the original function    | The derivative $f'(x)=\frac{dy}{dx}$ |
| CALCULUS (Integrals)   | $f(x)$ is function to integrate    | Function $F(x)=\int f(x)dx + C$      |
| DIFFERENTIAL EQUATIONS | $f(x)$ is the derivative of $y(x)$ | Solve for $y(x)$ by integrating      |

It's also important to understand the notation,

* $dx$ means infinitesimal step (small change) in x
* $dy$ means infinitesimal step (small change) in y
* $\frac{dy}{dx}$ means the rate of change of y with respect to x

<p align="center">
    <img src="svgs/delta-y-vs-dy.svg"
    alt="delta-y-vs-dy">
</p>

From this diagram, a secant line (not shown) connects the two points
on the curve where the slope is $\frac{\Delta y}{\Delta x}$.
As $\Delta x$ shrinks to zero the secant rotates into the
tangent line (green) and the slope is $\frac{dy}{dx}$, which is the derivative,
the instantaneous rate of change. Formally written as,

$$
\lim_{\Delta x \to 0} \frac{\Delta y}{\Delta x} = \frac{dy}{dx}
$$

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

## HOW TO SOLVE A DIFFERENTIAL EQUATION

> The goal of diffEQ is to find the function $y(x)$ that satisfies the
> differential equation.

The most common technique for simple differential equations is
**separation of variables** — rearrange the equation so all the $y$
terms are on one side and all the $x$ terms are on the other, then
integrate both sides.

Given this first-order ordinary differential equation, let's walk
through the steps to solve it.

$$
\frac{dy}{dx} = 2x
$$

Separate the variables,

$$
dy = 2x \, dx
$$

Integrate both sides,

$$
\int dy = \int 2x \, dx
$$

$$
y = x^2 + C
$$

This is the **general solution** — a family of curves, one for each
value of $C$.

To find a **specific solution**, we need an **initial condition**.
A differential equation paired with an initial condition is called
an **initial value problem (IVP)**.

If we're told that $y(0) = 0$ (meaning $y = 0$ when $x = 0$),

$$
y(0) = 0^2 + C = 0
$$

Hence

$$
C = 0
$$

Therefore, the specific solution to this initial value problem is

<table align="center"><tr><td>

$$\displaystyle y = x^2$$

</td></tr></table>

<table align="center"><tr><td>

$$y = x^2$$

</td></tr></table>

## UNDERSTANDING f(x)

Understanding the role of $f(x)$ is important because:

* In **calculus**, you typically **differentiate or integrate** $f(x)$
  to get new functions.
* In **differential equations**, you **start with $y'(x) = f(x) = \frac{dy}{dx}$
  and integrate** to recover $y(x)$.

The notation can be tricky, but knowing whether **$f(x)$ is the function
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
f(x) = y = x^2 + 3x + 5
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
F(x) = \int (2x +3 ) dx  = x^2 + 3x + C
$$

You may also see the integral written as,

$$
F(x) = y(x)
$$

### f(x) in DIFFERENTIAL EQUATIONS

In differential equations, $f(x)$ is often used to represent
the derivative of another function $y(x)$.
Like integrals, start with the rate of change and find the function
by **separation of variables**.

$$
f(x) = \frac{dy}{dx}
$$

Separate the variables,

$$
f(x){dx} = {dy}
$$

Integrate both sides of the equation,

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
y(x) = x^2 + 3x + C
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
\frac{\partial u}{\partial t} = \frac{\partial^2 u}{\partial x^2}
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
y''(x) = \frac{d^2y}{dx^2} = 2x
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

Using **separation of variables**, divide both sides by $P$,

Divide both sides by $P$,

$$
\frac{dP}{dt} = kP
$$

$$
\frac{1}{P} \cdot \frac{dP}{dt} = kP \cdot \frac{1}{P}
$$

$$
\frac{dP}{dtP} = k
$$

Multiply both sides by $dt$,

$$
dt \cdot \frac{dP}{dtP} = k \cdot dt
$$

$$
\frac{dP}{P} = k \, dt
$$

Integrate both sides,

$$
\int \frac{dP}{P} = \int k \, dt
$$

$$
\ln P = kt + C
$$

Solve for $P$ by exponentiating both sides,

$$
P(t) = e^{kt + C} = e^C \cdot e^{kt}
$$

Since $e^C$ is just a constant, we rename it $P_0$ (the initial
population at $t = 0$),

$$
P(t) = P_0 e^{kt}
$$

To find the growth rate constant $k$ given a known population at
some time $t$,

$$
k = \frac{1}{t} \ln \frac{P(t)}{P_0}
$$

As an example, if we have an initial population of $100$ bacteria
that doubles every $3$ hours,

$$
P_0 = 100 \quad \text{and} \quad P(3) = 2 P_0 = 200
$$

the constant $k$ would be,

$$
k = \frac{1}{3} \ln \frac{200}{100} = \frac{1}{3} \ln 2 = 0.231
$$

Therefore, the population at time $t$ would be,

<table align="center"><tr><td>

$$P(t) = 100 \, e^{0.231 t}$$

</td></tr></table>

### RADIOACTIVE DECAY

Consider a radioactive substance that decays at a rate proportional
to the current amount. The differential equation that models this
decay is structurally identical to population growth — just with a
negative rate constant.

$$
N'(t) = \frac{dN}{dt} = -kN
$$

where $N$ is the amount of substance at time $t$ and $k > 0$ is the
decay rate constant. The negative sign indicates the quantity is
decreasing.

Following the same separation of variables steps as population growth,
the solution is,

$$
N(t) = N_0 e^{-kt}
$$

where $N_0$ is the initial amount at time $t = 0$.

A common way to describe radioactive decay is the **half-life**
$T_{1/2}$ — the time it takes for half the substance to decay.
Setting $N(T_{1/2}) = \frac{N_0}{2}$ and solving for $T_{1/2}$,

$$
T_{1/2} = \frac{\ln 2}{k}
$$

As an example, if a substance has a half-life of $5$ years,

$$
k = \frac{\ln 2}{5} = 0.139
$$

Therefore, the amount remaining at time $t$ would be,

<table align="center"><tr><td>

$$N(t) = N_0 \, e^{-0.139 t}$$

</td></tr></table>

### EXPONENTIAL LEARNING RATE DECAY

In machine learning, the learning rate $\eta$ controls how big a
step gradient descent takes when updating weights. It's often
decreased over time during training so the model takes smaller
steps as it converges on a solution.

A common schedule is **exponential decay**, which is modeled by
the same differential equation as radioactive decay,

$$
\frac{d\eta}{dt} = -k\eta
$$

where $\eta$ is the learning rate at training step $t$ and $k > 0$
is the decay rate constant.

Following the same separation of variables steps as population growth,
the solution is,

$$
\eta(t) = \eta_0 \, e^{-kt}
$$

where $\eta_0$ is the initial learning rate at $t = 0$.

As an example, with an initial learning rate $\eta_0 = 0.1$ and a
decay rate $k = 0.01$, the learning rate at training step $t = 100$
would be,

<table align="center"><tr><td>

$$\eta(100) = 0.1 \cdot e^{-0.01 \cdot 100} = 0.1 \cdot e^{-1} = 0.0368$$

</td></tr></table>

This is the same exponential decay form as radioactive decay — just
applied to a different quantity. The learning rate starts at $0.1$
and decays toward $0$ as training progresses, letting the model
make finer and finer adjustments as it converges.

See the
[math behind training mlp neural networks cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/ai-fundamentals/math-behind-training-mlp-neural-networks-cheat-sheet)
for how the learning rate $\eta$ is used in gradient descent.

### GRADIENT FLOW (CONTINUOUS GRADIENT DESCENT)

In machine learning, gradient descent is typically written as a
discrete update rule,

$$
\theta_{new} = \theta - \eta \nabla f(\theta)
$$

But it has a continuous-time form expressed as a differential equation,

$$
\frac{d\theta}{dt} = -\nabla f(\theta)
$$

This is called **gradient flow**. The discrete update rule is what
you get when you approximate this differential equation with finite
steps — and the learning rate $\eta$ plays the role of the step size
in that approximation.

So neural network training is, fundamentally, **the discrete
approximation of a differential equation**. The learning rate
$\eta$ controls how closely the discrete steps follow the
continuous gradient flow path:

* Small $\eta$ → closer to the true gradient flow, but slower
* Large $\eta$ → faster but may overshoot or oscillate

This is why learning rate scheduling (decaying $\eta$ over time)
matters: you take large steps early to make fast progress, then
smaller steps later to closely track the gradient flow as you
approach a minimum.

See the
[math behind training mlp neural networks cheat sheet](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/ai-fundamentals/math-behind-training-mlp-neural-networks-cheat-sheet)
for the full gradient descent derivation.
