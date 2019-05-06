# ALGEBRA CHEAT SHEET

```
*** THIS CHEAT SHEET IS UNDER CONSTRUCTION - CHECK BACK SOON ***
```

`algebra` _tbd._

View my entire list of cheat sheets on
[my GitHub Webpage](https://jeffdecola.github.io/my-cheat-sheets/).

## OVERVIEW

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/fc6919e1cdaceba8467bb92e2ae5ef49.svg?invert_in_darkmode&sanitize=true" align=middle width=104.39689589999999pt height=16.438356pt/></p>

and 

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/03422e896c78b96b32028535ab59553e.svg?invert_in_darkmode&sanitize=true" align=middle width=63.09925874999999pt height=14.202794099999998pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/ceffe4dd6c1483b331e74a5ab1d9f9bb.svg?invert_in_darkmode&sanitize=true" align=middle width=141.22829819999998pt height=38.242408049999995pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/bcf11d902834d0fb9b2631c1ad2c0c74.svg?invert_in_darkmode&sanitize=true" align=middle width=124.93718159999999pt height=49.84491765pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/55298f326ecbd61ebb680520670162bb.svg?invert_in_darkmode&sanitize=true" align=middle width=176.6573589pt height=26.301595649999996pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/3ab852f5e7cecd438a7853e4b8955343.svg?invert_in_darkmode&sanitize=true" align=middle width=184.13079299999998pt height=39.61228755pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/487eac933217aafc0beb7d4173e75983.svg?invert_in_darkmode&sanitize=true" align=middle width=127.98480255pt height=39.452455349999994pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/446d4a8687e5be8933b9363a733a7f82.svg?invert_in_darkmode&sanitize=true" align=middle width=171.96327555pt height=18.439728449999997pt/></p>

and

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/3a0cfe0035d8bf51952f960d938e0a28.svg?invert_in_darkmode&sanitize=true" align=middle width=153.72738149999998pt height=39.61228755pt/></p>


Here's a display level formula
<p align="center"><img alt="<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/993ca78854d842d63d78390d9d51b2b1.svg?invert_in_darkmode&sanitize=true" align=middle width=174.7063263pt height=39.452455349999994pt/></p>" src="svgs/32737e0a8d5a4cf32ba3ab1b74902ab7.png?invert_in_darkmode&sanitize=true" align=middle width="127.89183pt" height="39.30498pt"/></p>

> It is well known that if <img alt="<img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/162f63774d8a882cc15ae1301cfd8ac0.svg?invert_in_darkmode&sanitize=true" align=middle width=119.34141284999998pt height=26.76175259999998pt/>" src="svgs/162f63774d8a882cc15ae1301cfd8ac0.png?invert_in_darkmode&sanitize=true" align=middle width="119.01186pt" height="26.70657pt"/>, then <img alt="<img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/584fa2612b78129d140fb208e9d76ae9.svg?invert_in_darkmode&sanitize=true" align=middle width=112.44128444999998pt height=33.20539200000001pt/>" src="svgs/584fa2612b78129d140fb208e9d76ae9.png?invert_in_darkmode&sanitize=true" align=middle width="112.3584pt" height="33.20526pt"/>.

This math is inline <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/6495ae5cdcbfe3f9de1ae7dba4839638.svg?invert_in_darkmode&sanitize=true" align=middle width=96.12240494999999pt height=26.76175259999998pt/>.

This is on a separate line
```math
a^2+b^2=c^2
```


# fractals
Mandelbulb fractal 3D in pure Javascript based on [GPU.JS](https://github.com/gpujs/gpu.js)

[Run Mandelbulb!](https://kamil-kielczewski.github.io/fractals/mandelbulb.html)

## Calc Ray

Input parameters ([left-handed coordinate system](https://en.wikipedia.org/wiki/Cartesian_coordinate_system#In_three_dimensions)): 
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/d62fbe219457fce60682a162b4ecbab4.svg?invert_in_darkmode&sanitize=true" align=middle width=124.40236709999998pt height=24.65753399999998pt/> - camera (eye) position at point 
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/aecdc767c97bdaf680b7c57d54dbe69d.svg?invert_in_darkmode&sanitize=true" align=middle width=119.63090204999997pt height=24.65753399999998pt/> - target point where camera looks  
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/356dfd3a8b76763cdf8121889b66694a.svg?invert_in_darkmode&sanitize=true" align=middle width=120.91704239999997pt height=24.65753399999998pt/> - camera vertical normalized vector which idicates where is up and were is down (not shown on picture, usually equal [0,1,0]). 
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/19f3ce3a8c773f5be82ab19a1c2a061a.svg?invert_in_darkmode&sanitize=true" align=middle width=64.7087859pt height=24.65753399999998pt/> - field of view (slacar value, for human eye <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/9e55bdb7fdbca783335bc66dc13b0ed2.svg?invert_in_darkmode&sanitize=true" align=middle width=40.52514509999999pt height=22.63850490000001pt/>)
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/63bb9849783d01d91403bc9a5fea12a2.svg?invert_in_darkmode&sanitize=true" align=middle width=9.075367949999992pt height=22.831056599999986pt/> - number of pixels on screen width 
* <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/0e51a2dede42189d77627c4d742822c3.svg?invert_in_darkmode&sanitize=true" align=middle width=14.433101099999991pt height=14.15524440000002pt/> - number of pixels screen in height 

<p align="center"><img src="/tex/raysMatrix.png" align=middle /></p>

**IDEA**: lets find position of center of each pixel <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/ea2306a8d1b4985c89ba5be73a9c95d9.svg?invert_in_darkmode&sanitize=true" align=middle width=21.309055349999994pt height=22.465723500000017pt/> which allows us to easily find ray which starts at <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/84df98c65d88c6adf15d4645ffa25e47.svg?invert_in_darkmode&sanitize=true" align=middle width=13.08219659999999pt height=22.465723500000017pt/> and go thought that pixel. To do it we find first <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/db44cdb37ee861207d561cdae6019ef0.svg?invert_in_darkmode&sanitize=true" align=middle width=28.77104339999999pt height=22.465723500000017pt/> and find others by move on vievports plane.

**ASSUMPTION**: <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/4cfe92e893a7541f68473ecb08419237.svg?invert_in_darkmode&sanitize=true" align=middle width=38.69280359999998pt height=22.831056599999986pt/> which simplify calculations but not change the result (because <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/92e0822b1528090efc2435d2ae60c9ee.svg?invert_in_darkmode&sanitize=true" align=middle width=18.17172884999999pt height=14.15524440000002pt/> is normalized and viewport size is determined by <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/2cc0a928b241d5552cd750dc8470c1ae.svg?invert_in_darkmode&sanitize=true" align=middle width=30.81434564999999pt height=22.831056599999986pt/> and <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/27e556cf3caa0673ac49a8f0de3c73ca.svg?invert_in_darkmode&sanitize=true" align=middle width=8.17352744999999pt height=22.831056599999986pt/>) 

**PRECALCULATIONS**: First we calculate normalized vectors <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/26bd4dc9b67fd60527afb9c639dc214a.svg?invert_in_darkmode&sanitize=true" align=middle width=39.40270784999999pt height=22.831056599999986pt/> from picutre (which are parallel to viewport plane and give as direction for shifting)

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/b3d0c1b791d706af329c9fd81a905c64.svg?invert_in_darkmode&sanitize=true" align=middle width=180.30961409999998pt height=14.611878599999999pt/></p>

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/5a1fa75c84e371498151e7e4e9670264.svg?invert_in_darkmode&sanitize=true" align=middle width=301.4880429pt height=37.9216761pt/></p>

notice: <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/f0395a1f114dcfd23b32a17a0ef8756d.svg?invert_in_darkmode&sanitize=true" align=middle width=91.45563569999999pt height=22.831056599999986pt/>, then we calculate viewport size divided by 2 and including aspect ratio <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/ef3b0585953dc3772b7ba3fb57cf79d5.svg?invert_in_darkmode&sanitize=true" align=middle width=11.664849899999997pt height=22.853275500000024pt/>

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/b0f38d3421473f7499124f1748be18ef.svg?invert_in_darkmode&sanitize=true" align=middle width=285.18904095pt height=33.81208709999999pt/></p>

and then we calculate shifting vectors <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/d97c5c89afb85662f7498de1024b7cf9.svg?invert_in_darkmode&sanitize=true" align=middle width=37.33838129999999pt height=14.15524440000002pt/> on viewport <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/0acac2a2d5d05a8394e21a70a71041b4.svg?invert_in_darkmode&sanitize=true" align=middle width=25.350096749999988pt height=14.15524440000002pt/> direction and viewport left upper pixel

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/3f33d3e81b32d7b4cb6fa00f939b2258.svg?invert_in_darkmode&sanitize=true" align=middle width=452.79895485pt height=34.3600389pt/></p>

**CALCULATIONS**: notice that <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/3981580a168bd9f393801557969972df.svg?invert_in_darkmode&sanitize=true" align=middle width=96.24790395pt height=22.465723500000017pt/> and ray <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/b3a4487e8fbec1e56edb389f47f51421.svg?invert_in_darkmode&sanitize=true" align=middle width=142.224291pt height=22.465723500000017pt/> so normalized ray <img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/92e0822b1528090efc2435d2ae60c9ee.svg?invert_in_darkmode&sanitize=true" align=middle width=18.17172884999999pt height=14.15524440000002pt/> is 

<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/251e3f493dba4940017c24fea49272e5.svg?invert_in_darkmode&sanitize=true" align=middle width=235.67744144999998pt height=17.031940199999998pt/></p>
<p align="center"><img src="/other/mathematics/pure/structures/algebra-cheat-sheet/tex/4561437ecc388c74c260300f7f069163.svg?invert_in_darkmode&sanitize=true" align=middle width=80.9965761pt height=34.177354199999996pt/></p>

