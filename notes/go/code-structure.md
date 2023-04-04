Reviewed: 2023-04-04

Good app structure

* makes it easier to reason about the code
* isolates what a programmer is working on w/o keeping the entire codebase in their head
* prevents bugs by decoupling components
* makes it easier to write useful tests

# Organizing by context

Some lanaguages, like Ruby, have a standard way of organizing code based on the architectural patterns being used. MVC, for example. In Go, usually we structure the code by reflecting the domain we're implemeting. We don't base the code structure on the scaffolding but on the specific types in the domain of the project we are working on. For example:

```
student/
  ... (some go files here)
teacher/
  ... (some go files here)
```

```go
package student

type Lesson struct {
  Name         string     // Name of the lesson, eg: "How to run a test"
  Video        string     // URL to the video for this lesson. Empty if the user
                          // doesn't have access to this.
  SourceCode   string     // URL to the source code for this lesson.
  CompletedAt  *time.Time // A boolean representing whether or not the lesson
                          // was completed by this user.
  // + more
}
```

```go
package teacher

// Using inline structs for brevity in this example
type Lesson struct {
  Name string
  // A video's URL can be constructed dynamically (and in some cases with time
  // limited access tokens) using this information.
  Video struct {
    Provider string // Youtube, Vimeo, etc
    ExternalID string
  }
  // Information needed to determine the URL of a repo/branch
  SourceCode struct {
    Provider string // Github, Gitlab, etc
    Repo     string // eg "gophercises/quiz"
    Branch   string // eg "solution-p1"
  }
  // Used to determine if a user has access to this lesson.
  // Usually a string like "twg-base", then when a user purchases
  // a course license they will have these permission strings linked to
  // their account. Prob not the most efficient way to do things, but works
  // well enough for now and makes it really easy to make packages down the
  // road that provide access to multiple courses.
  Requirement string
}
```

I am opting to go this route because I believe these two contexts will vary enough to justify the separation, but I also suspect that neither will grow to be large enough to justify any further organization.

Rather than overthinking it, I find it more useful to pick something that looks like a reasonably good fit and adapt it if needed.

# Packages as layers

This is very similar to approach called "hexagonal architecture".

You can mix-n-match it with the strategy described above.

At a high level, we define our resources and the services we use to interact with them:

```go
package teacher

type Lesson struct {
  // ... same as before
}

type LessonStore interface {
  Create(*Lesson) error
  // ...
}
```

Using `Lesson` and `LessonStore` we can't run our program, but we can write all of the core logic withour worrying about how it's implemented.

When we are ready to implement an interface, we add a new layer to our app:

```go
package sql

import "github.com/johndoe/my-app/teacher"

type TeacherLessonStore struct { ... }

func (ls *TeacherLessonStore) Create(lesson *teacher.Lesson) error { ... }
```

Using interfaces like this definitely makes it easier to test smaller pieces of code, but that only matters if it provides real benefits. Otherwise we end up writing interfaces, decoupling code, and creating new packages with no real benefit. Basically, we are creating busywork for ourselves.

---

Source: https://changelog.com/posts/on-go-application-structure

More: https://youtu.be/spKM5CyBwJA
