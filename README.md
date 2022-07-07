# bitebiten

Technically, it's `bit ebiten` not `bite biten`, though I usually prounounce it the latter way in my head. :)

`bit` for bit-101 - that's me.

`ebiten` for the `ebitengine` game engine (which used to be just `ebiten`). https://ebiten.org

`bitebiten` adds some useful features to `ebitengine`:

* A state machine, which you can think of as different scenes. For example, you can have an init state, a game state, a help state, high score state, etc. Each state acts as an ebitengine `Game` so you can swap states to go to the different screens or scenes.
* Asset manager/loader for images, sounds, fonts
* An actor model (think sprite), with additional actor types for images, sprite sheet animations, explosions. Probably more to come here.
* Some utils modules for stuff like color, points, and some geometry stuff.

This is *not* any kind of entity component system, in case you're wondering.

## Todo:

* Stacks for the state machine, so you can push a new state on, then pop it off and go back to the previous state and have it in the exact same ... state .. where you left it.
* More actor types
* Possibly integrate `gg` (https://github.com/fogleman/gg) to do some dynamic shape drawing. Not feasible to do dynamic drawing on every frame, but the ability to pre-create complex shapes at runtime would be nice.
