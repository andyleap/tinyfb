TinyFB
======

Inspired by TinyPTC, TinyFB offers a really simple way to put graphics on the screen.
TinyFB is ideal for simple "demo-like" apps and similar.

To use:

t := tinyfb.New("Title", 640, 480)

go t.Run()

then loop, calling t.Update() with a correctly sized image.RGBA{}