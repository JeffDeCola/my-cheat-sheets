# NEURAL NETWORKS CHEAT SHEET

_This is a high level overview of Neural Networks._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#overview)
* [PERCEPTRON (P)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#perceptron-p)
* [FEED-FORWARD (FF) / MULTI-LAYER PERCEPTRON (MLP)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#feed-forward-ff--multi-layer-perceptron-mlp)
* [AUTO ENCODER (AE)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#auto-encoder-ae)
* [RECURRENT NEURAL NETWORKS (RNN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#recurrent-neural-networks-rnn)
* [CONVOLUTIONAL NEURAL NETWORKS (CNN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#convolutional-neural-networks-cnn)
* [GENERATIVE ADVERSARIAL NETWORKS (GAN)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#generative-adversarial-networks-gan)

## OVERVIEW

A neural network is a series of algorithms that endeavors to recognize
underlying relationships in a set of data through a process that mimics the
way the human brain operates.

Nomenclature,

![IMAGE - neural networks nomenclature - IMAGE](../../../../docs/pics/neural-networks-nomenclature.svg)

## PERCEPTRON (P)

* DESCRIPTION:
  * The simplest form of a neural network
  * Known as a single-layer neural network
  * Has only 2 layers: An Input Layer and an Output Layer
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

![IMAGE - neural networks perceptron - IMAGE](../../../../docs/pics/neural-networks-perceptron.svg)

## FEED-FORWARD (FF) / MULTI-LAYER PERCEPTRON (MLP)

* DESCRIPTION:
  * Data moves in one direction
  * No loops in the network
  * Consists of three layers, an Input Layer, Hidden Layers, and an Output Layer
  * The hidden layers have no connection to the outside world
* FUNCTION:
  * It takes an input and calculates the weighted input for each node.
  * Afterward, it uses an activation function (mostly a `sigmoid function`)
    for classification purposes.
* TRAINING:
  * Adjusting the weights of the neurons to minimize the error between
    the predicted output and the actual output
  * This process is typically performed using the `backpropagation algorithm`
  * In backpropagation, the error is propagated back through the network
    to update the weights. The gradient of the loss function with
    respect to each weight is calculated, and the weights are
    adjusted using gradient descent.
  * `Gradient Descent` is an optimization algorithm used to minimize
    some function by iteratively moving in the direction of steepest
    descent as defined by the negative of the gradient
* APPLICATION:
  * Simple Classification tasks
  * Regression tasks
  * Pattern Recognition
  * Speech Recognition

![IMAGE - neural networks feed-forward - IMAGE](../../../../docs/pics/neural-networks-feed-forward.svg)

## AUTO ENCODER (AE)

* DESCRIPTION:
  * An unsupervised learning algorithm
  * The number of hidden cells is less than the number of input cells
  * The number of input cells is equal to the number of output cells
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

![IMAGE - neural networks auto encoder - IMAGE](../../../../docs/pics/neural-networks-auto-encoder.svg)

## RECURRENT NEURAL NETWORKS (RNN)

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

![IMAGE - neural networks recurrent - IMAGE](../../../../docs/pics/neural-networks-recurrent.svg)

## CONVOLUTIONAL NEURAL NETWORKS (CNN)

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

![IMAGE - neural networks convolutional - IMAGE](../../../../docs/pics/neural-networks-convolutional.svg)

## GENERATIVE ADVERSARIAL NETWORKS (GAN)

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

![IMAGE - neural networks generative adversarial - IMAGE](../../../../docs/pics/neural-networks-generative-adversarial.svg)
