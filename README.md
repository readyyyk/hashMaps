# Random Image API

## Examples

| Link                                           | Image                                                                                                 |
|------------------------------------------------|-------------------------------------------------------------------------------------------------------|
| `https://readyyyk-randimg.fly.dev/picsum`      | <img src="https://readyyyk-randimg.fly.dev/picsum" alt="example">                                     |
| `https://readyyyk-randimg.fly.dev/hashmap`     | <img src="https://readyyyk-randimg.fly.dev/hashmap" alt="example" width="64px">                       |
| `.../hashmap?seed=example`                     | <img src="https://readyyyk-randimg.fly.dev/hashmap?seed=example" alt="example" width="64px">          |
| `.../hashmap?seed=example&w=8&h=10` (w=8 h=10) | <img src="https://readyyyk-randimg.fly.dev/hashmap?seed=example&w=8&h=10" alt="example" width="64px"> |
| `.../picsum?seed=example&w=64&h=100`           | <img src="https://readyyyk-randimg.fly.dev/picsum?seed=example&w=64&h=100" alt="example">             |

## Usage
host: https://readyyyk-randimg.fly.dev/

**https://`host`/hashmap?`...[url params]`**

**https://`host`/picsum?`...[url params]`**

> returns:
> - hashmap - `svg` image
> - picsum - `jpeg` image

> `width` and `height` defaults:
>  - hashmap - 7x7
>  - picsum - 64x64

> `Seed` default value is `Unix time`

### URL parameters
| Parameter    | Type   | Value              |
|--------------|--------|--------------------|
| `w` (width)  | Int    | 1-100 for hashmaps |
| `h` (height) | Int    | 1-100 for hashmaps |
| `seed`       | String | any                |
