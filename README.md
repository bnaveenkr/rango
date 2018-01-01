# rango
Rango is a very simple raytracer written in golang. This is an attempt to learn go. 

The first pass is very similar to C raytracer I have written earlier https://github.com/bnaveenkr/raytracer


Some points for performance sake:
* Image Size: for quick testing, use something like 320x240
* Recursion Depth: Can be kept 0 for no reflections/refractions
* Cone and Cylinder can use smaller resolution, default image uses 32
* Sphere resolution, unlike Cone and Cylinder, is power of 4


### How to run
```
cd raytracer
go build
./raytracer
```

### Output Image
![Raytraced Image](https://raw.githubusercontent.com/bnaveenkr/rango/develop/raytracer/output.png)
 
