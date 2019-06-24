
# 👋 Meet Pathbot

Like its cousin Mazebot, Pathbot is always getting lost.

While exploring the deeper subterranean floors of the Noops, Inc. complex, Pathbot came upon a [TI-99/4A] with a collection of 1970s-era text adventure games. Pathbot eventually went back to work on the Noops assembly line, but it just couldn't stop thinking about hunting Wumpuses, avoiding pits, and escaping from the Grue.


# 🤖 API


## ✳️ How to play

`POST` to `https://api.noopschallenge.com/pathbot/start` to get started.

Every location in pathbot's many mazes will return you a JSON object with the following fields:

- **status** (string) either "in-progress" or "finished",
- **message** (string) a message for you from Pathbot.
- **exits** (string array) An array containing the available exits (N, S, E, or W).
- **description** A description of the room you are in.
- **mazeExitDirection** (string) The general direction toward the exit from the maze - one of (N, S, E, W, NW, NE, SW, SE).
- **mazeExitDistance** (number). The minimum number of rooms between your current location and the maze exit,
- **locationPath** (string) the API path for the location. POST your next move back to this path.

See the [API documentation](./API.md) for more information.

# Starter Kits

## Go interactive client

Pathbot has included a go client that will let you explore its mazes.
Can you write a program that can escape the maze?

# ✨ A few ideas

- **Create an automated solver**: Humans can be pretty good at solving mazes, but they'll never be as fast as a well-tuned computer. You could start from the [the included ruby script](./mazebot.rb) or start from scratch in another language.  If you create a solver in another language, please share it with the Noops!


Pathbot can't wait to see what you make!

More about Pathbot here: https://noopschallenge.com/challenges/pathbot