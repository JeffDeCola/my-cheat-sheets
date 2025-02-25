# CALCULUS CHEAT SHEET

_The study of continuous change._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#overview)
* [LIMITS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#limits)
* [DIFFERENTIAL CALCULUS (THE DERIVATIVE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#differential-calculus-the-derivative)
  * [USING LIMITS TO FIND THE DERIVATIVE](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#using-limits-to-find-the-derivative)
* [INTEGRAL CALCULUS (THE INTEGRAL)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#integral-calculus-the-integral)
  * [USING LIMITS TO FIND THE INTEGRAL](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#using-limits-to-find-the-integral)
* [AN EXAMPLE OF USING DERIVATIVES AND INTEGRALS](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#an-example-of-using-derivatives-and-integrals)

## OVERVIEW

Calculus is a branch of mathematics that studies continuous change.
It was developed independently by Isaac Newton and
Gottfried Leibniz in the 17th century, and it has become one of
the most powerful and widely-used mathematic tools ever created.

Two branches of Calculus,

* Differential Calculus
  * Rate of change and slopes
  * **Derivatives** - One quantity changes with respect to another
  * Example: If you have a function describing the position of a falling object,
    its derivative gives you the object's velocity, and the derivative of
    velocity gives you acceleration.

* Integral Calculus
  * Opposite of differential calculus
  * Finding area under curve.
  * **Integrals** - Accumulation of a quantity over an interval
  * Example: If you have a function describing the velocity of a falling object,
    its integral gives you the object's position.

## LIMITS

Limits are the foundation of calculus. They describe the value a function
approaches as the input approaches a certain value.

It is written as,

$$
lim_{x \to a} f(x) = L
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

The limit of `x² + 2` as `x` approaches `3` is,

$$
lim_{x \to 3} (x² + 2) = 11
$$

But why do we need limits?  Because some functions are not defined at certain points.

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
lim_{x \to 0} (\frac{1}{x}) = \infty
$$

So how does this apply to calculus?  Well, the derivative of a function is
defined as the limit of the average rate of change of the function as the
interval over which the rate of change is calculated approaches zero.

## DIFFERENTIAL CALCULUS (THE DERIVATIVE)

Derivatives are the rate of change of a function.
They describe how one quantity changes with respect to another.

The derivative of a function

$$
y = f(x)
$$

is written as

$$
\frac{dy}{dx} = f'(x)
$$

where `dy` is the change in `y`, `dx` is the change in `x`, and `f'(x)` is the
derivative of `f(x)`.

For example, given the function,

$$
y = f(x) = x² + 2
$$

The derivative of `x²` is `2x` (see next section on how we did this).

$$
\frac{dy}{dx} = f'(x) = 2x
$$

So if you're at `x = 2`, the slope of the tangent line is `4`.

<p align="center">
    <img src="svgs/f-of-x-equals-x-squared-plus-2-showing-a-tangent-at-x-equals-2.svg"
    align="middle"
</p>

The derivative of a function is the slope of the tangent line to the curve
at a given point.

### USING LIMITS TO FIND THE DERIVATIVE

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
f'(x) &= \lim_{{\Delta x \to 0}} \frac{((x+\Delta x)^2 + 2) - (x^2 + 2)}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{(x+\Delta x)^2 + 2 - x^2 - 2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{x^2 + 2x\Delta x + \Delta x^2 + 2 - x^2 - 2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} \frac{2x\Delta x + \Delta x^2}{\Delta x} \\
f'(x) &= \lim_{{\Delta x \to 0}} 2x + \Delta x \\
f'(x) &= 2x
\end{aligned}
$$

For example, we can see as x approaches 3, the derivative is 6.

## INTEGRAL CALCULUS (THE INTEGRAL)

Integrals are the area under a curve. They describe the accumulation of a
quantity over an interval.

The integral of a function

$$
y = f(x)
$$

is written as

$$
\int f(x) dx = F(x) + C
$$

where `F(x)` is the antiderivative of `f(x)` and `C` is the constant of
integration.

For example, given the function,

$$
y = f(x) = 2x
$$

The integral of `2x` is `x²` (see next section on how we did this).

$$
\int 2x dx = x² + C
$$

So if you're looking for the area under the curve of `2x`, you would find the
area of a triangle with base `x` and height `2x`.

### USING LIMITS TO FIND THE INTEGRAL

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
\int f(x) dx = \lim_{{\Delta x \to 0}} \sum_{i=1}^{n} f(x_i) \Delta x
$$

So giving the example above of $f(x) = 2x$, the integral is

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

And we can see the integral of `2x` is `x²`.

## AN EXAMPLE OF USING DERIVATIVES AND INTEGRALS

Lets take everything we learned and use an example of a car driving.
Instead of using `x` and `y`, we will use `t` for time and `s` for distance.
This will show us velocity at a given time and the distance traveled.

For example, given a car driving at a constant acceleration of `4 m/s²`,
starting at `1 m/s` at `t = 0`.

$$
a(t) = 4
$$

The velocity of the car at time `t` is the integral of acceleration.

$$
v(t) = \int a(t) dt = 4t + C
$$

Given the car starts at `1 m/s` at `t = 0`, we can solve for `C`.

$$
1 = 4(0) + C
$$

$$
C = 1
$$

So the velocity of the car at time `t` is `4t + 1`.

The distance traveled by the car at time `t` is the integral of velocity.

$$
s(t) = \int v(t) dt = \int (4t + 1) dt = 2t² + t + C
$$

Given the car starts at `0 m` at `t = 0`, we can solve for `C`.

$$
0 = 2(0)² + 0 + C
$$

$$
C = 0
$$

So the distance traveled by the car at time `t` is `2t² + t`.
