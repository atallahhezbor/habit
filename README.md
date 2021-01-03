# habit 

`habit` is a CLI tool for tracking and visualizing your habits.

What sets it apart from other task managers is its focus on habitual, repetitive tasks, as opposed to tasks that go away after completioni, offering a means to maintain, track, and visualize your progress on tasks that serve more long term, personal goals.

`habit` understands that some items need a little work every day, and seeing that progress grow is good motivation to keep that daily work up!

## Screenshots
![list-example](docs/images/list-example.png)


## Quick Start

### Install
You can download the latest release for your platform [here](https://github.com/atallahhezbor/habit/releases/latest) and place it somewhere on your `PATH`

or if you have go installed you can simply run
```
go get github.com/atallahhezbor/habit
```

### Create/Track
Once installed, create a habit to track a goal. For example:

`habit start 'Going for walks' --tag health --short walking`

Track your progress on this habit with the following

`habit tick walking`

Visualize your progress like so

```
$ habit hist
Task    |Ticks
walking |■|■|
```
and you'll see a symbol for each occurrence

## Features

### Create Habits

Habit are input via CLI and tagged with a category.

It will make the most semantic sense for you to use a [gerund phrase](https://en.wikipedia.org/wiki/Gerund) 
Example: `habit start 'Reaching out to friends' --tag social --short chatting`

A `tag` will group the habit with other habits in the same tag when displayed.
A `shortName` is an optional shorter way to refer to the habit when updating it. If not supplied, the first word will be used.


### Track Habit Progress 

Use a "tick" to track progress you've made toward a longterm goal / set of habits.
Example: `habit tick chatting`


### Visualize 

Beautiful colorized visualizations of your progress are at the heart of this project! What better way to be proud of the progress you've made?

There are two ways to visualize, `list` and `hist`

#### `list`
Output all your habits, grouped and colorized by the `tag` you have assigned them, along with the time since you last `tick`ed one.

Example: `habit list`

#### `hist`
Output a "histogram" of sorts to show your dilligence towards all your habits. This is very naive currently. Proper aggregation is soon to come.

Example: `habit hist`

### Get Prompted!

If you're at a loss for what to do, you can ask `habit` for a suggestion, and it'll output a random habit of yours

Example:
```
$ habit suggest
Hmm, you about you try listening to a new podcast?
```



