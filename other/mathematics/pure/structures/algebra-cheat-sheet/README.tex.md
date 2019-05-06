# ALGEBRA CHEAT SHEET

```
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
<p align="center"><img alt="$$&#10;\frac{n!}{k!(n-k)!} = {n \choose k}&#10;$$" src="svgs/32737e0a8d5a4cf32ba3ab1b74902ab7.png?invert_in_darkmode" align=middle width="127.89183pt" height="39.30498pt"/></p>

> It is well known that if <img alt="$ax^2 + bx + c =0$" src="svgs/162f63774d8a882cc15ae1301cfd8ac0.png?invert_in_darkmode" align=middle width="119.01186pt" height="26.70657pt"/>, then <img alt="$x = \frac{-b \pm \sqrt{b^2- 4ac}}{2a}$" src="svgs/584fa2612b78129d140fb208e9d76ae9.png?invert_in_darkmode" align=middle width="112.3584pt" height="33.20526pt"/>.

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

<p align="center"><img src="/tex/raysMatrix.png" align=middle /></p>

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

$$ p_{ij} = p_{1m} + q_x(i-1) + q_y(j-1)$$
$$ r_{ij} = \frac{p_{ij}}{||p_{ij}||} $$

