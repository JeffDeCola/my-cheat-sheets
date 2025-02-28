# NEURAL NETWORKS CHEAT SHEET

_This is a high level overview of Neural Networks._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#overview)
* [MATHEMATICAL MODEL OF A NEURON](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#mathematical-model-of-a-neuron)
  * [THE SUMMATION FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#the-summation-function)
  * [THE ACTIVATION FUNCTION](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#the-activation-function)
* [ARCHITECTURES](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#neural-architectures)
  * [PERCEPTRON (P)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#perceptron-p)
  * [MULTI-LAYER PERCEPTRON (MLP)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#multi-layer-perceptron-mlp)
  * [AUTO ENCODER (AE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#auto-encoder-ae)
  * [RECURRENT NEURAL NETWORKS (RNN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#recurrent-neural-networks-rnn)
  * [CONVOLUTIONAL NEURAL NETWORKS (CNN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#convolutional-neural-networks-cnn)
  * [GENERATIVE ADVERSARIAL NETWORKS (GAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#generative-adversarial-networks-gan)

Documentation and Reference

* [artificial intelligence](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet)
cheat sheet
* [my-neural-networks](https://github.com/JeffDeCola/my-neural-networks/tree/main)
  * [perceptron-simple-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/perceptron-simple-example)
  * [mlp-classification-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-classification-example)
  * [mlp-image-recognition-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-regression-example)
  * [mlp-regression-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-image-recognition-example)
* [the-math-behind-training-mlp-neural-networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/the-math-behind-training-mlp-neural-networks.md)

## OVERVIEW

A neural network is

* A series of algorithms that endeavors to recognize
underlying relationships in a set of data through a process that mimics the
way the human brain operates.
* The neural network itself is not an algorithm, but rather a framework for many
different machine learning algorithms to work together and process complex data inputs.
* Such systems "learn" to perform tasks by considering examples, generally without
being programmed with any task-specific rules.

## MATHEMATICAL MODEL OF A NEURON

The basic building block of a neural network is the neuron.
Each neuron is composed of two units, a summation unit and an activation unit.

An artificial neuron consists of the following components,

* Input(s) $x_{[i]}$
* Applies a weight to that input $w_{[i]}$
* **Summation Function**
  * Weighted sum of inputs
  * Adds a bias term to the weighted sums $b$
  * $s = f(x,w) = \sum_{i=1}^{n} x_i w_i + b$
* **Activation Function**
  * The activation function determines if the neuron will fire or not
  * Applies a non-linear activation function to the sum
  * $y = f(s)$
* Outputs the result $y$

![IMAGE mathematical model of a neuron - IMAGE](../../../../../docs/pics/neural-networks-mathematical-model-of-a-neuron.svg)

It’s important to stress that this model of a biological neuron is very coarse.
For example, there are many different types of neurons, each with different
properties. The dendrites in biological neurons perform complex nonlinear
computations. The synapses are not just a single weight, they’re a
complex non-linear dynamical system. The exact timing of the output
spikes in many systems is known to be important, suggesting that
the rate code approximation may not hold.

### THE SUMMATION FUNCTION

The summation function is the first step in the process of a neuron.
It takes the weighted inputs and sums them up. Then adds a bias term.

$$
s = f(x,w) = x_1 w_1 + x_2 w_2 + \cdots + x_n w_n + b
$$

or

$$
s = f(x) = \sum_{i=1}^{n} x_i w_i + b
$$

Where,

* $x_1, x_2, \ldots, x_n$ are the input values
* $w_1, w_2, \ldots, w_n$ are the weights
* $b$ is the bias

The main function of bias is to provide every node with a trainable
constant value in addition to the normal inputs
that the node receives.

### THE ACTIVATION FUNCTION

The activation function is the function that determines if the
neuron will fire or not. It is the second step in the process of a neuron.
The purpose is to introduce **non-linearity** into the output of a neuron.
This is important because most real-world data is non-linear
and we want our neural network to be able to model and understand real world data.

There are several activation functions you may encounter in practice
I like to use _**s**_ as the input to the activation function to
represent the weighted sum of the inputs plus bias.
Here are a few common activation functions,

* [Sigmoid Function](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-sigmoid-function-using-the-quotient-rule)
  * Output Range: (0, 1)
  * Usage: Often used in the output layer of neural networks for binary
    classification problems
  * The power of the sigmoid function is **once you know the value of
    the function, you already know the derivative of the function**

$$
\begin{aligned}
\sigma(x) &= \frac{1}{1 + e^{-x}} \\
\sigma'(x) &= \sigma(x) \cdot (1 - \sigma(x))
\end{aligned}
$$

<p align="center">
    <img src="svgs/sigmoid-function.svg"
    align="middle"
</p>

* [Tanh Function](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/other/stem/math/pure/changes/calculus-cheat-sheet#the-tanh-function-using-the-quotient-rule)
  * Output Range: (-1, 1)
  * Usage: Often used in hidden layers of neural networks
  * The tanh function is a scaled version of the sigmoid function
  * The tanh function is zero-centered, which helps the model learn
    faster and prevents the gradients from vanishing
  * Also, the derivative is just a simple function of the original function

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

## ARCHITECTURES

Nomenclature,

![IMAGE - neural networks nomenclature - IMAGE](../../../../../docs/pics/neural-networks-nomenclature.svg)

### PERCEPTRON (P)

My
[perceptron-simple-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/perceptron-simple-example)
neural network written in go.

* DESCRIPTION:
  * The simplest form of a neural network
  * Known as a single-layer neural network
  * Has only 2 layers: An Input Layer and an Output Layer
  * It gets its name from performing the human-like
    function of perception, seeing and recognizing images
* FUNCTION:
  * It takes an input and calculates the weighted input for each node
  * Afterward, it uses an activation function (mostly a sigmoid function)
    for classification purposes
* TRAINING:
  * Iteratively feed the network with our training data multiple times
  * Each time the network has seen the full set of training data,
    an epoch has passed.
  * It normally takes many epochs until a weight vector w can be learned
    to linearly separate our two classes of data
* APPLICATION:
  * Basic Classification tasks
  * Simple Linear Regression tasks
  * Example: Predicting if a student will pass or fail based on
    the number of hours they studied

![IMAGE - neural networks perceptron - IMAGE](../../../../../docs/pics/neural-networks-perceptron.svg)

### MULTI-LAYER PERCEPTRON (MLP)

Also known as a Feed-Forward Neural Network (FF).
My
[mlp-classification-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp-classification-example)
neural network written in go.

* DESCRIPTION:
  * Data moves in one direction
  * No loops in the network
  * Has three layers: An Input Layer, Hidden Layers, and an Output Layer
  * The Input Layer provides the initial data for the network
  * The Output Layer provides the final output of the network
  * The hidden layers have no connection to the outside world
  * The hidden Layers are where the network learns the patterns in the data
* FUNCTION:
  * It takes an input and calculates the weighted input for each node.
  * Afterward, it uses an activation function (mostly a sigmoid function)
    for classification purposes
* TRAINING:
  * Adjusting the weights of the neurons to minimize the error between
    the predicted output and the actual output
  * This process is typically performed using the backpropagation algorithm
  * In backpropagation, the error is propagated back through the network
    to update the weights. The gradient of the loss function with
    respect to each weight is calculated, and the weights are
    adjusted using gradient descent.
  * Gradient Descent is an optimization algorithm used to minimize
    some function by iteratively moving in the direction of steepest
    descent as defined by the negative of the gradient
* APPLICATION:
  * Simple Classification tasks
  * Regression tasks
  * Pattern Recognition
  * Speech Recognition
  * Example: Predicting student performance based on the number of hours
    they studied, their midterm grade, and their last test grade

![IMAGE - neural networks feed-forward - IMAGE](../../../../../docs/pics/neural-networks-feed-forward.svg)

### AUTO ENCODER (AE)

* DESCRIPTION:
  * An unsupervised learning algorithm
  * The number of hidden nodes is less than the number of input nodes
  * The number of input nodes is equal to the number of output nodes
* FUNCTION:
  * The Encoder compresses the input data
  * The Decoder decompresses the compressed data
* TRAINING:
  * The network is trained to minimize the difference between the input and output
  common patterns and generalize the data.
* APPLICATION:
  * AEs to find common patterns and generalize the data
  * Dimensionality Reduction
  * Image Compression
  * Example: Reducing the dimensionality of an image

![IMAGE - neural networks auto encoder - IMAGE](../../../../../docs/pics/neural-networks-auto-encoder.svg)

### RECURRENT NEURAL NETWORKS (RNN)

* DESCRIPTION:
  * A variation of a feed-forward neural network
  * A type of neural network where the output from the previous step is fed as input
    to the current step
  * This allows the network to have memory
* FUNCTION:
  * The network can remember important information about the input it received
    a few steps back
  * This is useful for tasks that depend on context
* TRAINING:
  * The network is trained using the backpropagation algorithm
  * The error is propagated back through the network to update the weights
* APPLICATION:
  * Time Series Prediction
  * Speech Recognition
  * Language Translation
  * Example: Predicting the next word in a sentence

![IMAGE - neural networks recurrent - IMAGE](../../../../../docs/pics/neural-networks-recurrent.svg)

### CONVOLUTIONAL NEURAL NETWORKS (CNN)

* DESCRIPTION:
  * A type of neural network that is primarily used for image classification
  * It is a feed-forward neural network
  * It is made up of neurons that have learnable weights and biases
* FUNCTION:
  * The network takes an input image, assigns importance to various aspects
    in the image and be able to differentiate one from the other
  * The network can learn to detect edges, shapes, and textures
* TRAINING:
  * The network is trained using the backpropagation algorithm
  * The error is propagated back through the network to update the weights
* APPLICATION:
  * Image Recognition
  * Image Classification
  * Object Detection
  * Image Segmentation
  * Example: Classifying images of cats and dogs

![IMAGE - neural networks convolutional - IMAGE](../../../../../docs/pics/neural-networks-convolutional.svg)

### GENERATIVE ADVERSARIAL NETWORKS (GAN)

* DESCRIPTION:
  * A type of neural network that is used to generate new data
  * It consists of two networks, a Generator and a Discriminator
  * The Generator generates new data instances
  * The Discriminator evaluates the data instances
* FUNCTION:
  * The Generator tries to generate data that is indistinguishable from real data
  * The Discriminator tries to distinguish between real data and fake data
  * The Generator and Discriminator are trained simultaneously
  * The Generator is trained to fool the Discriminator
  * The Discriminator is trained to detect the fake data
* TRAINING:
  * The Generator and Discriminator are trained simultaneously
  * The Generator is trained to fool the Discriminator
  * The Discriminator is trained to detect the fake data
* APPLICATION:
  * Image Generation
  * Image Enhancement
  * Image-to-Image Translation
  * Example: Generating new images of human faces

![IMAGE - neural networks generative adversarial - IMAGE](../../../../../docs/pics/neural-networks-generative-adversarial.svg)
