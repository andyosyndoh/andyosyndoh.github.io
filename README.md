## Groupie-tracker

### Introduction

* Groupie-tracker is a go project which consists on receiving a given [API]("https://groupietrackers.herokuapp.com/api") and manipulate the data contained in it, in order to create a site, displaying the information.

### Description

* The project consicts on building a user friendly website where the user can display the bands info through several data visualizations. (such as :cards)

* This project also focuses on the creation of events/actions and on their visualization.

### Features

* The given API consists of:

1. Artists (containing information about some bands and artists like their name(s), image, in which year they began their activity, the date of their first album and the members).

2. Locations (consists in their last and/or upcoming concert locations).

3. Dates (consists in their last and/or upcoming concert dates).

4. Relation (which does the link between all the other parts, artists, dates and locations).

### usage

  * To use the program, one needs to follow the steps below:

  1. Clone into the repository:
  ```bash
  git clone https://learn.zone01kisumu.ke/git/aosindo/groupie-tracker.git
  ```

  2. change directory to `groupie-tracker`

  ```bash
  cd groupie-tracker
  ```
  3. change directory to `cmd`

   ```bash
  cd cmd
  ```

4. run the command:
```bash
go run .
```
5. go to the web port given, @http://localhost:8080

## Authors

[aosindo](https://github.com/andyosyndoh)

[vandisi](https://github.com/Vinolia-E)