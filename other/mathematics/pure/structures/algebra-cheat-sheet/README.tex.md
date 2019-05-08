# ALGEBRA CHEAT SHEET

```text
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`algebra` _tbd._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

$$[ x^n + y^n = z^n]$$

and 

$$E=mc^2$$

and

```text
$$
E=mc^2
$$
```

and

$$E=mc^2 + \int_a^a x\, dx$$

and

$$\text{S}_1(N) = \sum_{p=1}^N \text{E}(p)$$

and

$$\text{S}1(N) = \sum{p=1}^N \text{E}(p)$$

and

$$f(x) = \int_{-\infty}^\infty\hat f(\xi)\,e^{2 \pi i \xi x}\,d\xi$$

and

$$\binom{n}{k} = \frac{n!}{k!(n-k)!}$$

and

$$
\cos (2\theta) = \cos^2 \theta - \sin^2 \theta
$$

and

$$
  \int_{-\infty}^\infty e^{-x^2/2} dx = \sqrt{2\pi}
$$

Here's a display level formula

This math is inline $`a^2+b^2=c^2`$.

This is on a separate line
```math
a^2+b^2=c^2
```

# fractals
Mandelbulb fractal 3D in pure Javascript based on [GPU.JS](https://github.com/gpujs/gpu.js)

[Run Mandelbulb!](https://kamil-kielczewski.github.io/fractals/mandelbulb.html)

## Calc Ray

Input parameters ([left-handed coordinate system](https://en.wikipedia.org/wiki/Cartesian_coordinate_system#In_three_dimensions)): 
* $E = [Ex,Ey,Ez]$ - camera (eye) position at point 
* $T= [Tx,Ty,Tz]$ - target point where camera looks  
* $w=[wx,wy,wz]$ - camera vertical normalized vector which idicates where is up and were is down (not shown on picture, usually equal [0,1,0]). 
* $\theta \in [0,\pi)$ - field of view (slacar value, for human eye $\approx 90^\circ$)
* $k$ - number of pixels on screen width 
* $m$ - number of pixels screen in height 

**IDEA**: lets find position of center of each pixel $P_{ij}$ which allows us to easily find ray which starts at $E$ and go thought that pixel. To do it we find first $P_{1m}$ and find others by move on vievports plane.

**ASSUMPTION**: $d=1$ which simplify calculations but not change the result (because $r_{ij}$ is normalized and viewport size is determined by $k,m$ and $\theta$) 

**PRECALCULATIONS**: First we calculate normalized vectors $v_n, b_n$ from picutre (which are parallel to viewport plane and give as direction for shifting)

$$t = T-E, \qquad b = w\times t  $$

$$ 
t_n = \frac{t}{||t||}, \qquad
b_n = \frac{b}{||b||}, \qquad
v_n = t_n\times b_n \\ 
$$

notice: $C=E+t_nd$, then we calculate viewport size divided by 2 and including aspect ratio $\frac{m}{k}$

$$g_x=\frac{h_x}{2} =d \tan \frac{\theta}{2}, \qquad  g_y =\frac{h_y}{2} = g_x \frac{m}{k}$$

and then we calculate shifting vectors $q_x,q_y$ on viewport $x,y$ direction and viewport left upper pixel

$$ q_x = \frac{2g_x}{k-1}b_n, \qquad
q_y = \frac{2g_y}{m-1}v_n, \qquad
p_{1m} = t_n d - g_xb_n - g_yv_n$$

**CALCULATIONS**: notice that $P_{ij} = E + p_{ij}$ and ray $R_{ij} = P_{ij} -E = p_{ij}$ so normalized ray $r_{ij}$ is 

$$ p_{ij} = p_{1m} + q_x(i-1) + q_y(j-1)$$ (1)
$$ r_{ij} = \frac{p_{ij}}{||p_{ij}||} $$ (2)


and,

```math
\cos (2\theta) = \cos^2 \theta - \sin^2 \theta
```
## COMMON MATH EQUATIONS

Note, dollar sign delimiters not shown,

Einsteins famous equation,

```txt
E=mc^2
```

$$
E=mc^2
$$

Pythagorean theorem,

```txt
x^n + y^n = z^n
```

$$
x^n + y^n = z^n
$$

A sample Integral,

```txt
\int_{a}^{b} x^2 dx
```

$$
\int_{a}^{b} x^2 dx
$$

Limits,

```txt
\lim_{x\to\infty} f(x)
```

$$
\lim_{x\to\infty} f(x)
$$

Some trigonometry,

```txt
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
```

$$
\sin(a + b ) = \sin(a)\cos(b) + \cos(a)\sin(b)
$$

Fractions (binomial coefficients),

```txt
\binom{n}{k} = \frac{n!}{k!(n-k)!}
```

$$
\binom{n}{k} = \frac{n!}{k!(n-k)!}
$$

Brackets,

```txt
\left( \frac{x}{y} \right)
```

$$
\left( \frac{x}{y} \right)
$$

Bracket array,

```txt
\left(
 \begin{array}{ccc}
  1 & 2 & 3\\
  4 & 4 & 9\\
  1 & -8 & 2
 \end{array}
\right)
```

$$
\left(
 \begin{array}{ccc}
  1 & 2 & 3\\
  4 & 4 & 9\\
  1 & -8 & 2
 \end{array}
\right)
$$

Arrays in Brackets with spacing (\qquad),

```txt
\left(
 \\begin{array}{ccc}
  1 & 2 & 3\\
  4 & 5 & 9\\
  1 & -8 & 2
 \\end{array}
\right)
\quad
\left\{
  \begin{array}{ccc}
  1 & 5 & 8\\

  0 & 2 & 4\\
  3 & 3 & -8
  \end{array}
\right\}
```

$$
\left(
 \begin{array}{ccc}
  1 & 2 & 3\\
  4 & 5 & 9\\
  1 & -8 & 2
 \end{array}
\right)
\quad
\left\{
  \begin{array}{ccc}
  1 & 5 & 8\\
  0 & 2 & 4\\
  3 & 3 & -8
  \end{array}
\right\}
$$