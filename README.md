# link creator for imgproxy
This is a quick scipt without any help or any helpful error checking to generate links given a set of inputs.
![image](https://user-images.githubusercontent.com/38912521/129556604-394f8bc0-2948-47ee-8294-a8061bd1a1d6.png)

## Docker
```
docker build -t some_random_name .
docker run some_random_name http://127.0.0.1:8000/test.jpg
```
![image](https://user-images.githubusercontent.com/38912521/129559400-11f0fd9d-6baa-4b87-bb84-3d2feda5b852.png)


## Usage
```bash
# configure env variables. They can also be passed as arguments to the program
export SERVER="192.168.60.4:9001" # server that runs imgproxy
export KEY="" # your key
export SALT="" # your salt
export EXTENSION="jpg" # jpg, png etc...
./imgproxy_link_creator --path "http://10.10.10.10/path/to/image.jpg" # path to image accessible only by imgproxy
```

## Build
```
go build .
```

## Options
Documentation for these options can be found [here](https://docs.imgproxy.net/generating_the_url_basic?id=generating-the-url-basic)
- Path      string
- Server    string
- Key       string 
- Salt      string 	
- Resize    string 
- Width     int    
- Height    int    
- Gravity   string 
- Enlarge   int    
- Extension string 
