# THE MATH BEHIND TRAINING MLP NEURAL NETWORKS

_The math behind training a multi-layer perceptron (MLP) neural network._

Table of Contents

* [OVERVIEW]()

Documentation and Reference

* [artificial intelligence](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet)
cheat sheet
* [neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md)
cheat sheet
* [my-neural-networks](https://github.com/JeffDeCola/my-neural-networks/tree/main)
  * [perceptron-simple-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/perceptron-simple-example)
  * [mlp-classification-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-classification-example)
  * [mlp-image-recognition-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-regression-example)
  * [mlp-regression-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-image-recognition-example)

## OVERVIEW

A neural networks is trained using a dataset. The dataset is split into two parts,

* Training data
  * Used to train the neural network
  * Adjusts the weights and biases of the neural network
  * Minimizes the error between the predicted output and the actual output
* Testing data
  * Used to test the neural network
  * See how well the neural network performs

## MATHEMATICAL MODEL OF NEURAL NETWORK

Since we are dealing with neural networks, let's quicky review our model of a neuron.
Check out a more detailed view in my cheat sheets
[mathematical model of a neuron](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#mathematical-model-of-a-neuron).

![IMAGE training-multi-layer-perceptron-neural-network-model-neuron IMAGE](../../../../../docs/pics/training-multi-layer-perceptron-neural-network-model-neuron.svg)

### THE SUMMATION FUNCTION

The summation function takes the weighted inputs and sums them up.
Then adds a bias term.

$$
\begin{aligned}
s = f(x,w,b) &= x_1 w_1 + x_2 w_2 + \cdots + x_n w_n + b \\
&= \sum_{i=1}^{n} x_i w_i + b
\end{aligned}
$$

Where,

* $x_1, x_2, \ldots, x_n$ are the input values
* $w_1, w_2, \ldots, w_n$ are the weights
* $b$ is the bias
* $n$ is the number of inputs

### THE ACTIVATION FUNCTION

The activation function $f(s)$ is a non-linear function that introduces non-linearity
to the output of a neuron.

Two common activation functions are, the sigmoid function and the tanh function.

$$
\begin{aligned}
\tanh(x) &= \frac{e^x - e^{-x}}{e^x + e^{-x}} \\
\tanh'(x) &= 1 - \tanh^2(x)
\end{aligned}
$$

<p align="center">
    <img src="svgs/tanh-function.svg"
    align="middle"
</p>

### THE OUTPUT

And the output of the neuron is simply the output of the activation function,

$$
f = f(s)
$$
