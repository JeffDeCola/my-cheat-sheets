# ARTIFICIAL INTELLIGENCE CHEAT SHEET

[![jeffdecola.com](https://img.shields.io/badge/website-jeffdecola.com-blue)](https://jeffdecola.com)
[![MIT License](https://img.shields.io/:license-mit-blue.svg)](https://jeffdecola.mit-license.org)

_An high level overview of artificial intelligence (AI)
and a quick dive into machine learning (ML)._

Table of Contents

* [OVERVIEW](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#overview)
* [TYPES OF AI](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#types-of-ai)
  * [CAPABILITY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#capability)
  * [FUNCTIONALITY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#functionality)
* [ARTIFICIAL INTELLIGENCE HIERARCHY](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#artificial-intelligence-hierarchy)
* [MACHINE LEARNING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#machine-learning)
  * [CLASSICAL LEARNING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#classical-learning)
  * [DEEP LEARNING (NEURAL NETWORKS)](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#deep-learning-neural-networks)
  * [ENSEMBLE LEARNING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#ensemble-learning)
  * [REINFORCEMENT LEARNING](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet#reinforcement-learning)

Documentation and Reference

* [neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md)
* [my-neural-networks](https://github.com/JeffDeCola/my-neural-networks/tree/main)
  * [the-math-behind-training-mlp-neural-networks](https://github.com/JeffDeCola/my-neural-networks/tree/main/math/the-math-behind-training-mlp-neural-networks)
  * [mlp-classification-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp/mlp-classification-example)
  * [mlp-image-recognition-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp/mlp-image-recognition-example)
  * [mlp-regression-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/mlp/mlp-regression-example)
  * [perceptron-simple-example](https://github.com/JeffDeCola/my-neural-networks/tree/main/perceptron/perceptron-simple-example)
* [the-math-behind-training-mlp-neural-networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/the-math-behind-training-mlp-neural-networks.md)

## OVERVIEW

Artificial Intelligence (AI) is the term used to describe the ability of a
machine to perform a cognitive processes.

## TYPES OF AI

AI can be looked at in two ways; How good it is (Capability) and what it can
do (Functionality).

### CAPABILITY

* **Machine Learning** (Narrow AI / Weak AI)
  * Designed for a specific task
* **Machine Intelligence** (General AI / Strong AI)
  * Can perform any intellectual task that a human can
* **Machine Consciousness** (Superintelligent AI)
  * AI that surpasses human intelligence

### FUNCTIONALITY

* **Reactive Machines**
  * Can only react to current situations
* **Limited Memory**
  * Can learn from the past
* **Theory of Mind**
  * Can understand human emotions
* **Self Awareness**
  * AI that has consciousness

## ARTIFICIAL INTELLIGENCE HIERARCHY

Artificial Intelligence can be broken down into many subcategories.
I will focus on Machine Learning.

* **Artificial Intelligence (AI)**
  * Machine Learning (ML)
  * Robotics
  * Natural Language Processing (NLP)
  * Computer Vision
  * Expert Systems
  * Fuzzy Logic
  * Evolutionary Algorithms

![IMAGE - artificial-intelligence-hierarchy - IMAGE](../../../../../docs/pics/software/development/artificial-intelligence-hierarchy.svg)

## MACHINE LEARNING

Machine Learning is a subset of AI that allows a system to learn from data
rather than through explicit programming.

There are many types of Machine Learning such as,

* **Machine Learning (ML)**
  * **Classical Learning**
    * Supervised Learning
      * Classification
      * Regression
    * Unsupervised Learning
      * Clustering
      * Pattern Recognition
      * Dimensionality Reduction Visualization
  * **Deep Learning** (Neural Networks)
    * [Perceptron (P)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#perceptron-p)
    * [Multi-Layer Perceptron (MLP)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#multi-layer-perceptron-mlp)
    * [Auto Encoders (AE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#auto-encoder-ae)
    * [Recurrent Neural Networks (RNN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#recurrent-neural-networks-rnn)
    * [Convolutional Neural Networks (CNN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#convolutional-neural-networks-cnn)
    * [Generative Adversarial Networks (GAN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#generative-adversarial-networks-gan)
    * etc...
  * **Ensemble Learning**
    * Boosting
    * Bagging
    * Stacking
  * **Reinforcement Learning**

### CLASSICAL LEARNING

Classical learning is the most common type of machine learning.
It can be broken down into two categories: Supervised and Unsupervised.
Both systems learn from mistakes and makes predictions on new data.

#### SUPERVISED LEARNING

In supervised learning, the system is given input data and the correct output
The system learns to predict the output from the input data.

* **Classification**
  * Used for predicting a category
  * Example: Is this email spam or not spam?
* **Regression**
  * Used for predicting a quantity
  * Example: What will the price of a house be?

#### UNSUPERVISED LEARNING

In unsupervised learning, the system is given input data without the correct
output and the system learns to find patterns in the data.

* **Clustering**
  * Used for grouping similar data
  * Example: Grouping customers by purchasing behavior
* **Pattern Recognition**
  * Used for finding patterns in data
  * Example: Finding patterns in stock market data
* **Dimensionality Reduction Visualization**
  * Used for reducing the number of features in data
  * Example: Reducing the number of features in an image

### DEEP LEARNING (NEURAL NETWORKS)

Deep learning is a subset of machine learning that uses neural networks.
It is the most popular type of machine learning today. A Neural Network is a
working system at the heart of a Deep Learning algorithm that helps it process
raw data.

* **[Neural networks](https://github.com/JeffDeCola/my-cheat-sheets/tree/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md)**
  * Like a human brain
  * A network of nodes that are interconnected
  * Each node is a neuron that is connected to other neurons
  * The network can learn from data, improve itself and make decisions

There are many types of neural networks such as,

* **[Perceptron (P)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#perceptron-p)**
  * The simplest form of a neural network
  * Used for simple classification tasks
  * Example: Is this email spam or not spam?
* **[Multi-Layer Perceptron (MLP)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#multi-layer-perceptron-mlp)**
  * Data moves in one direction
  * No loops in the network
  * Used for simple classification tasks
  * Example: Is this email spam or not spam?
* **[Auto Encoder (AE)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#auto-encoder-ae)**
  * Used for reducing dimensionality
  * Example: Reducing the number of features in an image
* **[Recurrent Neural Networks (RNN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#recurrent-neural-networks-rnn)**
  * Used for sequential data
  * Example: Predicting the next word in a sentence
* **[Convolutional Neural Networks (CNN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#convolutional-neural-networks-cnn)**
  * Used for image data
  * Example: Image recognition
* **[Generative Adversarial Networks (GAN)](https://github.com/JeffDeCola/my-cheat-sheets/blob/master/software/development/software-architectures/artificial-intelligence/artificial-intelligence-cheat-sheet/neural-networks.md#generative-adversarial-networks-gan)**
  * Used for generating new data
  * Example: Generating new images

### ENSEMBLE LEARNING

Ensemble learning is when you take multiple machine learning algorithms
and combine them to make a more powerful model. There are three types,

* **Boosting**
  * Used for reducing bias
  * Example: AdaBoost
* **Bagging**
  * Used for reducing variance
  * Example: Random Forest
* **Stacking**
  * Used for reducing bias and variance
  * Example: Stacked Generalization

### REINFORCEMENT LEARNING

Reinforcement learning is a type of machine learning that is used to
train an agent to make a sequence of decisions. Examples include
game playing and robotics.
