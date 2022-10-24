# react_app

This example was created using the steps described [here](https://bazelbuild.github.io/rules_nodejs/examples#react).

There are a few things that I still don't like about this:
- Adds another package.json which can be confusing (as deps are still defined in root package.json).
- Doesn't work well with ibazel.
- Tests aren't working yet.

I may spend some time to try to improve these things in the future, but for now.. This is what I've been able to do.
