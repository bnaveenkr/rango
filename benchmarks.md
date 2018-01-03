# Comparision GO vs C

From the below stats it appears Go is ~10X farter than C version



### Golang ([rango](https://github.com/bnaveenkr/rango/)) - 146.98s
 
```
➜  raytracer git:(develop) ✗ time ./raytracer
   Tracing started, Rays: 1228800, TriangleCount: 724, Recursion: 2
   Tracing done
   Writing image to disk
   Done writing the file to disk
   ./raytracer  146.98s user 6.28s system 99% cpu 2:33.44 total
```

### C ([bnaveenkr/raytracer](https://github.com/bnaveenkr/raytracer/)) - 1017.86s
```
➜  raytracer git:(master) time ./a.out
Image: output.ppm, Size: 1280x960
Initial Ray Count: 1228800, Triangle Count: 724, Recursion Depth: 2
Raytracing...
Done, writing image...
Done
./a.out  1017.86s user 1.57s system 99% cpu 17:01.15 total

```
