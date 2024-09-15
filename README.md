# Attack on Titan Multi-Agent Simulation

## Overview

This project, conducted over one mounth in an academic context for the Multi-Agent Simulation course at Université de Technologie de Compiègne implements a multi-agent simulation inspired by the popular anime **Attack on Titan**.

The simulation explores dynamic interactions between agents, focusing on environment perception and real-time decision-making. It aims to create a realistic and engaging experience that mirrors the strategic elements found in the anime.

## Key Features

- **Multi-Agent Framework**: Designed using a robust multi-threaded architecture that allows agents to operate independently while interacting with their environment.
- **Agent Behavior**: Each agent category (soldier, titan, civilian, or special character) has its own behavior based on what is perceived, allowing for diverse interactions and strategies.
- **Realistic Perception**: The simulation incorporates a realistic perception model that takes into account obstacles and angles where visibility is limited, depending on the relative position to the obstacles.
- **Real-Time Decision Making**: The simulation incorporates algorithms that allow agents to make decisions based on environmental cues and interactions with other agents.

## Installation

To set up the project, clone the repository and install any necessary dependencies:

```bash
git clone <repository-url>
cd <repository-directory>
```

## Usage

Once the environment is set up, run the main simulation script to initiate the multi-agent environment:

```bash
go run main.go
```
